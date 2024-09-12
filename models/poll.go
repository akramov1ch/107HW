package models

type Poll struct {
    ID      string   `json:"id" bson:"_id,omitempty"`
    Title   string   `json:"title" binding:"required"`
    Options []string `json:"options" binding:"required"`
    Votes   map[string]int `json:"votes" bson:"votes"`
}

type Vote struct {
    Option string `json:"option" binding:"required"`
}
