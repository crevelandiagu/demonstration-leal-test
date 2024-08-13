package handler

import (
	"demonstration-leal-test/internal/core/port"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LealHandler struct {
	svc port.LealHandler
}

func NewLealHandler(svc port.LealHandler) *LealHandler {
	return &LealHandler{
		svc,
	}
}

//-----------------------------------------------------------------------------------------

// @Summary      New User
// @Description  Save New User
// @Tags         User
// @Accept       json
// @Produce      json
// @Param user body domain.User true "User Data"
// @Success      200  {object}  domain.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/user [post]
func (lh *LealHandler) SaveUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary      Find User
// @Description  Find User
// @Tags         User
// @Accept       json
// @Produce      json
// @Param id path int true "User Data"
// @Success      200  {object}  domain.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/user/{id} [get]
func (lh *LealHandler) FindOneUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary      New Commerce
// @Description  Save New Commerce
// @Tags         Commerce
// @Accept       json
// @Produce      json
// @Param user body domain.Commerce true "User Data"
// @Success      200  {object}  domain.Commerce
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/commerce [post]
func (lh *LealHandler) SaveCommerce(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary      New Branche Commerce
// @Description  Save New Branche Commerce
// @Tags         Commerce
// @Accept       json
// @Produce      json
// @Param id path int true "User Data"
// @Param user body domain.Branches true "User Data"
// @Success      200  {object}  domain.Branches
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/commerce/{id}/branche [post]
func (lh *LealHandler) SaveBrancheCommerce(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary      Find Commerce
// @Description  Find Commerce
// @Tags         Commerce
// @Accept       json
// @Produce      json
// @Param id path int true "User Data"
// @Success      200  {object}  domain.Commerce
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/commerce/{id} [get]
func (lh *LealHandler) FindCommerce(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary      New Purchase Transaction
// @Description  Save Purchase Transaction
// @Tags         Commerce
// @Accept       json
// @Produce      json
// @Param id path int true "User Data"
// @Param user body domain.PurchaseTransaction true "User Data"
// @Success      200  {object}  domain.PurchaseTransaction
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/commerce/{id}/purchase-transaction [post]
func (lh *LealHandler) SavePurchaseTransaction(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary      Find Purchase Transaction
// @Description  Find Purchase Transaction
// @Tags         Commerce
// @Accept       json
// @Produce      json
// @Param id path int true "User Data"
// @Success      200  {object}  domain.PurchaseTransaction
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/commerce/{id}/purchase-transaction [get]
func (lh *LealHandler) FindPurchaseTransaction(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary      New Campaign Branche
// @Description  Save Campaign Branche
// @Tags         Commerce
// @Accept       json
// @Produce      json
// @Param id path int true "User Data"
// @Param idBranche path int true "User Data"
// @Param user body domain.Campaign true "User Data"
// @Success      200  {object}  domain.Campaign
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/commerce/{id}/branche/{idBranche}/campaign  [post]
func (lh *LealHandler) SaveCampaignBranche(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary      Find Campaign Branche
// @Description  Find Campaign Branche
// @Tags         Commerce
// @Accept       json
// @Produce      json
// @Param id path int true "User Data"
// @Param idBranche path int true "User Data"
// @Success      200  {object}  domain.Campaign
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/commerce/{id}/branche/{idBranche}/campaign [get]
func (lh *LealHandler) FindCampaignBranche(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary      New Redeem
// @Description  Save Redeem
// @Tags         Commerce
// @Accept       json
// @Produce      json
// @Param id path int true "User Data"
// @Param user body domain.LealPoints true "User Data"
// @Success      200  {object}  domain.LealPoints
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/commerce/{id}/redeem/ [post]
func (lh *LealHandler) SaveRedeem(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary      Find Redeem
// @Description  Find Redeem
// @Tags         Commerce
// @Accept       json
// @Produce      json
// @Param id path int true "User Data"
// @Success      200  {object}  domain.LealPoints
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/commerce/{id}/redeem/ [get]
func (lh *LealHandler) FindRedeem(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary      New Campaign Commerce
// @Description  Save Campaign Commerce
// @Tags         Commerce
// @Accept       json
// @Produce      json
// @Param id path int true "User Data"
// @Param user body domain.Campaign true "User Data"
// @Success      200  {object}  domain.Campaign
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/commerce/{id}/campaign [post]
func (lh *LealHandler) SaveCampaignCommerce(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary      Find Campaign Commerce
// @Description  Find Campaign Commerce
// @Tags         Commerce
// @Accept       json
// @Produce      json
// @Param id path int true "User Data"
// @Success      200  {object}  domain.Campaign
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/commerce/{id}/campaign [get]
func (lh *LealHandler) FindCampaignCommerce(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// response represents a response body format
type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

func handleSuccess(ctx *gin.Context, data any) {
	rsp := newResponse(true, "Success", data)
	ctx.JSON(http.StatusOK, rsp)
}

func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}
