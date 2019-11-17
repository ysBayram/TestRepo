package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/ysbayram/TestRepo/entities"
	"github.com/ysbayram/TestRepo/facades"
)

func CreateDBCon() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./testRepo.db")
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entities.Category{}, &entities.Campaign{}, &entities.Coupon{}, &entities.ProductQuantity{}, &entities.Product{}, &entities.Cart{}).Error
	if err != nil {
		panic(err)
	}

	return db
}

func SetupRouter(db *gorm.DB) *gin.Engine {

	r := gin.Default()

	//CATEGORY HANDLERS

	r.GET("/category", func(ctx *gin.Context) {
		categoryFacade := facades.NewCategoryFacade(entities.NewCategoryEntity(db))
		categoryFacade.GetCategory(ctx)
	})
	r.GET("/category/:id", func(ctx *gin.Context) {
		categoryFacade := facades.NewCategoryFacade(entities.NewCategoryEntity(db))
		categoryFacade.GetCategoryByID(ctx)
	})
	r.POST("/category", func(ctx *gin.Context) {
		categoryFacade := facades.NewCategoryFacade(entities.NewCategoryEntity(db))
		categoryFacade.CreateCategory(ctx)
	})
	r.PUT("/category/:id", func(ctx *gin.Context) {
		categoryFacade := facades.NewCategoryFacade(entities.NewCategoryEntity(db))
		categoryFacade.UpdateCategory(ctx)
	})
	r.DELETE("/category/:id", func(ctx *gin.Context) {
		categoryFacade := facades.NewCategoryFacade(entities.NewCategoryEntity(db))
		categoryFacade.DeleteCategory(ctx)
	})

	//CAMPAIGN HANDLERS

	r.GET("/campaign", func(ctx *gin.Context) {
		campaignFacade := facades.NewCampaignFacade(entities.NewCampaignEntity(db))
		campaignFacade.GetCampaign(ctx)
	})
	r.GET("/campaign/:id", func(ctx *gin.Context) {
		campaignFacade := facades.NewCampaignFacade(entities.NewCampaignEntity(db))
		campaignFacade.GetCampaignByID(ctx)
	})
	r.POST("/campaign", func(ctx *gin.Context) {
		campaignFacade := facades.NewCampaignFacade(entities.NewCampaignEntity(db))
		campaignFacade.CreateCampaign(ctx)
	})
	r.PUT("/campaign/:id", func(ctx *gin.Context) {
		campaignFacade := facades.NewCampaignFacade(entities.NewCampaignEntity(db))
		campaignFacade.UpdateCampaign(ctx)
	})
	r.DELETE("/campaign/:id", func(ctx *gin.Context) {
		campaignFacade := facades.NewCampaignFacade(entities.NewCampaignEntity(db))
		campaignFacade.DeleteCampaign(ctx)
	})

	//COUPON HANDLERS

	r.GET("/coupon", func(ctx *gin.Context) {
		couponFacade := facades.NewCouponFacade(entities.NewCouponEntity(db))
		couponFacade.GetCoupon(ctx)
	})
	r.GET("/coupon/:id", func(ctx *gin.Context) {
		couponFacade := facades.NewCouponFacade(entities.NewCouponEntity(db))
		couponFacade.GetCouponByID(ctx)
	})
	r.POST("/coupon", func(ctx *gin.Context) {
		couponFacade := facades.NewCouponFacade(entities.NewCouponEntity(db))
		couponFacade.CreateCoupon(ctx)
	})
	r.PUT("/coupon/:id", func(ctx *gin.Context) {
		couponFacade := facades.NewCouponFacade(entities.NewCouponEntity(db))
		couponFacade.UpdateCoupon(ctx)
	})
	r.DELETE("/coupon/:id", func(ctx *gin.Context) {
		couponFacade := facades.NewCouponFacade(entities.NewCouponEntity(db))
		couponFacade.DeleteCoupon(ctx)
	})

	//PRODUCT QUANTITY HANDLERS

	r.GET("/product-quantity", func(ctx *gin.Context) {
		productQuantityFacade := facades.NewProductQuantityFacade(entities.NewProductQuantityEntity(db))
		productQuantityFacade.GetProductQuantity(ctx)
	})
	r.GET("/product-quantity/:id", func(ctx *gin.Context) {
		productQuantityFacade := facades.NewProductQuantityFacade(entities.NewProductQuantityEntity(db))
		productQuantityFacade.GetProductQuantityByID(ctx)
	})
	r.POST("/product-quantity", func(ctx *gin.Context) {
		productQuantityFacade := facades.NewProductQuantityFacade(entities.NewProductQuantityEntity(db))
		productQuantityFacade.CreateProductQuantity(ctx)
	})
	r.PUT("/product-quantity/:id", func(ctx *gin.Context) {
		productQuantityFacade := facades.NewProductQuantityFacade(entities.NewProductQuantityEntity(db))
		productQuantityFacade.UpdateProductQuantity(ctx)
	})
	r.DELETE("/product-quantity/:id", func(ctx *gin.Context) {
		productQuantityFacade := facades.NewProductQuantityFacade(entities.NewProductQuantityEntity(db))
		productQuantityFacade.DeleteProductQuantity(ctx)
	})

	//PRODUCT HANDLERS

	r.GET("/product", func(ctx *gin.Context) {
		productFacade := facades.NewProductFacade(entities.NewProductEntity(db))
		productFacade.GetProduct(ctx)
	})
	r.GET("/product/:id", func(ctx *gin.Context) {
		productFacade := facades.NewProductFacade(entities.NewProductEntity(db))
		productFacade.GetProductByID(ctx)
	})
	r.POST("/product", func(ctx *gin.Context) {
		productFacade := facades.NewProductFacade(entities.NewProductEntity(db))
		productFacade.CreateProduct(ctx)
	})
	r.PUT("/product/:id", func(ctx *gin.Context) {
		productFacade := facades.NewProductFacade(entities.NewProductEntity(db))
		productFacade.UpdateProduct(ctx)
	})
	r.DELETE("/product/:id", func(ctx *gin.Context) {
		productFacade := facades.NewProductFacade(entities.NewProductEntity(db))
		productFacade.DeleteProduct(ctx)
	})

	//CART HANDLERS

	r.GET("/cart", func(ctx *gin.Context) {
		cartFacade := facades.NewCartFacade(entities.NewCartEntity(db), nil, nil, nil, nil, nil)
		cartFacade.GetCart(ctx)
	})
	r.GET("/cart/get-by-id/:id", func(ctx *gin.Context) {
		pqe := entities.NewProductQuantityEntity(db)
		pe := entities.NewProductEntity(db)
		ca := entities.NewCategoryEntity(db)
		cartFacade := facades.NewCartFacade(entities.NewCartEntity(db), nil, nil, &ca, &pe, &pqe)
		cartFacade.HandlerGetCartByID(ctx)
	})
	r.POST("/cart", func(ctx *gin.Context) {
		cartFacade := facades.NewCartFacade(entities.NewCartEntity(db), nil, nil, nil, nil, nil)
		cartFacade.CreateCart(ctx)
	})
	r.PUT("/cart/update-cart/:id", func(ctx *gin.Context) {
		cartFacade := facades.NewCartFacade(entities.NewCartEntity(db), nil, nil, nil, nil, nil)
		cartFacade.UpdateCart(ctx)
	})
	r.DELETE("/cart/delete-cart/:id", func(ctx *gin.Context) {
		cartFacade := facades.NewCartFacade(entities.NewCartEntity(db), nil, nil, nil, nil, nil)
		cartFacade.DeleteCart(ctx)
	})

	//CART Requested Endpoint

	r.POST("/cart/add-product", func(ctx *gin.Context) {
		pqe := entities.NewProductQuantityEntity(db)
		pe := entities.NewProductEntity(db)
		ca := entities.NewCategoryEntity(db)
		cartFacade := facades.NewCartFacade(entities.NewCartEntity(db), nil, nil, &ca, &pe, &pqe)
		cartFacade.HandlerApplyProduct(ctx)
	})

	r.PUT("/cart/apply-campaign/:cartid", func(ctx *gin.Context) {
		ce := entities.NewCampaignEntity(db)
		pqe := entities.NewProductQuantityEntity(db)
		pe := entities.NewProductEntity(db)
		ca := entities.NewCategoryEntity(db)
		cartFacade := facades.NewCartFacade(entities.NewCartEntity(db), &ce, nil, &ca, &pe, &pqe)
		cartFacade.HandlerApplyCampaign(ctx)
	})

	r.PUT("/cart/apply-discount", func(ctx *gin.Context) {
		pqe := entities.NewProductQuantityEntity(db)
		pe := entities.NewProductEntity(db)
		ca := entities.NewCategoryEntity(db)
		cartFacade := facades.NewCartFacade(entities.NewCartEntity(db), nil, nil, &ca, &pe, &pqe)
		cartFacade.HandlerApplyDiscount(ctx)
	})

	r.PUT("/cart/apply-coupon", func(ctx *gin.Context) {
		cu := entities.NewCouponEntity(db)
		pqe := entities.NewProductQuantityEntity(db)
		pe := entities.NewProductEntity(db)
		ca := entities.NewCategoryEntity(db)
		cartFacade := facades.NewCartFacade(entities.NewCartEntity(db), nil, &cu, &ca, &pe, &pqe)
		cartFacade.HandlerApplyCoupon(ctx)
	})

	r.POST("/cart/calculate-delivery-cost", func(ctx *gin.Context) {
		pqe := entities.NewProductQuantityEntity(db)
		pe := entities.NewProductEntity(db)
		ca := entities.NewCategoryEntity(db)
		cartFacade := facades.NewCartFacade(entities.NewCartEntity(db), nil, nil, &ca, &pe, &pqe)
		cartFacade.HandlerCalculateDeliveryCost(ctx)
	})

	r.GET("/cart/total-amount-after-discount/:cartid", func(ctx *gin.Context) {
		pqe := entities.NewProductQuantityEntity(db)
		pe := entities.NewProductEntity(db)
		ca := entities.NewCategoryEntity(db)
		cartFacade := facades.NewCartFacade(entities.NewCartEntity(db), nil, nil, &ca, &pe, &pqe)
		cartFacade.HandlerGetTotalAmountAfterDiscounts(ctx)
	})

	r.GET("/cart/coupon-discount/:cartid", func(ctx *gin.Context) {
		cu := entities.NewCouponEntity(db)
		pqe := entities.NewProductQuantityEntity(db)
		pe := entities.NewProductEntity(db)
		ca := entities.NewCategoryEntity(db)
		cartFacade := facades.NewCartFacade(entities.NewCartEntity(db), nil, &cu, &ca, &pe, &pqe)
		cartFacade.HandlerGetCouponDiscount(ctx)
	})

	r.GET("/cart/cart-print/:cartid", func(ctx *gin.Context) {
		pqe := entities.NewProductQuantityEntity(db)
		pe := entities.NewProductEntity(db)
		ca := entities.NewCategoryEntity(db)
		cartFacade := facades.NewCartFacade(entities.NewCartEntity(db), nil, nil, &ca, &pe, &pqe)
		cartFacade.HandlerCartPrint(ctx)
	})

	return r

}
