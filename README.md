# TestRepo

## Quick Start

```bash
git clone https://github.com/ysBayram/TestRepo.git
cd ./TestRepo
go run main.go
```

## To Do
- [ ] Add unit tests
- [ ] Refactoring and review code

# Requirements:

All Cases should have full unit test coverage.

Follow OO Solid Principles and write easy to understand code.

Be Ready to add new requirements, do pair programming and explain your code.

Do not copy the code from google. Provide your own implementation.



1. Case - Trading Card Game

Problem Description

In this Kata you will be implementing a rudimentary two-player trading card game. The rules are loosely based on Blizzard Hearthstone (http://us.battle.net/hearthstone/en/ ) which itself is an already much simpler and straight-forward game compared to other TCGs like Magic: The Gathering ( http://www.wizards.com/magic/ ).

Preparation

Each player starts the game with 30 Health and 0 Mana slots

Each player starts with a deck of 20 Damage cards with the following Mana costs: 0,0,1,1,2,2,2,3,3,3,3,4,4,4,5,5,6,6,7,8

From the deck each player receives 3 random cards has his initial hand

Gameplay

The active player receives 1 Mana slot up to a maximum of 10 total slots

The active player’s empty Mana slots are refilled

The active player draws a random card from his deck

The active player can play as many cards as he can afford. Any played card empties Mana slots and deals immediate damage to the opponent player equal to its Mana cost.

If the opponent player’s Health drops to or below zero the active player wins the game

If the active player can’t (by either having no cards left in his hand or lacking sufficient Mana to pay for any hand card) or simply doesn’t want to play another card, the opponent player becomes active

Special Rules

Bleeding Out: If a player’s card deck is empty before the game is over he receives 1 damage instead of drawing a card when it’s his turn.

Overload: If a player draws a card that lets his hand size become >5 that card is discarded instead of being put into his hand.

Dud Card: The 0 Mana cards can be played for free but don’t do any damage either. They are just annoyingly taking up space in your hand.
2. Case - Shopping Cart

Implement an e-commerce shopping cart class.

Rules:

Product has a title and price.

Product belong to a category.

Category may or may not have a parent category.

Category has a title.

Products are added to the Cart with quantity info.

Campaigns exist for product price discounts.

Campaigns are applicable to a product category.

Campaign discount vary based on the number of products in the cart

Coupons exist for cart discounts.

Coupons have minimum cart amount constraint. If Cart amount is less than minimum, discount is not applied.

Delivery Cost is Dynamic. Based on the number of deliveries and number of products.

Cart

Implement a class that takes a cart and calculates the delivery cost.

//sample creating a new category

Category food = new new Category(“food");

 

//products

Product apple = new Product(“Apple”, 100.0, category);

Product almond = new Product(“Almonds”, 150.0, category);

 

//Products can be added to a shopping cart with quantity

ShoppingCart cart = new ShoppingCart();

cart.addItem(apple,3);

cart.addItem(almond,1);

 

Discounts

//you can apply discounts to a category

//discount rules can be 20% on a category if bought more than 3 items

Campaign campaign1 = new Campaign(category,20.0,3,DiscountType.Rate);

//another campaign rule 50% on a category if bought more than 5 items

Campaign campaign2 = new Campaign(category,50.0,5,DiscountType.Rate);

//another campaign rule 5 TL amount discount on a category if bought more than  items

Campaign campaign3 = new Campaign(category,5.0,5,DiscountType.Amount);

 

//Cart should apply the maximum amount of discount to the cart.

cart.applyDiscounts(discount1,discount2,discount3);

 

//You can also apply a coupon to a cart

//Coupons have minimum amount. If the cart total is less than minimum amount, coupon is not applicable

//Coupon for 100 TL min purchase amount for a 10% discount

Coupon coupon = new Coupon(100, 10, DiscountType.Rate)

cart.applyCoupon(coupon)

 

Campaign Discounts are applied first, Then Coupons.

Delivery

To optimize our delivery Cost, we want to apply dynamic cargo pricing rules based on the number of deliveries and number of products in the cart.

Formula is ( CostPerDelivery * NumberOfDeliveries ) + (CostPerProduct * NumberOfProducts) + Fixed Cost

Fixed cost is 2.99 TL.

Implement a class that takes a cart and calculates the delivery cost.

 

DeliveryCostCalculator deliveryCostCalculator = new DeliveryCostCalculator(costPerDelivery,costPerProduct,fixedCost);

double deliveryCostCalculator.calculateFor(cart)

 

NumberOfDeliveries is calculated by the number of distinct categories in the cart.

If cart has products that belong to two distinct categories, number of deliveries is 2.

 

NumberOfProducts is the number of different products in the cart. It is not the quantity of products.

Cart Responsibilities

double cart.getTotalAmountAfterDiscounts()

double cart.getCouponDiscount()

double cart.getCampaignDiscount()

double cart.getDeliveryCost()

 

Group Products by Category and Print the CategoryName, ProductName, Quantity, Unit Price, Total Price, Total Discount(coupon,campaign) applied

cart.print()

 

At the final line print total amount and the delivery cost.

You will be asked additional requirements in the interview and how to implement them.

Please Drive your code with Tests and follow Object Oriented Design Principles (SRP, OCP) 

 