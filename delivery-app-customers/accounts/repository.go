package accounts

// Repository defines interface for repositories
type Repository interface {
	Create(data *Customer) (*Customer, error)
	FindByEmail(email string) (*Customer, error)
}
