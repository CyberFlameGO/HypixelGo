package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

}

func sendRequest(apiType string, apiKey string, uuid string) *http.Response {

	apiUrl := "https://api.hypixel.net/" + apiKey + "?key=[#type]"

	if len(apiKey) > 36 {
		fmt.Println("Your API key is invalid, please check it! (Exiting)")
		return
	}

	switch apiType {
	case "player":
		{
			strings.Replace(apiKey, "[#type]", "player", 1)

			resp, err := http.Get(apiUrl)

			if err != nil {
				fmt.Println("Could not sound API request, found error: ", err)
			}

			return resp
		}
		break

	case "guild":
		{
			strings.Replace(apiKey, "[#type]", "guild", 1)

			resp, err := http.Get(apiUrl)

			if err != nil {
				fmt.Println("Could not sound API request, found error: ", err)
			}

			return resp
		}

		break

	case "keys":
		{
			strings.Replace(apiKey, "[#type]", "keys", 1)

			resp, err := http.Get(apiUrl)

			if err != nil {
				fmt.Println("Could not sound API request, found error: ", err)
			}

			return resp
		}

		break

	case "boosters":
		{
			strings.Replace(apiKey, "[#type]", "boosters", 1)

			resp, err := http.Get(apiUrl)

			if err != nil {
				fmt.Println("Could not sound API request, found error: ", err)
			}

			return resp
		}

		break

	case "friends":
		{
			strings.Replace(apiKey, "[#type]", "friends", 1)

			resp, err := http.Get(apiUrl)

			if err != nil {
				fmt.Println("Could not sound API request, found error: ", err)
			}

			return resp
		}

		break

	}
}
