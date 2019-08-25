package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

}

func sendRequest(apiType string, apiKey string, uuid string) {

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

			fmt.Println(resp)
		}
		break

	case "guild":
		{

		}

		break

	case "keys":
		{

		}

		break

	case "boosters":
		{

		}

		break

	case "friends":
		{

		}

		break

	}
}
