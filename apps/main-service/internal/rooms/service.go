package rooms

import "github.com/IvanGelium/main-service/internal/domain"

type service struct {
	repo domain.RoomRepository
}

func NewService(r domain.RoomRepository) domain.RoomService {
	return &service{repo: r}
}

func (s *service) GetActiveRooms() ([]domain.Room, error) {
	return s.repo.FetchActive()
}
