package error

type NilPointerError string

func (e NilPointerError) Error() string{
	return string(e)
}
