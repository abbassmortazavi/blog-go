package validation

func ErrorMessage() map[string]string {
	return map[string]string{
		"required": "The Field is required",
		"min":      "The Field must be at least %d",
		"max":      "The Field must be at least %d",
		"len":      "The Field must be at least %d",
		"type":     "The type of the field should be %s",
		"email":    "The field must be a valid email address",
	}
}
