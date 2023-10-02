package user

import (
	"fmt"
	"phrase-back/src/db"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

type UserRequest struct {
	Username        string `json:"username"`
	Profile_picture string `json:"profile_picture"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Thematic_id     int64  `json:"thematic_id"`
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func VerifyPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func LoginCheck(email string, password string) (string, error) {

	var err error

	u := Users{}

	err = db.DB.Model(Users{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return "", err
	}
	fmt.Println(password)
	fmt.Println(u.Password)

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

	return tokenString, err
}

func Signin(c *gin.Context) {
	u := UserRequest{}

	if err := c.BindJSON(&u); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	token, err := LoginCheck(u.Email, u.Password)

	if err != nil {
		c.JSON(500, gin.H{"error": "email or password is incorrect."})
		return
	}

	c.JSON(200, gin.H{"token": token})
	c.Next()

}
func UserGetByEmail(c *gin.Context) {
	email := c.Param("email")
	var user Users
	db.DB.First(&user, "email", email)
	c.JSON(200, &user)
	return
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func UserCreate(c *gin.Context) {

	body := UserRequest{}

	c.BindJSON(&body)

	hashedPassword, err := HashPassword(body.Password)
	if err != nil {
		return
	}

	create := &Users{Username: body.Username, Profile_picture: body.Profile_picture, Name: body.Name, Description: body.Description, Email: body.Email, Password: hashedPassword, Thematic_id: body.Thematic_id}

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
