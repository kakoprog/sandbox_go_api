package server

import (
	"github.com/labstack/echo"
	"github.com/kakoprog/sandbox_go_api/config"
	"github.com/kakoprog/sandbox_go_api/api/todo/statuses"
  "github.com/kakoprog/sandbox_go_api/api/todo/tasks"
	"strconv"
)

func Run() {
	port := config.Read("server", "port").(int)

	e := echo.New()
	setRouting(e)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))
}

func setRouting(e *echo.Echo) {
  routeApi := e.Group("/api")
	routeApiTodo := routeApi.Group("/todo")
	routeApiTodoTasks := routeApiTodo.Group("/tasks")
	routeApiTodoTasks.GET("", tasks.GetContents)
	routeApiTodoTasks.GET("/:id", tasks.GetContent)
	routeApiTodoTasks.POST("", tasks.PostContent)
	routeApiTodoTasks.PUT("/:id", tasks.PutContent)
	routeApiTodoTasks.DELETE("/:id", tasks.DeleteContent)
	routeApiTodoStatuses := routeApiTodo.Group("/statuses")
	routeApiTodoStatuses.GET("", statuses.GetContents)
	routeApiTodoStatuses.GET("/:id", statuses.GetContent)
}
