package util

import (
	"strings"
	"time"
)

func SplitSting(input, delimiter string) []string {
	return strings.Split(input, delimiter) // []string{"ETHBTC", "BTCUSDC"}
}

// Name return
func GetWaitTime() (waitTime time.Duration) {
	waitTime = time.Second * 30
	return
}
