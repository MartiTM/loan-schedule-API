package scheduler

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/MartiTM/loan-schedule-API/utils"
)

type SchedulerInput struct {
	BorrowedCapital 	int 	// `json:"capital emprunté"` // in cents
	AnnualInterestRate 	float64 // `json:"taux d'intérêt annuel"`
	Terms 				int		// `json:"nombre d'échéance"`
}

type SchedulerOutput struct {
	Terms []TermsOutput
}

type TermsOutput struct {
	DueDateAmount 				int `json:"montant de l échéance"`
	RemainingCapitalDue 		int `json:"capital restant dû"`
	InterestAmount 				int `json:"montant des intérêts"`
	CapitalAmount 				int `json:"montant du capital"`
	RemainingCapitalToBePaid 	int `json:"capital restant à rembourser"`
}

func (input SchedulerInput) GetSchedulerOutput() (SchedulerOutput, error) {
	if ok, err := input.IsValid(); !ok {
		return SchedulerOutput{}, fmt.Errorf("SchedulerInput are not valid, err : %v", err)
	}

	output := SchedulerOutput{make([]TermsOutput, 0)}

	remainingCapital := input.BorrowedCapital

	interestRatePerTerm := utils.GetInterestRatePerTerm(input.AnnualInterestRate)

	for i:=0; i<input.Terms; i++ {
		term := TermsOutput{}

		term.DueDateAmount = utils.GetDueDateAmount(remainingCapital, interestRatePerTerm, input.Terms-i)
		term.RemainingCapitalDue = remainingCapital
		term.InterestAmount = utils.GetInterestAmountByTerm(term.RemainingCapitalDue, interestRatePerTerm)
		term.CapitalAmount = term.DueDateAmount - term.InterestAmount
		term.RemainingCapitalToBePaid = term.RemainingCapitalDue - term.CapitalAmount
		
		remainingCapital = term.RemainingCapitalToBePaid

		output.Terms = append(output.Terms, term)
	}

	return output, nil
}

func (input SchedulerInput) IsValid() (bool, error) {
	if input.BorrowedCapital <= 0 {
		return false, fmt.Errorf("borrowedCapital <= 0")
	}
	
	if input.AnnualInterestRate <= 0 || input.AnnualInterestRate > 1 {
		return false, fmt.Errorf("annualInterestRate <= 0 or > 1")
	}
	
	if input.Terms <= 0 {
		return false, fmt.Errorf("terms <= 0")
	}

	return true, nil
}

func (input *SchedulerInput) UnmarshalJSON(b []byte) error {
    var testObject map[string]interface{}

    err := json.Unmarshal(b, &testObject)
    if err != nil {
        return err
    }

	if len(testObject) != 3 {
		return fmt.Errorf("expected 3 parameters (\"capital emprunté\", \"taux d'intérêt annuel\" and \"nombre d'échéance\")")
	}

    borrowedCapital, ok := testObject["capital emprunté"].(float64)
    if !ok {
        return fmt.Errorf("\"capital emprunté\" is not a int")
    }
    annualInterestRate, ok := testObject["taux d'intérêt annuel"].(float64)
    if !ok {
        return fmt.Errorf("\"taux d'intérêt annuel\" is not a float64")
    }
    terms, ok := testObject["nombre d'échéance"].(float64)
    if !ok {
        return fmt.Errorf("\"nombre d'échéance\" is not a int")
    }

    input.BorrowedCapital = int(borrowedCapital)
    input.AnnualInterestRate = annualInterestRate
    input.Terms = int(terms)

    return nil
}

func (output SchedulerOutput) MarshalJSON() ([]byte, error) {
	var termsStrings []string
	for _, term := range output.Terms {
	  innerObjectBytes, err := json.Marshal(term)
	  if err != nil {
		return nil, err
	  }
	  termsStrings = append(termsStrings, string(innerObjectBytes))
	}
	return []byte("[" + strings.Join(termsStrings, ",") + "]"), nil
}