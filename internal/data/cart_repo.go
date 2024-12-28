package data

import (
	"context"
	"errors"
	"gorm.io/gorm"

	"ShoppingCartService/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// Cart 定义嵌套的 Cart 结构体
// 一个用户对应一个购物车，购物车中有多个商品
type Cart struct {
	UserId string `gorm:"primaryKey;type:varchar(64);column:user_id;index:idx_user_item"` // 设置联合索引
	Items  []Item `gorm:"foreignKey:CartID"`                                              // 外键字段
}

// Item 定义嵌套的 Item 结构体
// CartID 为外键字段，关联 Cart 结构体的 UserId 字段
type Item struct {
	ItemId   string `gorm:"type:varchar(64);column:item_id;index:idx_user_item"` // 设置联合索引
	Quantity int32  `gorm:"column:quantity"`                                     // 个数
	CartID   string `gorm:"column:cart_id"`                                      // 外键字段
}

type cartRepo struct {
	data *Data
	log  *log.Helper
}

// NewCartRepo .
func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// FindCart 判断购物车是否存在
func (r *cartRepo) FindCart(ctx context.Context, userId string) (bool, error) {
	var cart biz.Cart
	if err := r.data.mysql.WithContext(ctx).Where("user_id = ?", userId).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *cartRepo) CreateCart(ctx context.Context, c *biz.Cart) error {
	return r.data.mysql.WithContext(ctx).Create(c).Error
}

func (r *cartRepo) GetCart(ctx context.Context, userId string) (*biz.Cart, error) {
	var cart biz.Cart
	if err := r.data.mysql.WithContext(ctx).Preload("Items").Where("user_id = ?", userId).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &cart, nil
}

func (r *cartRepo) SaveCart(ctx context.Context, c *biz.Cart) error {
	return r.data.mysql.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(c).Error; err != nil {
			return err
		}
		for _, item := range c.Items {
			item.CartID = c.UserId
			if err := tx.Save(&item).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *cartRepo) ClearCart(ctx context.Context, userId string) error {
	return r.data.mysql.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("cart_id = ?", userId).Delete(&biz.Item{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *cartRepo) DeleteCart(ctx context.Context, userId string) error {
	return r.data.mysql.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", userId).Delete(&biz.Cart{}).Error; err != nil {
			return err
		}
		if err := tx.Where("cart_id = ?", userId).Delete(&biz.Item{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *cartRepo) AddItem(ctx context.Context, userId string, item *biz.Item) error {
	item.CartID = userId
	return r.data.mysql.WithContext(ctx).Create(item).Error
}

func (r *cartRepo) DeleteItem(ctx context.Context, userId string, itemId string) error {
	return r.data.mysql.WithContext(ctx).Where("cart_id = ? AND item_id = ?", userId, itemId).Delete(&biz.Item{}).Error
}

func (r *cartRepo) UpdateItem(ctx context.Context, userId string, item *biz.Item) error {
	item.CartID = userId
	return r.data.mysql.WithContext(ctx).Where("cart_id = ? AND item_id = ?", userId, item.ItemId).Updates(item).Error
}
