package office

import (
	"net/http"

	"github.com/eibrahimarisoy/car_rental/pkg/errorHandler"
	mw "github.com/eibrahimarisoy/car_rental/pkg/middlewares"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"github.com/gin-gonic/gin"
)

type OfficeHandler struct {
	officeService OfficeServiceInterface
}

func NewOfficeHandler(r *gin.RouterGroup, officeService OfficeServiceInterface) {
	handler := &OfficeHandler{officeService: officeService}

	r.GET("/", mw.PaginationMiddleware(), handler.GetOffices)
	r.POST("/", handler.CreateOffice)
}

// GetAllOffices is a handler to get all office
// @Summary      List all offices
// @Description  List all offices with pagination and search
// @Tags         offices
// @Accept       json
// @Produce      json
// @Param        page  query    int     false  "Page number"
// @Param        limit query    int     false  "Page limit"
// @Success      200  {object}  pagination.Pagination
// @Failure      500  {object}  _type.APIErrorResponse
// @Router       /offices/    [get]
func (h *OfficeHandler) GetOffices(c *gin.Context) {
	pagination := c.MustGet("pagination").(*pgHelper.Pagination)

	pagination, err := h.officeService.GetOffices(pagination)
	if err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, pagination)
}

// CreateOffice is a handler to create a office
// @Summary      Create a office
// @Description  Create a office with payload
// @Tags         offices
// @Accept       json
// @Produce      json
// @Param        body body    office.OfficeRequest  true  "Office payload"
// @Success      200  {object}  office.OfficeResponse
// @Failure	     400  {object}  _type.APIErrorResponse
// @Failure      500  {object}  _type.APIErrorResponse
// @Router       /offices/    [post]
func (h *OfficeHandler) CreateOffice(c *gin.Context) {
	reqBody := OfficeRequest{}

	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	if err := reqBody.Validate(); err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	office, err := h.officeService.CreateOffice(reqBody.ToOffice())

	if err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	officeResponse := OfficeResponse{}
	officeResponse.FromOffice(office)
	c.JSON(http.StatusCreated, officeResponse)
}
