package compare

type Compare interface {
	Compare(value interface{}) bool
}