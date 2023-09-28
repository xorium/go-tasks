package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	ID   int
	Name string
}

// ApiClient позволяет взаимодействовать с удаленным сервисом по сети.
type ApiClient struct{}

func (c *ApiClient) GetUser(userID int) User {
	// симуляция рандомной сетевой задержки
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	return User{
		ID:   userID,
		Name: fmt.Sprintf("Name-%d", userID),
	}
}
