package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	ID   int
	Hash string
}

// ApiClient позволяет взаимодействовать с удаленным сервисом по сети.
type ApiClient struct{}

func (c *ApiClient) GetUser(userID int) User {
	// симуляция рандомной сетевой задержки
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

	userHash := sha256.Sum256([]byte(fmt.Sprintf("user-%d", userID)))
	hashStr := string(userHash[:])

	return User{
		ID:   userID,
		Hash: hashStr,
	}
}
