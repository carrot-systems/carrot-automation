package rest

import (
	"github.com/carrot-systems/carrot-automation/src/core/domain"
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

	id := c.Param("id")

	err := rH.Usecases.DeleteWorkflow(user, id)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func (rH RoutesHandler) UpdateWorkflowActionsHandler(c *gin.Context) {
	user := rH.getAuthenticatedUser(c)
	if user == nil {
		return
	}

	id := c.Param("id")
	_ = id

	var content domain.WorkflowContent
	err := c.BindJSON(&content)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	err = rH.Usecases.SetWorkflowData(user, id, content)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.Status(http.StatusOK)

}
