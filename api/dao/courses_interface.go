package dao

import "GoGin/internal/model"

type CourseRepository interface {
	PickCourse(StudentID, CourseID int) error
	DropCourse(StudentID, CourseID int) error
	CheckEnrollment(studentID int) ([]model.Enrollment, error)
	CheckInfo() ([]model.Course, error)
	AddCourse(Course model.Course) error
	CheckCourse(courseID int) (model.Course, error)
}
