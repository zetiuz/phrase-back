package thematics

type Thematics struct {
	Id      string
	Name_en string
	Name_es string
}

func (Thematics) TableName() string {
	return "thematics"
}
