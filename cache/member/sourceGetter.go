package member

type SourceGetter interface {
	Get(key string) ([]byte, error)
}

type SourceGetterFunc func(key string) ([]byte, error)

func (f SourceGetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}
