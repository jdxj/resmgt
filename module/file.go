package module

import "time"

type File struct {
	ID       int
	Owner    int
	Category *int
	Content  string
	Title    string
	DateTime time.Time
}
