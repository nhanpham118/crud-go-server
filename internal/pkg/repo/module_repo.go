package repo

import (
	"crud-go-server/internal/pkg/entity"

	"gorm.io/gorm"
)

func NewModuleRepo(db *gorm.DB) ModuleRepo {
	return &moduleRepo{db: db}
}

type ModuleRepo interface {
	GetModules() ([]entity.Module, error)
	GetModuleByID(module_code string) (*entity.Module, error)
	CreateModule(module *entity.Module) error
	Update(module *entity.Module) error
	Delete(id string) error
}

type moduleRepo struct {
	db *gorm.DB
}

func (repo *moduleRepo) GetModules() ([]entity.Module, error) {
	var modules []entity.Module

	// result := repo.db.Where("student_no = ?", id).First(&module)
	result := repo.db.Find(&modules)

	if result.Error != nil {
		return nil, result.Error
	}
	return modules, nil
}

func (repo *moduleRepo) GetModuleByID(id string) (*entity.Module, error) {
	var module entity.Module

	// result := repo.db.Where("student_no = ?", id).First(&module)
	result := repo.db.Where("module_code = ?", id).First(&module)

	if result.Error != nil {
		return nil, result.Error
	}
	return &module, nil
}

func (repo *moduleRepo) CreateModule(module *entity.Module) error {
	tx := repo.db.Begin()
	row := new(ModuleDao).fromStruct(module)

	// Execute if error then rollback before this
	if err := tx.Save(&row).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (repo *moduleRepo) Update(module *entity.Module) error {
	row := new(ModuleDao).fromStruct(module)
	q := repo.db.Where("module_code = ?", row.ModuleCode).Updates(&row)

	return q.Error
}

func (repo *moduleRepo) Delete(id string) error {
	q := repo.db.Where("module_code = ?", id).Delete(ModuleDao{})
	return q.Error
}

type ModuleDao struct {
	ModuleCode string `gorm:"primaryKey;not null;column:module_code;type:varchar(8)"`
	ModuleName string `gorm:"column:module_name;type:varchar(30)"`
}

func (dao *ModuleDao) TableName() string {
	return "modules"
}

func (dao *ModuleDao) fromStruct(item *entity.Module) *ModuleDao {
	dao.ModuleCode = item.ModuleCode
	dao.ModuleName = item.ModuleName
	return dao
}
