package utils

import (
	"log"
	"strconv"
)

func ResponseView(data string, message string, status bool) string {
	res := `{"responseData":` + data +
		`,"message":"` + message +
		`","status":` + strconv.FormatBool(status) + `}`
	log.Println(res)
	return res
}
