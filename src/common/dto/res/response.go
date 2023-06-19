package response

type ValidationErrorResponse struct {
	Error struct {
		Password string `json:"password" example:"Field 'password' validation failed on the rule 'required'"`
	} `json:"error"`
}
