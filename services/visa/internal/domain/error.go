package domain

import "errors"

var (
	ErrDecodingReqBody     = errors.New("error decoding request body")
	ErrApply               = errors.New("error occured while creating application")
	ErrHashingPassword     = errors.New("error while hashing password")
	ErrEncoding            = errors.New("error encoding json")
	ErrDecodingCartId      = errors.New("provided cartId is not a number or non-positive")
	ErrDecodingItemId      = errors.New("provided itemId is not a number or non-positive")
	ErrInvalidProductName  = errors.New("product name can not be blanck")
	ErrInvalidQuantity     = errors.New("quantiy can not be non-positive")
	ErrCartNotFound        = errors.New("there is no cart with such cartID")
	ErrNotFound            = errors.New("invalid cartID or itemID")
	ErrAddingItem          = errors.New("item wasn't added to cart")
	ErrUnauthorized        = errors.New("invalid phone number or password")
	ErrCreatingUser        = errors.New("error creating user")
	ErrPhoneNumberNotFound = errors.New("phone number not found")
	ErrLoadEnvVars         = errors.New("error loading .env variables")
)
