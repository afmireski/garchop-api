package ports

type UserRepositoryPort interface {
	Create(input interface{}) error

	FindById(input interface{}) (interface{}, error)

	Update(input interface{}) (interface{}, error)

	Delete(input interface{}) error
}