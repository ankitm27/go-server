package www

import (
	"go-server/services"
	utility "go-server/utility"
	"time"
)

var config = utility.GetConfig()
var cronTime = config.CronTime

func Scheduler() {
	services.Schedule(cronTime * time.Second)
}
