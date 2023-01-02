package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MartiTM/loan-schedule-API/scheduler"
	"github.com/MartiTM/loan-schedule-API/utils"
)

func CalcScheduler(res http.ResponseWriter, req *http.Request) {
	if isErr(res, req) {
		return
	}
	
	var input scheduler.SchedulerInput
	payload := req.Body
	defer req.Body.Close()

	err := json.NewDecoder(payload).Decode(&input)
	
	if err != nil {
		HandlerMessage := []byte(fmt.Sprintf(`{
			"success": false,
			"message": "Error parsing the input data, err : %v",
			}`, err))
		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}
	
	if ok, err := input.IsValid(); !ok {
		HandlerMessage := []byte(fmt.Sprintf(`{
			"success": false,
			"message": "Error parsing the input data, err : %v",
			}`, err))
		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	output, err := input.GetSchedulerOutput()

	if err != nil {
		HandlerMessage := []byte(fmt.Sprintf(`{
			"success": false,
			"message": "Error calculation the output data, err : %v",
			}`, err))
		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	outputJSON, err := json.Marshal(output)

	if err != nil {
		HandlerMessage := []byte(fmt.Sprintf(`{
			"success": false,
			"message": "Error parsing the output data, err : %v",
			}`, err))
		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	utils.ReturnJsonResponse(res, http.StatusOK, outputJSON)
}

func isErr(res http.ResponseWriter, req *http.Request) bool {
	if req.URL.Path != "/" {

		HandlerMessage := []byte(`{
			"success": false,
		 	"message": "Not Found",
		}`)
	  
		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return true
	}

	if req.Method != "POST" {

		HandlerMessage := []byte(`{
			"success": false,
		 	"message": "Invalid HTTP method executed. Please use : POST",
		}`)
	  
		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return true
	}
	
	if req.Header.Get("Content-type") != "application/json" {

		HandlerMessage := []byte(`{
			"success": false,
		 	"message": "Unsupported media type. Please use : application/json",
		}`)
	  
		utils.ReturnJsonResponse(res, http.StatusUnsupportedMediaType, HandlerMessage)
		return true
	}

	return false
}