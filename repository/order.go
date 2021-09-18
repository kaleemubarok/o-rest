package repository

import (
	"github.com/jinzhu/gorm"
	"o-rest/entity"
)
type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

type OrderRepository interface {
	CreateOrder(oder entity.Order) (entity.Order, error)
	GetOrder(oder entity.Order) (entity.Order, error)
	GetOrders() ([]entity.Order, error)
	UpdateOrder(oder entity.Order) (entity.Order, error)
	DeleteOrder(oder entity.Order) error
}

func (r *orderRepository) CreateOrder(oder entity.Order) (entity.Order, error) {
	err := r.db.Create(&oder).Error
	if err != nil {
		return oder, err
	}

	return oder, nil
}

func (r *orderRepository) GetOrder(order entity.Order) (entity.Order, error) {
	var result entity.Order
	err := r.db.Find(&result,"order_id = ?",order.OrderID).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *orderRepository) GetOrders() ([]entity.Order, error) {
	var ords []entity.Order
	err := r.db.Find(&ords).Error
	if err != nil {
		return ords, err
	}

	return ords, nil
}

func (r *orderRepository) UpdateOrder(order entity.Order) (entity.Order, error) {
	ords := entity.Order{}
	err := r.db.Model(&ords).Update(order).Where("order_id = ?",order.OrderID).Error
	if err != nil {
		return ords, err
	}

	return ords, nil
}

func (r *orderRepository) DeleteOrder(oder entity.Order) error {
	var ords entity.Order
	err := r.db.Delete(&ords, "order_id = ?",oder.OrderID).Error
	if err != nil {
		return err
	}

	return nil
}