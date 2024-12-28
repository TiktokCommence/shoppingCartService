package data

import (
	"ShoppingCartService/internal/conf"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Data struct {
	mysql *gorm.DB
}

// NewData .
func NewData(c *conf.Mysql, logger log.Logger) (*Data, func(), error) {
	db, err := NewMysql(c)
	if err != nil || db == nil {
		return nil, nil, err
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		mysql: db,
	}, cleanup, nil
}

func NewMysql(c *conf.Mysql) (*gorm.DB, error) {
	if c == nil || c.Database == nil || c.Database.Dsn == "" {
		return nil, fmt.Errorf("internal/data: invalid database configuration")
	}

	db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("internal/data: failed to open database: %w", err)
	}

	if err := db.AutoMigrate(&Cart{}, &Item{}); err != nil {
		return nil, fmt.Errorf("internal/data: failed to migrate models: %w", err)
	}

	return db, nil
}
