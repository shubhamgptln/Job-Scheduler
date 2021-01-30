package model

import "time"

type Task struct{
	ID int64 `json:"id"`
	TaskType string `json:"task_type"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	Priority int32 `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status int32 `json:"status"`
	Owner string `json:owner`
}

type Request struct{

}

type Response struct{

}