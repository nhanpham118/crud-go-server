package repo

import (
	"crud-go-server/internal/pkg/entity"

	"gorm.io/gorm"
)

func NewMarkRepo(db *gorm.DB) MarkRepo {
	return &markRepo{db: db}
}

type MarkRepo interface {
	GetMarks(string, string) ([]MarkDao, error)
	CreateMark(mark *entity.Mark) error
	Update(mark *entity.Mark) error
	Delete(studentID string, moduleID string) error
}

type markRepo struct {
	db *gorm.DB
}

func (repo *markRepo) GetMarks(studentID string, moduleCode string) ([]MarkDao, error) {
	var marks []MarkDao

	// ! Preload foreignkey first, if not cant get data
	result := repo.db.
		Preload("Student").Preload("Module").
		Where("student_no REGEXP ? AND module_code REGEXP ?", studentID, moduleCode).
		Find(&marks)

	if result.Error != nil {
		return nil, result.Error
	}
	return marks, nil
}

func (repo *markRepo) CreateMark(mark *entity.Mark) error {
	tx := repo.db.Begin()
	row := new(MarkDao).fromStruct(mark)

	// Execute if error then rollback before this
	if err := tx.Create(&row).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (repo *markRepo) Update(mark *entity.Mark) error {
	row := new(MarkDao).fromStruct(mark)
	q := repo.db.Where("student_no = ? AND mark_code = ?", row.StudentID, row.ModuleID).Updates(&row)

	return q.Error
}

func (repo *markRepo) Delete(studentID string, moduleID string) error {
	q := repo.db.Where("student_no = ? AND mark_code = ?", studentID, moduleID).Delete(MarkDao{})
	return q.Error
}

type MarkDao struct {
	StudentID string     `gorm:"column:student_no;type:varchar(10)" json:"student_no"`
	Student   StudentDao `gorm:"foreignKey:StudentID" json:"-"`
	ModuleID  string     `gorm:"column:module_code;type:varchar(8)" json:"module_code"`
	Module    ModuleDao  `gorm:"foreignKey:ModuleID" json:"-"`
	Mark      int        `gorm:"column:mark;type:decimal" json:"mark"`
}

func (dao *MarkDao) TableName() string {
	return "marks"
}

func (dao *MarkDao) fromStruct(item *entity.Mark) *MarkDao {
	dao.StudentID = item.StudentID
	dao.ModuleID = item.ModuleID
	dao.Mark = item.Mark
	return dao
}
