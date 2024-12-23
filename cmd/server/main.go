package main

import (
	"fmt"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/logger"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/tracer"
	v1AuthController "github.com/guilhermealegre/slot-games-api/internal/auth/controller/v1"
	"github.com/guilhermealegre/slot-games-api/internal/helpers"
	"github.com/guilhermealegre/slot-games-api/internal/infrastructure/rate_limiter"
	v1UserController "github.com/guilhermealegre/slot-games-api/internal/user/controller/v1"

	"os"

	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/validator"

	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/database"

	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/app"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/http"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/redis"

	v1AliveController "github.com/guilhermealegre/slot-games-api/internal/alive/controller/v1"
	v1AliveModel "github.com/guilhermealegre/slot-games-api/internal/alive/model/v1"
	v1AuthModel "github.com/guilhermealegre/slot-games-api/internal/auth/model/v1"
	v1AuthRepository "github.com/guilhermealegre/slot-games-api/internal/auth/repository/v1"
	v1UserModel "github.com/guilhermealegre/slot-games-api/internal/user/model/v1"
	v1UserRepository "github.com/guilhermealegre/slot-games-api/internal/user/repository/v1"

	v1Middleware "github.com/guilhermealegre/slot-games-api/internal/middleware/v1"
	v1SlotController "github.com/guilhermealegre/slot-games-api/internal/slot/controller/v1"
	v1SlotModel "github.com/guilhermealegre/slot-games-api/internal/slot/model/v1"
	v1SlotRepository "github.com/guilhermealegre/slot-games-api/internal/slot/repository/v1"
	v1SwaggerController "github.com/guilhermealegre/slot-games-api/internal/swagger/controller/v1"
	_ "github.com/lib/pq" // postgres driver
)

func main() {
	// app initialization

	newApp := app.New(nil)
	newHttp := http.New(newApp, nil)
	newLogger := logger.New(newApp, nil)
	newTracer := tracer.New(newApp, nil)
	newValidator := validator.New(newApp).
		AddFieldValidators().
		AddStructValidators()
	newRedis := redis.New(newApp, nil)
	newDatabase := database.New(newApp, nil)

	//helpers
	randomSymbols := helpers.NewRandomSymbols(3)

	// repository
	authRepository := v1AuthRepository.NewRepository(newApp)
	userRepository := v1UserRepository.NewRepository(newApp)
	slotRepository := v1SlotRepository.NewRepository(newApp)

	// models
	aliveModel := v1AliveModel.NewModel(newApp)
	authModel := v1AuthModel.NewModel(newApp, authRepository, userRepository)
	userModel := v1UserModel.NewModel(newApp, userRepository)
	slotModel := v1SlotModel.NewModel(newApp, slotRepository, userRepository, randomSymbols)

	// 4 request per second and 2 of burst
	spinRateLimiter := rate_limiter.NewRateLimiter(4, 2)

	newHttp.
		//middlewares
		WithMiddleware(v1Middleware.NewAuthenticateMiddleware(newApp)).
		WithMiddleware(v1Middleware.NewPrintRequestMiddleware(newApp)).
		WithMiddleware(v1Middleware.NewSpinRateLimiterMiddleware(newApp, spinRateLimiter)).
		//controllers
		WithController(v1SwaggerController.NewController(newApp)).
		WithController(v1AliveController.NewController(newApp, aliveModel)).
		WithController(v1AuthController.NewController(newApp, authModel)).
		WithController(v1UserController.NewController(newApp, userModel)).
		WithController(v1SlotController.NewController(newApp, slotModel))

	newApp.
		WithValidator(newValidator).
		WithDatabase(newDatabase).
		WithRedis(newRedis).
		WithLogger(newLogger).
		WithTracer(newTracer).
		WithHttp(newHttp)

	// start app
	if err := newApp.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
