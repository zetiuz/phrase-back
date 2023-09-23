package user

type Users struct {
	Id          string
	Username    string
	Name        string
	Description string
	Email       string
	Password    string
	Language_id int64
	Status      bool
}

func (Users) TableName() string {
	return "user"
}
