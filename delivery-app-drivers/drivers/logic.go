package drivers

import (
	"fmt"
	"time"

	"github.com/SantiagoBedoya/delivery-app-drivers/utils/httperrors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository Repository
	publicKey  []byte
	privateKey []byte
}

// NewService creates an instance for services
func NewService(repository Repository, privateKey, publicKey []byte) Service {
	return &service{
		repository: repository,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (s *service) SignIn(data *Driver) (*AccessToken, *httperrors.HTTPError) {
	if err := data.ValidateSignIn(); err != nil {
		return nil, err
	}
	existingAccount, err := s.repository.FindByEmail(data.Email)
	if err != nil {
		if err == ErrAccountNotFound {
			return nil, httperrors.NewUnauthorizedError("invalid credentials")
		}
		return nil, httperrors.NewUnexpectedError(err)
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(existingAccount.Password),
		[]byte(data.Password),
	); err != nil {
		return nil, httperrors.NewUnauthorizedError("invalid credentials")
	}
	token, err := s.generateJwtToken(fmt.Sprint(existingAccount.ID))
	if err != nil {
		return nil, httperrors.NewUnexpectedError(err)
	}
	return &AccessToken{AccessToken: token}, nil
}

func (s *service) SignUp(data *Driver) (*Driver, *httperrors.HTTPError) {
	if err := data.ValidateSignUp(); err != nil {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, httperrors.NewUnexpectedError(err)
	}
	data.Password = string(hash)
	newAccount, err := s.repository.Create(data)
	if err != nil {
		if err == ErrDuplicateEmail {
			return nil, httperrors.NewBadRequestError(err.Error())
		}
		return nil, httperrors.NewUnexpectedError(err)
	}
	return newAccount, nil
}

func (s *service) generateJwtToken(userID string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(s.privateKey)
	if err != nil {
		return "", nil
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.StandardClaims{
		Id:        userID,
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	})
	return jwtToken.SignedString(key)
}
