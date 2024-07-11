package core

type Value interface {
	Len() int
}

type Cache interface {
	Add(key string, value Value)
	Get(key string) (value Value, ok bool)
	UseBytes() int64
	MaxBytes() int64
}
