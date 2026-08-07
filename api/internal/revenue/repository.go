package revenue

type IRevenueRepository interface {
	Create(model Revenue) error
	GetAll(userId int) ([]Revenue, error)
	GetById(id uint) (Revenue, error)
	Update(model Revenue) error
	Delete(id int64, userId int64) error
}
