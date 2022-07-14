package doer

//go:generate go run github.com/golang/mock/mockgen -destination mocks/mock_doer.go -package mocks github.com/basebandit/mocking/doer Doer
type Doer interface {
	DoSomething(int, string) error
}
