package firestore

type Database interface {
	Save() error
}

func New() Database {
	return &Firestore{}
}
