package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"testing"
)

//unit testing


//test fetching quotes

func TestIFetchQuotesSuccessfully(t *testing.T){
	quotes := []InputQuote{{
		Author: "Einstein",
		Text:   "Physics",
	}}
	resp, err := http.Get(GetEnv("API_BASE_URL"))
	if err != nil {
		log.Println(err.Error())
	}
	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body,&quotes); err != nil{
		log.Println("Can not parse body into quotes struct ")
	}
	for index, quote := range quotes {
		log.Println(fmt.Sprintf("%d-%s-%s",index,quote.Author,quote.Text))
	}
	quot := OutputQuotes{}
	quot.ConstructWithGivenSlice(quotes)
	quot.Print()
}
//test grouping by authors
func TestIGroupQuotesByAuthorsSuccessfully(t *testing.T){

	testData := []InputQuote{
		{
			Author: "Einstein",
			Text:   "Physics",
		},
		{
			Author: "Einstein",
			Text:   "Music",
		},
		{
			Author: "Einstein",
			Text:   "Physics",
		},
		{
			Author: "Einstein",
			Text:   "Silence",
		},
		{
			Author: "Edward",
			Text:   "Physics",
		},
		{
			Author: "Edward",
			Text:   "Music",
		},
	}
	quotes := OutputQuotes{}
	quotes.ConstructWithGivenSlice(testData)

	assert.Equalf(t,2,len(quotes["Edward"].Quotes),"Edward has 2 sentences if the grouping operation is done in truth")
	assert.Equalf(t,4,len(quotes["Einstein"].Quotes),"Einstein has 3 sentences if the grouping operation is done in truth")

}
//test reverse operation

func TestIReverseSentencesSuccessfully(t *testing.T){

	textToReverse := "Physics"
	reversedText := Reverse(textToReverse)
	assert.Equalf(t, "scisyhP",reversedText,"If the reverse operation done , Physics should be scisyhP")
}