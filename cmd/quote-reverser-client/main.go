package main

import (
	"fmt"
	"github.com/eneskzlcn/quote-reverser-client/internal/config"
	"github.com/eneskzlcn/quote-reverser-client/internal/quote"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
func run() error {
	env, exists := os.LookupEnv("DEPLOY_ENV")
	if !exists {
		env = "local"
	}
	conf, err := config.LoadConfig("./.dev", env, "yaml")
	if err != nil {
		return err
	}

	quoteClient := quote.NewClient(conf.Client)

	quotes, err := quoteClient.ConsumeQuotes()
	if err != nil {
		return err
	}
	quotes.ReverseAndGroupByAuthor().PrintInJSONFormat()

	return nil
}
