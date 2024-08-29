package apis

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/super_secret_app/database"
)

func GenericHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"esto": "functiona",
	})
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Birthday string `json:"birthday"`
}

type Error struct {
	Message string `json:"message"`
}

func NewError(msg string) Error {
	return Error{Message: msg}
}

func hash(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func InsertUser(c *gin.Context) {
	username := c.PostForm("username")
	password := hash(c.PostForm("password"))
	birthday := c.PostForm("birthday")
	cellphone := c.PostForm("cellphone")

	reg, _ := regexp.Compile(`^\d+$`)
	valid := reg.Match([]byte(cellphone))

	now := time.Now().UTC()
	birthdayParsed, err := time.Parse("2006-01-02", birthday)
	dateComp := now.Compare(birthdayParsed)

	if dateComp == -1 {
		c.HTML(http.StatusOK, "error.html", NewError("Fecha de nacimiento inválida"))
		return
	}

	if !valid {
		c.HTML(http.StatusOK, "error.html", NewError("Número de telefono inválido"))
		return
	}

	user := database.NewUser(username, password, birthday, cellphone)
	err = database.Insert(user)
	if err != nil {
		c.JSON(500, gin.H{
			"ocurrio un error al ingresar a la base de datos": err,
		})
	}
	c.HTML(http.StatusOK, "user.html", user)
}

func ViewLastUser(c *gin.Context) {
	user, err := database.SelectLastUser()
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.HTML(http.StatusOK, "view_user.html", user)
}
