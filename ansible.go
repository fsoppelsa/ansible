package ansible

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Response struct {
	Msg        string `json:"msg"`
	ConnString string `json:"connstring"`
	Cmd        string `json:"command"`
	Changed    bool   `json:"changed"`
	Failed     bool   `json:"failed"`
}

func ExitJson(responseBody Response) {
	returnResponse(responseBody)
}

func FailJson(responseBody Response) {
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

func ParseVariables(args []string) []byte {
	var response Response

	if len(args) != 2 {
		response.Msg = "Not enough arguments provided!"
		FailJson(response)
	}

	argsFile := args[1]

	text, err := ioutil.ReadFile(argsFile)
	if err != nil {
		response.Msg = "Could not read configuration file: " + argsFile
		FailJson(response)
	}

	return text
}
