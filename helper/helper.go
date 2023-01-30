package helper

import (
	"reflect"
	"strconv"
)

type formatResponseAPI struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Status    string `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	TotalData int    `json:"total_data"`
}

func ChangeFromInterface(data interface{}) (result int) {
	result = 0

	dataType := reflect.TypeOf(data)
	if dataType.Name() == "string" {
		tmp := string(data.(string))
		res, _ := strconv.Atoi(tmp)
		result = res

	} else if dataType.Name() == "float64" {
		result = int(data.(float64))
	}

	return result
}

func ResponseAPIFormat(status, message string, code, totalData int, data interface{}) formatResponseAPI {
	result := formatResponseAPI{
		Meta: meta{
			Status:    status,
			Code:      code,
			Message:   message,
			TotalData: totalData,
		},
		Data: data,
	}

	return result
}
