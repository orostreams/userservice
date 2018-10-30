package services


//ServiceInterface an interface that all services must satisify
type ServiceInterface interface {
	Validate(attributes map[string]interface) error
	Create(attributes map[string]interface) (map[string]interface, error)
	GetAll() ([]map[string]interface, error)
	GetByID(id int) (map[string]interface, error)
	Update(id int,attributes map[string]interface) (map[string]interface, error)
	Delete(id int) (*models.Interface, error)
}
