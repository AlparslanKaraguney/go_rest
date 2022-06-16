package user

// Business logic for user

type Service interface {
	GetAll() ([]Model, error)
	GetByID(id uint) (*Model, error)
	Create(model Model) (uint, error)
	// Update(id uint, model Model) (Model, error)
	// Delete(id uint) error
}

type service struct {
	data Data
}

// Compiled time check for interface implementation
var _ Service = service{}

func NewService(data Data) Service {
	return service{data: data}
}

func (s service) GetAll() ([]Model, error) {
	return s.data.GetAll()
}

func (s service) GetByID(id uint) (*Model, error) {
	return s.data.GetByID(id)
}

func (s service) Create(model Model) (uint, error) {
	return s.data.Create(model)
}
