package exception

type BadRequest struct {
	Message string
}

func (e *BadRequest) Error() string {
	return e.Message
}
