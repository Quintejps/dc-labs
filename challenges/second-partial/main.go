package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var activeUsers = make(map[string]string)

func loginVerification(user string) bool {
	for _, value := range activeUsers {
		if user == value {
			return false
		}
	}
	return true
}

func login(c *gin.Context) {

	user := c.MustGet(gin.AuthUserKey).(string)
	temp := loginVerification(user)

	if !temp {
		c.JSON(http.StatusOK, gin.H{
			"message": "You already in " + user,
		})
	} else {
		token := (xid.New()).String()
		activeUsers[token] = user

		c.JSON(http.StatusOK, gin.H{
			"message": "Hi " + user + " welcome to the DPIP System",
			"token":   token,
		})
	}

}

func logout(c *gin.Context) {
	token := strings.Trim(strings.TrimLeft(c.GetHeader("authorization"), "Bearer"), " ")

	for key, value := range activeUsers {
		if token == key {
			user := value
			delete(activeUsers, key)
			c.JSON(http.StatusOK, gin.H{
				"message": "Bye " + user + ", your token has been revoked",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Token " + token + " does not exist",
	})
}

func status(c *gin.Context) {
	token := strings.Trim(strings.TrimLeft(c.GetHeader("authorization"), "Bearer"), " ")

	for key, value := range activeUsers {
		if token == key {
			user := value
			c.JSON(http.StatusOK, gin.H{
				"message": "Hi " + user + ", the DPIP System is Up and Running",
				"time":    time.Now().UTC().String(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Token " + token + " does not exist",
	})
}

func upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("data")
	if err != nil {
		log.Fatal(err)
	}
	token := strings.Trim(strings.TrimLeft(c.GetHeader("authorization"), "Bearer"), " ")

	for key := range activeUsers {
		if token == key {
			fileName := header.Filename
			fileSize := header.Size

			out, err := os.Create("C:\\Users\\quint\\AppData\\Local\\Temp\\" + fileName)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()
			_, err = io.Copy(out, file)
			if err != nil {
				log.Fatal(err)
			}
			fileSize = fileSize / 1000
			str := strconv.FormatInt(fileSize, 10)
			c.JSON(http.StatusOK, gin.H{
				"message":  "An image has been successfully uploaded",
				"filename": fileName,
				"size":     str + "kb",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Token " + token + " does not exist",
	})
}

func main() {

	r := gin.Default()

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"user1": "pass1",
		"user2": "pass2",
		"user3": "pass3",
	}))

	authorized.GET("/login", login)
	r.GET("/logout", logout)
	r.GET("/status", status)
	r.POST("/upload", upload)

	r.Run() // listen and serve on 0.0.0.0:8080
}

