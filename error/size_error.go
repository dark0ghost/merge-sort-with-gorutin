package error

type SizeError string

func (e  SizeError) Error() string{
	return string(e)
}
