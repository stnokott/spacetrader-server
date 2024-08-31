package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/stnokott/spacetrader-server/internal/log"
)

func main() {
	for i := uint8(16); i <= 231; i++ {
		fmt.Println(i, aurora.Index(i, "pew-pew"), aurora.BgIndex(i, "pew-pew"))
	}

	s := fmt.Sprintf(
		"%s 2024-08-31 19:58:59 %s creating caches",
		aurora.BgBlue("[INFO]").White(),
		aurora.BgIndex(20, "[server]"),
	)
	fmt.Println(s)
	logger := log.ForComponent("Foo")
	logger.Info(s)
	logger.Debug(s)
}
