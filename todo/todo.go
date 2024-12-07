package todo

import (
	"apidemo/auth"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Todo struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"text"`
	gorm.Model
}

func (Todo) TableName() string {
	return "todolist"
}

type TodoHandler struct {
	db *gorm.DB
}

func NewTodoHandler(db *gorm.DB) *TodoHandler {
	return &TodoHandler{db: db}
}

func (t *TodoHandler) NewTask(c *gin.Context) {

	s := c.Request.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(s, "Bearer ")

	if err := auth.Protect((tokenString)); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var todo Todo

	log.Println("New Task", &todo.Title)
	if err := c.ShouldBindJSON(&todo); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	r := t.db.Create(&todo)
	if err := r.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID": todo.Model.ID,
	})

}
