/**
 * Most of this file is borrowed from: https://gist.github.com/sivel/ccd81bdfb31ca0c0e05d
 *
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// ModuleArgs are the module inputs
type ModuleArgs struct {
	Name string
	Time string
}

// Response are the values returned from the module
type Response struct {
	Msg     string `json:"msg"`
	Busy    bool   `json:"busy"`
	Changed bool   `json:"changed"`
	Failed  bool   `json:"failed"`
}

// ExitJSON is ...
func ExitJSON(responseBody Response) {
	returnResponse(responseBody)
}

// FailJSON is ...
func FailJSON(responseBody Response) {
	responseBody.Failed = true
	returnResponse(responseBody)
}

func returnResponse(responseBody Response) {
	var response []byte
	var err error
	response, err = json.Marshal(responseBody)
	if err != nil {
		response, _ = json.Marshal(Response{Msg: "Invalid response object"})
	}
	fmt.Println(string(response))
	if responseBody.Failed {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

// Temp: "go run main.go args.json"
func main() {
	var response Response

	if len(os.Args) != 2 {
		response.Msg = "No argument file provided"
		FailJSON(response)
	}

	argsFile := os.Args[1]

	text, err := ioutil.ReadFile(argsFile)
	if err != nil {
		response.Msg = "Could not read configuration file: " + argsFile
		FailJSON(response)
	}

	var moduleArgs ModuleArgs
	err = json.Unmarshal(text, &moduleArgs)
	if err != nil {
		response.Msg = "Configuration file not valid JSON: " + argsFile
		FailJSON(response)
	}

	// Current time as a sane default if none is passed in the module
	b, err := time.Now().MarshalText()
	if err != nil {
		response.Msg = "Failed to marshal current time: " + argsFile
		FailJSON(response)
	}

	t := string(b)

	if moduleArgs.Time != "" {
		t = moduleArgs.Time
	}

	busy, err := isItBusy(t)
	response.Busy = busy
	if err != nil {
		response.Msg = "ERROR: " + err.Error() + argsFile
		FailJSON(response)
	}

	response.Msg = fmt.Sprintf("The timeslot %v is busy: %v", t, busy)
	ExitJSON(response)
}
