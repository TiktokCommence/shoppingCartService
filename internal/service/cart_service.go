package service

import (
	"context"

	v1 "ShoppingCartService/api/cart/v1"
	"ShoppingCartService/internal/biz"
)

// CartService is a cart service.
type CartService struct {
	v1.UnimplementedCartServer
	uc *biz.CartUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.CartUsecase) *CartService {
	return &CartService{uc: uc}
}

func (s *CartService) GetCart(ctx context.Context, req *v1.GetCartRequest) (reply *v1.GetCartReply, err error) {
	cart, err := s.uc.GetCart(ctx, req.User.GetUserId())
	if err != nil {
		return nil, err
	}
	reply = &v1.GetCartReply{
		Commodity: make([]*v1.Commodity, 0),
	}
	for _, item := range cart.Items {
		commodity := &v1.Commodity{
			CommodityId:  item.ItemId,
			CommodityNum: item.Quantity,
		}
		reply.Commodity = append(reply.Commodity, commodity)
	}
	return reply, nil
}

func (s *CartService) ClearCart(ctx context.Context, req *v1.ClearCartRequest) (reply *v1.ClearCartReply, err error) {
	err = s.uc.ClearCart(ctx, req.User.GetUserId())
	if err != nil {
		return &v1.ClearCartReply{
			Message: "clear cart failed",
		}, err
	}
	return &v1.ClearCartReply{
		Message: "clear cart success",
	}, nil
}

func (s *CartService) AddItem(ctx context.Context, req *v1.AddItemRequest) (reply *v1.AddItemReply, err error) {
	err = s.uc.AddItem(ctx, req.User.GetUserId(), req.Commodity.GetCommodityId(), req.Commodity.GetCommodityNum())
	if err != nil {
		return &v1.AddItemReply{
			Message: "add item failed",
		}, err
	}
	reply = &v1.AddItemReply{
		Message: "add item success",
	}
	return reply, nil
}

func (s *CartService) UpdateItem(ctx context.Context, req *v1.UpdateItemRequest) (reply *v1.UpdateItemReply, err error) {
	err = s.uc.UpdateItem(ctx, req.User.GetUserId(), req.Commodity.GetCommodityId(), req.Commodity.GetCommodityNum())
	if err != nil {
		return &v1.UpdateItemReply{
			Message: "update item failed",
		}, err
	}
	reply = &v1.UpdateItemReply{
		Message: "update item success",
	}
	return reply, nil
}

func (s *CartService) DeleteItem(ctx context.Context, req *v1.DeleteItemRequest) (reply *v1.DeleteItemReply, err error) {
	err = s.uc.DeleteItem(ctx, req.User.GetUserId(), req.Commodity.GetCommodityId())
	if err != nil {
		return &v1.DeleteItemReply{
			Message: "delete item failed",
		}, err
	}
	reply = &v1.DeleteItemReply{
		Message: "delete item success",
	}
	return reply, nil
}

func (s *CartService) CreateCart(ctx context.Context, req *v1.CreateCartRequest) (reply *v1.CreateCartReply, err error) {
	err = s.uc.CreateCart(ctx, req.User.GetUserId())
	if err != nil {
		return &v1.CreateCartReply{
			Message: "create cart failed",
		}, err
	}
	reply = &v1.CreateCartReply{
		Message: "create cart success",
	}
	return reply, nil
}
