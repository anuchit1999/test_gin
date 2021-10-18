package service

import (
    "blog/api/repository"
    "blog/models"
)

//PostService PostService struct
type InstructorService struct {
    repository repository.InstructorRepository
}

//NewPostService : returns the PostService struct instance
func NewInstructorService(r repository.InstructorRepository) InstructorService {
    return InstructorService{
        repository: r,
    }
}

//Save -> calls post repository save method
func (i InstructorService) Save(instructor models.Instructor) error {
    return i.repository.Save(instructor)
}

//FindAll -> calls post repo find all method
func (i InstructorService) FindAll(instructor models.Instructor, keyword string) (*[]models.Instructor, int64, error) {
    return i.repository.FindAll(instructor, keyword)
}

// Update -> calls postrepo update method
func (i InstructorService) Update(instructor models.Instructor) error {
    return i.repository.Update(instructor)
}

// Delete -> calls post repo delete method
func (i InstructorService) Delete(id int64) error {
    var instructor models.Instructor
    instructor.Instructor_ID = id
    return i.repository.Delete(instructor)
}

// Find -> calls post repo find method
func (i InstructorService) Find(instructor models.Instructor) (models.Instructor, error) {
    return i.repository.Find(instructor)
}
