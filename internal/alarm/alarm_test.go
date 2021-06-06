// !test
package alarm_test

import (
	"github.com/ozoncp/ocp-project-api/internal/alarm"
	"testing"
	"time"
)

func TestAlarm(t *testing.T) {
	a := alarm.NewAlarm(time.Second * 3)

	alarms := a.Alarms()

	closeAlarm := func(d time.Duration) {
		c := time.After(d)
		<-c
		a.Close()
	}

	go closeAlarm(time.Second * 5)

	count := 0
	wanted := 1
	for {
		_, ok := <-alarms
		if ok {
			count++
		} else {
			break
		}
	}
	if count != wanted {
		t.Errorf("Alarmer is very fast or lazy, got count of alarms = %d, wanted = %d\n", count, wanted)
	}

	a = alarm.NewAlarm(time.Second * 1)
	a.ResetTimeout(time.Second * 3)

	alarms = a.Alarms()

	go closeAlarm(time.Second * 7)

	count = 0
	wanted = 2
	for {
		_, ok := <-alarms
		if ok {
			count++
		} else {
			break
		}
	}
	if count != wanted {
		t.Errorf(
			"ResetTimeout: Alarmer is very fast or lazy, got count of alarms = %d, wanted = %d\n",
			count, wanted)
	}
}
