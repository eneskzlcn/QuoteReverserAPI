package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func ConsumeQuoteApi(inputQuotesSlice *[]InputQuote) error{
	resp, err := http.Get(GetEnv("API_BASE_URL"))
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body,inputQuotesSlice); err != nil{
		return err
	}
	return nil
}


func main(){
	var inputQuoteSlice = []InputQuote{}
	ConsumeQuoteApi(&inputQuoteSlice)
	outputQuotes := OutputQuotes{}
	outputQuotes.ConstructWithGivenSlice(inputQuoteSlice)
	outputQuotes.PrintAsWantedJSON()
}
