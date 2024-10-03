package auth

import (
	"github.com/MuhammadIbraAlfathar/app-tweet/internal/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MockAuthRepository struct {
	mock.Mock
}

type MockUserRepository struct {
	mock.Mock
}

func (r *MockUserRepository) Create(user *schema.User) error {
	args := r.Called(user)
	return args.Error(0)
}

type AuthUseCaseTestSuite struct {
	suite.Suite

	userRepo *MockUserRepository
	useCase  *UseCase
}

func (s *AuthUseCaseTestSuite) SetupTest() {
	s.userRepo = new(MockUserRepository)
	s.useCase = NewUseCase(s.userRepo)
}

func (s *AuthUseCaseTestSuite) TestRegister_Success() {
	req := &RegisterRequest{
		Email:    "ibraalfathar121@gmail.com",
		Name:     "asdasa",
		Password: "1212121",
		Role:     "admin",
		Gender:   "male",
	}
	s.userRepo.On("Create", mock.Anything).Return(nil)

	err := s.useCase.Register(req)
	assert.NoError(s.T(), err)

}

func TestAuthUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthUseCaseTestSuite))
}
