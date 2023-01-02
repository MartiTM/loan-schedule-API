package test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MartiTM/loan-schedule-API/handler"
)

func TestBasic(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.CalcScheduler(w, r)
	}))
	defer ts.Close()

	input := strings.NewReader(`{
		"capital emprunté": 5000000,
		"taux d'intérêt annuel": 0.0425,
		"nombre d'échéance": 4
	}`)

	res, err := http.Post(ts.URL, "application/json", input)
	if err != nil {
		t.Errorf("Post Error, err : %v",err)
	}
	result, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("Read body error, err : %v", err)
	}
	
	if res.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Wrong Content-Type responce header, found : %v, wanted \"application/json\"", res.Header.Get("Content-Type"))
	}
	
	if res.Header.Get("Access-Control-Allow-Origin") != "*" {
		t.Errorf("Wrong Access-Control-Allow-Origin responce header, found : %v, wanted \"application/json\"", res.Header.Get("Access-Control-Allow-Origin"))
	}
	
	wanted := `[{"montant de l échéance":1261087,"capital restant dû":5000000,"montant des intérêts":17708,"montant du capital":1243379,"capital restant à rembourser":3756621},{"montant de l échéance":1261087,"capital restant dû":3756621,"montant des intérêts":13305,"montant du capital":1247782,"capital restant à rembourser":2508839},{"montant de l échéance":1261087,"capital restant dû":2508839,"montant des intérêts":8885,"montant du capital":1252202,"capital restant à rembourser":1256637},{"montant de l échéance":1261088,"capital restant dû":1256637,"montant des intérêts":4451,"montant du capital":1256637,"capital restant à rembourser":0}]`
	
	if !bytes.Equal(result, []byte(wanted)) {
		t.Errorf("Responce body wrong,\nwanted : \n%v,\nfound :\n%q", wanted, result)
	}
}