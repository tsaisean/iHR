package response

import "iHR/repositories"

type AutoCompleteResponse struct {
	Suggestions []repositories.Suggestion `json:"suggestions"`
}
