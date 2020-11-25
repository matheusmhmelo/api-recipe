package handler

import (
	"encoding/json"
	"net/http"
)

//Heartbeat only to check the health of the API
func Heartbeat(w http.ResponseWriter, r *http.Request) {
	ret, _ := json.Marshal(map[string]string{
		"msg": "Recipes Working!",
	})
	_, _ = w.Write(ret)

}
