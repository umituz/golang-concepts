package User

import (
	"github.com/golang/mock/gomock"
	"golangconcepts/mockgen/mocks"
	"testing"
)

func TestUser_Use(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := mocks.NewMockUserInterface(ctrl)
	testUser := &User{User: mockUser}

	mockUser.EXPECT().AddUser(1, "Kenan UZ").Return(nil).Times(1)
	testUser.Use()
}
