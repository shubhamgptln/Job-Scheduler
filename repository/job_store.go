package repository

import (
	"job-Scheduler/model"
	"context"
	"time"
)

type JobCreator interface{
	Create(ctx context.Context,job *model.Task) error
}

type JobUpdater interface{
	UpdatePriority(ctx context.Context,jobID uint64,priority int64)
	UpdateSchduledRuntime(ctx context.Context,jobID uint64,startTime time.Time)
}