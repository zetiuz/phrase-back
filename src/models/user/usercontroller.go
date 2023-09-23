package user

import (
	"phrase-back/src/db"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

type UserRequest struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Language_id int64  `json:"language"`
	Status      bool   `json:"status"`
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {

	var err error

	u := Users{}

	err = db.DB.Model(Users{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	// Crear y firmar el token
	claims := &Claims{
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Signin(c *gin.Context) {
	u := UserRequest{}

	token, err := LoginCheck(u.Email, u.Password)

	if err != nil {
		c.JSON(500, gin.H{"error": "email or password is incorrect."})
		return
	}

	c.JSON(200, gin.H{"token": token})

}

func UserCreate(c *gin.Context) {

	body := UserRequest{}

	c.BindJSON(&body)

	create := &Users{Username: body.Username, Name: body.Name, Description: body.Description, Email: body.Email, Password: body.Password, Language_id: body.Language_id, Status: body.Status}

	result := db.DB.Create(&create)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &create)
}
func UserGet(c *gin.Context) {
	var user []Users
	db.DB.Table("user").Select("*").Scan(&user)
	c.JSON(200, &user)
	return
}

func UserGetByUser(c *gin.Context) {
	Username := c.Param("username")
	var user Users
	db.DB.First(&user, "username", Username)
	c.JSON(200, &user)
	return
}

func UserUpdate(c *gin.Context) {

	username := c.Param("username")
	var user Users
	db.DB.First(&user, "username", username)

	body := UserRequest{}
	c.BindJSON(&body)
	data := &Users{Name: body.Name, Description: body.Description, Email: body.Email, Password: body.Password}

	result := db.DB.Model(&user).Where("username", username).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &user)
}
