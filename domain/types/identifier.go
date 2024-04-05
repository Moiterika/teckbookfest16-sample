package types

type Identifier[T any] interface {
	Id() T
	Equals(Identifier[T]) bool
}
