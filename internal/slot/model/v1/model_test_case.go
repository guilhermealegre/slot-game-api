package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gocraft/dbr/v2"
	"github.com/guilhermealegre/go-clean-arch-core-lib/test"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
	"github.com/guilhermealegre/slot-games-api/internal"
	v1 "github.com/guilhermealegre/slot-games-api/internal/slot/domain/v1"
	v1UserDomain "github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
)

type TestCase struct {
	test.BaseTestCase
	Repository       test.MapCall
	UserRepository   test.MapCall
	SymbolsGenerator test.MapCall
}

func testCaseSpinSlotMachine1(tx dbr.SessionRunner) *TestCase {
	ctx := context.NewContext(&gin.Context{})

	userDetails := &v1UserDomain.User{
		UUID:      "fd6c83dc-bfde-11ef-b641-32cab69c541a",
		Email:     "guilherme@gmail.com",
		FirstName: "Guilherme",
		LastName:  "Alegre",
		Avatar:    "test.jpg",
		Wallet: v1UserDomain.Wallet{
			WalletID: 1,
			Balance:  100,
		},
	}
	modelArg := []any{
		ctx,
		10.0,
		1,
	}

	spinResult := &v1.SpinSlotMachine{
		SpinResult: v1.SpinResult{
			Symbols:   []int32{1, 2, 3},
			BetAmount: 10.0,
			Payout:    0.0,
			ResultType: v1.SpinResultType{
				Id:   3,
				Key:  "loss",
				Name: "Loss",
			},
			Winning: false,
		},
		Balance: 90.0,
	}

	modelReturn := []any{
		spinResult,
		nil,
	}

	getSpinResultTypeArg := []any{
		ctx,
		"loss",
	}

	getSpinResultTypeReturn := []any{
		&v1.SpinResultType{
			Id:   3,
			Key:  "loss",
			Name: "Loss",
		},
		nil,
	}

	getUserDetailsArg := []any{
		ctx,
		1,
	}

	getUserDetailsReturn := []any{
		userDetails,
		nil,
	}

	updateWalletCreditsArg := []any{
		ctx,
		tx,
		1,
		90.0,
	}

	updateWalletCreditsReturn := []any{
		90.0,
		nil,
	}

	saveSpinSlotResultArg := []any{
		ctx,
		tx,
		1,
		spinResult,
	}

	saveSpinSlotResultReturn := []any{
		nil,
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Spin slot machine with loss as result",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
			"GetSpinResultType": test.CallList{
				test.Call{
					Arguments: getSpinResultTypeArg,
					Expected:  getSpinResultTypeReturn,
				},
			},
			"SaveSpinSlotResult": test.CallList{
				test.Call{
					Arguments: saveSpinSlotResultArg,
					Expected:  saveSpinSlotResultReturn,
				},
			},
		},
		UserRepository: test.MapCall{
			"GetUserDetails": test.CallList{
				test.Call{
					Arguments: getUserDetailsArg,
					Expected:  getUserDetailsReturn,
				},
			},
			"UpdateWalletCredits": test.CallList{
				test.Call{
					Arguments: updateWalletCreditsArg,
					Expected:  updateWalletCreditsReturn,
				},
			},
		},
		SymbolsGenerator: test.MapCall{
			"Generate": test.CallList{
				test.Call{
					Expected: []any{[]int32{1, 2, 3}},
				},
			},
		},
	}

}

func testCaseSpinSlotMachine2(tx dbr.SessionRunner) *TestCase {
	ctx := context.NewContext(&gin.Context{})

	userDetails := &v1UserDomain.User{
		UUID:      "fd6c83dc-bfde-11ef-b641-32cab69c541a",
		Email:     "guilherme@gmail.com",
		FirstName: "Guilherme",
		LastName:  "Alegre",
		Avatar:    "test.jpg",
		Wallet: v1UserDomain.Wallet{
			WalletID: 1,
			Balance:  100,
		},
	}
	modelArg := []any{
		ctx,
		10.0,
		1,
	}

	spinResult := &v1.SpinSlotMachine{
		SpinResult: v1.SpinResult{
			Symbols:   []int32{1, 1, 3},
			BetAmount: 10.0,
			Payout:    20.0,
			ResultType: v1.SpinResultType{
				Id:   2,
				Key:  "x2",
				Name: "Profit x2",
			},
			Winning: true,
		},
		Balance: 110.0,
	}

	modelReturn := []any{
		spinResult,
		nil,
	}

	getSpinResultTypeArg := []any{
		ctx,
		"x2",
	}

	getSpinResultTypeReturn := []any{
		&v1.SpinResultType{
			Id:   2,
			Key:  "x2",
			Name: "Profit x2",
		},
		nil,
	}

	getUserDetailsArg := []any{
		ctx,
		1,
	}

	getUserDetailsReturn := []any{
		userDetails,
		nil,
	}

	updateWalletCreditsArg := []any{
		ctx,
		tx,
		1,
		110.0,
	}

	updateWalletCreditsReturn := []any{
		110.0,
		nil,
	}

	saveSpinSlotResultArg := []any{
		ctx,
		tx,
		1,
		spinResult,
	}

	saveSpinSlotResultReturn := []any{
		nil,
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        2,
			Description: "Spin slot machine with Two identical symbols as x2 profit",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
			"GetSpinResultType": test.CallList{
				test.Call{
					Arguments: getSpinResultTypeArg,
					Expected:  getSpinResultTypeReturn,
				},
			},
			"SaveSpinSlotResult": test.CallList{
				test.Call{
					Arguments: saveSpinSlotResultArg,
					Expected:  saveSpinSlotResultReturn,
				},
			},
		},
		UserRepository: test.MapCall{
			"GetUserDetails": test.CallList{
				test.Call{
					Arguments: getUserDetailsArg,
					Expected:  getUserDetailsReturn,
				},
			},
			"UpdateWalletCredits": test.CallList{
				test.Call{
					Arguments: updateWalletCreditsArg,
					Expected:  updateWalletCreditsReturn,
				},
			},
		},
		SymbolsGenerator: test.MapCall{
			"Generate": test.CallList{
				test.Call{
					Expected: []any{[]int32{1, 1, 3}},
				},
			},
		},
	}

}

func testCaseSpinSlotMachine3(tx dbr.SessionRunner) *TestCase {
	ctx := context.NewContext(&gin.Context{})

	userDetails := &v1UserDomain.User{
		UUID:      "fd6c83dc-bfde-11ef-b641-32cab69c541a",
		Email:     "guilherme@gmail.com",
		FirstName: "Guilherme",
		LastName:  "Alegre",
		Avatar:    "test.jpg",
		Wallet: v1UserDomain.Wallet{
			WalletID: 1,
			Balance:  100,
		},
	}
	modelArg := []any{
		ctx,
		10.0,
		1,
	}

	spinResult := &v1.SpinSlotMachine{
		SpinResult: v1.SpinResult{
			Symbols:   []int32{1, 1, 1},
			BetAmount: 10.0,
			Payout:    100.0,
			ResultType: v1.SpinResultType{
				Id:   1,
				Key:  "x10",
				Name: "Profit x10",
			},
			Winning: true,
		},
		Balance: 190.0,
	}

	modelReturn := []any{
		spinResult,
		nil,
	}

	getSpinResultTypeArg := []any{
		ctx,
		"x10",
	}

	getSpinResultTypeReturn := []any{
		&v1.SpinResultType{
			Id:   1,
			Key:  "x10",
			Name: "Profit x10",
		},
		nil,
	}

	getUserDetailsArg := []any{
		ctx,
		1,
	}

	getUserDetailsReturn := []any{
		userDetails,
		nil,
	}

	updateWalletCreditsArg := []any{
		ctx,
		tx,
		1,
		190.0,
	}

	updateWalletCreditsReturn := []any{
		190.0,
		nil,
	}

	saveSpinSlotResultArg := []any{
		ctx,
		tx,
		1,
		spinResult,
	}

	saveSpinSlotResultReturn := []any{
		nil,
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        3,
			Description: "Spin slot machine with Three identical symbols as x10 profit",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
			"GetSpinResultType": test.CallList{
				test.Call{
					Arguments: getSpinResultTypeArg,
					Expected:  getSpinResultTypeReturn,
				},
			},
			"SaveSpinSlotResult": test.CallList{
				test.Call{
					Arguments: saveSpinSlotResultArg,
					Expected:  saveSpinSlotResultReturn,
				},
			},
		},
		UserRepository: test.MapCall{
			"GetUserDetails": test.CallList{
				test.Call{
					Arguments: getUserDetailsArg,
					Expected:  getUserDetailsReturn,
				},
			},
			"UpdateWalletCredits": test.CallList{
				test.Call{
					Arguments: updateWalletCreditsArg,
					Expected:  updateWalletCreditsReturn,
				},
			},
		},
		SymbolsGenerator: test.MapCall{
			"Generate": test.CallList{
				test.Call{
					Expected: []any{[]int32{1, 1, 1}},
				},
			},
		},
	}

}

func testCaseSpinSlotMachine4() *TestCase {
	ctx := context.NewContext(&gin.Context{})

	userDetails := &v1UserDomain.User{
		UUID:      "fd6c83dc-bfde-11ef-b641-32cab69c541a",
		Email:     "guilherme@gmail.com",
		FirstName: "Guilherme",
		LastName:  "Alegre",
		Avatar:    "test.jpg",
		Wallet: v1UserDomain.Wallet{
			WalletID: 1,
			Balance:  100,
		},
	}
	modelArg := []any{
		ctx,
		110.0,
		1,
	}

	modelReturn := []any{
		nil,
		internal.ErrInsufficientFunds(),
	}

	getUserDetailsArg := []any{
		ctx,
		1,
	}

	getUserDetailsReturn := []any{
		userDetails,
		nil,
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        4,
			Description: "Spin slot machine with insufficient funds error",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},

		UserRepository: test.MapCall{
			"GetUserDetails": test.CallList{
				test.Call{
					Arguments: getUserDetailsArg,
					Expected:  getUserDetailsReturn,
				},
			},
		},
	}

}

func testCaseSpinSlotMachine5() *TestCase {
	ctx := context.NewContext(&gin.Context{})

	modelArg := []any{
		ctx,
		110.0,
		1,
	}

	modelReturn := []any{
		nil,
		internal.ErrorGeneric(),
	}

	getUserDetailsArg := []any{
		ctx,
		1,
	}

	getUserDetailsReturn := []any{
		nil,
		internal.ErrorGeneric(),
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        5,
			Description: "Spin slot machine with generic error",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},

		UserRepository: test.MapCall{
			"GetUserDetails": test.CallList{
				test.Call{
					Arguments: getUserDetailsArg,
					Expected:  getUserDetailsReturn,
				},
			},
		},
	}

}

func testCaseGetSpinSlotHistory1() *TestCase {
	ctx := context.NewContext(&gin.Context{})

	spinHistory := v1.SpinSlotMachineHistory{
		{
			SpinResult: v1.SpinResult{
				Symbols:   []int32{1, 2, 3},
				BetAmount: 10.0,
				Payout:    0.0,
				ResultType: v1.SpinResultType{
					Id:   3,
					Key:  "loss",
					Name: "Loss",
				},
				Winning: false,
			},
			Balance: 90.0,
		},
		{
			SpinResult: v1.SpinResult{
				Symbols:   []int32{1, 1, 3},
				BetAmount: 10.0,
				Payout:    20.0,
				ResultType: v1.SpinResultType{
					Id:   2,
					Key:  "x2",
					Name: "Profit x2",
				},
				Winning: false,
			},
			Balance: 90.0,
		},
	}

	modelArg := []any{
		ctx,
		1,
	}

	modelReturn := []any{
		spinHistory,
		nil,
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Get spin slot history with success",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
			"GetSpinSlotHistory": test.CallList{
				test.Call{
					Arguments: modelArg,
					Expected:  modelReturn,
				},
			},
		},
	}
}

func testCaseGetSpinSlotHistory2() *TestCase {
	ctx := context.NewContext(&gin.Context{})

	modelArg := []any{
		ctx,
		1,
	}

	modelReturn := []any{
		nil,
		internal.ErrorGeneric(),
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        2,
			Description: "Get spin slot history with generic error",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
			"GetSpinSlotHistory": test.CallList{
				test.Call{
					Arguments: modelArg,
					Expected:  modelReturn,
				},
			},
		},
	}
}
