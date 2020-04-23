package routes

import(
	UT "Golang-Social-Network/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Check_HashTag_Exist(hashtag_name string) (int, bool){ // return hashtag id and if hashtag exists
	db := UT.Conn_DB()
	defer db.Close()
	var (
		hashtagCount int
		hashtagID int
	)
	db.QueryRow("SELECT COUNT(hashtag_id), hashtag_id FROM Hashtags WHERE hashtag_name = ?", hashtag_name).Scan(&hashtagCount, &hashtagID)
	if hashtagCount == 0{
		_, err := db.Exec("INSERT INTO Hashtags (hashtag_name) VALUES(?)", hashtag_name)
		UT.Err(err)
		db.QueryRow("SELECT COUNT(hashtag_id), hashtag_id FROM Hashtags WHERE hashtag_name = ?", hashtag_name).Scan(&hashtagCount, &hashtagID)
		if hashtagCount == 1 {
			return hashtagID, true
		}else{panic("Database Errors")}
	}else if hashtagCount == 1{
		return hashtagID, true
	}else{
		return 0, false
	}
}

func Create_Follow_HashTag(post_id interface{}, hashtag_name string) (int, bool){ // This function is called as part of the CreatePost function
	db := UT.Conn_DB()
	defer db.Close()
	var hashtagCount int
	if hashtag_name != ""{ // hashtag is present
		hashtag_id, hashtag_err := Check_HashTag_Exist(hashtag_name)
		if hashtag_err == false{panic("Database Errors")}else{
			db.QueryRow("SELECT COUNT(*) FROM Posts_Hashtags WHERE hashtag_id = ? AND post_id = ?", hashtag_id, post_id).Scan(&hashtagCount)
			if hashtagCount == 0{
				_, err := db.Exec("INSERT INTO Posts_Hashtags (hashtag_id, post_id) VALUES(?, ?)", hashtag_id, post_id)
				UT.Err(err)
				return hashtag_id, true
			}else{return hashtag_id, true}
		}
	}else{return 0, false}
}

func FollowHashTag(c *gin.Context) {
	is_loggedin(c, "")
	var (
		hashtag_id int
		hashtag_count int
	)
	db := UT.Conn_DB()
	defer db.Close()
	hashtag_name := c.Param("hashtagName")
	db.QueryRow("SELECT COUNT(hashtag_id), hashtag_id FROM Hashtags WHERE hashtag_name = ?", hashtag_name).Scan(&hashtag_count, &hashtag_id)
	if hashtag_count != 1 || hashtag_id == 0 {panic("Invalid hashtag name")}
	my_id, _ := UT.Get_Id_and_Username(c)
	if my_id == 0 {panic("Invalid user id")}
	
	_, err := db.Exec("INSERT INTO Users_Hashtags (user_id, hashtag_id) VALUES(?, ?)", my_id, hashtag_id)
	UT.Err(err)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Followed hashtag "+ hashtag_name +" successfully",
		"success": true,
	})
}

func UnFollowHashTag(c *gin.Context){
	is_loggedin(c, "")
	var (
		hashtag_id int
		hashtag_count int
	)
	db := UT.Conn_DB()
	defer db.Close()
	hashtag_name := c.Param("hashtagName")
	db.QueryRow("SELECT COUNT(hashtag_id), hashtag_id FROM Hashtags WHERE hashtag_name = ?", hashtag_name).Scan(&hashtag_count, &hashtag_id)
	if hashtag_count != 1 || hashtag_id == 0 {panic("Invalid hashtag name")}
	my_id, _ := UT.Get_Id_and_Username(c)
	if my_id == 0 {panic("Invalid user id")}
	_, err := db.Exec("DELETE FROM Users_Hashtags WHERE user_id = ? AND hashtag_id = ?", my_id, hashtag_id)
	UT.Err(err)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Unfollowed hashtag "+ hashtag_name +" successfully",
		"success": true,
	})

}

func ShowHottestHashtags(c *gin.Context){
	is_loggedin(c, "")
	var (
		hashtag_id int
		hashtag_name string
		followers_num int
		posts_num int
		created_date string
	)
	hottest_hashtags := []interface{}{}
	db := UT.Conn_DB()
	defer db.Close()
	stmt, err := db.Prepare("SELECT hashtag_id, hashtag_name, followers_num, posts_num, DATE(created_date) FROM Hashtags ORDER BY followers_num DESC, posts_num DESC, created_date DESC LIMIT 10")
	UT.Err(err)
	rows, err := stmt.Query()
	UT.Err(err)
	for rows.Next(){
		rows.Scan(&hashtag_id, &hashtag_name, &followers_num, &posts_num, &created_date)
		hashtag := map[string]interface{}{
			"hashtag_id": hashtag_id,
			"hashtag_name": hashtag_name,
			"followers_num": followers_num,
			"posts_num": posts_num,
			"created_date": created_date,
		}
		hottest_hashtags = append(hottest_hashtags, hashtag)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hottest Hashtags List",
		"success": true,
		"hashtags": hottest_hashtags,
	})
}

func GetFollowingHashtags(c *gin.Context){
	is_loggedin(c, "")
	my_id, _ := UT.Get_Id_and_Username(c)
	var (
		hashtag_id int
		hashtag_name string
		followers_num int
		posts_num int
	)
	following_hashtags := []interface{}{}
	db := UT.Conn_DB()
	defer db.Close()
	stmt, err := db.Prepare("SELECT Users_Hashtags.hashtag_id, Hashtags.hashtag_name, Hashtags.followers_num, Hashtags.posts_num FROM Users_Hashtags INNER JOIN Hashtags USING(hashtag_id) WHERE user_id = ? LIMIT 10")
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "DB Error","success": false,})}
	rows, err := stmt.Query(my_id)
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "DB Error","success": false,})}
	for rows.Next(){
		rows.Scan(&hashtag_id, &hashtag_name, &followers_num, &posts_num)
		hashtag := map[string]interface{}{
			"hashtag_id": hashtag_id,
			"hashtag_name": hashtag_name,
			"followers_num": followers_num,
			"posts_num": posts_num,
		}
		following_hashtags = append(following_hashtags, hashtag)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Following Hashtags List",
		"success": true,
		"hashtags": following_hashtags,
	})
}

func GetHashtagPosts(c *gin.Context){
	is_loggedin(c, "")
	hashtag_id := strings.TrimSpace(c.PostForm("hashtag_id"))
	hashtag_name := strings.TrimSpace(c.PostForm("hashtag_name"))

	posts := []interface{}{}
	var (
		post_id int
		user_id int
		user_name string
		title string
		content string
		likes int
		allow_comments bool
		comments_number int
		created_date string
	)

	db := UT.Conn_DB()
	defer db.Close()
	stmt, _ := db.Prepare("SELECT Users.username, Posts.created_by, Posts.post_id, Posts.title, Posts.content, Posts.likes, Posts.allow_comments, Posts.comments_num, DATE(Posts.created_date) FROM Posts_Hashtags INNER JOIN Posts USING (post_id) INNER JOIN Users ON Posts.created_by = Users.user_id WHERE Posts_Hashtags.hashtag_id = ?")
	rows, _ := stmt.Query(hashtag_id)
	for rows.Next(){
		rows.Scan(&user_name, &user_id, &post_id, &title, &content, &likes, &allow_comments, &comments_number, &created_date)
		if allow_comments == true{
			post := map[string]interface{}{
				"post_id": post_id,
				"user_id": user_id,
				"user_name": user_name,
				"title": title, 
				"content": content,
				"likes": likes,
				"allow_comments": allow_comments,
				"comments": ShowComments(c, post_id),
				"comments_num": comments_number,
				"images": ShowPostImages(c, post_id, user_id),
				"created_date": created_date,
			}
			posts = append(posts, post)
		}else{
			post := map[string]interface{}{
				"post_id": post_id,
				"user_id": user_id,
				"user_name": user_name,
				"title": title, 
				"content": content,
				"likes": likes,
				"allow_comments": allow_comments,
				"comments": allow_comments,
				"comments_num": 0,
				"images": ShowPostImages(c, post_id, user_id),
				"created_date": created_date,
			}
			posts = append(posts, post)
		}
	
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Retrieved posts of hashtag "+hashtag_name,
		"success": true,
		"hashtag_name": hashtag_name,
		"hashtag_id": hashtag_id,
		"posts": posts,
	})
}