package scheduler

import (
	"fmt"
	"encoding/json"

	"github.com/MartiTM/loan-schedule-API/utils"
)

type SchedulerInput struct {
	BorrowedCapital int // in cents
	AnnualInterestRate float64
	Terms int
}

type SchedulerOutput struct {
	Terms []TermsOutput
}

type TermsOutput struct {
	DueDateAmount 				int `json:"montant de l'échéance"`
	RemainingCapitalDue 		int `json:"capital restant dû"`
	InterestAmount 				int `json:"montant des intérêts"`
	CapitalAmount 				int `json:"montant du capital"`
	RemainingCapitalToBePaid 	int `json:"capital restant à rembourser"`
}

func (input SchedulerInput) GetSchedulerOutput() (SchedulerOutput, error) {
	if ok, err := input.isValid(); !ok {
		return SchedulerOutput{}, fmt.Errorf("SchedulerInput are not valid, err : %v", err)
	}

	output := SchedulerOutput{make([]TermsOutput, input.Terms)}

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

func (input SchedulerInput) isValid() (bool, error) {
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

func (output SchedulerOutput) ToJson() ([]byte, error) {
	return json.Marshal(output)
}