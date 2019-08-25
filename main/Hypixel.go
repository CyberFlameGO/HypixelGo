package main

import "fmt"

func main() {

}

func sendRequest(apiType string, apiKey string) {

	if len(apiKey) > 36 {
		fmt.Println("Your API key is invalid, please check it! (Exiting)")
		return
	}

	switch apiType {
	case "player":
		{

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
