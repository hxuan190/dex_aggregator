package utils

import "errors"

var (
	ErrDivideByZero                   = errors.New("divide by zero")
	ErrTokenMaxExceeded               = errors.New("token amount exceeds u64 max")
	ErrTokenMinSubceeded              = errors.New("sqrt price below minimum")
	ErrSqrtPriceOutOfBounds           = errors.New("sqrt price out of bounds")
	ErrMultiplicationOverflow         = errors.New("multiplication overflow")
	ErrLiquidityUnderflow             = errors.New("liquidity underflow")
	ErrLiquidityOverflow              = errors.New("liquidity overflow")
	ErrInvalidSqrtPriceLimitDirection = errors.New("invalid sqrt price limit direction")
	ErrZeroTradableAmount             = errors.New("zero tradable amount")
	ErrPartialFillError               = errors.New("partial fill not allowed")
	ErrInvalidTickArraySequence       = errors.New("invalid tick array sequence")
	ErrInvalidAccountData             = errors.New("invalid account data")
)
