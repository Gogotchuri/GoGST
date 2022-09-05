package vayanaTypes

import (
	"encoding/json"
	"fmt"
	ewbConsts "github.com/gogotchuri/GoGST/consts"
	"strconv"
	"strings"
)

var ErrorTokenExpired = fmt.Errorf("err-expired-token")

type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Message string    `json:"message"`
	Type    string    `json:"type"`
	Args    ErrorArgs `json:"args"`
}

type ErrorArgs struct {
	ParameterName string `json:"parameter-name"`
	Status        string `json:"status"`
	ErrorText     string `json:"error-text"`
	ErrorCodes    []int  `json:"error-codes"`
}

func (e *Error) IsEWBError() bool {
	return e.Type == "Ewb" || e.Message == "err-ewb-returned-error"
}

func (e *Error) IsTokenExpired() bool {
	return e.Type == "Authorization" && e.Message == ErrorTokenExpired.Error()
}

func (e *Error) GetErrorMessages() []string {
	codes := e.GetErrorCodes()
	if len(codes) == 0 {
		return []string{}
	}
	messages := make([]string, len(codes))
	for _, code := range codes {
		if msg, ok := ewbConsts.ErrorCodes[code]; ok {
			messages = append(messages, msg)
		}
	}
	return messages
}

func (e *Error) GetErrorCodes() []int {
	if !e.IsEWBError() || e.Args.ErrorText == "" {
		return []int{}
	}
	if e.Args.ErrorCodes == nil {
		e.Args.ErrorCodes = []int{}
	}
	if len(e.Args.ErrorCodes) != 0 {
		return e.Args.ErrorCodes
	}
	e.parseErrorCodes()
	return e.Args.ErrorCodes
}

func (e *Error) parseErrorCodes() {
	if e.Args.ErrorText == "" {
		return
	}
	e.Args.ErrorCodes = []int{}
	e.Args.ErrorText = strings.ReplaceAll(e.Args.ErrorText, "\\", "")
	errCodes := struct {
		ErrorCodes string `json:"errorCodes"`
	}{}
	err := json.Unmarshal([]byte(e.Args.ErrorText), &errCodes)
	if err != nil {
		return
	}
	for _, code := range strings.Split(errCodes.ErrorCodes, ",") {
		c, err := strconv.ParseInt(code, 10, 32)
		if err != nil {
			continue
		}
		e.Args.ErrorCodes = append(e.Args.ErrorCodes, int(c))
	}
}
