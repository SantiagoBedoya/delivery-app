package drivers

// Repository defines interface for repositories
type Repository interface {
	Create(data *Driver) (*Driver, error)
	FindByEmail(email string) (*Driver, error)
}
