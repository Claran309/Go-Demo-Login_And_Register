package services

import (
	"GoGin/api/dao"
	"GoGin/internal/model"
)

type CourseService struct {
	CourseRepo dao.CourseRepository
}

func NewCourseService(courseRepo dao.CourseRepository) *CourseService {
	return &CourseService{CourseRepo: courseRepo}
}

func (s *CourseService) GetInfo() ([]model.Course, error) {
	courses, err := s.CourseRepo.CheckInfo()
	return courses, err
}

func (s *CourseService) GetEnrollmentInfo(userID int) ([]model.Course, error) {
	enrollments, err := s.CourseRepo.CheckEnrollment(userID)
	if err != nil {
		return nil, err
	}

	courses := make([]model.Course, len(enrollments))
	for _, enrollment := range enrollments {
		courses = append(courses, enrollment.Course)
	}

	return courses, nil
}

func (s *CourseService) PickCourse(studentID, courseID int) (model.Course, error) {
	err := s.CourseRepo.PickCourse(studentID, courseID)
	if err != nil {
		return model.Course{}, err
	}
	course, err := s.CourseRepo.CheckCourse(courseID)
	if err != nil {
		return model.Course{}, err
	}

	return course, nil
}

func (s *CourseService) DropCourse(studentID, courseID int) (model.Course, error) {
	err := s.CourseRepo.DropCourse(studentID, courseID)
	if err != nil {
		return model.Course{}, err
	}
	course, err := s.CourseRepo.CheckCourse(courseID)
	if err != nil {
		return model.Course{}, err
	}

	return course, nil
}

func (s *CourseService) AddCourse(name string, capital int) (model.Course, error) {
	course := model.Course{
		Name:    name,
		Capital: capital,
	}

	err := s.CourseRepo.AddCourse(course)
	if err != nil {
		return model.Course{}, err
	}

	return course, nil
}
