package repository

import (
	"blog/infrastructure"
	"blog/models"
)

//PostRepository -> PostRepository
type InstructorRepository struct {
	db infrastructure.Database
}

// NewPostRepository : fetching database
func NewInstructorRepository(db infrastructure.Database) InstructorRepository {
	return InstructorRepository{
		db: db,
	}
}

//Save -> Method for saving post to database
func (i InstructorRepository) Save(instructor models.Instructor) error {
	return i.db.DB.Create(&instructor).Error
}

//FindAll -> Method for fetching all posts from database
func (i InstructorRepository) FindAll(instructor models.Instructor, keyword string) (*[]models.Instructor, int64, error) {
	var posts []models.Instructor
	var totalRows int64 = 0

	queryBuider := i.db.DB.Order("created_at desc").Model(&models.Instructor{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			i.db.DB.Where("post.title LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(instructor).
		Find(&posts).
		Count(&totalRows).Error
	return &posts, totalRows, err
}

//Update -> Method for updating Post
func (i InstructorRepository) Update(instructor models.Instructor) error {
	return i.db.DB.Save(&instructor).Error
}

//Find -> Method for fetching post by id
func (i InstructorRepository) Find(instructor models.Instructor) (models.Instructor, error) {
	var instructors models.Instructor
	err := i.db.DB.
		Debug().
		Model(&models.Instructor{}).
		Where(&instructor).
		Take(&instructors).Error
	return instructors, err
}

//Delete Deletes Post
func (i InstructorRepository) Delete(instructor models.Instructor) error {
	return i.db.DB.Delete(&instructor).Error
}
