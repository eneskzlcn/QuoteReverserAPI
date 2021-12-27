package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type InputQuote struct{
	Author string	`json:"author"`
	Text string	`json:"text"`
}

type OutputQuote struct{
	Author string	`json:"author"`
	Quotes  []string `json:"quotes"`
}


type OutputQuotes map[string]OutputQuote

func(oq *OutputQuotes) Add(quote InputQuote){
	(*oq)[quote.Author] = OutputQuote{
		Author: quote.Author,
		Quotes: append((*oq)[quote.Author].Quotes,Reverse(quote.Text)),
	}
}
func(oq *OutputQuotes) ConstructWithGivenSlice(quoteSlice []InputQuote ){
	for _, quote := range quoteSlice {
		oq.Add(quote)
	}
}
func(oq *OutputQuotes) Print(){
	for author, outputQuote := range *oq {
		log.Println(fmt.Sprintf("Author : %s said these sentences",author))
		for i, sentence := range outputQuote.Quotes {
			log.Println(fmt.Sprintf("\t || %d-%s ||",i,sentence))
		}
		log.Println("")
	}
}
func(oq *OutputQuotes) TransformToSlice() []OutputQuote{
	var outputQoute = []OutputQuote{}
	for _, quote := range *oq {
		outputQoute = append(outputQoute,quote)
	}
	return outputQoute
}
func(oq *OutputQuotes) PrintAsWantedJSON(){
	quotesAsSlice := oq.TransformToSlice()
	quotesAsSliceAsByte,err := json.Marshal(quotesAsSlice)
	if err != nil {
		log.Println("There is an error occured in marshalling")
	}
	log.Println(string(quotesAsSliceAsByte))
}
