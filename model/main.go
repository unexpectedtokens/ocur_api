package model

type LogicError struct{
	ModelErrorMessage string
}


func (e LogicError) Error() string{
	return e.ModelErrorMessage
}

func NewLogicError(msg string) LogicError{
	return LogicError{
		ModelErrorMessage: msg,
	}
}