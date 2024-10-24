package models

type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Bio        string `json:"bio"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

type Visit struct {
	VisitorID string `json:"visitor_id"`
	VisitTime string `json:"visit_time"`
}

type Link struct {
	LinkID     string  `json:"link_id"`
	UserID     string  `json:"user_id"`
	Platform   string  `json:"platform"`
	Url        string  `json:"url"`
	ClickCount int     `json:"click_count"`
	Visits     []Visit `json:"visits"`
}
