package utils

import (
	"github.com/an-halim/golang-advanced/assignment-1/constant"
)

func GetRiskProfileDefinition(score int) string {
	for _, profileRisk := range constant.RiskMapping {
		if score >= profileRisk.MinScore && score <= profileRisk.MaxScore {
			return profileRisk.Definition
		}
	}
	return ""
}
