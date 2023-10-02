package post

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"phrase-back/src/db"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type PostRequest struct {
	Id          string    `json:"id"`
	User_id     string    `json:"user_id"`
	Image       string    `json:"imagen"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"createdAt"`
	Thematic_id int64     `json:"thematic_id"`
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(401, gin.H{"Error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func TokenValid(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Comprueba que el algoritmo de token sea el que esperas:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("my_secret_key"), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return err
	}
	return nil
}

func ExtractToken(r *http.Request) string {

	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func PostCreate(c *gin.Context) {
	body := PostRequest{}

	c.ShouldBindJSON(&body)

	create := &Posts{Id: body.Id, User_id: body.User_id, Image: body.Image, Body: body.Body, CreatedAt: time.Now(), Thematic_id: body.Thematic_id}

	result := db.DB.Create(&create)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &create)
}
func PostGet(c *gin.Context) {
	var post []Posts
	db.DB.Table("posts").Select("*").Scan(&post)
	c.JSON(200, &post)
	return
}
func PostGetByUsername(c *gin.Context) {
	Username := c.Param("user_id")
	var post []Posts
	db.DB.Raw("SELECT * FROM posts WHERE user_id = ?", Username).Scan(&post)
	c.JSON(200, &post)
	return
}
func PostGetByThematics(c *gin.Context) {
	Thematics := c.Param("thematic_id")
	var post []Posts
	db.DB.Where("thematic_id= ?", Thematics).Table("posts").Select("*").Find(&post)
	c.JSON(200, &post)
	return
}
func PostGetByWord(c *gin.Context) {
	word := c.Param("word")
	var post []Posts
	db.DB.Where("body LIKE ? ", "%"+word+"%").Table("posts").Select("*").Find(&post)
	c.JSON(200, &post)
	return
}
func PostUpdate(c *gin.Context) {

	id := c.Param("id")
	var post Posts
	db.DB.First(&post, "id", id)

	body := PostRequest{}
	c.BindJSON(&body)
	data := &Posts{Image: body.Image, Body: body.Body}

	result := db.DB.Model(&post).Where("id", id).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &post)
}
func PostDelete(c *gin.Context) {

	id := c.Param("id")
	var post Posts

	db.DB.Delete(&post, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}
