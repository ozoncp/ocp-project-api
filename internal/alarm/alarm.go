package alarm

import "time"

type Alarm interface {
	Alarms() <-chan struct{}
	ResetTimeout(d time.Duration)
	Close()
}

func NewAlarm(d time.Duration) Alarm {
	c := make(chan struct{}, 1)
	done := make(chan struct{})

	a := &alarm{
		timeout: d,
		c:       c,
		done:    done,
	}

	go a.startAlarm()
	return a
}

type alarm struct {
	timeout time.Duration
	c       chan struct{}
	done    chan struct{}
}

func (a *alarm) startAlarm() {
	timer := time.After(a.timeout)
	for {
		select {
		case <-timer:
			// if channel is not empty, skip writing
			select {
			case a.c <- struct{}{}:
			default:
			}
			timer = time.After(a.timeout)
		case <-a.done:
			close(a.c)
			close(a.done)
			return
		}
	}
}

func (a *alarm) Alarms() <-chan struct{} {
	return a.c
}

func (a *alarm) ResetTimeout(d time.Duration) {
	a.timeout = d
}

func (a *alarm) Close() {
	a.done <- struct{}{}
}
