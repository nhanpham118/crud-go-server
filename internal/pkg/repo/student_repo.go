package repo

import (
	"crud-go-server/internal/pkg/entity"

	"gorm.io/gorm"
)

func NewStudentRepo(db *gorm.DB) StudentRepo {
	return &studentRepo{db: db}
}

type StudentRepo interface {
	GetStudents() ([]*entity.FullStudent, error)
	GetStudentByID(id int) (*entity.FullStudent, error)
	CreateStudent(student *entity.Student) error
	Update(student *entity.Student) error
	Delete(id string) error
}

type studentRepo struct {
	db *gorm.DB
}

func (repo *studentRepo) GetStudents() ([]*entity.FullStudent, error) {
	var student []*FullStudentDao

	// result := repo.db.Where("student_no = ?", id).First(&student)
	result := repo.db.Model(&StudentDao{}).
		Select("*").
		Joins("INNER JOIN marks ON marks.student_no = students.student_no").
		Joins("INNER JOIN modules ON modules.module_code = marks.module_code").
		Scan(&student)

	if result.Error != nil {
		return nil, result.Error
	}
	return toStructFullStudentList(student)
}

func (repo *studentRepo) GetStudentByID(id int) (*entity.FullStudent, error) {
	var student []*FullStudentDao

	// result := repo.db.Where("student_no = ?", id).First(&student)
	result := repo.db.Model(&StudentDao{}).
		Select("*").
		Joins("INNER JOIN marks ON marks.student_no = students.student_no").
		Joins("INNER JOIN modules ON modules.module_code = marks.module_code").
		Where("students.student_no = ?", id).
		Scan(&student)

	if result.Error != nil {
		return nil, result.Error
	}
	return toStructFullStudent(student)
}

func (repo *studentRepo) CreateStudent(student *entity.Student) error {
	tx := repo.db.Begin()
	row := new(StudentDao).fromStruct(student)

	// Execute if error then rollback before this
	if err := tx.Save(&row).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (repo *studentRepo) Update(student *entity.Student) error {
	row := new(StudentDao).fromStruct(student)
	q := repo.db.Where("student_no = ?", row.StudentID).Updates(&row)

	return q.Error
}

func (repo *studentRepo) Delete(id string) error {
	q := repo.db.Where("student_no = ?", id).Delete(StudentDao{})
	return q.Error
}

// DAO - data access object
type StudentDao struct {
	StudentID string `gorm:"primaryKey;not null;column:student_no;type:varchar(10)"`
	SurName   string `gorm:"column:surname;type:varchar(20)"`
	ForeName  string `gorm:"column:forename;type:varchar(20)"`
}

func (dao *StudentDao) TableName() string {
	return "students"
}

func (dao *StudentDao) toStruct() (*entity.Student, error) {
	return &entity.Student{
		StudentID: dao.StudentID,
		SurName:   dao.SurName,
		ForeName:  dao.ForeName,
	}, nil
}

func (dao *StudentDao) fromStruct(item *entity.Student) *StudentDao {
	dao.StudentID = item.StudentID
	dao.ForeName = item.ForeName
	dao.SurName = item.SurName
	return dao
}

type FullStudentDao struct {
	StudentID  string `gorm:"not null;column:student_no;type:varchar(10)"`
	SurName    string `gorm:"column:surname;type:varchar(20)"`
	ForeName   string `gorm:"column:forename;type:varchar(20)"`
	ModuleCode string `gorm:"not null;column:module_code;type:varchar(8)"`
	ModuleName string `gorm:"column:module_name;type:varchar(30)"`
	Mark       int    `gorm:"column:mark;type:decimal"`
}

func toStructFullStudent(students []*FullStudentDao) (*entity.FullStudent, error) {
	score := make([]entity.Score, len(students))
	for i, record := range students {
		score[i].Mark = record.Mark
		score[i].ModuleCode = record.ModuleCode
		score[i].ModuleName = record.ModuleName
	}

	return &entity.FullStudent{
		StudentID: students[0].StudentID,
		SurName:   students[0].SurName,
		ForeName:  students[0].ForeName,
		Scores:    score,
	}, nil
}

func toStructFullStudentList(students []*FullStudentDao) ([]*entity.FullStudent, error) {
	studentList := make([]*entity.FullStudent, 0)
	existID := make(map[string]bool)

	for _, student := range students {
		if !existID[student.StudentID] {
			// Create list score for this ID
			scores := make([]entity.Score, 0)
			for _, record := range students {
				if record.StudentID == student.StudentID {
					var score entity.Score
					score.Mark = record.Mark
					score.ModuleCode = record.ModuleCode
					score.ModuleName = record.ModuleName
					scores = append(scores, score)
				}
			}

			// Add this student to final list
			studentList = append(studentList, &entity.FullStudent{
				StudentID: student.StudentID,
				SurName:   student.SurName,
				ForeName:  student.ForeName,
				Scores:    scores,
			})
			existID[student.StudentID] = true
		}
	}

	return studentList, nil
}
