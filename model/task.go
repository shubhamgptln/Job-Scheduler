package model

type Task struct{
	ID int64 `json:"id"`
	TaskType string `json:"task_type"`
	StartDate int64 `json:"start_date"`
	EndDate int64 `json:"end_date"`
	Priority int32 `json:"priority"`
}