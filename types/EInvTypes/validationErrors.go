package EInvTypes

import "strings"

type ValidationErrors []error

func (v ValidationErrors) Errors() []error {
	return v
}

func (v ValidationErrors) Error() string {
	strs := make([]string, len(v))
	for i, err := range v {
		strs[i] = err.Error()
	}
	return strings.Join(strs, ", ")
}

func FromError(err error) ValidationErrors {
	if err == nil {
		return nil
	}
	if errs, ok := err.(ValidationErrors); ok {
		return errs
	}
	return []error{err}
}
func FromErrors(errs []error) ValidationErrors {
	if len(errs) == 0 {
		return nil
	}
	return errs
}
