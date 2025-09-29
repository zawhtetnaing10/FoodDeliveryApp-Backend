package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(writer http.ResponseWriter, code int, payload any) {
	payloadData, err := json.Marshal(payload)
	if err != nil {
		RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	_, err = writer.Write(payloadData)
	if err != nil {
		log.Printf("Unable to write the response %v", err)
	}
}

func RespondWithError(writer http.ResponseWriter, code int, msg string) {
	type errorVals struct {
		Error string `json:"error"`
	}

	errorStruct := errorVals{
		Error: msg,
	}

	errData, err := json.Marshal(errorStruct)
	if err != nil {
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		writer.WriteHeader(http.StatusInternalServerError)
		_, err = writer.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Unable to write the response %v", err)
			return
		}
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	_, err = writer.Write(errData)
	if err != nil {
		log.Printf("Unable to write the response %v", err)
		return
	}
}
