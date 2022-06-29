package drivers

import "github.com/SantiagoBedoya/delivery-app-drivers/utils/httperrors"

// Service defines interface for services
type Service interface {
	SignIn(data *Driver) (*AccessToken, *httperrors.HTTPError)
	SignUp(data *Driver) (*Driver, *httperrors.HTTPError)
}
