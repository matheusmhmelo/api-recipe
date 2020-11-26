package utils

import (
	"fmt"
	"net/http"
)

func CreateBadRequestResponse(w http.ResponseWriter, err error)  {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(fmt.Sprintf(`{ "error": "%s" }`, err.Error())))
	return
}
