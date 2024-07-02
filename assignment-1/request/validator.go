package request

import "github.com/an-halim/golang-advanced/assignment-1/constant"

// FindRiskProfile calculates the risk profile category based on the provided answers.
func ValidateProfile(answers []Answers) (result constant.ProfileRisk, totalScore int) {
	// Initialize the total score

	// Iterate through each answer to find its weight and sum them up
	for _, answer := range answers {
		for _, question := range constant.Questions {
			if question.ID == answer.QuestionID {
				for _, option := range question.Options {
					if option.Answer == answer.Answer {
						totalScore += option.Weight
						break
					}
				}
			}
		}
	}

	// Determine the risk category based on the total score
	for _, profileRisk := range constant.RiskMapping {
		if totalScore >= profileRisk.MinScore && totalScore <= profileRisk.MaxScore {
			result.MinScore = profileRisk.MinScore
			result.MaxScore = profileRisk.MaxScore
			result.Category = profileRisk.Category
			result.Definition = profileRisk.Definition

			return result, totalScore
		}
	}

	// Return Conservative as the default category if no match is found
	return result, totalScore
}
