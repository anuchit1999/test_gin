package routes

import (
    "blog/api/controller"
    "blog/infrastructure"
)

//PostRoute -> Route for question module
type InstructorRoute struct {
    Controller controller.InstructorController
    Handler    infrastructure.GinRouter
}

//NewPostRoute -> initializes new choice rouets
func NewInstructorRoute(
    controller controller.InstructorController,
    handler infrastructure.GinRouter,

) InstructorRoute {
    return InstructorRoute{
        Controller: controller,
        Handler:    handler,
    }
}

//Setup -> setups new choice Routes
func (i InstructorRoute) Setup() {
    instructor := i.Handler.Gin.Group("/instructor") //Router group
    {
        instructor.GET("/", i.Controller.GetPosts)
        instructor.POST("/", i.Controller.AddPost)
        instructor.GET("/:id", i.Controller.GetPost)
        instructor.DELETE("/:id", i.Controller.DeletePost)
        instructor.PUT("/:id", i.Controller.UpdatePost)
    }
}
