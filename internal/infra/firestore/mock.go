package firestore

import "github.com/stretchr/testify/mock"

type FirestoreMock struct {
	mock.Mock
}

func (f *FirestoreMock) Save() error {
	args := f.Called()
	return args.Error(0)
}

func NewFirestoreMock() *FirestoreMock {
	return &FirestoreMock{}
}

func (m *FirestoreMock) DoSomething() error {

	args := m.Called()
	return args.Error(0)

}
