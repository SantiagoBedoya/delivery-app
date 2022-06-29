package vendors

import (
	"strings"

	"github.com/SantiagoBedoya/delivery-app-vendors/utils/httperrors"
)

// Vendor define data struct for Vendors
type Vendor struct {
	ID               int     `json:"id" db:"id"`
	Name             string  `json:"name" db:"name"`
	OwnerName        string  `json:"owner_name" db:"owner_name"`
	FoodType         string  `json:"food_type" db:"food_type"`
	Address          string  `json:"address" db:"address"`
	Phone            string  `json:"phone" db:"phone"`
	Email            string  `json:"email" db:"email"`
	ServiceAvailable bool    `json:"service_available" db:"service_available" default:"false"`
	CoverImage       string  `json:"cover_image" db:"cover_image"`
	Rating           int     `json:"rating" db:"rating"`
	Lat              float64 `json:"lat" db:"lat"`
	Lng              float64 `json:"lng" db:"lng"`
	Password         string  `json:"password" db:"password"`
}

// ValidateSignUp make validations for vendor data struct
func (c *Vendor) ValidateSignUp() *httperrors.HTTPError {
	if strings.TrimSpace(c.Name) == "" {
		return httperrors.NewBadRequestError("name should not be emtpy")
	}
	if strings.TrimSpace(c.OwnerName) == "" {
		return httperrors.NewBadRequestError("owner name should not be emtpy")
	}
	if strings.TrimSpace(c.Address) == "" {
		return httperrors.NewBadRequestError("address should not be emtpy")
	}
	if strings.TrimSpace(c.Phone) == "" {
		return httperrors.NewBadRequestError("phone should not be emtpy")
	}
	if strings.TrimSpace(c.Email) == "" {
		return httperrors.NewBadRequestError("email should not be emtpy")
	}
	if strings.TrimSpace(c.Password) == "" {
		return httperrors.NewBadRequestError("password should not be emtpy")
	}
	return nil
}

// ValidateSignIn make validations for customer data struct
func (c *Vendor) ValidateSignIn() *httperrors.HTTPError {
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
