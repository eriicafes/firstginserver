package schemas

type CreateItem struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"desc"`
	Labels      []string `json:"labels"`
}

type UpdateItem struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"desc" binding:"required"`
	Labels      []string `json:"labels" binding:"required"`
	IsActive    bool     `json:"is_active" binding:"required"`
}
