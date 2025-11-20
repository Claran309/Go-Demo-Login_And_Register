package mysql

import (
	"GoGin/api/dao"
	"GoGin/internal/model"
	"errors"
	"log"

	"gorm.io/gorm"
)

type mysqlCourseRepo struct {
	db *gorm.DB
}

func NewMysqlCourseRepo(db *gorm.DB) dao.CourseRepository {
	err := db.AutoMigrate(&model.Student{}, &model.Course{})
	if err != nil {
		log.Fatal("Failed to migrate student & course table:", err)
	}
	err = db.AutoMigrate(&model.Enrollment{})
	if err != nil {
		log.Fatal("Failed to migrate enrollment table:", err)
	}

	return &mysqlCourseRepo{db: db}
}

func (repo *mysqlCourseRepo) PickCourse(StudentID, CourseID int) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		// 检查学生是否存在
		var student model.Student
		if err := tx.First(&student, StudentID).Error; err != nil {
			return errors.New("student Not Found")
		}

		// 检查课程是否存在
		var course model.Course
		if err := tx.First(&course, CourseID).Error; err != nil {
			return errors.New("course Not Found")
		}

		// 检查课程是否已满
		if course.Enroll >= course.Capital {
			return errors.New("course is full")
		}

		// 是否重复选择
		var exists int64
		if err := tx.Model(&model.Enrollment{}).
			Where("student_id = ? AND course_id = ?", StudentID, CourseID).
			Count(&exists).Error; err != nil {
			return err
		}
		if exists >= 1 {
			return errors.New("enrollment exists")
		}

		// 创建选课关系
		enrollment := model.Enrollment{
			StudentID: StudentID,
			CourseID:  CourseID,
		}
		if err := tx.Create(&enrollment).Error; err != nil {
			return errors.New("enrollment create failed")
		}

		// 更新选课人数
		if err := tx.Model(&model.Course{}).
			Where("course_id = ?", CourseID).
			Update("enroll", gorm.Expr("enroll + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}

func (repo *mysqlCourseRepo) DropCourse(StudentID, CourseID int) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		//是否存在记录
		var enrollment model.Enrollment
		if err := tx.Where("student_id = ? AND course_id = ?", StudentID, CourseID).
			First(&enrollment).Error; err != nil {
			return errors.New("enrollment Not Found")
		}

		//删除
		if err := tx.Delete(&enrollment).Error; err != nil {
			return errors.New("delete failed")
		}

		//更新人数
		if err := tx.Model(&model.Course{}).
			Where("course_id = ?", CourseID).
			Update("enroll", gorm.Expr("enroll - ?", 1)).Error; err != nil {
			return errors.New("update failed")
		}

		return nil
	})
}

func (repo *mysqlCourseRepo) CheckEnrollment(studentID int) ([]model.Enrollment, error) {
	var enrollment []model.Enrollment
	if err := repo.db.Where("student_id = ?", studentID).First(&enrollment).Error; err != nil {
		return nil, errors.New("enrollment select failed")
	}
	return enrollment, nil
}

func (repo *mysqlCourseRepo) CheckInfo() ([]model.Course, error) {
	var course []model.Course
	if err := repo.db.Find(&course).Error; err != nil {
		return nil, errors.New("course select failed")
	}
	return course, nil
}

func (repo *mysqlCourseRepo) AddCourse(Course model.Course) error {
	if err := repo.db.Create(&Course).Error; err != nil {
		return errors.New("course create failed")
	}
	return nil
}

func (repo *mysqlCourseRepo) CheckCourse(courseID int) (model.Course, error) {
	var course model.Course
	if err := repo.db.First(&course, courseID).Error; err != nil {
		return model.Course{}, errors.New("course not found")
	}

	return course, nil
}
