package vendor

import (
	"net/http"

	"github.com/eibrahimarisoy/car_rental/pkg/errorHandler"
	mw "github.com/eibrahimarisoy/car_rental/pkg/middlewares"
	paginationHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"github.com/gin-gonic/gin"
)

type VendorHandler struct {
	vendorService VendorServiceInterface
}

func NewVendorHandler(r *gin.RouterGroup, vendorService VendorServiceInterface) {
	handler := &VendorHandler{
		vendorService: vendorService,
	}

	r.POST("/", handler.CreateVendor)
	r.GET("/", mw.PaginationMiddleware(), handler.GetVendors)
}

// CreateVendor is a handler to create a vendor
// @Summary      Create a vendor
// @Description  Create a vendor with payload
// @Tags         vendor
// @Accept       json
// @Produce      json
// @Param        body body      vendor.VendorRequest  true  "Vendor payload"
// @Success      200  {object}  vendor.VendorResponse
// @Failure	     400  {object}  _type.APIErrorResponse
// @Failure      500  {object}  _type.APIErrorResponse
// @Router       /vendors/    [post]
func (h *VendorHandler) CreateVendor(c *gin.Context) {
	reqBody := &VendorRequest{}

	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	if err := reqBody.Validate(); err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	vendor, err := h.vendorService.CreateVendor(reqBody.ToVendor())

	if err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	resBody := &VendorResponse{}
	resBody.FromVendor(vendor)

	c.JSON(200, resBody)
	return
}

// GetVendors is a handler to get all vendors
// @Summary      List all vendors
// @Description  List all vendors with pagination and search
// @Tags         vendor
// @Accept       json
// @Produce      json
// @Param        q     query    string  false  "Search query"
// @Param        page  query    int     false  "Page number"
// @Param        limit query    int     false  "Page limit"
// @Success      200  {object}  pagination.Pagination
// @Failure      500  {object}  _type.APIErrorResponse
// @Router       /vendors/    [get]
func (h *VendorHandler) GetVendors(c *gin.Context) {
	pagination := c.MustGet("pagination").(*paginationHelper.Pagination)

	pagination, err := h.vendorService.GetVendors(pagination)

	if err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, pagination)
}
