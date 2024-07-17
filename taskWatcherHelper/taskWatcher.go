package taskwatcher

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
)

type TaskWatcher struct {
	Cron      *cron.Cron
	CronId    cron.EntryID
	RedisConn *redis.Client

	TaskPrefixKey          string // to seperate different task watcher
	UpdateOrphanedTaskCron string

	TickerDuration    time.Duration
	HeartbeatDuration time.Duration // heartbeat should longer than ticker to prevent redis transport speed too slow

	GetRunningTasksFunc  func() []string
	UpdateTaskToFailFunc func(id string, redisErr error)

	ErrorAfterStartTaskWatcher error
}

func NewTaskWatcher(
	redisConn *redis.Client,
	updateOrphanedTaskCron,
	taskPrefixKey string,
	tickerDuration,
	heartbeatDuration time.Duration,
	getRunningTasksFunc func() []string,
	updateTaskToFailFunc func(id string, redisErr error),
) (taskWatcher TaskWatcher) {
	taskWatcher = TaskWatcher{
		RedisConn:              redisConn,
		UpdateOrphanedTaskCron: updateOrphanedTaskCron,
		TaskPrefixKey:          taskPrefixKey,
		TickerDuration:         tickerDuration,
		HeartbeatDuration:      heartbeatDuration,
		GetRunningTasksFunc:    getRunningTasksFunc,
		UpdateTaskToFailFunc:   updateTaskToFailFunc,
	}
	return
}

func (a *TaskWatcher) StartTaskWatcher(uniqueKey string) (stopFunc func(), err error) {
	ticker := time.NewTicker(a.TickerDuration)
	tickerChan := make(chan bool)
	ctx := context.TODO()
	redisKey := fmt.Sprintf("%s%s", a.TaskPrefixKey, uniqueKey)

	err = a.RedisConn.Set(ctx, redisKey, "", a.HeartbeatDuration).Err()
	if err != nil {
		return
	}

	go func() {
		for {
			select {
			case <-tickerChan:
				return
			// interval task
			case <-ticker.C:
				err := a.RedisConn.Set(ctx, redisKey, "", a.HeartbeatDuration).Err()
				if err != nil {
					fmt.Println(err.Error())
					a.ErrorAfterStartTaskWatcher = err
				}
			}
		}
	}()

	stopFunc = func() { //stop ticker
		tickerChan <- true
		ticker.Stop()
	}

	return
}

func (a *TaskWatcher) StartUpdateOrphanedTask(uniqueKey string) (err error) {
	if a.CronId != 0 {
		a.Cron.Start()
		return
	}
	a.CronId, err = a.Cron.AddFunc(a.UpdateOrphanedTaskCron, func() {
		ids := a.GetRunningTasksFunc()
		for _, id := range ids {
			redisKey := fmt.Sprintf("%s%s", a.TaskPrefixKey, id)
			_, redisErr := a.RedisConn.Get(context.Background(), redisKey).Result()
			if redisErr != nil { //job not in redis
				a.UpdateTaskToFailFunc(id, redisErr)
			}
		}
	})
	a.Cron.Start()
	return
}

func (a *TaskWatcher) CloseUpdateOrphanedTask(uniqueKey string) (err error) {
	a.Cron.Stop()
	return
}

// stopFunc, err := StartTaskWatcher(ID)
// if err != nil {
//  //	fail task
// }
// defer stopFunc()
