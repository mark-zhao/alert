package daemon

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNameFieldRequired = errors.New("Config.Name field is required.")
	ErrNoServiceSystemDetected = errors.New("No service system detected.")
	ErrNotInstalled = errors.New("service is not installed")
)

// KeyValue provides a list of platform specific options.
type KeyValue map[string]interface{}

var tf = map[string]interface{}{
	"cmd": func(s string) string {
		return `"` + strings.Replace(s, `"`, `\"`, -1) + `"`
	},
	"cmdEscape": func(s string) string {
		return strings.Replace(s, " ", `\x20`, -1)
	},
}

// System represents the service manager that is available
type System interface {
	// String returns a description of the system
	String() string

	// Detect returns true if the system is available to use
	Detect() bool

	// Interactive returns false if running under the system service manager
	// and true otherwise
	Interactive() bool

	// New creates a new service of this system
	New(i Interface, c *Config) (Service, error)
}

//Config provides the setup for service. The Name field is required
type Config struct {
	Name 		string
	DisplayName string
	Description string
	UserName 	string
	Arguments 	[]string

	Executable  string
	Dependencies []string
	WorkingDirectory string
	ChRoot 		 string

	//System specific options
	//*POSIX
	//    - SystemdScript string ()                 - Use custom systemd script
	//    - UpstartScript string ()                 - Use custom upstart script
	//    - SysvScript    string ()                 - Use custom sysv script
	//    - RunWait       func() (wait for SIGNAL)  - Do not install signal but wait for this function to return.
	//    - ReloadSignal  string () [USR1, ...]     - Signal to send on reaload.
	//    - PIDFile       string () [/run/prog.pid] - Location of the PID file.
	//    - LogOutput     bool   (false)            - Redirect StdErr & StdOut to files.
	Option KeyValue
}

func (kv KeyValue) bool(name string, defaultValue bool) bool{
	if v, ok := kv[name]; ok{
		if castValue, is := v.(bool); is{
			return castValue
		}
	}
	return defaultValue
}

// int returns the value of the given name, assuming the value is an int.
// If the value isn't found or is not of the type, the defaultValue is returned.
func (kv KeyValue) int(name string, defaultValue int) int{
	if v, ok := kv[name]; ok{
		if castValue, is := v.(int); is{
			return  castValue
		}
	}
	return defaultValue
}

// string returns the value of the given name, assuming the value is a string.
// If the value isn't found or is not of the type, the defaultValue is returned.
func (kv KeyValue) string(name string, defaultValue string) string {
	if v, ok := kv[name]; ok{
		if castValue, is := v.(string); is{
			return castValue
		}
	}
	return defaultValue
}

// float64 returns the value of the given name, assuming the value is a float64.
// If the value isn't found or is not of the type, the defaultValue is returned.
func (kv KeyValue)float64(name string, defaultValue float64) float64 {
	if v, ok := kv[name]; ok{
		if castValue, is := v.(float64); is{
			return castValue
		}
	}
	return defaultValue
}

// funcSingle returns the value of the given name, assuming the value is a func.
// If the value isn't found or is not of the type, the defaultValue is returned.
func (kv KeyValue) funcSingle(name string, defaultValue func()) func(){
	if v, ok := kv[name]; ok{
		if castValue, is := v.(func()); is{
			return castValue
		}
	}
	return defaultValue
}


// Interface represents the service interface for a program. Start runs before
// the hosting process is granted control and Stop runs when control is returned.
//
//	1. OS service manager executes user program.
//	2. User program sees it is executed from a service manager (IsInteractive is false).
//  3. User program calls Service.Run() which blocks.
//  4. Interface.Start() is called and quickly returns.
//  5. User program runs.
//  6. OS service manager signals the user program to stop.
//  7. Interface.Stop() is called and quickly returns.
//		- For a successful exit, os.Exit should not be called in Interface.Stop().
//	8. Service.Run returns.
//	9. User program should quickly exit.
type Interface interface{
	// Start provides a place to initiate the service. The service doesn't not
	// signal a completed start until after this function returns, so the
	// Start function must not take more then a few seconds at most.
	Start(s Service) error

	// Stop provides a place to clean up program execution before it is terminated.
	// It should not take more then a few seconds to execute
	// Stop should not call os.Exit directly in the function.
	Stop(s Service) error
}


// Service represents a service that can be run or controlled.
type Service interface {
	// Run should be called shortly after the program entry point.
	// After Interface.Stop has finished running, Run will stop blocking.
	// After Run stops blocking, the program must exit shortly after.
	Run() error

	// Start signals to the OS service manager the given service should start.
	Start() error

	// Stop signals to the OS service manager the given service should stop.
	Stop() error

	// Restart signals to the OS service manager the given service should stop then start.
	Restart() error

	// Install setups up the given service in the OS service manager. This may require
	// greater rights. Will return an error if it is already installed.
	Install() error

	// Uninstall removes the given service from the OS service manager. This may require
	// greater rights. Will return an error if the service is not present.
	Uninstall() error

	// String displays the name of the service. The display name if present,
	// otherwise the name.
	String() string
}

// ControlAction list valid string texts to use in Control.
var ControlAction = [5]string{"start", "stop", "restart", "install", "uninstall"}

// Control issues control functions to the service from a given action string.
func Control(s Service, action string) error {
	var err error
	switch action {
	case ControlAction[0]:
		err = s.Start()
	case ControlAction[1]:
		err = s.Stop()
	case ControlAction[2]:
		err = s.Restart()
	case ControlAction[3]:
		err = s.Install()
	case ControlAction[4]:
		err = s.Uninstall()
	default:
		err = fmt.Errorf("Unknown action %s", action)
	}
	if err != nil {
		return fmt.Errorf("Failed to %s %v: %v", action, s, err)
	}
	return nil
}

