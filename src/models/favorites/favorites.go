package favorites

type Favorites struct {
	User_id string
	Post_id int64
}

func (Favorites) TableName() string {
	return "favorites"
}
