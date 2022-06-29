package foods

import "github.com/SantiagoBedoya/delivery-app-foods/utils/httperrors"

type service struct {
	repository Repository
}

// NewService creates and implements Service interface
func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Create(data *Food) (*Food, *httperrors.HTTPError) {
	if err := data.Validate(); err != nil {
		return nil, err
	}
	createdFood, err := s.repository.Create(data)
	if err != nil {
		return nil, httperrors.NewUnexpectedError(err)
	}
	return createdFood, nil
}
func (s *service) UpdateByID(foodID string, data *Food) *httperrors.HTTPError {
	if err := data.ValidateUpdate(); err != nil {
		return err
	}
	if err := s.repository.UpdateByID(foodID, data); err != nil {
		if err == ErrFoodNotFound {
			return httperrors.NewNotFoundError(err.Error())
		}
		return httperrors.NewUnexpectedError(err)
	}
	return nil
}
func (s *service) DeleteByID(foodID string) *httperrors.HTTPError {
	if err := s.repository.DeleteByID(foodID); err != nil {
		return httperrors.NewUnexpectedError(err)
	}
	return nil
}
func (s *service) GetAll() ([]Food, *httperrors.HTTPError) {
	foods, err := s.repository.GetAll()
	if err != nil {
		return nil, httperrors.NewUnexpectedError(err)
	}
	return foods, nil
}
func (s *service) GetByID(foodID string) (*Food, *httperrors.HTTPError) {
	food, err := s.repository.GetByID(foodID)
	if err != nil {
		if err == ErrFoodNotFound {
			return nil, httperrors.NewNotFoundError(err.Error())
		}
		return nil, httperrors.NewUnexpectedError(err)
	}
	return food, nil
}
