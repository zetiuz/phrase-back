package user

type Users struct {
	Username        string
	Profile_picture string
	Name            string
	Description     string
	Email           string
	Password        string
	Thematic_id     int64
}

func (Users) TableName() string {
	return "users"
}
