package location

import (
	mw "github.com/eibrahimarisoy/car_rental/pkg/middlewares"
	paginationHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	locationService LocationServiceInterface
}

func NewLocationHandler(r *gin.RouterGroup, locationService LocationServiceInterface) {
	handler := &LocationHandler{
		locationService: locationService,
	}

	r.GET("/", mw.PaginationMiddleware(), handler.GetAllLocations)

}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  models.Location
// @Failure      400  {object}  _type.ErrorType
// @Failure      404  {object}  _type.ErrorType
// @Failure      500  {object}  _type.ErrorType
// @Router       /accounts/{id} [get]
func (h *LocationHandler) GetAllLocations(c *gin.Context) {
	pagination := c.MustGet("pagination").(*paginationHelper.Pagination)

	pagination, err := h.locationService.GetAllLocations(pagination)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"locations": pagination,
	})
}
