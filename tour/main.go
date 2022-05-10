package main

import (
	"GoProgrammingTourBook/tour/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd execute error: %v\n", err)
	}

	/* 	location, _ := time.LoadLocation("Asia/Shanghai")
	   	inputTime := "2029-09-04 12:02:33"
	   	layout := "2006-01-02 15:04:05"
	   	t, _ := time.ParseInLocation(layout, inputTime, location)
	   	dateTime := time.Unix(t.Unix(), 0).In(location).Format(layout)
	   	log.Printf("输入时间 %s, 输出时间 %s", inputTime, dateTime) */
}
