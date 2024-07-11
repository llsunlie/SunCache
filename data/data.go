package data

type Data interface {
	Get(key string) ([]byte, error)
}
