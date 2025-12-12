package helper

import (
	"encoding/json"
	"net/http"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}

	if validationError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		WebResponse := WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Errors: exception,
		}

		json.NewEncoder(w).Encode(WebResponse)
		return true
	}
	return false
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(error)
	if ok {
		if exception.Error() == "category not found" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)

			WebResponse := WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Errors: exception.Error(),
			}

			json.NewEncoder(w).Encode(WebResponse)
			return true
		}
	}
	return false
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	WebResponse := WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Errors: err,
	}

	json.NewEncoder(w).Encode(WebResponse)
}