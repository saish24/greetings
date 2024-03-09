package greetings

import "testing"

func TestRandomQuote(t *testing.T) {
	quote, err := RandomQuote()
	if len(quote) == 0 || err != nil {
		t.Error(`err: `, err)
	}
}
