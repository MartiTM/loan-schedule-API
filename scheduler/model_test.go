package scheduler_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/MartiTM/loan-schedule-API/scheduler"
)

func TestFromJsonToSchedulerInputBasic(t *testing.T) {
	data := []byte(`{"capital emprunté": 5000000,"taux d intérêt annuel": 0.0425,"nombre d échéance": 4}`)
	input, err := scheduler.FromJsonToInput(data)

	if err != nil {
		t.Errorf("Error %v", err)
		return
	}

	wanted := scheduler.SchedulerInput{5000000,0.0425,4}

	inputJson, _ := json.Marshal(input)
	wantedJson, _ := json.Marshal(wanted)

	if !bytes.Equal(inputJson, wantedJson) {
		t.Errorf("found :\n %q;\n want :\n %q", inputJson, wantedJson)
	}
}

func TestGetSchedulerOutputBasic(t *testing.T) {
	output, err := scheduler.SchedulerInput{5000000,0.0425,4}.GetSchedulerOutput()

	ans := scheduler.SchedulerOutput{
		[]scheduler.TermsOutput{
			{
				1261087,
				5000000,
				17708,
				1243379,
				3756621,
			},
			{
				1261087,
				3756621,
				13305,
				1247782,
				2508839,
			},
			{
				1261087,
				2508839,
				8885,
				1252202,
				1256637,
			},
			{
				1261088,
				1256637,
				4451,
				1256637,
				0,
			}}}
	
	if err != nil {
		t.Errorf("Error %v", err)
		return
	}

	if len(output.Terms) != 4 {
		t.Errorf("Error : want 4 terms, found %v", len(output.Terms))
		return
	}

	for i, term := range output.Terms {
		if ! reflect.DeepEqual(term, ans.Terms[i]) {
			t.Errorf("Error : terms : %v; want : %v, found : %v", i+1, ans.Terms[i], term)
			return
		}
	}
}