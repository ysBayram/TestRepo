package facades

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ysbayram/TestRepo/entities"
	"github.com/ysbayram/TestRepo/utils"
)

type CouponFacade struct {
	ce entities.CouponEntity
}

func NewCouponFacade(ce entities.CouponEntity) CouponFacade {
	return CouponFacade{ce}
}

func (h *CouponFacade) DeleteCoupon(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, h.ce.DeleteCoupon(id))
}

func (h *CouponFacade) UpdateCoupon(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	var tmpObj entities.Coupon
	if err := ctx.BindJSON(&tmpObj); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, h.ce.UpdateCoupon(tmpObj))
	}
}

func (h *CouponFacade) CreateCoupon(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	var c entities.Coupon
	if err := ctx.BindJSON(&c); err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, h.ce.CreateCoupon(c))
}

func (h *CouponFacade) GetCouponByID(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, h.ce.GetCouponByID(id))
}

func (h *CouponFacade) GetCoupon(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	ctx.JSON(http.StatusOK, h.ce.GetCoupon())
}
