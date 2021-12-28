package main

import (
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
)
//GetEnv function docs from https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66

func GetEnv(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

//Reverse function reverses given string
func Reverse(stringToReverse string) string{
	stringToReverseAsByte := []rune(stringToReverse)
	for i, j := 0, len(stringToReverseAsByte)-1; i < j; i, j = i+1, j-1 {
		stringToReverseAsByte[i], stringToReverseAsByte[j] = stringToReverseAsByte[j], stringToReverseAsByte[i]
	}
	return string(stringToReverseAsByte)
}
func CreateResponseFileFromOutputQuotes(quotes OutputQuotes) error {
	quotesAsSlice:= quotes.TransformToSlice()
	file, _ := json.MarshalIndent(quotesAsSlice, "", " ")
	err := ioutil.WriteFile("test.json", file, 0644)
	return err
}