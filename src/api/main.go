package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Printf("Here! \n")
	var offset time.Duration = 8
	// fmt.Printf("[offset]: %v \n", offset)
	var startTime uint64 = 0
	startTime = 1533916800000
	fmt.Printf("startTime: %v \n", startTime)
	// 這一行”無敵重要”，這一行是uint64轉時間的重要程式
	t := time.Unix(0, int64(startTime)*int64(time.Millisecond)).UTC()

	date := t.Add(offset * time.Hour).Format("2006-01-02 15:04:05")
	startTimeStr := strings.Split(date, " ")[0]

	fmt.Printf("todayDate: %v \n", startTimeStr)

}
