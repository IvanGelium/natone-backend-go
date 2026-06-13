package rooms

import "github.com/IvanGelium/main-service/internal/domain"

type repository struct {
}

func NewRepository() domain.RoomRepository {
	return &repository{}
}

func (r *repository) FetchActive() ([]domain.Room, error) {
	return []domain.Room{
		{ID: 1, Name: "Комната симуляции огня"},
		{ID: 2, Name: "Мультиплеерный холст"},
	}, nil
}
