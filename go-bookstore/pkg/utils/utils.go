package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, target any) {
	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		panic(err)
	}
}
