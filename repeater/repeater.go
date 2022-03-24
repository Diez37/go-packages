package repeater

import (
	"context"
	"github.com/diez37/go-packages/log"
	"sync"
	"time"
)

type Repeater interface {
	AddProcess(name string, delay time.Duration, process Process) Repeater
	AddTasks(tasks ...Task) Repeater
	Serve(ctx context.Context)
}

type repeater struct {
	logger log.Logger

	wg *sync.WaitGroup

	rwMutex *sync.RWMutex
	tasks   []Task
}

func New(logger log.Logger) Repeater {
	return &repeater{
		logger:  logger,
		wg:      &sync.WaitGroup{},
		rwMutex: &sync.RWMutex{},
	}
}

func (repeater *repeater) AddProcess(name string, delay time.Duration, process Process) Repeater {
	repeater.rwMutex.Lock()
	defer repeater.rwMutex.Unlock()

	repeater.tasks = append(repeater.tasks, NewTask(name, delay, process))

	repeater.logger.Infof("repeater: add new task '%s'", name)

	return repeater
}

func (repeater *repeater) AddTasks(tasks ...Task) Repeater {
	repeater.rwMutex.Lock()
	defer repeater.rwMutex.Unlock()

	repeater.tasks = append(repeater.tasks, tasks...)

	return repeater
}

func (repeater *repeater) Serve(ctx context.Context) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	repeater.wg.Add(1)
	go func(ctx context.Context) {
		defer repeater.wg.Done()

		timestamp := time.Now().In(time.UTC).Unix()

		repeater.logger.Info("repeater: started")

		for {
			select {
			case <-ctx.Done():
				repeater.logger.Info("repeater: shutdown")
				return
			case <-time.After(time.Second):
				timestamp++

				repeater.rwMutex.RLock()

				for _, task := range repeater.tasks {
					if task.NextTimestamp() <= timestamp {
						if task.IsLock() {
							continue
						}

						task.Lock()
						repeater.wg.Add(1)
						go func(ctx context.Context, task Task) {
							defer repeater.wg.Done()
							defer task.UnLock()

							startTime := time.Now().In(time.UTC)

							repeater.logger.Infof("repeater: run task '%s'", task.Name())

							if err := task.Process().Process(ctx); err != nil {
								repeater.logger.Errorf("repeater: task '%s', error - %s", task.Name(), err)
								return
							}

							repeater.logger.Infof(
								"repeater: task '%s' completed in '%s'",
								task.Name(),
								time.Now().In(time.UTC).Sub(startTime),
							)
						}(ctx, task)
					}
				}

				repeater.rwMutex.RUnlock()
			}
		}
	}(ctx)

	repeater.wg.Wait()
}
