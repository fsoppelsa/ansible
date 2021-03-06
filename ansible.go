/*
Copyright (c) 2016, Fabrizio Soppelsa
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of ansible-swarm nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

// Packaged https://gist.github.com/sivel/ccd81bdfb31ca0c0e05d
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

	if response, err = json.Marshal(responseBody); err != nil {
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
