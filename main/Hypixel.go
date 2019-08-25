package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

	//general testing
}

func sendRequest(apiType string, apiKey string, uuid string) *http.Response {

	apiUrl := "https://api.hypixel.net/" + apiKey + "?key=[#type]&byUuid=[#uuid]"

	if len(apiKey) > 36 {
		fmt.Println("Your API key is invalid, please check it! (Exiting)")
		return nil
	}

	switch apiType {
	case "player":
		{
			strings.Replace(apiKey, "[#type]", "player", 1)
			strings.Replace(apiKey, "[#uuid]", uuid, 1)

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
			strings.Replace(apiKey, "[#uuid]", uuid, 1)

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
			strings.Replace(apiKey, "[#uuid]", uuid, 1)

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
			strings.Replace(apiKey, "[#uuid]", uuid, 1)

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
			strings.Replace(apiKey, "[#uuid]", uuid, 1)

			resp, err := http.Get(apiUrl)

			if err != nil {
				fmt.Println("Could not sound API request, found error: ", err)
			}

			return resp
		}

		break

	}

	return nil
}
