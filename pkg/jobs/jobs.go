package jobs

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func RunJobs() {
	fmt.Println("вызов Джоба")
	// вызов сервис
	//err := gocron.Every(24).Hours().Do(service.Notification)
	//if err != nil {
	//	log.Println("ошибка ")
	//	return
	//}
	<-gocron.Start()
	//

}
