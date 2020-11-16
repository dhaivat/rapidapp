package stocks_test

import (
	"encoding/json"
	"hello-run/stocks"
	"testing"
)

var (
	testData = struct {
		wantResults        int
		wantPrice          float64
		sampleJsonResponse string
	}{
		wantResults: 1,
		wantPrice:   412.55,
		sampleJsonResponse: `{
		"quoteResponse":{
		   "result":[
			  {
				 "language":"en-US",
				 "region":"US",
				 "quoteType":"EQUITY",
				 "currency":"USD",
				 "bid":412.55,
				 "ask":412.5,
				 "bidSize":9,
				 "askSize":8,
				 "shortName":"Tesla, Inc.",
				 "longName":"Tesla, Inc.",
				 "messageBoardId":"finmb_27444752",
				 "exchangeTimezoneName":"America/New_York",
				 "exchangeTimezoneShortName":"EST",
				 "gmtOffSetMilliseconds":-18000000,
				 "market":"us_market",
				 "esgPopulated":false,
				 "marketState":"REGULAR",
				 "displayName":"Tesla",
				 "symbol":"TSLA"
			  }
		   ]
		}
	 }`}
)

func TestStockParser(t *testing.T) {
	result := stocks.QuoteResponseWrapper{}
	if err := json.Unmarshal([]byte(testData.sampleJsonResponse), &result); err != nil {
		t.Fatalf("unmarshal failed with %v", err)
	}

	if len(result.QuoteResponse.Result) != 1 {
		t.Fatalf("want %v result, got %v", testData.wantResults, len(result.QuoteResponse.Result))
	}

	quoteResult := result.QuoteResponse.Result[0]
	if quoteResult.Bid != 412.55 {
		t.Fatalf("wanted %v, got %v", testData.wantPrice, quoteResult.Bid)
	}

}
