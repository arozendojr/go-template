package firestore

type Register interface {
	Create() error
	Delete() error
	Read() error
}

func New() Register {
	return &User{}
}
