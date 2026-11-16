package exception

type CollectionElementLimitExceeded struct {
	Message string
}

func (e *CollectionElementLimitExceeded) Error() string {
	return e.Message
}
