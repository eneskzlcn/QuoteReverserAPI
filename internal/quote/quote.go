package quote

import (
	"encoding/json"
	"github.com/eneskzlcn/quote-reverser-client/internal/stringutil"
	"log"
)

type Quote struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}
type AuthorsQuotesResponse struct {
	Author string   `json:"author"`
	Quotes []string `json:"quotes"`
}

type AuthorsQuotesMap map[string][]string

type Quotes []Quote

func (q *Quotes) ReverseAndGroupByAuthor() AuthorsQuotesMap {
	authorQuotesMap := make(AuthorsQuotesMap, 0)
	for _, quote := range *q {
		authorQuotesMap[quote.Author] = append(authorQuotesMap[quote.Author], stringutil.Reverse(quote.Text))
	}
	return authorQuotesMap
}
func (aqm AuthorsQuotesMap) PrintInJSONFormat() {
	authorQuotesResponses := aqm.ToAuthorQuotesResponses()
	authorsQuotesSliceAsByte, err := json.Marshal(authorQuotesResponses)
	if err != nil {
		return
	}
	log.Println(string(authorsQuotesSliceAsByte))
}
func (aqm AuthorsQuotesMap) ToAuthorQuotesResponses() []AuthorsQuotesResponse {
	var authorsQuotesSlice []AuthorsQuotesResponse
	for author, authorsQuotes := range aqm {
		authorsQuotesSlice = append(authorsQuotesSlice, AuthorsQuotesResponse{
			Author: author,
			Quotes: authorsQuotes,
		})
	}
	return authorsQuotesSlice
}
