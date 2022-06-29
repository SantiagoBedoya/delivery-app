package foods

import "github.com/SantiagoBedoya/delivery-app-foods/utils/httperrors"

// Service define interface for services
type Service interface {
	Create(data *Food) (*Food, *httperrors.HTTPError)
	UpdateByID(foodID string, data *Food) *httperrors.HTTPError
	DeleteByID(foodID string) *httperrors.HTTPError
	GetAll() ([]Food, *httperrors.HTTPError)
	GetByID(foodID string) (*Food, *httperrors.HTTPError)
}
