package main

import (
	"fmt"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var printer = message.NewPrinter(language.English)

func fmtInt(n int) string {
	return printer.Sprintf("%d", n)
}

func fmtDuration(d time.Duration) string {
	var seconds, minutes, hours, days int
	seconds = int(d.Seconds())
	minutes, seconds = seconds/60, seconds%60
	hours, minutes = minutes/60, minutes%60
	days, hours = hours/24, hours%24

	return fmt.Sprintf("%02dd %02dh:%02dm:%02ds", days, hours, minutes, seconds)
}
