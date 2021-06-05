package alarm_test

import (
	"fmt"
	"github.com/ozoncp/ocp-project-api/internal/alarm"
	"testing"
	"time"
)

func TestAlarm(t *testing.T) {
	a := alarm.NewAlarm(time.Second * 1)

	alarms := a.Alarms()

	for i := 0; i < 3; i++ {
		select {
		case <-alarms:
			fmt.Println("Alarm ticked")
		}
	}

	a.Close()

	_, ok := <-alarms
	fmt.Printf("Alarm %v\n", ok)

	select {
	case <-alarms:
		fmt.Println("Alarm ticked")
	}
}
