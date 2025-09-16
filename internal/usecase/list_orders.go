package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type ListOrderItemDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersOutputDTO []ListOrderItemDTO

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {
	orders, err := l.OrderRepository.FindAll()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}

	dto := ListOrdersOutputDTO{}

	for _, o := range orders {
		itemDTO := ListOrderItemDTO{
			ID:         o.ID,
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
		}

		dto = append(dto, itemDTO)
	}

	return dto, nil
}
