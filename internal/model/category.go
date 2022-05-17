package model

import (
	"context"
	"fmt"
	"stuff/pkg/gorm"

	log "github.com/sirupsen/logrus"
	originOrm "gorm.io/gorm"
)

// Category 商品种类，在数据库中固定几个：游戏、pc硬件等，只提供查询接口
type Category struct {
	ID           int    `json:"id" gorm:"column:id"`
	BannerURL    string `json:"banner_url" gorm:"column:banner_url"`
	FrontDesc    string `json:"front_desc"`
	FrontName    string `json:"front_name"`
	IconURL      string `json:"icon_url" gorm:"column:icon_url"`
	ImgURL       string `json:"img_url" gorm:"column:img_url"`
	IsShow       int    `json:"is_show"`
	Keywords     string `json:"keywords"`
	Level        string `json:"level"`
	Name         string `json:"name"`
	ParentID     int    `json:"parent_id" gorm:"column:parent_id"`
	ShowIndex    int    `json:"show_index"`
	SortOrder    int    `json:"sort_order"`
	Type         int    `json:"type"`
	WapBannerURL string `json:"wap_banner_url" gorm:"column:wap_banner_url"`
}

// GetCategoryList 传入name则模糊查询
func GetCategoryList(ctx context.Context, name string, limit, page int) ([]Category, int64, error) {
	cates := make([]Category, 0)
	session := gorm.DB.Session(&originOrm.Session{Context: ctx}).Where(1)

	if len(name) > 0 {
		nameQuery := fmt.Sprintf("%%%s%%", name)
		session = session.Where("name=?", nameQuery)
	}

	result := session.Limit(limit).Offset(limit * (page - 1)).Find(&cates)
	if result.Error != nil {
		log.Error("get category list error: ", result.Error)
		return nil, 0, result.Error
	}

	return cates, result.RowsAffected, nil
}
