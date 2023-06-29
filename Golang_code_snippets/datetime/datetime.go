package main

import (
	"fmt"
	"time"
)

func dateTimeToMilliSec() {
	now := time.Now().UTC()
	y, m, d := time.Now().Date()
	fmt.Println(now.UnixMilli())
	fmt.Println(y, m, d)
	today := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	fmt.Println(today)
	fmt.Println(today.UnixMilli())

	//convert back to utc datetime format
	utcDate := time.Unix(0, today.UnixMilli()*int64(time.Millisecond)).UTC()
	fmt.Println(utcDate)

	//get current time in millisec
	now = time.Now()
	fmt.Println("current time in milliseconds", now.UnixMilli())
}

func main() {
	dateTimeToMilliSec()
}
