package http

import (
	"net/http"

	infra "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/http"
)

var (
	Api      = infra.NewGroup("api")
	GroupV1  = Api.Group("v1")
	GroupV1P = Api.Group("v1").Group("p")

	//documentation
	GroupV1PDocumentation = GroupV1P.Group("documentation")

	SwaggerDocs    = GroupV1PDocumentation.NewEndpoint("/docs", http.MethodGet)
	SwaggerSwagger = GroupV1PDocumentation.NewEndpoint("/swagger", http.MethodGet)

	Alive       = GroupV1.NewEndpoint("/alive", http.MethodGet)
	PublicAlive = GroupV1P.NewEndpoint("/alive", http.MethodGet)

	//auth
	Register = GroupV1P.NewEndpoint("/register", http.MethodPost)
	Login    = GroupV1P.NewEndpoint("/login", http.MethodPost)

	//user
	Profile         = GroupV1.NewEndpoint("/profile", http.MethodGet)
	DepositCredits  = GroupV1.NewEndpoint("/wallet/deposit", http.MethodPost)
	WithdrawCredits = GroupV1.NewEndpoint("/wallet/withdraw", http.MethodPost)

	//slot
	SpinSlotMachine = GroupV1.NewEndpoint("/slot/spin", http.MethodPost)
	SpinSlotHistory = GroupV1.NewEndpoint("/slot/history", http.MethodGet)
)
