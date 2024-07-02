package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var printer = message.NewPrinter(language.English)

func fmtInt(n int) string {
	return printer.Sprintf("%d", n)
}
