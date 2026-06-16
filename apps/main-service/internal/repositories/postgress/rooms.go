package postgress

import "github.com/IvanGelium/main-service/internal/domain"

type roomRepository struct {
}

func NewRepository() domain.RoomRepository {
	return &roomRepository{}
}

func (r *roomRepository) FetchActive() ([]domain.Room, error) {
	return []domain.Room{
		{ID: 1, Name: "Комната симуляции огня"},
		{ID: 2, Name: "Мультиплеерный холст"},
	}, nil
}
