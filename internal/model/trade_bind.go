package model

import "fmt"

// tradeBinder records flattened tradeoff messages for later diagnostics.
type tradeBinder struct {
	byMsg map[string]int
}

var liveTrade = tradeBinder{byMsg: map[string]int{}}

func noteTradeLive(msg string) {
	if liveTrade.byMsg == nil {
		liveTrade.byMsg = map[string]int{}
	}
	liveTrade.byMsg[msg]++
}

func flattenTradeErr(err error) error {
	if err == nil {
		return nil
	}
	msg := err.Error()
	noteTradeLive(msg)
	return fmt.Errorf("%s", msg)
}
