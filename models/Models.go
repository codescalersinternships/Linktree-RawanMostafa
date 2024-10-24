package models

type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Bio        string `json:"bio"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

type Link struct {
	UserID     string `json:"user_id"`
	Platform   string `json:"platform"`
	Url        string `json:"url"`
	ClickCount int    `json:"click_count"`
	VisitorID  string `json:"visitor_id"`
}
