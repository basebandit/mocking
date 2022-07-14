package user_test

import (
	"errors"
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

func TestUseReturnsErrorFromDo(t *testing.T) {
	mockCtlr := gomock.NewController(t)
	defer mockCtlr.Finish()

	dummyError := errors.New("dummy")
	mockDoer := mocks.NewMockDoer(mockCtlr)
	testUser := &user.User{Doer: mockDoer}

	//Expect Do to be called once with 123 and "Hello GoMock" a parameters and return dummyError from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError).Times(1)

	err := testUser.Use()

	if err != dummyError{
		t.Fail()
	}
}
