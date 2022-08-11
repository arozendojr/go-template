package firestore

type Repository interface {
	Save() error
}

func New() Repository {
	return &Client{}
}
