package timer

import (
	"alert"
	"sync"
	"time"
)

var alertHandler *alert.WeChatAlert

func init() {
	alertHandler = alert.NewWeChatAlert()
}

type Timer struct {
	doneChan chan struct{}
	mu       sync.Mutex
	timer    *time.Timer
	doChan   chan bool
	interval int64
}

func NewTimer(do chan bool, interval int64) *Timer {
	ch := make(chan struct{}, 1)
	return &Timer{
		doneChan: ch,
		doChan:   do,
		interval: interval,
		timer:    time.NewTimer(time.Second * time.Duration(interval)),
	}
}

func (t *Timer) Cancel() {
	t.doneChan <- struct{}{}
	t.timer.Stop()
}

func (t *Timer) UpdateInterval(newInterval int64) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.interval = newInterval
	t.timer.Reset(time.Duration(t.interval))
}

func (t *Timer) Schedule() {
	for {
		select {
		case <-t.timer.C:
			t.doChan <- true
			t.timer.Reset(time.Second * time.Duration(t.interval))
		case <-t.doneChan:
			t.doChan <- false
			return
		}
	}
}

func NewSimpleTimer(c int) *time.Timer {
	now := time.Now()

	onTimer := time.Date(now.Year(), now.Month(), now.Day(), c, 0, 0, 0, now.Location())
	if now.Before(onTimer) {
		return time.NewTimer(onTimer.Sub(now))
	}

	//the next day
	onTimer = time.Date(now.Year(), now.Month(), now.Day()+1, c, 0, 0, 0, now.Location())
	return time.NewTimer(onTimer.Sub(now))
}
