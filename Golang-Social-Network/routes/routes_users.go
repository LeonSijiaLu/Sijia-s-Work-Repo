package routes

import (
	UT "Golang-Social-Network/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"strconv"
)

func ToLogout(c *gin.Context){
	is_loggedin(c, "")
	session := UT.GetSession(c)
	delete(session.Values, "id")
	delete(session.Values, "username")
	session.Save(c.Request, c.Writer)
	c.Redirect(http.StatusFound, "/login")
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "You have logged out",
		"success": true,
	})
}

func ToSignUp(c *gin.Context){
	username := strings.TrimSpace(c.PostForm("username"))
	password := strings.TrimSpace(c.PostForm("password"))
	password_repeated := strings.TrimSpace(c.PostForm("password_repeated"))
	email := strings.TrimSpace(c.PostForm("email"))

	if username == "" || password == "" || password_repeated == "" || email == "" {
		panic("You forgot some values")
	}else if len(username) < 3 || len(username) > 32{
		panic("Username length needs to be between 3 and 32")
	}else if checkmail.ValidateFormat(email) != nil{
		panic("Incorrect email")
	}else if password != password_repeated{
		panic("Passwords need to match")
	}else{
		db := UT.Conn_DB()
		defer db.Close()
		rs, err := db.Exec("INSERT INTO Users(username, password, email) VALUES (?, ?, ?)", username, hash(password), email)
		UT.Err(err)
		user_id, _ := rs.LastInsertId()
		session := UT.GetSession(c)
		session.Values["id"] = strconv.FormatInt(user_id,10)
		session.Values["username"] = username
		session.Save(c.Request, c.Writer) 
		c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"message": "Welcome, " + username + " !!",
		})
	}
}

func ToLogin(c *gin.Context){
	login_username := strings.TrimSpace(c.PostForm("username"))
	login_password := strings.TrimSpace(c.PostForm("password"))
	if login_username == "" || login_password == ""{
		panic("Please enter username and password")
	}else{
		var id int
		var count_id int
		var username string
		var password string
		db := UT.Conn_DB()
		defer db.Close()
		db.QueryRow("SELECT COUNT(user_id), user_id, username, password FROM Users WHERE username = ?", login_username).Scan(&count_id, &id, &username, &password)
		if count_id != 1{
			panic("Incorrect username or password")
		}else{
			err := bcrypt.CompareHashAndPassword([]byte(password), []byte(login_password)) // check if hashed password match
			if err != nil{
				panic("Incorrect password")
			}else{
				session := UT.GetSession(c)
				session.Values["id"] = strconv.FormatInt(int64(id), 10)
				session.Values["username"] = username
				session.Save(c.Request, c.Writer)
				c.JSON(http.StatusOK, map[string]interface{}{
					"success": true,
					"message": username + ", you have logged in",
				})
 			}
		}
	}
}

func FollowUser(c *gin.Context){
	var (
		target_id int
	)
	db := UT.Conn_DB()
	defer db.Close()
	db.QueryRow("SELECT user_id FROM Users WHERE username = ?", strings.TrimSpace(c.PostForm("username"))).Scan(&target_id)
	if target_id == 0 {panic("Invalid username")}
	my_id, _ := UT.Get_Id_and_Username(c)
	if my_id == 0 {panic("Invalid user id")}
	if my_id != target_id{
		_, err := db.Exec("INSERT INTO Follow (follow_by, follow_to) VALUES(?, ?)", my_id, target_id)
		UT.Err(err)
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Followed user successfully",
			"success": true,
		})
	}else{panic("You cannot follow yourself")}
}

func UnFollowUser(c *gin.Context){
	var (
		target_id int
	)
	db := UT.Conn_DB()
	defer db.Close()
	db.QueryRow("SELECT user_id FROM Users WHERE username = ?", strings.TrimSpace(c.PostForm("username"))).Scan(&target_id)
	if target_id == 0 {panic("Invalid username")}
	my_id, _ := UT.Get_Id_and_Username(c)
	if my_id == 0 {panic("Invalid user id")}
	if my_id != target_id{
		_, err := db.Exec("DELETE FROM Follow WHERE follow_by = ? AND follow_to = ?", my_id, target_id)
		UT.Err(err)
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Unfollowed user successfully",
			"success": true,
		})
	}else{panic("You cannot unfollow yourself")}
}

func BlockUser(c *gin.Context){
	var (
		target_id int
	)
	db := UT.Conn_DB()
	defer db.Close()
	username := strings.TrimSpace(c.PostForm("username"))
	db.QueryRow("SELECT user_id FROM Users WHERE username = ?", username).Scan(&target_id)
	if target_id == 0 {panic("Invalid username")}
	my_id, _ := UT.Get_Id_and_Username(c)
	if my_id == 0 {panic("Invalid user id")}
	if my_id != target_id{
		_, err := db.Exec("INSERT INTO Blacklist (black_by, black_to) VALUES(?, ?)", my_id, target_id)
		UT.Err(err)
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Blocked "+username+" successfully",
			"success": true,
		})
	}else{panic("You cannot block yourself")}
}

func UnBlockUser(c *gin.Context){
	var (
		target_id int
	)
	db := UT.Conn_DB()
	defer db.Close()
	username := strings.TrimSpace(c.PostForm("username"))
	db.QueryRow("SELECT user_id FROM Users WHERE username = ?", username).Scan(&target_id)
	if target_id == 0 {panic("Invalid username")}
	my_id, _ := UT.Get_Id_and_Username(c)
	if my_id == 0 {panic("Invalid user id")}
	if my_id != target_id{
		_, err := db.Exec("DELETE FROM Blacklist WHERE black_by = ? AND black_to = ?", my_id, target_id)
		UT.Err(err)
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Unblocked "+username+" successfully",
			"success": true,
		})
	}else{panic("You cannot unblock yourself")}	
}