package accounts

import "github.com/SantiagoBedoya/delivery-app-customers/utils/httperrors"

// Service defines interface for services
type Service interface {
	SignIn(data *Customer) (*AccessToken, *httperrors.HTTPError)
	SignUp(data *Customer) (*Customer, *httperrors.HTTPError)
}
