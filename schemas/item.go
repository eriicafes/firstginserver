package schemas

type CreateItem struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"desc"`
	Labels      []string `json:"labels"`
}
