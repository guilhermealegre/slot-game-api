package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gocraft/dbr/v2"
	"github.com/guilhermealegre/go-clean-arch-core-lib/test"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
	"github.com/guilhermealegre/slot-games-api/internal"
	v1 "github.com/guilhermealegre/slot-games-api/internal/auth/domain/v1"
	v1UserDomain "github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
)

type TestCase struct {
	test.BaseTestCase
	Repository     test.MapCall
	UserRepository test.MapCall
}

func testCaseLogin1() *TestCase {
	ctx := context.NewContext(&gin.Context{})
	email := "guilhermealegre@gmail.com"
	password := "testtest"
	uuid := "fd6c83dc-bfde-11ef-b641-32cab69c541a"

	modelArg := []any{
		ctx,
		email,
		password,
	}

	modelReturn := []any{
		map[string]any{
			"user_id":   float64(1),
			"email":     email,
			"user_uuid": uuid,
		},
		nil,
	}

	getAuthDetailsByEmailArg := []any{
		ctx,
		email,
	}

	getAuthDetailsByEmailReturn := []any{
		&v1.Auth{
			UserID:   1,
			Email:    email,
			Password: []byte("$2a$10$KJvcNca5SdPhqjCa.l3Ice3aer0eXE.USD50zGY0PGYsDLRZE4uAS"),
		},
		nil,
	}

	getUserDetailsArg := []any{
		ctx,
		1,
	}

	getUserDetailsReturn := []any{
		&v1UserDomain.User{
			UUID:      uuid,
			Email:     email,
			FirstName: "Guilherme",
			LastName:  "Alegre",
			Avatar:    "test.jpg",
			Wallet: v1UserDomain.Wallet{
				WalletID: 1,
				Balance:  1000,
			},
		},
		nil,
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Test login with success",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
			"GetAuthDetailsByEmail": test.CallList{
				test.Call{
					Arguments: getAuthDetailsByEmailArg,
					Expected:  getAuthDetailsByEmailReturn,
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
		},
	}
}

func testCaseLogin2() *TestCase {
	ctx := context.NewContext(&gin.Context{})
	email := "guilhermealegre@gmail.com"
	password := "testtest"

	modelArg := []any{
		ctx,
		email,
		password,
	}

	modelReturn := []any{
		nil,
		internal.ErrorGeneric(),
	}

	getAuthDetailsByEmailArg := []any{
		ctx,
		email,
	}

	getAuthDetailsByEmailReturn := []any{
		nil,
		internal.ErrorGeneric(),
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        2,
			Description: "Test login with error",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
			"GetAuthDetailsByEmail": test.CallList{
				test.Call{
					Arguments: getAuthDetailsByEmailArg,
					Expected:  getAuthDetailsByEmailReturn,
				},
			},
		},
	}
}

func testCaseSignup1(tx dbr.SessionRunner) *TestCase {

	ctx := context.NewContext(&gin.Context{})
	email := "guilhermealegre@gmail.com"
	password := "testtest"

	createUser := &v1UserDomain.CreateUser{
		FirstName: "guilherme",
		LastName:  "alegre",
		Avatar:    "test.jpg",
	}

	createAuth := &v1.CreateAuth{
		Email:    email,
		Password: password,
	}

	modelArg := []any{
		ctx,
		createUser,
		createAuth,
	}

	modelReturn := []any{
		nil,
	}

	emailExistArg := []any{
		ctx,
		email,
	}

	emailExistReturn := []any{
		false,
		nil,
	}

	createUserArg := []any{
		ctx,
		tx,
		createUser,
	}

	createUserReturn := []any{
		1,
		nil,
	}

	createWalletArg := []any{
		ctx,
		tx,
		1,
	}

	createWalletReturn := []any{
		nil,
	}

	createAuthenticationArg := []any{
		ctx,
		tx,
		createAuth,
	}

	createAuthenticationReturn := []any{
		nil,
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Test login with success",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
			"EmailExist": test.CallList{
				test.Call{
					Arguments: emailExistArg,
					Expected:  emailExistReturn,
				},
			},
			"CreateAuthentication": test.CallList{
				test.Call{
					Arguments: createAuthenticationArg,
					Expected:  createAuthenticationReturn,
				},
			},
		},
		UserRepository: test.MapCall{
			"CreateUser": test.CallList{
				test.Call{
					Arguments: createUserArg,
					Expected:  createUserReturn,
				},
			},
			"CreateWallet": test.CallList{
				test.Call{
					Arguments: createWalletArg,
					Expected:  createWalletReturn,
				},
			},
		},
	}
}

func testCaseSignup2() *TestCase {

	ctx := context.NewContext(&gin.Context{})
	email := "guilhermealegre@gmail.com"
	password := "testtest"

	createUser := &v1UserDomain.CreateUser{
		FirstName: "guilherme",
		LastName:  "alegre",
		Avatar:    "test.jpg",
	}

	createAuth := &v1.CreateAuth{
		Email:    email,
		Password: password,
	}

	modelArg := []any{
		ctx,
		createUser,
		createAuth,
	}

	modelReturn := []any{
		internal.ErrorInvalidEmail(),
	}

	emailExistArg := []any{
		ctx,
		email,
	}

	emailExistReturn := []any{
		true,
		nil,
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        2,
			Description: "Test signup without success because email already exists",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
			"EmailExist": test.CallList{
				test.Call{
					Arguments: emailExistArg,
					Expected:  emailExistReturn,
				},
			},
		},
	}
}

func testCaseSignup3() *TestCase {

	ctx := context.NewContext(&gin.Context{})
	email := "guilhermealegre@gmail.com"
	password := "testtest"

	createUser := &v1UserDomain.CreateUser{
		FirstName: "guilherme",
		LastName:  "alegre",
		Avatar:    "test.jpg",
	}

	createAuth := &v1.CreateAuth{
		Email:    email,
		Password: password,
	}

	modelArg := []any{
		ctx,
		createUser,
		createAuth,
	}

	modelReturn := []any{
		internal.ErrorGeneric(),
	}

	emailExistArg := []any{
		ctx,
		email,
	}

	emailExistReturn := []any{
		false,
		internal.ErrorGeneric(),
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        3,
			Description: "Test signup without success with generic error",
			Call: test.Call{
				Arguments: modelArg,
				Expected:  modelReturn,
			},
		},
		Repository: test.MapCall{
			"EmailExist": test.CallList{
				test.Call{
					Arguments: emailExistArg,
					Expected:  emailExistReturn,
				},
			},
		},
	}
}
