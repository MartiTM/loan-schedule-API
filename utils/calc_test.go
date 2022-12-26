package utils_test

import (
	"testing"

	"github.com/MartiTM/loan-schedule-API/utils"
)

func TestGetInterestRatePerTermBasic(t *testing.T) {
	ans := utils.GetInterestRatePerTerm(0.06)
	wanted := 0.06/12
	if ans != wanted {
		t.Errorf("GetInterestRatePerTerm(0.0425) = %v; want %v", ans, wanted)
	}
}

func TestGetDueDateAmountBasic(t *testing.T) {
	ans := utils.GetDueDateAmount(5000000, utils.GetInterestRatePerTerm(0.0425), 4)
	wanted := 1261087
	if ans != wanted {
		t.Errorf("GetDueDateAmount(5000000, GetInterestRatePerTerm(0.0425), 4) = %v; want %v", ans, wanted)
	}
}

func TestGetDueDateAmountBasic2(t *testing.T) {
	ans := utils.GetDueDateAmount(1256637, utils.GetInterestRatePerTerm(0.0425), 1)
	wanted := 1261088
	if ans != wanted {
		t.Errorf("GetDueDateAmount(1256637, GetInterestRatePerTerm(0.0425), 4) = %v; want %v", ans, wanted)
	}
}

func TestGetInterestAmountByTermBasic(t *testing.T) {
	ans := utils.GetInterestAmountByTerm(5000000, 0.0425/12)
	wanted := 17708
	if ans != wanted {
		t.Errorf("GetInterestAmountByTerm(5000000, 0.0425/12) = %v; want %v", ans, wanted)
	}
}