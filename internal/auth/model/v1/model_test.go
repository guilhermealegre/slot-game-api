package v1

import (
	"github.com/golang-jwt/jwt"
	coreDbSession "github.com/guilhermealegre/go-clean-arch-core-lib/database/session"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/app"
	infraDatabase "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/database"
	contextDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/http"
	httpConfig "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/http/config"
	v1 "github.com/guilhermealegre/slot-games-api/internal/auth/domain/v1"
	v1Repo "github.com/guilhermealegre/slot-games-api/internal/auth/repository/v1"
	v1UserDomain "github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
	v1UserRepo "github.com/guilhermealegre/slot-games-api/internal/user/repository/v1"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TModel struct {
	model      v1.IModel
	repository *v1Repo.RepositoryMock
	userRepo   *v1UserRepo.RepositoryMock
}

func newModel(appMock *app.AppMock) *TModel {
	if appMock == nil {
		appMock = app.NewAppMock()
	}

	repo := v1Repo.NewRepositoryMock()
	userRepo := v1UserRepo.NewRepositoryMock()
	model := NewModel(appMock, repo, userRepo)

	return &TModel{
		model:      model,
		repository: repo,
		userRepo:   userRepo,
	}
}

// Login test
func TestLogin(t *testing.T) {
	jwtSecret := "test"
	newApp := app.NewAppMock()
	config := &httpConfig.Config{
		JwtSecret: jwtSecret,
	}
	newHttpMock := http.NewHttpMock()
	newHttpMock.Mock.On("Config").Return(config)
	newApp.Mock.On("Http").Return(newHttpMock)

	testCases := []*TestCase{
		testCaseLogin1(),
		testCaseLogin2(),
	}

	for _, test := range testCases {
		test.Log(t)

		m := newModel(newApp)

		// setup
		test.Repository.Setup(m.repository)
		test.UserRepository.Setup(m.userRepo)

		// model
		result, err := m.model.Login(
			test.Arguments[0].(contextDomain.IContext),
			test.Arguments[1].(string),
			test.Arguments[2].(string),
		)

		assert.Equal(t, test.Expected[0] == nil, result == nil)
		assert.Equal(t, test.Expected[1] == nil, err == nil)
		if test.Expected[0] != nil {

			if result != nil {
				accessToken, _ := jwt.Parse(result.AccessToken, func(token *jwt.Token) (any, error) {
					_, _ = token.Method.(*jwt.SigningMethodHMAC)
					return []byte(jwtSecret), nil
				})

				refreshToken, _ := jwt.Parse(result.RefreshToken, func(token *jwt.Token) (any, error) {
					_, _ = token.Method.(*jwt.SigningMethodHMAC)
					return []byte(jwtSecret), nil
				})

				accessClaims, _ := accessToken.Claims.(jwt.MapClaims)
				refreshClaims, _ := refreshToken.Claims.(jwt.MapClaims)
				expMap := test.Expected[0].(map[string]any)
				assert.Equal(t, expMap["user_uuid"], accessClaims["user_uuid"])
				assert.Equal(t, expMap["email"], accessClaims["email"])
				assert.Equal(t, expMap["user_id"], accessClaims["user_id"])
				assert.Equal(t, expMap["user_uuid"], refreshClaims["user_uuid"])
				assert.Equal(t, expMap["email"], refreshClaims["email"])
				assert.Equal(t, expMap["user_id"], refreshClaims["user_id"])
			}
		}
	}
}

// Login test
func TestSignup(t *testing.T) {
	txMock := coreDbSession.NewTxMock()
	txMock.On("RollbackUnlessCommitted").Return(nil)
	txMock.On("Commit").Return(nil)
	sessionMock := coreDbSession.NewSessionMock()
	sessionMock.Mock.On("Begin").Return(txMock, nil)
	dbMock := infraDatabase.NewDatabaseMock()
	dbMock.Mock.On("Write").Return(sessionMock)
	newApp := app.NewAppMock()
	newApp.Mock.On("Database").
		Return(dbMock)

	testCases := []*TestCase{
		testCaseSignup1(txMock),
		testCaseSignup2(),
		testCaseSignup3(),
	}

	for _, test := range testCases {
		test.Log(t)

		m := newModel(newApp)

		// setup
		test.Repository.Setup(m.repository)
		test.UserRepository.Setup(m.userRepo)

		// model
		err := m.model.Signup(
			test.Arguments[0].(contextDomain.IContext),
			test.Arguments[1].(*v1UserDomain.CreateUser),
			test.Arguments[2].(*v1.CreateAuth),
		)

		assert.Equal(t, test.Expected[0] == nil, err == nil)
	}
}
