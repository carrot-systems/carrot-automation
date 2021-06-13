package rest

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine, routesHandler RoutesHandler) {
	r.Use(routesHandler.fetchingUserMiddleware())
	r.NoRoute(routesHandler.endpointNotFound)
	servicesGroup := r.Group("/services")
	servicesGroup.GET("/", routesHandler.GetServicesHandler)
	servicesGroup.GET("/:id", routesHandler.GetActionsHandler)

	workflowsGroup := r.Group("/workflows")
	workflowsGroup.POST("/", routesHandler.CreateWorkflowHandler)
	workflowsGroup.GET("/", routesHandler.GetWorkflowHandler)
	workflowsGroup.DELETE("/:id", routesHandler.DeleteWorkflowHandler)
	workflowsGroup.POST("/:id", routesHandler.UpdateWorkflowActionsHandler)
}
