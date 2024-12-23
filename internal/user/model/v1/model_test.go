package v1

import (
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/app"
	contextDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	v1 "github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
	v1Repo "github.com/guilhermealegre/slot-games-api/internal/user/repository/v1"

	"github.com/stretchr/testify/assert"
	"testing"
)

type TModel struct {
	model      v1.IModel
	repository *v1Repo.RepositoryMock
}

func newModel(appMock *app.AppMock) *TModel {
	if appMock == nil {
		appMock = app.NewAppMock()
	}

	repo := v1Repo.NewRepositoryMock()
	model := NewModel(appMock, repo)

	return &TModel{
		model:      model,
		repository: repo,
	}
}

func TestGetProfile(t *testing.T) {
	testCases := []*TestCase{
		testCaseGetProfile1(),
		testCaseGetProfile2(),
	}

	for _, test := range testCases {
		test.Log(t)

		m := newModel(nil)

		// setup
		test.Repository.Setup(m.repository)

		// model
		result, err := m.model.GetProfile(
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

func TestDepositCredits(t *testing.T) {
	testCases := []*TestCase{
		testCaseDepositCredits1(),
		testCaseDepositCredits2(),
	}

	for _, test := range testCases {
		test.Log(t)

		m := newModel(nil)

		// setup
		test.Repository.Setup(m.repository)

		// model
		result, err := m.model.DepositCredits(
			test.Arguments[0].(contextDomain.IContext),
			test.Arguments[1].(int),
			test.Arguments[2].(float64),
		)

		assert.Equal(t, test.Expected[0], result)
		assert.Equal(t, test.Expected[1] == nil, err == nil)
	}

}

func TestWithdrawCredits(t *testing.T) {
	testCases := []*TestCase{
		testCaseWithdrawCredits1(),
		testCaseWithdrawCredits2(),
		testCaseWithdrawCredits3(),
	}

	for _, test := range testCases {
		test.Log(t)

		m := newModel(nil)

		// setup
		test.Repository.Setup(m.repository)

		// model
		result, err := m.model.WithdrawCredits(
			test.Arguments[0].(contextDomain.IContext),
			test.Arguments[1].(int),
			test.Arguments[2].(float64),
		)

		assert.Equal(t, test.Expected[0], result)
		assert.Equal(t, test.Expected[1] == nil, err == nil)
	}

}
