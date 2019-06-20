package Models

import "time"

type Data struct {
	DataType string
	Data     interface{}
	time     time.Time
	UserId   string
	ReqId     string
	IP     string
}
