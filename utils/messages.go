package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/iselegant/cnappdemo/domain/model"
	"github.com/labstack/echo"
)

var messageConfig map[string]interface{}

func init() {
	err := json.Unmarshal([]byte(MessagesConfig), &messageConfig)
	if err != nil {
		panic(err)
	}
}

// GetMessageStatusCode is ...
func GetMessageStatusCode(messageCode string) int {
	return int((messageConfig[messageCode].(map[string]interface{}))["statusCode"].(float64))
}

// GetMessageMessageCode is ...
func GetMessageMessageCode(messageCode string) string {
	return (messageConfig[messageCode].(map[string]interface{}))["messageCode"].(string)
}

// GetMessageMessage is ...
func GetMessageMessage(lang string, messageCode string, args ...interface{}) string {
	llang := lang

	if !(llang == "ja" || llang == "en") {
		llang = "ja"
	}
	return fmt.Sprintf(((messageConfig[messageCode].(map[string]interface{}))["message"].(map[string]interface{}))[llang].(string), args...)
}

// GetErrorMassage is ...
func GetErrorMassage(context interface{}, lang string, err1 error) (err error) {
	c := context.(echo.Context)
	messageCode := err1.Error()

	errStatusCode := GetMessageStatusCode(messageCode)
	errMessageCode := GetMessageMessageCode(messageCode)
	errMessage := GetMessageMessage(lang, messageCode)

	errorMessages := &model.ErrorMessages{
		Code:    errMessageCode,
		Message: errMessage,
	}

	return c.JSON(errStatusCode, errorMessages)
}

// SetErrorMassage is ...
func SetErrorMassage(messageCode string) (err error) {
	err = errors.New(messageCode)
	return
}
