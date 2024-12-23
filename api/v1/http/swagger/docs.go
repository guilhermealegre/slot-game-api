/*
	 Slot Games Service

	 # Slot Games Service API

	 Schemes: http, https
	 BasePath: /api/v1
	 Version: 1.0

	 Consumes:
	 - application/json

	 Produces:
	 - application/json

	 SecurityDefinitions:
		Bearer:
		  type: apiKey
		  name: Authorization
		  in: header

	 swagger:meta
*/
package swagger

import (
	_ "github.com/guilhermealegre/slot-games-api/internal/alive/controller/v1" // alive controller
	_ "github.com/guilhermealegre/slot-games-api/internal/auth/controller/v1"  // auth controller
	_ "github.com/guilhermealegre/slot-games-api/internal/slot/controller/v1"  // alive controller
	_ "github.com/guilhermealegre/slot-games-api/internal/user/controller/v1"  // user controller
)
