package domain

type ErrorRequiredField struct {
    Message string
}

func (e *ErrorRequiredField) Error() string {
    return e.Message + " field is required"
}
