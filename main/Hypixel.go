package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiUrl string = "https://api.hypixel.net/"

func main() {

	resp, err := http.Get(apiUrl)

	if err != nil {
		fmt.Errorf("Could not send request, got error: %s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(body)
}
