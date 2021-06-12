package rest

import (
	"github.com/carrot-systems/carrot-automation/src/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rH RoutesHandler) GetServicesHandler(c *gin.Context) {
	if rH.getAuthenticatedUser(c) == nil {
		return
	}

	services, err := rH.Usecases.GetAllServices()

	if err != nil {
		rH.handleError(c, err)
	}

	c.JSON(http.StatusOK, domain.Status{
		Success: true,
		Data:    services,
	})
}

func (rH RoutesHandler) GetActionsHandler(c *gin.Context) {
	if rH.getAuthenticatedUser(c) == nil {
		return
	}

	serviceId := c.Param("id")

	if serviceId == "" {
		rH.handleError(c, ErrFormValidation)
		return
	}

	actions, err := rH.Usecases.GetServiceActions(serviceId)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, domain.Status{
		Success: true,
		Data:    actions,
	})
}
