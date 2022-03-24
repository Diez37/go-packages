package repeater

import (
	"sync/atomic"
	"time"
)

type Task interface {
	Name() string
	Process() Process
	NextTimestamp() int64
	Lock()
	UnLock()
	IsLock() bool
}

type task struct {
	name          string
	delay         int64
	process       Process
	nextTimestamp int64
	lock          uint32
}

func NewTask(name string, delay time.Duration, process Process) Task {
	return &task{
		name:          name,
		delay:         int64(delay.Seconds()),
		process:       process,
		nextTimestamp: time.Now().In(time.UTC).Unix(),
	}
}

func (task *task) Name() string {
	return task.name
}

func (task *task) Process() Process {
	atomic.AddInt64(&task.nextTimestamp, task.delay)

	return task.process
}

func (task *task) NextTimestamp() int64 {
	return atomic.LoadInt64(&task.nextTimestamp)
}

func (task *task) Lock() {
	atomic.StoreUint32(&task.lock, 1)
}

func (task *task) UnLock() {
	atomic.StoreUint32(&task.lock, 0)
}

func (task *task) IsLock() bool {
	return atomic.LoadUint32(&task.lock) == 1
}
