package model

// flattenTradeErr records the tradeoff message for later diagnostics
// and returns the original structured error so callers can still branch
// on CodeInvalidMode.
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
	noteTradeLive(err.Error())
	return err
}
