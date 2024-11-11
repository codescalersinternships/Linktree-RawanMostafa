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

type MsgResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LinkRequest struct {
	Url      string `json:"url"`
}

type LinkResponse struct {
	Message string `json:"message"`
	LinkID  string `json:"linkid"`
}
