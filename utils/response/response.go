package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type HTTPError struct {
	Error string `json:"error"`
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Printf("could not encode data: %v\nerror: %v", data, err)
		}
	}
}

func OK(w http.ResponseWriter, data interface{}) {
	jsonResponse(w, http.StatusOK, data)
}

func Created(w http.ResponseWriter, data interface{}) {
	jsonResponse(w, http.StatusCreated, data)
}

func HTTPErrorResp(w http.ResponseWriter, statusCode int, errMsg string) {
	jsonResponse(w, statusCode, &HTTPError{Error: errMsg})
}

func BadReq(w http.ResponseWriter, errMsg string) {
	HTTPErrorResp(w, http.StatusBadRequest, errMsg)
}

func Conflict(w http.ResponseWriter, errMsg string) {
	HTTPErrorResp(w, http.StatusConflict, errMsg)
}

func Forbidden(w http.ResponseWriter, errMsg string) {
	HTTPErrorResp(w, http.StatusForbidden, errMsg)
}

func InternalErr(w http.ResponseWriter, errMsg string) {
	HTTPErrorResp(w, http.StatusInternalServerError, errMsg)
}

func NotFound(w http.ResponseWriter, errMsg string) {
	HTTPErrorResp(w, http.StatusNotFound, errMsg)
}

func Unauthorized(w http.ResponseWriter, errMsg string) {
	HTTPErrorResp(w, http.StatusUnauthorized, errMsg)
}
