package model

import "strings"

type Task struct {
	ID          string `json:"id" extensions:"x-order=0"`          // タスクID（自動で生成されるUUID）
	UserID      string `json:"user_id" extensions:"x-order=1"`     // ユーザID
	Title       string `json:"title" extensions:"x-order=2"`       // タスクタイトル
	Description string `json:"description" extensions:"x-order=3"` // タスク説明文
	Status      int    `json:"status" extensions:"x-order=4"`      // タスクのステータス（0: 未着手, 1: 完了）
}

type TaskFetchParam struct {
	UserID string  `form:"-"`
	Status *int    `form:"status"`
	Sort   *string `form:"sort"`
}

func (param *TaskFetchParam) GetSort() []string {
	if sort := param.Sort; sort != nil {
		return strings.Split(*sort, ",")
	}
	return []string{}
}
