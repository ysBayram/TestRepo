package facades

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ysbayram/TestRepo/entities"
	"github.com/ysbayram/TestRepo/utils"
)

type CartFacade struct {
	cartEnt            entities.CartEntity
	campaignEnt        *entities.CampaignEntity
	couponEnt          *entities.CouponEntity
	categoryEnt        *entities.CategoryEntity
	productEnt         *entities.ProductEntity
	productQuantityEnt *entities.ProductQuantityEntity
}

func NewCartFacade(ce entities.CartEntity, ca *entities.CampaignEntity, cu *entities.CouponEntity, co *entities.CategoryEntity, pe *entities.ProductEntity, pqe *entities.ProductQuantityEntity) CartFacade {
	return CartFacade{ce, ca, cu, co, pe, pqe}
}

func (h *CartFacade) DeleteCart(ctx *gin.Context) {
	defer utils.ErrorFuncForDefer(ctx)
	paramID := ctx.Params.ByName("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		logrus.Panic(err)
	}
	ctx.JSON(http.StatusOK, h.cartEnt.DeleteCart(id))
}

func (h *CartFacade) UpdateCart(ctx *gin.Context) {
	defer utils.ErrorFuncForDefer(ctx)
	var tmpObj entities.Cart
	if err := ctx.BindJSON(&tmpObj); err != nil {
		logrus.Panic(err)
	} else {
		ctx.JSON(http.StatusOK, h.cartEnt.UpdateCart(tmpObj))
	}
}

func (h *CartFacade) CreateCart(ctx *gin.Context) {
	defer utils.ErrorFuncForDefer(ctx)
	var c entities.Cart
	if err := ctx.BindJSON(&c); err != nil {
		logrus.Panic(err)
	}
	ctx.JSON(http.StatusOK, h.cartEnt.CreateCart(c))
}

func (h *CartFacade) HandlerGetCartByID(ctx *gin.Context) {
	defer utils.ErrorFuncForDefer(ctx)

	paramID := ctx.Params.ByName("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		logrus.Panic(err)
	}

	result := h.GetCartByID(id)

	ctx.JSON(http.StatusOK, result)
}

func (h *CartFacade) GetCartByID(cartID int) entities.Cart {
	result := h.cartEnt.GetCartByID(cartID)
	pqArr := h.productQuantityEnt.GetProductQuantityByCartID(cartID)
	for i := 0; i < len(pqArr); i++ {
		tmpProduct := h.productEnt.GetProductByID(pqArr[i].ProductID)
		tmpCategory := h.categoryEnt.GetCategoryByID(tmpProduct.CategoryID)
		tmpProduct.Category = tmpCategory
		pqArr[i].Product = tmpProduct
	}
	result.Products = pqArr
	return result
}

func (h *CartFacade) GetCart(ctx *gin.Context) {
	defer utils.ErrorFuncForDefer(ctx)
	cartArr := h.cartEnt.GetCart()
	for _, item := range cartArr {
		item = h.GetCartByID(int(item.ID))
	}
	ctx.JSON(http.StatusOK, cartArr)
}

func (h *CartFacade) HandlerApplyProduct(ctx *gin.Context) {
	defer utils.ErrorFuncForDefer(ctx)
	var reqInput entities.ProductQuantity
	if err := ctx.BindJSON(&reqInput); err != nil {
		logrus.Panic(err)
	}
	h.productQuantityEnt.CreateProductQuantity(reqInput)
	tmpCart := h.GetCartByID(reqInput.CartID)
	var tmpTotalAmount float64
	for _, item := range tmpCart.Products {
		tmpTotalAmount += item.Product.Price
	}
	tmpCart.TotalAmount = tmpTotalAmount
	h.cartEnt.UpdateCart(tmpCart)
	result := h.GetCartByID(reqInput.CartID)
	ctx.JSON(http.StatusOK, result)
}

func (h *CartFacade) HandlerApplyDiscount(ctx *gin.Context) {
	defer utils.ErrorFuncForDefer(ctx)
	type dtoApplyDiscounts struct {
		CartID       int
		Discount     float64
		DiscountType entities.DiscountType
	}
	var reqInput dtoApplyDiscounts

	if err := ctx.BindJSON(&reqInput); err != nil {
		logrus.Panic(err)
	}
	result := h.applyDiscount(reqInput.CartID, reqInput.Discount, reqInput.DiscountType)
	ctx.JSON(http.StatusOK, result)
}

func (h *CartFacade) applyDiscount(CartID int, discount float64, discountType entities.DiscountType) (c entities.Cart) {
	tmpCart := h.GetCartByID(CartID)
	if discountType == entities.AMOUNT {
		tmpCart.TotalDiscounts += discount
	} else if discountType == entities.RATE {
		tmpDiscount := tmpCart.TotalAmount * (discount / 100)
		tmpCart.TotalDiscounts += tmpDiscount
	}
	return h.cartEnt.UpdateCart(tmpCart)
}

func (h *CartFacade) HandlerApplyCampaign(ctx *gin.Context) {
	defer utils.ErrorFuncForDefer(ctx)
	paramID := ctx.Params.ByName("cartid")
	CartID, err := strconv.Atoi(paramID)
	if err != nil {
		logrus.Panic(err)
	}
	tmpCart := h.GetCartByID(CartID)
	var campaigns []entities.Campaign
	var productCount int
	for _, v := range tmpCart.Products {
		campaigns = h.campaignEnt.GetCampaignByCategory(v.Product.CategoryID)
		productCount += v.Quantity
	}
	for _, v := range campaigns {
		if productCount > v.ProductCount {
			h.applyDiscount(CartID, v.Discount, v.DiscountType)
		}
	}
	result := h.GetCartByID(CartID)
	ctx.JSON(http.StatusOK, result)
}

func (h *CartFacade) HandlerApplyCoupon(ctx *gin.Context) {
	defer utils.ErrorFuncForDefer(ctx)
	type dtoApplyCoupon struct {
		CartID      int
		CouponValue string
	}
	var reqInput dtoApplyCoupon
	if err := ctx.BindJSON(&reqInput); err != nil {
		logrus.Panic(err)
	}

	tmpCart := h.GetCartByID(reqInput.CartID)
	coupon := h.couponEnt.GetCouponByValue(reqInput.CouponValue)
	if coupon.ID != 0 {
		if tmpCart.TotalAmount > coupon.MinAmount {
			h.applyDiscount(reqInput.CartID, coupon.Discount, coupon.DiscountType)
		} else {
			logrus.Panic("Need more total amount for apply coupon!")
		}
	} else {
		logrus.Panic("Please enter valid coupon Value!")
	}
	ctx.JSON(http.StatusOK, tmpCart)
}

func (h *CartFacade) HandlerCalculateDeliveryCost(ctx *gin.Context) {
	defer utils.ErrorFuncForDefer(ctx)
	type dtoCalculateDeliveryCost struct {
		CartID          int
		CostPerDelivery float64
		CostPerProduct  float64
	}
	var reqInput dtoCalculateDeliveryCost
	if err := ctx.BindJSON(&reqInput); err != nil {
		logrus.Panic(err)
	}
	tmpCart := h.GetCartByID(reqInput.CartID)
	categoryCountMap := make(map[string]int)
	for _, v := range tmpCart.Products {
		if val, ok := categoryCountMap[v.Product.Category.Title]; ok {
			categoryCountMap[v.Product.Category.Title] = val + 1
		} else {
			categoryCountMap[v.Product.Category.Title] = 1
		}
	}

	NumberOfDeliveries := len(categoryCountMap)
	NumberOfProducts := len(tmpCart.Products)

	deliveryCost := (reqInput.CostPerDelivery * float64(NumberOfDeliveries)) + (reqInput.CostPerProduct * float64(NumberOfProducts)) + utils.FIXEDCOST

	result := make(map[string]float64)
	result["DeliveryCost"] = deliveryCost
	tmpCart.DeliveryCost = deliveryCost
	h.cartEnt.UpdateCart(tmpCart)

	ctx.JSON(http.StatusOK, result)
}

func (h *CartFacade) HandlerGetTotalAmountAfterDiscounts(ctx *gin.Context) {
	defer utils.ErrorFuncForDefer(ctx)
	paramID := ctx.Params.ByName("cartid")
	cartID, err := strconv.Atoi(paramID)
	if err != nil {
		logrus.Panic(err)
	}
	tmpCart := h.GetCartByID(cartID)
	result := make(map[string]float64)
	result["TotalAmountAfterDiscounts"] = tmpCart.TotalAmount - tmpCart.TotalDiscounts
	ctx.JSON(http.StatusOK, result)
}

func (h *CartFacade) HandlerGetCouponDiscount(ctx *gin.Context) {
	defer utils.ErrorFuncForDefer(ctx)
	paramID := ctx.Params.ByName("cartid")
	cartID, err := strconv.Atoi(paramID)
	if err != nil {
		logrus.Panic(err)
	}
	tmpCart := h.GetCartByID(cartID)
	tmpCoupon := h.couponEnt.GetCouponByID(tmpCart.CouponID)
	result := make(map[string]float64)
	if tmpCoupon.DiscountType == entities.AMOUNT {
		result["CouponDiscount"] = tmpCoupon.Discount
	} else if tmpCoupon.DiscountType == entities.RATE {
		tmpDiscount := tmpCart.TotalAmount * (tmpCoupon.Discount / 100)
		result["CouponDiscount"] = tmpDiscount
	}
	ctx.JSON(http.StatusOK, result)
}

func (h *CartFacade) HandlerCartPrint(ctx *gin.Context) {
	defer utils.ErrorFuncForDefer(ctx)
	paramID := ctx.Params.ByName("cartid")
	cartID, err := strconv.Atoi(paramID)
	if err != nil {
		logrus.Panic(err)
	}

	result := make(map[string]interface{})
	tmpCart := h.GetCartByID(cartID)

	for _, pq := range tmpCart.Products {
		if _, ok := result[pq.Product.Category.Title]; !ok {
			result[pq.Product.Category.Title] = make([]map[string]interface{}, 0)
		}
		productDetail := make(map[string]interface{})
		productDetail["ProductTitle"] = pq.Product.Title
		productDetail["ProductTitle"] = pq.Product.Title
		productDetail["ProductQuantity"] = pq.Quantity
		productDetail["UnitPrice"] = pq.Product.Price
		result[pq.Product.Category.Title] = append(result[pq.Product.Category.Title].([]map[string]interface{}), productDetail)
	}
	result["TotalDiscount"] = tmpCart.TotalDiscounts
	result["TotalPrice"] = tmpCart.TotalAmount

	ctx.JSON(http.StatusOK, result)
}
