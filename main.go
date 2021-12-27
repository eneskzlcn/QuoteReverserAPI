package main

import "log"

func main(){
	apiURl := viperEnvVariable("API_BASE_URL")
	log.Println(apiURl)
}
