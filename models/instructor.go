package models
import "time"

type Instructor struct {
    Instructor_ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
    Name     string    `gorm:"size:200" json:"name"`
    Surname      string    `gorm:"size:3000" json:"surname" `
	Profile      string    `gorm:"size:3000" json:"Profile" `
	Description      string    `gorm:"size:3000" json:"Description" `
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (instructor *Instructor) TableName() string {
    return "instructor"
}

func (instructor *Instructor) ResponseMap() map[string]interface{} {
    resp := make(map[string]interface{})
    resp["instructor_id"] = instructor.Instructor_ID
    resp["name"] = instructor.Name
    resp["surname"] = instructor.Surname
	resp["profile"] = instructor.Profile
	resp["description"] = instructor.Description
    resp["created_at"] = instructor.CreatedAt
    resp["updated_at"] = instructor.UpdatedAt
    return resp
}