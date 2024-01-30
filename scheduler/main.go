package scheduler

import (
	"errors"
	"time"

	temp "github.com/debeando/go-common/time"
)

type Scheduler struct {
	Enable   bool   `yaml:"enable"`
	Start    string `yaml:"start"`
	End      string `yaml:"end"`
	Interval uint16 `yaml:"interval"`
	start    time.Time
	end      time.Time
}

func (s *Scheduler) Do(start func(), end func()) (err error) {
	if !s.Enable {
		return errors.New("Scheduler is disable")
	}

	if s.start, err = temp.StringToTime(s.Start); err != nil {
		return err
	}

	if s.end, err = temp.StringToTime(s.End); err != nil {
		return err
	}

	for {
		time.Sleep(DeltaSeconds() * time.Second)

		if temp.BetweenNow(s.start, s.end) {
			start()
		} else {
			end()
		}

		time.Sleep(time.Duration(s.Interval) * time.Second)
	}
}

func DeltaSeconds() time.Duration {
	return time.Duration(60 - temp.Now().Second())
}
