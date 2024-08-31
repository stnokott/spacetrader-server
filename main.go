package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/sirupsen/logrus"
)

func main() {
	for i := uint8(16); i <= 231; i++ {
		fmt.Println(i, aurora.Index(i, "pew-pew"), aurora.BgIndex(i, "pew-pew"))
	}

	s := fmt.Sprintf(
		"%s 2024-08-31 19:58:59 %s creating caches",
		aurora.BgBlue("[INFO]").White(),
		aurora.Index(89, "[server]"),
	)
	logrus.Info(s)
}
