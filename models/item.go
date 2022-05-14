package models

import (
	"crypto/rand"
	"fmt"

	"github.com/eriicafes/fisrtginserver/schemas"
)

type Item struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"desc"`
	Labels      []string `json:"labels"`
	IsActive    bool     `json:"is_active"`
}

func generateId() string {
	bytes := make([]byte, 16)

	rand.Read(bytes)

	return fmt.Sprintf("%x", bytes)
}

func NewItem(input schemas.CreateItem) *Item {
	makeLabels := func() []string {
		if input.Labels != nil {
			return input.Labels
		}
		return []string{}
	}

	return &Item{
		Id:          generateId(),
		Name:        input.Name,
		Description: input.Description,
		Labels:      makeLabels(),
	}
}
