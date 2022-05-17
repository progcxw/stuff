package controller

import (
	"stuff/internal/errors"
	"stuff/internal/model"
	"stuff/internal/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CategoryController struct{}

func (CategoryController) ListCategory(ctx *gin.Context) {
	var req model.ListParam
	c := utils.DefaultGinContext(ctx)
	traceID := utils.ShouldGetTraceID(c)
	log.Debug(traceID)

	resp := utils.DefaultResponse()
	defer utils.HTTPJSONResponse(c, ctx, &resp)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("ListCategory get request param error: %+v, traceID: %s", err, traceID)
		utils.SetErrorResponse(&resp, errors.CodeInvalidParams, err.Error())
		return
	}

	lists, count, err := model.GetCategoryList(c, req.Name, req.Limit, req.Page)
	if err != nil {
		log.Errorf("ListCategory get list error: %+v, traceID: %s", err, traceID)
		utils.SetErrorResponse(&resp, errors.CodeInternalServerError, err.Error())
		return
	}

	utils.SetSuccessfulResponse(&resp, errors.CodeOK, model.ListResp{
		Lists: lists,
		Total: count,
	})
}
