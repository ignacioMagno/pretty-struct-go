package callbackend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func SendStruct(port int, str interface{}) {

	js, err := json.Marshal(str)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
	_, err = http.Post("http://localhost:"+strconv.Itoa(port), "application/json", bytes.NewReader(js))
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())

	}
}
