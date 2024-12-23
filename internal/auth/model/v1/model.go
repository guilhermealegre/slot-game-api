package v1

import (
	"github.com/golang-jwt/jwt"
	uuid2 "github.com/google/uuid"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	ctxDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	"github.com/guilhermealegre/slot-games-api/internal"
	"github.com/guilhermealegre/slot-games-api/internal/auth/domain/v1"
	v1UserDomain "github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
	v1UserRepository "github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Model struct {
	app      domain.IApp
	repo     v1.IRepository
	userRepo v1UserRepository.IRepository
}

func NewModel(
	app domain.IApp,
	repository v1.IRepository,
	userRepo v1UserRepository.IRepository,
) v1.IModel {
	return &Model{
		app:      app,
		repo:     repository,
		userRepo: userRepo,
	}
}

func (m *Model) Login(ctx ctxDomain.IContext, email, password string) (tokenPair *v1.TokenPair, err error) {

	authDetails, err := m.repo.GetAuthDetailsByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(authDetails.Password, []byte(password))
	if err != nil {
		return nil, err
	}

	userDetails, err := m.userRepo.GetUserDetails(ctx, authDetails.UserID)
	if err != nil {
		return nil, err
	}
	tokenJWTDetails := &v1.TokenJWTDetails{
		UserUUID: userDetails.UUID,
		UserID:   authDetails.UserID,
		Email:    authDetails.Email,
	}

	tokenPair, err = m.generateTokenPair(tokenJWTDetails)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func (m *Model) Signup(ctx ctxDomain.IContext, createUser *v1UserDomain.CreateUser, createAuth *v1.CreateAuth) error {
	exist, err := m.repo.EmailExist(ctx, createAuth.Email)
	if err != nil {
		return err
	}

	if exist {
		return internal.ErrorInvalidEmail()
	}

	tx, err := m.app.Database().Write().Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		return m.app.Logger().DBLog(err)
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(createAuth.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	uuid, err := uuid2.NewUUID()
	if err != nil {
		return err
	}

	createAuth.Password = string(encryptedPassword)
	createUser.UserUUID = uuid.String()

	createAuth.UserID, err = m.userRepo.CreateUser(ctx, tx, createUser)
	if err != nil {
		return err
	}

	if err = m.userRepo.CreateWallet(ctx, tx, createAuth.UserID); err != nil {
		return err
	}

	if err := m.repo.CreateAuthentication(ctx, tx, createAuth); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return m.app.Logger().DBLog(err)
	}

	return nil
}

func (m *Model) generateTokenFromClaim(claims map[string]interface{}, tokenTTL time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims[v1.ExpirationTime] = time.Now().Add(tokenTTL).Unix()
	for key, value := range claims {
		token.Claims.(jwt.MapClaims)[key] = value
	}

	tokenString, err := token.SignedString([]byte(m.app.Http().Config().JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (m *Model) generateTokenPair(userDetails *v1.TokenJWTDetails) (*v1.TokenPair, error) {
	accessTokenClaims := map[string]interface{}{
		v1.UserUUID:       userDetails.UserUUID,
		v1.Email:          userDetails.Email,
		v1.UserID:         userDetails.UserID,
		v1.ExpirationTime: time.Now().Add(v1.AccessTokenTTL).Unix(),
	}

	accessToken, err := m.generateTokenFromClaim(accessTokenClaims, v1.AccessTokenTTL)
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := map[string]interface{}{
		v1.UserUUID:       userDetails.UserUUID,
		v1.Email:          userDetails.Email,
		v1.UserID:         userDetails.UserID,
		v1.ExpirationTime: time.Now().Add(v1.RefreshTokenTTL).Unix(),
	}
	refreshToken, err := m.generateTokenFromClaim(refreshTokenClaims, v1.RefreshTokenTTL)
	if err != nil {
		return nil, err
	}

	tokenPair := &v1.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokenPair, nil
}
