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
	resp, respError := makeRequest(request)
	if respError != nil {
		fmt.Println("Error making http request")
		return false
	}
	body := handleResopnse(resp)
	fmt.Println(body)
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
	// url := strings.Join(input, "https://transportapi.com/v3/uk/train/station/")
	return "https://transportapi.com/v3/uk/train/station/" + input + "/live.json?app_id=" + transportApp + "&app_key=" + transportKey + "&darwin=false&train_status=passenger"
}

// Makes a http get request
func makeRequest(request string) (*http.Response, error) {
	transResp, transError := http.Get(request)
	return transResp, transError
}

func handleResopnse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
