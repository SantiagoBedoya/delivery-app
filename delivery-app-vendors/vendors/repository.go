package vendors

// Repository defines interface for repositories
type Repository interface {
	Create(data *Vendor) (*Vendor, error)
	FindByEmail(email string) (*Vendor, error)
}
