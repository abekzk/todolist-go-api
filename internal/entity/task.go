package entity

type Task struct {
	ID          uint
	UserID      int
	User        User
	Title       string
	Description string
	Status      string // todo: 未着手, wip: 進行中, done: 完了
}