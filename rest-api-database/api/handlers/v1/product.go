package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golanguzb70/tracing-examples/rest-api-database/models"
)

// @Router		/product [POST]
// @Summary		Create product
// @Tags        Product
// @Description	Here product can be created.
// @Security    BearerAuth
// @Accept      json
// @Produce		json
// @Param       post   body       models.ProductCreateReq true "post info"
// @Success		200 	{object}  models.ProductApiResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) ProductCreate(c *gin.Context) {
	body := &models.ProductCreateReq{}
	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.storage.Postgres().ProductCreate(context.Background(), body)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "ProductCreate: h.storage.Postgres().ProductCreate()") {
		return
	}

	c.JSON(http.StatusOK, &models.ProductApiResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/product/{id} [GET]
// @Summary		Get product by key
// @Tags        Product
// @Description	Here product can be got.
// @Accept      json
// @Produce		json
// @Param       id       path     int true "id"
// @Success		200 	{object}  models.ProductApiResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) ProductGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "ProductGet: strconv.Atoi()") {
		return
	}

	res, err := h.storage.Postgres().ProductGet(context.Background(), &models.ProductGetReq{
		Id: id,
	})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "ProductGet: h.storage.Postgres().ProductGet()") {
		return
	}

	c.JSON(http.StatusOK, models.ProductApiResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/product/list [GET]
// @Summary		Get products list
// @Tags        Product
// @Description	Here all products can be got.
// @Accept      json
// @Produce		json
// @Param       filters query models.ProductFindReq true "filters"
// @Success		200 	{object}  models.ProductApiFindResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) ProductFind(c *gin.Context) {
	var (
		dbReq = &models.ProductFindReq{}
		err   error
	)
	dbReq.Page, err = ParsePageQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "ProductFind: helper.ParsePageQueryParam(c)") {
		return
	}
	dbReq.Limit, err = ParseLimitQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "ProductFind: helper.ParseLimitQueryParam(c)") {
		return
	}

	dbReq.Search = c.Query("search")
	dbReq.OrderByCreatedAt, _ = strconv.ParseUint(c.Query("order_by_created_at"), 10, 8)

	res, err := h.storage.Postgres().ProductFind(context.Background(), dbReq)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "ProductFind: h.storage.Postgres().ProductFind()") {
		return
	}

	c.JSON(http.StatusOK, &models.ProductApiFindResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/product [PUT]
// @Summary		Update product
// @Tags        Product
// @Description	Here product can be updated.
// @Security    BearerAuth
// @Accept      json
// @Produce		json
// @Param       post   body       models.ProductUpdateReq true "post info"
// @Success		200 	{object}  models.ProductApiResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) ProductUpdate(c *gin.Context) {
	body := &models.ProductUpdateReq{}
	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "ProductUpdate: c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.storage.Postgres().ProductUpdate(context.Background(), body)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "ProductUpdate: h.storage.Postgres().ProductUpdate()") {
		return
	}

	c.JSON(http.StatusOK, &models.ProductApiResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/product/{id} [DELETE]
// @Summary		Delete product
// @Tags        Product
// @Description	Here product can be deleted.
// @Security    BearerAuth
// @Accept      json
// @Produce		json
// @Param       id       path     int true "id"
// @Success		200 	{object}  models.DefaultResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) ProductDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "ProductDelete: strconv.Atoi()") {
		return
	}

	err = h.storage.Postgres().ProductDelete(context.Background(), &models.ProductDeleteReq{Id: id})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "ProductDelete: h.storage.Postgres().ProductDelete()") {
		return
	}

	c.JSON(http.StatusOK, models.DefaultResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "Successfully deleted",
	})
}
