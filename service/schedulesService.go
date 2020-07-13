package service

import (
	"fmt"
	"time"
)

func doArticleCountData() {
	ticker := time.NewTicker(time.Second * 5 * 60)
	go func() {
		for {
			<-ticker.C
			fmt.Println("start exec doArticleCountData.............")
			NewCountAllDataService().scheduleCountAll()
		}
	}()
}

func StartSchedules() {
	// 统计所有数据
	doArticleCountData()
}
