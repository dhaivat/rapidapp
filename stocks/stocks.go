package stocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	yahooStockURL = `https://query1.finance.yahoo.com/v7/finance/quote?symbols=%s&range1h&includeTimestamps=false&includePrePost=false&corsDomain=finance.yahoo.com&.tsrc=finance`
)

// GetQuotes takes a vararg of stock tickers and returns a QuoteResponse or Error
func GetQuotes(ticker ...string) ([]Quote, error) {

	resp, err := http.Get(fmt.Sprintf(yahooStockURL, strings.Join(ticker, ",")))
	if err != nil {
		log.Printf("unable to fetch stocks %v", err)
		return nil, fmt.Errorf("uanble to fetch stocks %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	qResp := &QuoteResponseWrapper{}
	if err := json.Unmarshal(body, &qResp); err != nil {
		return nil, fmt.Errorf("unable to parse quote reponse %v", err)
	}

	res := qResp.QuoteResponse.Result
	return res, nil
}

// GetPrice takes a single stock ticker and returns just the price as float64
func GetPrice(ticker string) float64 {
	resp, err := GetQuotes(ticker)
	if err != nil || len(resp) != 1 {
		log.Println(err)
		log.Println(resp)
		return -1.0
	}
	return resp[0].Bid
}
