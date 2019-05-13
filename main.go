package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const transportApp = "8cd3c6eb"
const transportKey = "35ffde015d8d54f31a68c48b74ac8a71"

func main() {
	trainTimes()
}

func trainTimes() bool {
	input := getInput()
	valid, message := validateInput(input)

	if !valid {
		fmt.Println(message)
		return false
	}

	request := formatTrainStationRequest(input)
	jsonBody, respError := makeRequest(request)
	if respError {
		fmt.Println("Error making http request")
		return false
	}

	data, _ := formatJson(jsonBody)
	//fmt.Println(body)
	return true

}

// Gets the station the user wahts times for
func getInput() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Train Station code: ")
	input, _ := reader.ReadString('\n')

	return strings.TrimRight(input, "\n")
}

// Validate the users input
func validateInput(input string) (bool, string) {

	valid := validateStationCode(input)
	if !valid {
		return false, "Invalid station code"
	}

	return true, ""
}

// Validates the input is a station code
// TODO: Add regex
func validateStationCode(input string) bool {
	if input == "BMC" {
		return true
	}
	return false
}

// Formats the request into a url with path parameters
func formatTrainStationRequest(input string) string {
	return "https://transportapi.com/v3/uk/train/station/" + input + "/live.json?app_id=" + transportApp + "&app_key=" + transportKey + "&darwin=false&train_status=passenger"
}

// Makes a http get request
func makeRequest(request string) ([]byte, bool) {
	transResp, transError := http.Get(request)

	if transError != nil {
		return []byte(""), true
	} else {
		data, _ := ioutil.ReadAll(transResp.Body)
		fmt.Println(string(data))
		return data, false
	}
}

type TrainTime struct {
	Time      string
	Location  string
	Transport string
}

func formatJson(json []byte) (*TrainTime, error) {
	fmt.Println(string(json))

	return body, err
}
