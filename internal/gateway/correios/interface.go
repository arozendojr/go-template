package firestore

type Address interface {
	GetAddress() error
}

func New() Address {
	return &Correios{}
}
