package hero

import "gorm.io/gorm"

type IRepository interface {
	Save(hero Hero) (Hero, error)
	FindAll() ([]Hero, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(hero Hero) (Hero, error) {
	if err := r.db.Create(&hero).Error; err != nil {
		return hero, err
	}

	return hero, nil
}

func (r *repository) FindAll() ([]Hero, error) {
	var heroes []Hero
	if err := r.db.Find(&heroes).Error; err != nil {
		return heroes, err
	}

	return heroes, nil
}
