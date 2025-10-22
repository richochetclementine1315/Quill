package models

type Blog struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Image  string `json:"image"`
	UserID uint   `json:"user_id"`
	// creating relationship between blog and user by foreign key
	User User `json:"user" gorm:"foreignkey:UserID"`
}
