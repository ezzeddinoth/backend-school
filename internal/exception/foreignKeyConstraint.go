package exception

type ForeignKeyConstraint struct {
	Message string
}

func (e *ForeignKeyConstraint) Error() string {
	return e.Message
}
