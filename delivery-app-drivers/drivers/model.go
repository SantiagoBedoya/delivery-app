package drivers

import (
	"strings"

	"github.com/SantiagoBedoya/delivery-app-drivers/utils/httperrors"
)

// Driver define data struct for customers
type Driver struct {
	ID          int     `json:"id" db:"id"`
	FirstName   string  `json:"first_name" db:"first_name"`
	LastName    string  `json:"last_name" db:"last_name"`
	Email       string  `json:"email" db:"email"`
	Address     string  `json:"address" db:"address"`
	Phone       string  `json:"phone" db:"phone"`
	Verified    bool    `json:"verified" db:"verified" default:"false"`
	IsAvailable bool    `json:"is_available" db:"is_available" default:"false"`
	Lat         float64 `json:"lat" db:"lat"`
	Lng         float64 `json:"lng" db:"lng"`
	Password    string  `json:"password" db:"password"`
}

// ValidateSignUp make validations for customer data struct
func (c *Driver) ValidateSignUp() *httperrors.HTTPError {
	if strings.TrimSpace(c.FirstName) == "" {
		return httperrors.NewBadRequestError("first name should not be emtpy")
	}
	if strings.TrimSpace(c.LastName) == "" {
		return httperrors.NewBadRequestError("last name should not be emtpy")
	}
	if strings.TrimSpace(c.Email) == "" {
		return httperrors.NewBadRequestError("email should not be emtpy")
	}
	if strings.TrimSpace(c.Password) == "" {
		return httperrors.NewBadRequestError("password should not be emtpy")
	}
	if strings.TrimSpace(c.Address) == "" {
		return httperrors.NewBadRequestError("address should not be emtpy")
	}
	if strings.TrimSpace(c.Phone) == "" {
		return httperrors.NewBadRequestError("phone should not be emtpy")
	}
	return nil
}

// ValidateSignIn make validations for customer data struct
func (c *Driver) ValidateSignIn() *httperrors.HTTPError {
	if strings.TrimSpace(c.Email) == "" {
		return httperrors.NewBadRequestError("email should not be emtpy")
	}
	if strings.TrimSpace(c.Password) == "" {
		return httperrors.NewBadRequestError("password should not be emtpy")
	}
	return nil
}

// AccessToken define data struct for access token
type AccessToken struct {
	AccessToken string `json:"access_token"`
}
