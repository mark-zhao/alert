package timer

var minTimeAlert = 3
var maxTimeAlert = 120

type HeartBeat struct {
	doWork chan bool
	task   func() string
	//number of times about alert to client
	times int
	timer *Timer
}

func NewHeartBeat(t int64, hbtask func() string) *HeartBeat {
	ch := make(chan bool, 1)
	return &HeartBeat{
		times:  0,
		doWork: ch,
		task:   hbtask,
		timer:  NewTimer(ch, t),
	}
}

func (hb *HeartBeat) Cancel() {
	hb.timer.Cancel()
}

// timeForAlert control how often messages are sent
func (hb *HeartBeat) timeForAlert(msg string) {
	//heartBeat ok
	if msg == "" {
		return
	}

	if hb.times < minTimeAlert {
		hb.times++
		alertHandler.Alert(msg)
	} else if hb.times >= maxTimeAlert {
		hb.times = 0
	} else {
		hb.times++
	}
}

func (hb *HeartBeat) Run() {
	go hb.timer.Schedule()
	for {
		select {
		case ok := <-hb.doWork:
			if ok {
				hb.timeForAlert(hb.task())
			} else {
				close(hb.doWork)
				return
			}
		default:

		}
	}
}
