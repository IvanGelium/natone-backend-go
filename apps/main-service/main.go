package main

import (
	"fmt"
	"net/http"

	"github.com/IvanGelium/main-service/internal/rooms"
)

func main() {
	roomRepo := rooms.NewRepository()
	roomSvc := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomSvc)

	http.HandleFunc("/rooms", roomHandler.GetRooms)

	fmt.Println("Бэкенд по Чистой Архитектуре запущен на http://localhost:5001")
	if err := http.ListenAndServe(":5001", nil); err != nil {
		panic(err)
	}
}
