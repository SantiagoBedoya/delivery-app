package foods

import (
	"strings"

	"github.com/SantiagoBedoya/delivery-app-foods/utils/httperrors"
)

// Food defines data struct
type Food struct {
	ID          int     `json:"id" db:"id"`
	VendorID    string  `json:"vendor_id" db:"vendor_id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Category    string  `json:"category" db:"category"`
	Type        string  `json:"type" db:"type"`
	ReadyTime   string  `json:"ready_time" db:"ready_time"`
	Price       float64 `json:"price" db:"price"`
	Rating      int     `json:"rating" db:"rating"`
}

// ValidateUpdate make validation for Food struct for update case
func (f *Food) ValidateUpdate() *httperrors.HTTPError {
	if strings.TrimSpace(f.Name) == "" {
		return httperrors.NewBadRequestError("name should not be empty")
	}
	if strings.TrimSpace(f.Description) == "" {
		return httperrors.NewBadRequestError("description should not be empty")
	}
	if strings.TrimSpace(f.Category) == "" {
		return httperrors.NewBadRequestError("category should not be empty")
	}
	if strings.TrimSpace(f.Type) == "" {
		return httperrors.NewBadRequestError("type should not be empty")
	}
	if strings.TrimSpace(f.ReadyTime) == "" {
		return httperrors.NewBadRequestError("ready time should not be empty")
	}
	if f.Price == 0 {
		return httperrors.NewBadRequestError("price should not be 0")
	}
	return nil
}

// Validate make validation for Food struct
func (f *Food) Validate() *httperrors.HTTPError {
	if strings.TrimSpace(f.VendorID) == "" {
		return httperrors.NewBadRequestError("vendor ID should not be empty")
	}
	if strings.TrimSpace(f.Name) == "" {
		return httperrors.NewBadRequestError("name should not be empty")
	}
	if strings.TrimSpace(f.Description) == "" {
		return httperrors.NewBadRequestError("description should not be empty")
	}
	if strings.TrimSpace(f.Category) == "" {
		return httperrors.NewBadRequestError("category should not be empty")
	}
	if strings.TrimSpace(f.Type) == "" {
		return httperrors.NewBadRequestError("type should not be empty")
	}
	if strings.TrimSpace(f.ReadyTime) == "" {
		return httperrors.NewBadRequestError("ready time should not be empty")
	}
	if f.Price == 0 {
		return httperrors.NewBadRequestError("price should not be 0")
	}
	return nil
}
