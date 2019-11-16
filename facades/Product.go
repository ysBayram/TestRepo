package facades

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ysbayram/TestRepo/entities"
	"github.com/ysbayram/TestRepo/utils"
)

type ProductFacade struct {
	ce entities.ProductEntity
}

func NewProductFacade(ce entities.ProductEntity) ProductFacade {
	return ProductFacade{ce}
}

func (h *ProductFacade) DeleteProduct(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, h.ce.DeleteProduct(id))
}

func (h *ProductFacade) UpdateProduct(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	var tmpObj entities.Product
	if err := ctx.BindJSON(&tmpObj); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, h.ce.UpdateProduct(tmpObj))
	}
}

func (h *ProductFacade) CreateProduct(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	var c entities.Product
	if err := ctx.BindJSON(&c); err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, h.ce.CreateProduct(c))
}

func (h *ProductFacade) GetProductByID(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, h.ce.GetProductByID(id))
}

func (h *ProductFacade) GetProduct(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	ctx.JSON(http.StatusOK, h.ce.GetProduct())
}
