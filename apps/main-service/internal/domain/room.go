package domain

type Room struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RoomService interface {
	GetActiveRooms() ([]Room, error)
}

type RoomRepository interface {
	FetchActive() ([]Room, error)
}
