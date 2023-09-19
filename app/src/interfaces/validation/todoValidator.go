package validation

import "github.com/go-ozzo/ozzo-validation"

type TodoValidationResult struct {
	Title   error
	Content error
}

func TodoValidate(title string, content string) TodoValidationResult {
	res := TodoValidationResult{
		Title:   validation.Validate(title, validation.Required),
		Content: validation.Validate(content, validation.Required),
	}

	return res
}
