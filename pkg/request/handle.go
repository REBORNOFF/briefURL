package request

import (
	"app/breafURL/pkg/response"
	"errors"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	if r.ContentLength == 0 {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return nil, errors.New("NO_CONTENT_LENGTH")
	}

	body, err := Decode[T](r.Body)
	if err != nil {
		response.ResponseWithJson(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	err = IsValidate[T](body)
	if err != nil {
		response.ResponseWithJson(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	return body, nil
}
