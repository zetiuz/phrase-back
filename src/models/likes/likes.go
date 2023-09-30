package likes

type Likes struct {
	User_id string
	Post_id int64
}

func (Likes) TableName() string {
	return "likes"
}
