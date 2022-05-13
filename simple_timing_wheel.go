package timingwheel

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type SimpleTimingWheel struct {
	wheel  [][]Task
	curPos int
	max    int
}

func NewSimpleTimingWheel(max int) *SimpleTimingWheel {
	wheel := make([][]Task, max)
	return &SimpleTimingWheel{
		wheel:  wheel,
		curPos: 0,
		max:    max,
	}
}

func (tw *SimpleTimingWheel) Start() {
	go func() {
		cTicker := time.Tick(time.Second)
		for {
			select {
			case t := <-cTicker:
				log.Printf("t:%v, cur:%d\n", t, tw.curPos)
				tw.curPos++
				if tw.curPos > len(tw.wheel)-1 {
					tw.curPos = 0
				}

				// execute task
				for _, task := range tw.wheel[tw.curPos] {
					go task.Execute(task)
				}
				tw.wheel[tw.curPos] = []Task{}
			}
		}
	}()
}

func (tw *SimpleTimingWheel) Add(delay int, t Task) error {
	if delay > tw.max {
		return errors.New(fmt.Sprintf("delay:%d is bigger than max:%d", delay, tw.max))
	}

	tw.wheel[(tw.curPos+delay)%tw.max] = append(tw.wheel[(tw.curPos+delay)%tw.max], t)
	return nil
}
