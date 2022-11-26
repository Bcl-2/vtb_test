package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"main/service-1/models"
	"net/http"
)

const db = "postgres://postgres:serp2000@localhost:5432/postgres"

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true}))

	router.POST("/sendMessage", func(c *gin.Context) {
		var message models.Message
		err := c.ShouldBind(&message)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Number must be int type. Action must be string type",
			})
			return
		}

		if err != nil {
			log.Println("Error on:", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": message.Action,
		})
		err = executeQuery(db, message.Number, message.Action)
		if err != nil {
			log.Println(err)
		}
	})
	err := router.Run()
	if err != nil {
		return
	}
}

func executeQuery(host string, num int, action string) error {
	ctx := context.Background()
	cursor, err := pgx.Connect(ctx, host)
	defer cursor.Close(ctx)
	err = cursor.Ping(context.Background())
	if err != nil {
		return err
	}
	_, err = cursor.Exec(ctx, `INSERT INTO test(row_num,action) values($1,$2)`, num, action)
	if err != nil {
		return err
	}
	return nil
}
