package foods

// Repository defines interface for repositories
type Repository interface {
	Create(data *Food) (*Food, error)
	UpdateByID(foodID string, data *Food) error
	DeleteByID(foodID string) error
	GetAll() ([]Food, error)
	GetByID(foodID string) (*Food, error)
}
