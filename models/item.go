package models

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"

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
	time := strconv.FormatInt(time.Now().Unix(), 10)

	c := md5.New()
	c.Write([]byte(time))

	return fmt.Sprintf("%x", c.Sum(nil))
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
