package database

import "errors"

var (
	ErrCantFindProduct        = errors.New("cant find the product")
	ErrCantDecodeProduct      = errors.New("cant find the product")
	ErrUserIdIsNotValid       = errors.New("this user is not valid")
	ErrCantUpdateUser         = errors.New("cant add this product to the cart")
	ErrCantRemoveItemFromCart = errors.New("cant remove this item from the cart")
	ErrCantGetItem            = errors.New("was unable to get the item from the cart")
	ErrCantBuyCartItem        = errors.New("cant update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
