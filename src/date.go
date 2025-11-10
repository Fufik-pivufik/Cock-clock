package main

import (
	"time"
)

func ParseDate() (string, string, string) {

	t := time.Now()
	var date string = t.Format("Monday | 02 January 2006")
	var tim string = t.Format("15:04:05")
	location := time.Now().Format("MST")

	return tim, date, location
}
