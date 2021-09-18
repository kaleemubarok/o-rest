package service

import (
	"o-rest/entity"
	"o-rest/repository"
)

type service struct {
	rOrder repository.OrderRepository
	rItem  repository.ItemRepository
}

func NewService(rOrder repository.OrderRepository, rItem repository.ItemRepository) *service {
	return &service{rOrder, rItem}
}

type OrderService interface {
	CreateOrder(req entity.OrderRqResponse) (entity.OrderRqResponse, error)
	GetOrders() ([]entity.OrderRqResponse, error)
	UpdateOrder(req entity.OrderRqResponse) (entity.OrderRqResponse, error)
	DeleteOrder(req entity.OrderRqResponse) error
}

func (s *service) CreateOrder(req entity.OrderRqResponse) (entity.OrderRqResponse, error) {
	order := entity.Order{
		CustomerName: req.CustomerName,
		OrderAt:      req.OrderAt,
	}
	order, err := s.rOrder.CreateOrder(order)
	if err != nil {
		return entity.OrderRqResponse{}, err
	}

	for _, v := range req.Items {
		item := entity.Item{
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Quantity:    v.Quantity,
			OrderID:     order.OrderID,
		}
		item, err = s.rItem.CreateItem(item)
		if err != nil {
			return entity.OrderRqResponse{}, err
		}
	}

	return entity.OrderRqResponse{
		OrderAt:      order.OrderAt,
		CustomerName: order.CustomerName,
		Items:        req.Items,
	}, nil
}

func (s *service) GetOrders() ([]entity.OrderRqResponse, error) {
	var orders []entity.Order
	orders, err := s.rOrder.GetOrders()
	if err != nil {
		return []entity.OrderRqResponse{}, err
	}

	var opts []entity.OrderRqResponse
	for _, v := range orders {
		sItem := entity.Item{
			OrderID: v.OrderID,
		}
		items, err := s.rItem.GetItems(sItem)
		if err != nil {
			return []entity.OrderRqResponse{}, err
		}

		opts = append(opts, entity.OrderRqResponse{
			OrderID:      v.OrderID,
			OrderAt:      v.OrderAt,
			CustomerName: v.CustomerName,
			Items:        items,
		})

	}

	return opts, err
}

func (s *service) UpdateOrder(req entity.OrderRqResponse) (entity.OrderRqResponse, error) {
	order := entity.Order{
		OrderID:      req.OrderID,
		CustomerName: req.CustomerName,
		OrderAt:      req.OrderAt,
	}

	order, err := s.rOrder.UpdateOrder(order)
	if err != nil {
		return entity.OrderRqResponse{}, err
	}

	for _, v := range req.Items {
		item := entity.Item{
			ItemID:      v.ItemID,
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Quantity:    v.Quantity,
			OrderID:     order.OrderID,
		}

		item, err = s.rItem.UpdateItem(item)
		if err != nil {
			return entity.OrderRqResponse{}, err
		}
	}

	return entity.OrderRqResponse{
		OrderID:      req.OrderID,
		OrderAt:      order.OrderAt,
		CustomerName: order.CustomerName,
		Items:        req.Items,
	}, nil
}

func (s *service) DeleteOrder(req entity.OrderRqResponse) error {
	order := entity.Order{OrderID: req.OrderID}
	err := s.rOrder.DeleteOrder(order)
	if err != nil {
		return err
	}

	item := entity.Item{OrderID: req.OrderID}
	err = s.rItem.DeleteItem(item)
	if err != nil {
		return err
	}

	return nil
}
