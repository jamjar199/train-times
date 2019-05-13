package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const transportApp = "8cd3c6eb"
const transportKey = "35ffde015d8d54f31a68c48b74ac8a71"

// Train Times struct
type TrainTime struct {
	Time      string `json:"aimed_departure_time"`
	Location  string `json:"destination_name"`
	Transport string `json:"mode"`
}

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
	times, respError := makeRequest(request)
	if respError {
		fmt.Println("Error making http request")
		return false
	}

	fmt.Println(times)
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
func makeRequest(request string) (*TrainTime, bool) {
	transResp, transError := http.Get(request)

	if transError != nil {
		return &TrainTime{}, true
	}
	defer transResp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(transResp.Body)
	newStr := buf.String()

	// fmt.Println(transResp.Body)

	times, jsonError := formatJson(newStr)
	if jsonError {
		return &TrainTime{}, true
	}
	return times, false

}

// Format response to Train Times struct
func formatJson(r io.Reader) (*TrainTime, bool) {
	times := new(TrainTime)

	err := json.NewDecoder(r).Decode(times)
	if err != nil {
		return &TrainTime{}, true
	}
	fmt.Println(times)
	return times, false
}
