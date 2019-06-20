package Models

import "time"

type Data struct {
	DataType string
	Data     interface{}
	Time     time.Time
	UserId   string
	ReqId    string
	IP       string
}
