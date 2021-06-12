package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rH RoutesHandler) CreateWorkflowHandler(c *gin.Context) {
	user := rH.getAuthenticatedUser(c)
	if user == nil {
		return
	}

	name := "test"

	err := rH.Usecases.CreateWorkflow(user, name)
	if err != nil {
		rH.handleError(c, err)
		return
	}
	c.Status(http.StatusCreated) //TODO: send workflow id
}

func (rH RoutesHandler) GetWorkflowHandler(c *gin.Context) {
	user := rH.getAuthenticatedUser(c)
	if user == nil {
		return
	}

	workflows, err := rH.Usecases.GetUserWorkflows(user)

	if err != nil {
		rH.handleError(c, err)
	}

	c.JSON(http.StatusOK, workflows)
}

func (rH RoutesHandler) DeleteWorkflowHandler(c *gin.Context) {
	user := rH.getAuthenticatedUser(c)
	if user == nil {
		return
	}

	id := "0ac1e794-cde1-4cd4-a475-be802e552fdf"

	err := rH.Usecases.DeleteWorkflow(user, id)

	if err != nil {
		rH.handleError(c, err)
	}

	c.Status(http.StatusOK)
}
