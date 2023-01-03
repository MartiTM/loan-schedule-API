package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MartiTM/loan-schedule-API/scheduler"
	"github.com/MartiTM/loan-schedule-API/utils"
)

func CalcScheduler(res http.ResponseWriter, req *http.Request) {
	if ok, httpCode, errorMsg := isRequestValid(res, req); !ok {
		sendError(res, httpCode, errorMsg)
		return
	}
	
	var input scheduler.SchedulerInput
	payload := req.Body
	defer req.Body.Close()

	err := json.NewDecoder(payload).Decode(&input)
	if err != nil {
		sendError(res, http.StatusInternalServerError, fmt.Sprintf("Error parsing the input data, err : %v", err))
		return
	}
	
	if ok, err := input.IsValid(); !ok {
		sendError(res, http.StatusInternalServerError, fmt.Sprintf("Error input data are not valid, err : %v", err))
		return
	}

	output, err := input.GetSchedulerOutput()
	if err != nil {
		sendError(res, http.StatusInternalServerError, fmt.Sprintf("Error calculation the output data, err : %v", err))
		return
	}

	outputJSON, err := json.Marshal(output)
	if err != nil {
		sendError(res, http.StatusInternalServerError, fmt.Sprintf("Error parsing the output data, err : %v", err))
		return
	}

	utils.ReturnJsonResponse(res, http.StatusOK, outputJSON)
}

func isRequestValid(res http.ResponseWriter, req *http.Request) (bool, int, string) {
	if req.URL.Path != "/" {
		return true, http.StatusNotFound, "Not Found"
	}

	if req.Method != "POST" {
		return true, http.StatusMethodNotAllowed, "Invalid HTTP method executed. Please use : POST"
	}
	
	if req.Header.Get("Content-type") != "application/json" {
		return true, http.StatusUnsupportedMediaType, "Unsupported media type. Please use : application/json"
	}

	return false, 0, ""
}

func sendError(res http.ResponseWriter, httpCode int, errorMessage string) {
	HandlerMessage := []byte(fmt.Sprintf(`{"message": "%v"}`, errorMessage))
	utils.ReturnJsonResponse(res, httpCode, HandlerMessage)
}