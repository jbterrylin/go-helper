package sequencetaskhelper

import (
	"fmt"
	"sync"

	ophelper "github.com/jbterrylin/go-helper/opHelper"
)

type SequenceTask struct {
	StopWhileMeetErr   bool
	Ch                 chan map[string]interface{}
	SeqNum             int
	SeqNumMutex        sync.RWMutex
	CurrentRunningTask int
	ChStatus           string
	TaskFunc           func(map[string]interface{}) error
	RecordSuccessFunc  func(map[string]interface{})
	RecordFailFunc     func(map[string]interface{})
	DeleteTaskCh       chan func(map[string]interface{}) bool // this is insecure when task is executing, so need to stop execute when end current, then delete
}

func NewSequenceTask(
	stopWhileMeetErr bool,
	seqNum int,
	taskFunc func(map[string]interface{}) error,
	recordSuccessFunc func(map[string]interface{}),
	recordFailFunc func(map[string]interface{}),
) (sequenceTask SequenceTask) {
	sequenceTask = SequenceTask{
		StopWhileMeetErr:  stopWhileMeetErr,
		SeqNum:            ophelper.Or(seqNum, 1),
		ChStatus:          END_SEQUENCE,
		TaskFunc:          taskFunc,
		RecordSuccessFunc: recordSuccessFunc,
		RecordFailFunc:    recordFailFunc,
	}
	sequenceTask.Ch = make(chan map[string]interface{})
	sequenceTask.DeleteTaskCh = make(chan func(map[string]interface{}) bool)

	sequenceTask.taskHandler()
	return
}

func (a *SequenceTask) addCurrentRunningTask() {
	a.SeqNumMutex.Lock()
	a.CurrentRunningTask++
	a.SeqNumMutex.Unlock()
}

func (a *SequenceTask) reduceCurrentRunningTask() {
	a.SeqNumMutex.Lock()
	a.CurrentRunningTask--
	a.SeqNumMutex.Unlock()
}

func (a *SequenceTask) getCurrentRunningTask() int {
	a.SeqNumMutex.RLock()
	defer a.SeqNumMutex.RUnlock()
	return a.CurrentRunningTask
}

func (a *SequenceTask) taskHandler() {
	go func() {
		for {
			if a.ChStatus == START_SEQUENCE {
				if a.getCurrentRunningTask() < a.SeqNum {
					select {
					case value, ok := <-a.Ch:
						if !ok {
							fmt.Println("Ch closed")
							return
						}
						go func(v map[string]interface{}) {
							a.addCurrentRunningTask()
							defer a.reduceCurrentRunningTask()
							err := a.TaskFunc(v)
							if err != nil {
								if a.RecordFailFunc != nil {
									a.RecordFailFunc(v)
								}
								if a.StopWhileMeetErr {
									a.ChStatus = END_SEQUENCE
								}
							} else {
								if a.RecordSuccessFunc != nil {
									a.RecordSuccessFunc(v)
								}
							}
							a.deleteTaskHandler()
						}(value)
					default:
						continue
					}
				}
			}
		}
	}()
}

func (a *SequenceTask) AddTask(data map[string]interface{}) {
	a.Ch <- data
}

func (a *SequenceTask) deleteTaskHandler() {
	select {
	case value, ok := <-a.DeleteTaskCh:
		if !ok {
			fmt.Println("DeleteTaskCh closed")
			return
		}
		a.deleteTask(value)
	default:
		return
	}
}

func (a *SequenceTask) deleteTask(filter func(map[string]interface{}) bool) {
	newCh := make(chan map[string]interface{})
	go func() {
		defer close(newCh)
		for value := range a.Ch {
			if !filter(value) {
				newCh <- value
			}
		}
	}()
	a.Ch = newCh
}

func (a *SequenceTask) DeleteTask(filter func(map[string]interface{}) bool) {
	a.DeleteTaskCh <- filter
}

func (a *SequenceTask) Start() {
	a.ChStatus = START_SEQUENCE
}

func (a *SequenceTask) End() {
	a.ChStatus = END_SEQUENCE
}

func (a *SequenceTask) ReloadTask(data []map[string]interface{}) {
	a.Ch = make(chan map[string]interface{})
	for _, d := range data {
		a.Ch <- d
	}
}
