package facades

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ysbayram/TestRepo/entities"
	"github.com/ysbayram/TestRepo/utils"
)

type CampaignFacade struct {
	ce entities.CampaignEntity
}

func NewCampaignFacade(ce entities.CampaignEntity) CampaignFacade {
	return CampaignFacade{ce}
}

func (h *CampaignFacade) DeleteCampaign(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, h.ce.DeleteCampaign(id))
}

func (h *CampaignFacade) UpdateCampaign(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	var tmpObj entities.Campaign
	if err := ctx.BindJSON(&tmpObj); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, h.ce.UpdateCampaign(tmpObj))
	}
}

func (h *CampaignFacade) CreateCampaign(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	var c entities.Campaign
	if err := ctx.BindJSON(&c); err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, h.ce.CreateCampaign(c))
}

func (h *CampaignFacade) GetCampaignByID(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, h.ce.GetCampaignByID(id))
}

func (h *CampaignFacade) GetCampaign(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, utils.GetStackFrame(4, 6))
			ctx.JSON(http.StatusInternalServerError, r)
		}
	}()
	ctx.JSON(http.StatusOK, h.ce.GetCampaign())
}
