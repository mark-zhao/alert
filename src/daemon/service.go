package daemon

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

const (
	optionKeepAlive            = "KeepAlive"
	optionKeepAliveDefault     = true
	optionRunAtLoad            = "RunAtLoad"
	optionRunAtLoadDefault     = false
	optionUserService          = "UserService"
	optionUserServiceDefault   = false
	optionSessionCreate        = "SessionCreate"
	optionSessionCreateDefault = false
	optionLogOutput            = "LogOutput"
	optionLogOutputDefault     = false

	optionRunWait      = "RunWait"
	optionReloadSignal = "ReloadSignal"
	optionPIDFile      = "PIDFile"

	optionSystemdScript = "SystemdScript"
	optionSysvScript    = "SysvScript"
	optionUpstartScript = "UpstartScript"
	optionLaunchdConfig = "LaunchdConfig"
)

// Status represents service status as an byte value
type Status byte

// Status of service represented as an byte
const (
	StatusUnknown Status = iota // Status is unable to be determined due to an error or it was not installed.
	StatusRunning
	StatusStopped
)

type linuxSystemService struct {
	name        string
	detect      func() bool
	interactive func() bool
	new         func(i Interface, platform string, c *Config) (Service, error)
}

func (sc linuxSystemService) String() string {
	return sc.name
}
func (sc linuxSystemService) Detect() bool {
	return sc.detect()
}
func (sc linuxSystemService) Interactive() bool {
	return sc.interactive()
}
func (sc linuxSystemService) New(i Interface, c *Config) (Service, error) {
	return sc.new(i, sc.String(), c)
}

var (
	system  System
	systemRegistry []System
)

func detectSystem() System {
	for _, choice := range systemRegistry{
		if choice.Detect()  == false{
			continue
		}
		return choice
	}
	return nil
}

func ChooseSystem(a ...System){
	systemRegistry = a
	system = detectSystem()
}


func init() {
	ChooseSystem(linuxSystemService{
		name:   "linux-systemd",
		detect: isSystemd,
		interactive: func() bool {
			is, _ := isInteractive()
			return is
		},
		new: newSystemdService,
	},
		linuxSystemService{
			name:   "linux-upstart",
			detect: isUpstart,
			interactive: func() bool {
				is, _ := isInteractive()
				return is
			},
			new: newUpstartService,
		},
		linuxSystemService{
			name:   "unix-systemv",
			detect: func() bool { return true },
			interactive: func() bool {
				is, _ := isInteractive()
				return is
			},
			new: newSystemVService,
		},
	)
}

func isInteractive() (bool, error) {
	// TODO: This is not true for user services.
	return os.Getppid() != 1, nil
}

// New creates a new service based on a service interface and configruation
func New(i Interface, c *Config) (Service, error){
	if len(c.Name) == 0{
		return nil, ErrNameFieldRequired
	}
	if system == nil{
		return nil, ErrNoServiceSystemDetected
	}
	return system.New(i, c)
}

func run(command string, arguments ...string)error{
	_, _, err := runCommand(command, false, arguments...)
	return err
}

func runWithOutput(command string, arguments ...string)(int, string, error){
	return runCommand(command, true, arguments...)
}

func runCommand(command string, readStdout bool, arguments ...string) (int, string, error) {
	cmd := exec.Command(command, arguments...)

	var output string
	var stdout io.ReadCloser
	var err error

	if readStdout {
		// Connect pipe to read Stdout
		stdout, err = cmd.StdoutPipe()

		if err != nil {
			// Failed to connect pipe
			return 0, "", fmt.Errorf("%q failed to connect stdout pipe: %v", command, err)
		}
	}

	// Connect pipe to read Stderr
	stderr, err := cmd.StderrPipe()

	if err != nil {
		// Failed to connect pipe
		return 0, "", fmt.Errorf("%q failed to connect stderr pipe: %v", command, err)
	}

	// Do not use cmd.Run()
	if err := cmd.Start(); err != nil {
		// Problem while copying stdin, stdout, or stderr
		return 0, "", fmt.Errorf("%q failed: %v", command, err)
	}

	// Zero exit status
	// Darwin: launchctl can fail with a zero exit status,
	// so check for emtpy stderr
	if command == "launchctl" {
		slurp, _ := ioutil.ReadAll(stderr)
		if len(slurp) > 0 {
			return 0, "", fmt.Errorf("%q failed with stderr: %s", command, slurp)
		}
	}

	if readStdout {
		out, err := ioutil.ReadAll(stdout)
		if err != nil {
			return 0, "", fmt.Errorf("%q failed while attempting to read stdout: %v", command, err)
		} else if len(out) > 0 {
			output = string(out)
		}
	}

	if err := cmd.Wait(); err != nil {
		exitStatus, ok := isExitError(err)
		if ok {
			// Command didn't exit with a zero exit status.
			return exitStatus, output, err
		}

		// An error occurred and there is no exit status.
		return 0, output, fmt.Errorf("%q failed: %v", command, err)
	}

	return 0, output, nil
}

func isExitError(err error) (int, bool) {
	if exiterr, ok := err.(*exec.ExitError); ok {
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			return status.ExitStatus(), true
		}
	}

	return 0, false
}


func (c *Config)execPath() (string, error){
	if len(c.Executable) == 0{
		return filepath.Abs(c.Executable)
	}
	return os.Executable()
}

// versionAtMost will return true if the provided version is less than or equal to max
func versionAtMost(version, max []int) (bool, error) {
	if comp, err := versionCompare(version, max); err != nil {
		return false, err
	} else if comp == 1 {
		return false, nil
	}
	return true, nil
}

// versionCompare take to versions split into integer arrays and attempts to compare them
// An error will be returned if there is an array length mismatch.
// Return values are as follows
// -1 - v1 is less than v2
// 0  - v1 is equal to v2
// 1  - v1 is greater than v2
func versionCompare(v1, v2 []int) (int, error) {
	if len(v1) != len(v2) {
		return 0, errors.New("version length mismatch")
	}

	for idx, v2S := range v2 {
		v1S := v1[idx]
		if v1S > v2S {
			return 1, nil
		}

		if v1S < v2S {
			return -1, nil
		}
	}
	return 0, nil
}

// parseVersion will parse any integer type version seperated by periods.
// This does not fully support semver style versions.
func parseVersion(v string) []int {
	version := make([]int, 3)

	for idx, vStr := range strings.Split(v, ".") {
		vS, err := strconv.Atoi(vStr)
		if err != nil {
			return nil
		}
		version[idx] = vS
	}

	return version
}



