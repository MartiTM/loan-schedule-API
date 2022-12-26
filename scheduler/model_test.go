package scheduler_test

import (
	"bytes"
	"testing"

	"github.com/MartiTM/loan-schedule-API/scheduler"
)

func TestGetSchedulerOutputBasic(t *testing.T) {
	output, err := scheduler.SchedulerInput{5000000,0.0425,4}.GetSchedulerOutput()
	
	if err != nil {
		t.Errorf("Error %v", err)
	}

	ans, err := output.ToJson()

	if err != nil {
		t.Errorf("Error %v", err)
	}

	wanted := []byte(`[
		{
			"montant de l'échéance": 1261087,
			"capital restant dû": 5000000,
			"montant des intérêts": 17708,
			"montant du capital": 1243379,
			"capital restant à rembourser": 3756621,
		},
		{
			"montant de l'échéance": 1261087,
			"capital restant dû": 3756621,
			"montant des intérêts": 13305,
			"montant du capital": 1247782,
			"capital restant à rembourser": 2508839,
		},
		{
			"montant": 1261087,
			"capital restant dû": 2508839,
			"montant des intérêts": 8885,
			"montant du capital": 1252202,
			"capital restant à rembourser": 1256637,
		},
		{
			"montant": 1261088,
			"capital restant dû": 1256637,
			"montant des intérêts": 4451,
			"montant du capital": 1256637,
			"capital restant à rembourser": 0,
		}
]`)

	if bytes.Equal(ans, wanted) {
		t.Errorf("found %q; want %q", ans, wanted)
	}
}