package timer

import (
	"strings"
	"time"
	"glog"
)


var tenClock = 10
var nightClock = 21
var cephStatusMap = make(map[string]string)
const normalState = "HEALTH_OK"

type StatusTimer struct {
	doWork         chan bool
	updateInterval chan int64
	task           func(map[string]string)
	timer          *Timer
	clock10        *time.Timer
	clock21        *time.Timer
	cephStatusMap  map[string]string
}

func NewTimerStatus(t int64, upCh chan int64, task func(map[string]string)) *StatusTimer {
	ch := make(chan bool, 1)
	return &StatusTimer{
		doWork:         ch,
		updateInterval: upCh,
		task:           task,
		timer:          NewTimer(ch, t),
		clock10:        NewSimpleTimer(tenClock),
		clock21:        NewSimpleTimer(nightClock),
	}
}

func (t *StatusTimer) Cancel() {
	t.timer.Cancel()
}

func (t *StatusTimer) UpdateInterval(newInterval int64) {
	t.timer.UpdateInterval(newInterval)
}

func (t *StatusTimer)checkStatus(){
	var msg string
	for _, m := range cephStatusMap{
		if !strings.Contains(m, normalState){
			msg = msg + m
		}
	}
	if msg != "" {
		alertHandler.Alert(msg)
	}
}

//alertDaily process send ceph cluster status to wechat client right now
func (t *StatusTimer) alertDaily() {
	var msg string
	glog.Info("it's time for alertDaily; ", time.Now().Format("2006-01-02 15:04:05"))

	for r, m := range cephStatusMap {
		if m == "" { //just start progma
			msg = msg + "region:" + r + " on line \n"
		} else { //not first alertDaily
			msg = msg + "\n" + m
		}
	}
	glog.Info("it's time for alertDaily; ", time.Now().Format("2006-01-02 15:04:05"), "; msg:", msg)
	alertHandler.Alert(msg)
}

func (t *StatusTimer) Run() {
	go t.timer.Schedule()
	for {
		select {
		case <-t.clock10.C:
			t.alertDaily()
			t.clock10.Reset(time.Hour * 24)
		case <-t.clock21.C:
			t.alertDaily()
			t.clock21.Reset(time.Hour * 24)
		default:
			select {
			case ok := <-t.doWork:
				if ok {
					t.task(cephStatusMap)
					t.checkStatus()
					cephStatusMap = make(map[string]string)
				} else {
					close(t.doWork)
					t.clock10.Stop()
					t.clock21.Stop()
					return
				}
			case interval := <-t.updateInterval:
				t.UpdateInterval(interval)
			}
		}
	}
}
