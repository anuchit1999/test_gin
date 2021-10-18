package controller

import (
    "blog/api/service"
    "blog/models"
    "blog/util"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

//PostController -> PostController
type InstructorController struct {
    service service.InstructorService
}

//NewPostController : NewPostController
func NewInstructorController(s service.InstructorService) InstructorController {
    return InstructorController{
        service: s,
    }
}

// GetPosts : GetPosts controller
func (i InstructorController) GetPosts(ctx *gin.Context) {
    var instructors models.Instructor

    keyword := ctx.Query("keyword")

    data, total, err := i.service.FindAll(instructors, keyword)

    if err != nil {
        util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
        return
    }
    respArr := make([]map[string]interface{}, 0, 0)

    for _, n := range *data {
        resp := n.ResponseMap()
        respArr = append(respArr, resp)
    }

    ctx.JSON(http.StatusOK, &util.Response{
        Success: true,
        Message: "Post result set",
        Data: map[string]interface{}{
            "rows":       respArr,
            "total_rows": total,
        }})
}

// AddPost : AddPost controller
func (i *InstructorController) AddPost(ctx *gin.Context) {
    var instructor models.Instructor
    ctx.ShouldBindJSON(&instructor)

    if instructor.Name == "" {
        util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
        return
    }
    if instructor.Surname == "" {
        util.ErrorJSON(ctx, http.StatusBadRequest, "Surnume is required")
        return
    }
	if instructor.Profile == "" {
        util.ErrorJSON(ctx, http.StatusBadRequest, "Profile is required")
        return
    }
    if instructor.Description == "" {
        util.ErrorJSON(ctx, http.StatusBadRequest, "Description is required")
        return
    }

    err := i.service.Save(instructor)

    if err != nil {
        util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create post")
        return
    }
    util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Post")
}

//GetPost : get post by id
func (i *InstructorController) GetPost(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
    if err != nil {
        util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
        return
    }
    var instructor models.Instructor
    instructor.Instructor_ID = id
    foundPost, err := i.service.Find(instructor)
    if err != nil {
        util.ErrorJSON(c, http.StatusBadRequest, "Error Finding Post")
        return
    }
    response := foundPost.ResponseMap()

    c.JSON(http.StatusOK, &util.Response{
        Success: true,
        Message: "Result set of Post",
        Data:    &response})

}

func (i *InstructorController) DeletePost(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
    if err != nil {
        util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
        return
    }
    err = i.service.Delete(id)

    if err != nil {
        util.ErrorJSON(c, http.StatusBadRequest, "Failed to delete Post")
        return
    }
    response := &util.Response{
        Success: true,
        Message: "Deleted Sucessfully"}
    c.JSON(http.StatusOK, response)
}

//UpdatePost : get update by id
func (i InstructorController) UpdatePost(ctx *gin.Context) {
    idParam := ctx.Param("id")

    id, err := strconv.ParseInt(idParam, 10, 64)

    if err != nil {
        util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
        return
    }
    var instructor models.Instructor
    instructor.Instructor_ID = id

    instructorRecord, err := i.service.Find(instructor)

    if err != nil {
        util.ErrorJSON(ctx, http.StatusBadRequest, "Post with given id not found")
        return
    }
    ctx.ShouldBindJSON(&instructorRecord)

    if instructorRecord.Name == "" {
        util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
        return
    }
    if instructorRecord.Surname == "" {
        util.ErrorJSON(ctx, http.StatusBadRequest, "Surname is required")
        return
    }
	if instructorRecord.Profile == "" {
        util.ErrorJSON(ctx, http.StatusBadRequest, "Profile is required")
        return
    }
    if instructorRecord.Description == "" {
        util.ErrorJSON(ctx, http.StatusBadRequest, "Description is required")
        return
    }

    if err := i.service.Update(instructorRecord); err != nil {
        util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Post")
        return
    }
    response := instructorRecord.ResponseMap()

    ctx.JSON(http.StatusOK, &util.Response{
        Success: true,
        Message: "Successfully Updated Post",
        Data:    response,
    })
}
