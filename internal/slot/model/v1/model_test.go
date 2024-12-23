package v1

import (
	coreDbSession "github.com/guilhermealegre/go-clean-arch-core-lib/database/session"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/app"
	infraDatabase "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/database"
	contextDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	"github.com/guilhermealegre/slot-games-api/internal/helpers"
	v1 "github.com/guilhermealegre/slot-games-api/internal/slot/domain/v1"
	v1Repo "github.com/guilhermealegre/slot-games-api/internal/slot/repository/v1"
	v1UserRepo "github.com/guilhermealegre/slot-games-api/internal/user/repository/v1"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TModel struct {
	model            v1.IModel
	repository       *v1Repo.RepositoryMock
	userRepo         *v1UserRepo.RepositoryMock
	symbolsGenerator *helpers.RandomSymbolsMock
}

func newModel(appMock *app.AppMock) *TModel {
	if appMock == nil {
		appMock = app.NewAppMock()
	}

	repo := v1Repo.NewRepositoryMock()
	userRepo := v1UserRepo.NewRepositoryMock()
	symbolsGenerator := helpers.NewRandomSymbolsMock()
	model := NewModel(appMock, repo, userRepo, symbolsGenerator)

	return &TModel{
		model:            model,
		repository:       repo,
		userRepo:         userRepo,
		symbolsGenerator: symbolsGenerator,
	}
}

func TestSpinSlotMachine(t *testing.T) {
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
		testCaseSpinSlotMachine1(txMock),
		testCaseSpinSlotMachine2(txMock),
		testCaseSpinSlotMachine3(txMock),
		testCaseSpinSlotMachine4(),
		testCaseSpinSlotMachine5(),
	}

	for _, test := range testCases {
		test.Log(t)

		m := newModel(newApp)

		// setup
		test.SymbolsGenerator.Setup(m.symbolsGenerator)
		test.Repository.Setup(m.repository)
		test.UserRepository.Setup(m.userRepo)

		// model
		result, err := m.model.SpinSlotMachine(
			test.Arguments[0].(contextDomain.IContext),
			test.Arguments[1].(float64),
			test.Arguments[2].(int),
		)

		assert.Equal(t, test.Expected[0] == nil, result == nil)
		assert.Equal(t, test.Expected[1] == nil, err == nil)
		if test.Expected[0] != nil {
			assert.Equal(t, test.Expected[0], result)
		}
	}
}

func TestGetSpinSlotHistory(t *testing.T) {

	testCases := []*TestCase{
		testCaseGetSpinSlotHistory1(),
		testCaseGetSpinSlotHistory2(),
	}

	for _, test := range testCases {
		test.Log(t)

		m := newModel(nil)

		// setup
		test.Repository.Setup(m.repository)
		test.UserRepository.Setup(m.userRepo)

		// model
		result, err := m.model.GetSpinSlotHistory(
			test.Arguments[0].(contextDomain.IContext),
			test.Arguments[1].(int),
		)

		assert.Equal(t, test.Expected[0] == nil, result == nil)
		assert.Equal(t, test.Expected[1] == nil, err == nil)
		if test.Expected[0] != nil {
			assert.Equal(t, test.Expected[0], result)
		}
	}
}
