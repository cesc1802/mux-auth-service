package common

type Hasher interface {
	Hash(data string) string
}
