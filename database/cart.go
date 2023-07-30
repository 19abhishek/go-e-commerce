package database

import "errors"

var (

	ErrCantFindProduct = errors.New("Can't find Product")
	ErrCantDecodeProducts = errors.New("Can't Decode Products")
	ErrUserIdIsNotValid = errors.New("User ID is not valid")
	ErrCantUpdateUser = errors.New("Can't Update User")
	ErrCantRemoveItemCart = errors.New("Can't Remove Item from cart")
	ErrCantGetItem = errors.New("Can't Get Item from the cart")
	ErrCantBuyCartItem = errors.New("Can't Buy Item from cart")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}