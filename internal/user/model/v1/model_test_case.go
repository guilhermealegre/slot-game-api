package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-core-lib/test"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
	"github.com/guilhermealegre/slot-games-api/internal"
	"github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
)

type TestCase struct {
	test.BaseTestCase
	Repository test.MapCall
}

func testCaseGetProfile1() *TestCase {
	ctx := context.NewContext(&gin.Context{})

	modelArg := []any{
		ctx,
		1,
	}

	modelReturn := []any{
		&v1.User{
			UUID:      "fd6c83dc-bfde-11ef-b641-32cab69c541a",
			Email:     "guilherme@gmail.com",
			FirstName: "Guilherme",
			LastName:  "Alegre",
			Avatar:    "test.jpg",
			Wallet: v1.Wallet{
				WalletID: 1,
				Balance:  1000,
			},
		},
		nil,
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Get user profile with success",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
			"GetUserDetails": test.CallList{
				test.Call{
					Arguments: modelArg,
					Expected:  modelReturn,
				},
			},
		},
	}
}

func testCaseGetProfile2() *TestCase {
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
			Description: "Get user profile with error",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
			"GetUserDetails": test.CallList{
				test.Call{
					Arguments: modelArg,
					Expected:  modelReturn,
				},
			},
		},
	}
}

func testCaseDepositCredits1() *TestCase {
	ctx := context.NewContext(&gin.Context{})

	modelArg := []any{
		ctx,
		1,
		100.0,
	}

	modelReturn := []any{
		1100.0,
		nil,
	}

	getUserDetailsArg := []any{
		ctx,
		1,
	}

	getUserDetailsReturn := []any{
		&v1.User{
			UUID:      "fd6c83dc-bfde-11ef-b641-32cab69c541a",
			Email:     "guilherme@gmail.com",
			FirstName: "Guilherme",
			LastName:  "Alegre",
			Avatar:    "test.jpg",
			Wallet: v1.Wallet{
				WalletID: 1,
				Balance:  1000,
			},
		},
		nil,
	}
	updateWalletCreditsArg := []any{
		ctx,
		nil,
		1,
		1100.0,
	}

	updateWalletCreditsReturn := []any{
		1100.0,
		nil,
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Deposit credits with success",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
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
	}
}

func testCaseDepositCredits2() *TestCase {
	ctx := context.NewContext(&gin.Context{})

	modelArg := []any{
		ctx,
		1,
		100.0,
	}

	modelReturn := []any{
		0.0,
		internal.ErrorGeneric(),
	}

	getUserDetailsArg := []any{
		ctx,
		1,
	}

	getUserDetailsReturn := []any{
		&v1.User{
			UUID:      "fd6c83dc-bfde-11ef-b641-32cab69c541a",
			Email:     "guilherme@gmail.com",
			FirstName: "Guilherme",
			LastName:  "Alegre",
			Avatar:    "test.jpg",
			Wallet: v1.Wallet{
				WalletID: 1,
				Balance:  1000,
			},
		},
		nil,
	}
	updateWalletCreditsArg := []any{
		ctx,
		nil,
		1,
		1100.0,
	}

	updateWalletCreditsReturn := []any{
		0.0,
		internal.ErrorGeneric(),
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Deposit credits with error",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
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
	}
}

func testCaseWithdrawCredits1() *TestCase {
	ctx := context.NewContext(&gin.Context{})

	modelArg := []any{
		ctx,
		1,
		100.0,
	}

	modelReturn := []any{
		900.0,
		nil,
	}

	getUserDetailsArg := []any{
		ctx,
		1,
	}

	getUserDetailsReturn := []any{
		&v1.User{
			UUID:      "fd6c83dc-bfde-11ef-b641-32cab69c541a",
			Email:     "guilherme@gmail.com",
			FirstName: "Guilherme",
			LastName:  "Alegre",
			Avatar:    "test.jpg",
			Wallet: v1.Wallet{
				WalletID: 1,
				Balance:  1000,
			},
		},
		nil,
	}
	updateWalletCreditsArg := []any{
		ctx,
		nil,
		1,
		900.0,
	}

	updateWalletCreditsReturn := []any{
		900.0,
		nil,
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Withdraw credits with success",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
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
	}
}

func testCaseWithdrawCredits2() *TestCase {
	ctx := context.NewContext(&gin.Context{})

	modelArg := []any{
		ctx,
		1,
		1100.0,
	}

	modelReturn := []any{
		0.0,
		internal.ErrInsufficientFunds(),
	}

	getUserDetailsArg := []any{
		ctx,
		1,
	}

	getUserDetailsReturn := []any{
		&v1.User{
			UUID:      "fd6c83dc-bfde-11ef-b641-32cab69c541a",
			Email:     "guilherme@gmail.com",
			FirstName: "Guilherme",
			LastName:  "Alegre",
			Avatar:    "test.jpg",
			Wallet: v1.Wallet{
				WalletID: 1,
				Balance:  1000,
			},
		},
		nil,
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        2,
			Description: "Withdraw credits with insufficient funds error",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
			"GetUserDetails": test.CallList{
				test.Call{
					Arguments: getUserDetailsArg,
					Expected:  getUserDetailsReturn,
				},
			},
		},
	}
}

func testCaseWithdrawCredits3() *TestCase {
	ctx := context.NewContext(&gin.Context{})

	modelArg := []any{
		ctx,
		1,
		100.0,
	}

	modelReturn := []any{
		0.0,
		internal.ErrorGeneric(),
	}

	getUserDetailsArg := []any{
		ctx,
		1,
	}

	getUserDetailsReturn := []any{
		&v1.User{
			UUID:      "fd6c83dc-bfde-11ef-b641-32cab69c541a",
			Email:     "guilherme@gmail.com",
			FirstName: "Guilherme",
			LastName:  "Alegre",
			Avatar:    "test.jpg",
			Wallet: v1.Wallet{
				WalletID: 1,
				Balance:  1000,
			},
		},
		nil,
	}
	updateWalletCreditsArg := []any{
		ctx,
		nil,
		1,
		900.0,
	}

	updateWalletCreditsReturn := []any{
		0.0,
		internal.ErrorGeneric(),
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        3,
			Description: "Withdraw credits with generic error",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
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
	}
}
