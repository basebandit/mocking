package user_test

import (
	"testing"

	"github.com/basebandit/mocking/user"

	"github.com/basebandit/mocking/mocks"
	"github.com/golang/mock/gomock"
)

func TestUse(t *testing.T) {
	mockCtlr := gomock.NewController(t)
	defer mockCtlr.Finish()

	mockDoer := mocks.NewMockDoer(mockCtlr)
	testUser := &user.User{Doer: mockDoer}

	//Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

	testUser.Use()
}
