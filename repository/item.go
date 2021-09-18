package repository

import (
	"github.com/jinzhu/gorm"
	"o-rest/entity"
)

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{db}
}

type ItemRepository interface {
	CreateItem(item entity.Item) (entity.Item, error)
	GetItems(item entity.Item) ([]entity.ItemRqResponse, error)
	UpdateItem(item entity.Item) (entity.Item, error)
	DeleteItem(item entity.Item) error
}

func (r *itemRepository) CreateItem(item entity.Item) (entity.Item, error) {
	err := r.db.Create(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *itemRepository) GetItems(item entity.Item) ([]entity.ItemRqResponse, error) {
	var response []entity.Item
	var result []entity.ItemRqResponse

	err := r.db.Find(&response, "order_id = ?", item.OrderID).Error
	if err != nil {
		return result, err
	}

	for _, v := range response {
		data := entity.ItemRqResponse{
			ItemID:      v.ItemID,
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Quantity:    v.Quantity,
		}
		result = append(result, data)
	}
	return result, nil
}

func (r *itemRepository) UpdateItem(item entity.Item) (entity.Item, error) {
	var result entity.Item
	err := r.db.Model(&result).Update(item).Where("item_id = ?", item.ItemID).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *itemRepository) DeleteItem(item entity.Item) error {
	var itm entity.Item
	err := r.db.Delete(&itm, "order_id = ?", item.OrderID).Error
	if err != nil {
		return err
	}

	return nil
}
