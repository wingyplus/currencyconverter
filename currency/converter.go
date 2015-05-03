package currency

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type QueryResult struct {
	Query struct {
		Count int `json:"count"`
	} `json:"query"`
	Result map[string]struct {
		From  string  `json:"fr"`
		To    string  `json:"to"`
		ID    string  `json:"id"`
		Value float64 `json:"val"`
	} `json:"results"`
}

func Convert(from, to string) (result QueryResult, err error) {
	url := fmt.Sprintf("http://www.freecurrencyconverterapi.com/api/v3/convert?q=%s_%s", from, to)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}
