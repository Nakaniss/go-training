package entity

import "time"

type Todo struct {
	ID        int64
	User      User
	Title     string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Todo) SetDone(done bool) {
	t.Done = done
}

func ParseTime(value []byte) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", string(value))
}
