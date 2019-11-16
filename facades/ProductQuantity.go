package facades

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ysbayram/TestRepo/entities"
	"github.com/ysbayram/TestRepo/utils"
)

type ProductQuantityFacade struct {
	ce entities.ProductQuantityEntity
}

func NewProductQuantityFacade(ce entities.ProductQuantityEntity) ProductQuantityFacade {
	return ProductQuantityFacade{ce}
}

func (h *ProductQuantityFacade) DeleteProductQuantity(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, h.ce.DeleteProductQuantity(id))
}

func (h *ProductQuantityFacade) UpdateProductQuantity(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	var tmpObj entities.ProductQuantity
	if err := ctx.BindJSON(&tmpObj); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, h.ce.UpdateProductQuantity(tmpObj))
	}
}

func (h *ProductQuantityFacade) CreateProductQuantity(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	var c entities.ProductQuantity
	if err := ctx.BindJSON(&c); err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, h.ce.CreateProductQuantity(c))
}

func (h *ProductQuantityFacade) GetProductQuantityByID(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, h.ce.GetProductQuantityByID(id))
}

func (h *ProductQuantityFacade) GetProductQuantity(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	ctx.JSON(http.StatusOK, h.ce.GetProductQuantity())
}
