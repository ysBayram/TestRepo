package facades

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ysbayram/TestRepo/entities"
	"github.com/ysbayram/TestRepo/utils"
)

type CategoryFacade struct {
	ce entities.CategoryEntity
}

func NewCategoryFacade(ce entities.CategoryEntity) CategoryFacade {
	return CategoryFacade{ce}
}

func (h *CategoryFacade) DeleteCategory(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	paramID := ctx.Params.ByName("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, h.ce.DeleteCategory(id))
}

func (h *CategoryFacade) UpdateCategory(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	var tmpObj entities.Category
	if err := ctx.BindJSON(&tmpObj); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, h.ce.UpdateCategory(tmpObj))
	}
}

func (h *CategoryFacade) CreateCategory(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	var c entities.Category
	if err := ctx.BindJSON(&c); err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, h.ce.CreateCategory(c))
}

func (h *CategoryFacade) GetCategoryByID(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	paramID := ctx.Params.ByName("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, h.ce.GetCategoryByID(id))
}

func (h *CategoryFacade) GetCategory(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	ctx.JSON(http.StatusOK, h.ce.GetCategory())
}
