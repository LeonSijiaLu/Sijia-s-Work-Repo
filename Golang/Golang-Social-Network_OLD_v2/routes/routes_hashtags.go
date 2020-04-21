package routes

import(
	UT "Golang-Social-Network/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
	stmt, err := db.Prepare("SELECT hashtag_id, hashtag_name, followers_num, posts_num, DATE(created_date) FROM Hashtags ORDER BY posts_num DESC, followers_num DESC, created_date DESC LIMIT 10")
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