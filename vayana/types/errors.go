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

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

type ErrorArgs struct {
	IRPErrors     IRPError `json:"irp-err"`
	ErrorLocation string   `json:"errorLocation"`
	ParameterName string   `json:"parameter-name"`
	Status        string   `json:"status"`
	ErrorText     string   `json:"error-text"`
	ErrorCodes    []int    `json:"error-codes"`
}

type IRPError struct {
	Details []IRPErrorDetail `json:"details"`
	Info    []IRPErrorInfo   `json:"info"`
}

type IRPErrorDetail struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type IRPErrorInfo struct {
	InfCd    string      `json:"InfCd"`
	Desc     interface{} `json:"Desc"`
	InfoDesc struct {
		AckNo int64  `json:"AckNo"`
		AckDt string `json:"AckDt"`
		Irn   string `json:"Irn"`
	} `json:"InfoDesc"`
}

func (e Error) IsEWBError() bool {
	return e.Type == "Ewb" || e.Message == "err-ewb-returned-error"
}

func (e Error) IsIRPError() bool {
	return e.Type == "Irp" || e.Message == "err-irp-returned-error"
}

func (e Error) GetIRPErrorMessages() []string {
	var messages []string
	for _, detail := range e.Args.IRPErrors.Details {
		messages = append(messages, detail.ErrorMessage)
	}
	return messages
}

func (e Error) IsInvalidBodyError() bool {
	return e.Type == "ClientRequest" && e.Message == "err-invalid-request-body"
}

func (e Error) GetInvalidBodyErrorMessages() []string {
	return strings.Split(e.Args.ErrorLocation, ";")
}
func (e Error) IsTokenExpired() bool {
	return e.Type == "Authorization" && e.Message == ErrorTokenExpired.Error()
}

func (e *Error) GetEwbErrorMessages() []string {
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
