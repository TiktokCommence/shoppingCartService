package biz

import (
	"context"

	v1 "ShoppingCartService/api/cart/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
	ErrUserExist    = errors.NotFound(v1.ErrorReason_ErrorReason_USER_EXIST.String(), "user exist")
)

type Item struct {
	ItemId   string
	Quantity int32
	CartID   string
}

// Cart is a Cart model.
type Cart struct {
	UserId string
	Items  []Item
}

// CartRepo is a Greater repo.
type CartRepo interface {
	GetCart(ctx context.Context, userId string) (*Cart, error)
	SaveCart(ctx context.Context, c *Cart) error
	DeleteCart(ctx context.Context, userId string) error
	ClearCart(ctx context.Context, userId string) error
	CreateCart(ctx context.Context, c *Cart) error
	FindCart(ctx context.Context, userId string) (bool, error)
	AddItem(ctx context.Context, userId string, item *Item) error
	DeleteItem(ctx context.Context, userId string, itemId string) error
	UpdateItem(ctx context.Context, userId string, item *Item) error
}

// CartUsecase is a Cart usecase.
type CartUsecase struct {
	repo CartRepo
	log  *log.Helper
}

// NewCartUsecase new a usecase.
func NewCartUsecase(repo CartRepo, logger log.Logger) *CartUsecase {
	return &CartUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CartUsecase) GetCart(ctx context.Context, userId string) (*Cart, error) {
	if ok, err := uc.repo.FindCart(ctx, userId); err != nil {
		return nil, err
	} else if !ok {
		return nil, ErrUserNotFound
	}
	return uc.repo.GetCart(ctx, userId)
}

func (uc *CartUsecase) ClearCart(ctx context.Context, userId string) error {
	if ok, err := uc.repo.FindCart(ctx, userId); err != nil {
		return err
	} else if !ok {
		return ErrUserNotFound
	}
	return uc.repo.ClearCart(ctx, userId)
}

func (uc *CartUsecase) AddItem(ctx context.Context, userId string, itemId string, quantity int32) error {
	if ok, err := uc.repo.FindCart(ctx, userId); err != nil {
		return err
	} else if !ok {
		return ErrUserNotFound
	}
	return uc.repo.AddItem(ctx, userId, &Item{ItemId: itemId, Quantity: quantity})
}

func (uc *CartUsecase) UpdateItem(ctx context.Context, userId string, itemId string, quantity int32) error {
	if ok, err := uc.repo.FindCart(ctx, userId); err != nil {
		return err
	} else if !ok {
		return ErrUserNotFound
	}
	return uc.repo.UpdateItem(ctx, userId, &Item{ItemId: itemId, Quantity: quantity})
}

func (uc *CartUsecase) DeleteItem(ctx context.Context, userId string, itemId string) error {
	if ok, err := uc.repo.FindCart(ctx, userId); err != nil {
		return err
	} else if !ok {
		return ErrUserNotFound
	}
	return uc.repo.DeleteItem(ctx, userId, itemId)
}

func (uc *CartUsecase) DeleteCart(ctx context.Context, userId string) error {
	if ok, err := uc.repo.FindCart(ctx, userId); err != nil {
		return err
	} else if !ok {
		return ErrUserNotFound
	}
	return uc.repo.DeleteCart(ctx, userId)
}

func (uc *CartUsecase) CreateCart(ctx context.Context, userId string) error {
	ok, _ := uc.repo.FindCart(ctx, userId)
	if ok {
		return ErrUserExist
	}
	return uc.repo.CreateCart(ctx, &Cart{UserId: userId})
}
