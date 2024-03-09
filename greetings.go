package greetings

import (
	"bytes"
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"time"
)

func RandomQuote() (string, error) {
	var quoteData struct {
		Quotes []string `json:"greetings,omitempty"`
	}

	if data, err := os.ReadFile("/Users/saish/go/src/github.com/saish24/greetings/greetings-list.json"); err != nil {
		return "", err
	} else {
		reader := bytes.NewReader(data)
		json.NewDecoder(reader).Decode(&quoteData)
	}

	if len(quoteData.Quotes) == 0 {
		return "", errors.New("no quotes loaded in library")
	}

	r := rand.New(rand.NewSource(time.Now().Unix())) // initialize local pseudorandom generator
	i := r.Intn(len(quoteData.Quotes))
	return quoteData.Quotes[i], nil
}
