package vendors

import "github.com/SantiagoBedoya/delivery-app-vendors/utils/httperrors"

// Service defines interface for services
type Service interface {
	SignIn(data *Vendor) (*AccessToken, *httperrors.HTTPError)
	SignUp(data *Vendor) (*Vendor, *httperrors.HTTPError)
}
