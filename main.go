package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)
type Quote struct{
	Author string `json:"author"`
	Text string `json:"text"`
}

type AuthorsQuotesMap map[string][]string

func(aqm *AuthorsQuotesMap) ConstructWithGivenSlice(quotes []Quote){
	for _, quote := range quotes {
		aqm.Add(quote)
	}
}
func(aqm *AuthorsQuotesMap) Add(quote Quote ){
	(*aqm)[quote.Author] = append((*aqm)[quote.Author],Reverse(quote.Text))
}
func(aqm AuthorsQuotesMap)  PrintAsWantedJSON() {
	type AuthorsQuotes struct {
		Author string   `json:"author"`
		Quotes []string `json:"quotes"`
	}
	var authorsQuotesSlice []AuthorsQuotes
	for author, authorsQuotes := range aqm {
		authorsQuotesSlice = append(authorsQuotesSlice, AuthorsQuotes{
			Author: author,
			Quotes: authorsQuotes,
		})
	}
	authorsQuotesSliceAsByte, err := json.Marshal(authorsQuotesSlice)
	if err != nil {
		log.Println("There is an error occured in marshalling")
	}
	log.Println(string(authorsQuotesSliceAsByte))
}
func ConsumeQuoteApi(quotes *[]Quote) error{
	resp, err := http.Get(GetEnv("API_BASE_URL"))
	if err != nil {
		return err
	}
	body,_ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body,quotes); err != nil{
		return err
	}
	return nil
}

func main(){
	var readingQuotes []Quote
	err := ConsumeQuoteApi(&readingQuotes)
	if err != nil {
		log.Println("Error consuming the api")
	}
	authorsQuotesMap :=AuthorsQuotesMap{}
	authorsQuotesMap.ConstructWithGivenSlice(readingQuotes)
	authorsQuotesMap.PrintAsWantedJSON()
}
