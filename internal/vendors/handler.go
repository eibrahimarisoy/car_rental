package vendors

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

// NewVendorHandler creates a new vendor handler
func NewVendorHandler(r *gin.RouterGroup, vendorService VendorServiceInterface) {
	handler := &VendorHandler{
		vendorService: vendorService,
	}

	r.GET("/", mw.PaginationMiddleware(), handler.GetVendors)
	r.POST("/", handler.CreateVendor)
}

// GetVendors is a handler to get all vendors
// @Summary      List all vendors
// @Description  List all vendors with pagination and search
// @Tags         vendors
// @Accept       json
// @Produce      json
// @Param        q     query    string  false  "Search query"
// @Param        page  query    int     false  "Page number"
// @Param        limit query    int     false  "Page limit"
// @Success      200  {object}  vendors.VendorListResponse
// @Failure      500  {object}  _type.APIErrorResponse
// @Router       /vendors/    [get]
func (h *VendorHandler) GetVendors(c *gin.Context) {
	pagination := c.MustGet("pagination").(*paginationHelper.Pagination)

	vendors, err := h.vendorService.GetVendors(pagination)

	if err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, VendorsToVendorListResponse(vendors, pagination))
}

// CreateVendor is a handler to create a vendor
// @Summary      Create a vendor
// @Description  Create a vendor with payload
// @Tags         vendors
// @Accept       json
// @Produce      json
// @Param        body body      vendors.VendorRequest  true  "Vendor payload"
// @Success      201  {object}  vendors.VendorResponse
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

	c.JSON(200, VendorToResponse(vendor))
	return
}
