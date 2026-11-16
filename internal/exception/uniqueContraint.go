package exception

type UniqueContraint struct {
	Message string
}

func (e *UniqueContraint) Error() string {
	return e.Message
}
