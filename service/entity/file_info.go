package entity

import "time"

type FileInfo struct {
	FileName string
	Modify   time.Time
	Size     int64
}
