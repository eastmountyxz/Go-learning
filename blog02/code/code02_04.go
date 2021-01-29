package main

import (
   "fmt"
   "time"
)

func main() {
	//获取当前日期
    today := time.Now()
	fmt.Println("当前日期为:", today)
	//分别获取年月日
	year := today.Year()
    month := today.Month()
	day := today.Day()
	fmt.Println("年:", year)
	fmt.Println("月:", month)
	fmt.Println("日:", day)
	fmt.Printf("%d-%02d-%02d", year, month, day)
}