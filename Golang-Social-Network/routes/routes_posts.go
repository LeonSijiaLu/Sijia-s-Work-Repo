package routes

import (
	UT "Golang-Social-Network/utils"
	"strings"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePost(c *gin.Context) {
	title := strings.TrimSpace(c.PostForm("title"))
	content := strings.TrimSpace(c.PostForm("content"))
	hashtags, mentions := extractTags_Mentions(content)
	id, _ := UT.Get_Id_and_Username(c)

	if title == "" || content == ""{
		panic("Please enter title and content")
	}else{
		db := UT.Conn_DB()
		defer db.Close()
		stmt, _ := db.Prepare("INSERT INTO Posts(title, content, created_by) VALUES (?, ?, ?)")
		rs, err := stmt.Exec(title, content, id)
		UT.Err(err)
		new_postid, _ := rs.LastInsertId()
		if len(hashtags) != 0{
			for _, eachHashTag := range hashtags{
				Create_Follow_HashTag(new_postid, eachHashTag)
			}
		}
		if len(mentions) != 0{
			for _, eachMentionUser := range mentions{
				Create_Mention(new_postid, eachMentionUser)
			}
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Posts successfully created",
			"success": true,
			"postID": new_postid,
			"hastags": hashtags,
			"mentions": mentions,
		})
	}
}

func DeletePost(c *gin.Context){
	post_id := strings.TrimSpace(c.PostForm("post_id"))
	my_id, _ := UT.Get_Id_and_Username(c)
	if post_id == "" {
		panic("Please select a post to delete")
	}else{
		db := UT.Conn_DB()
		defer db.Close()
		_, err := db.Exec("DELETE FROM Posts WHERE post_id = ? AND created_by = ?", post_id, my_id)
		UT.Err(err)
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Deleted the post successfully",
			"success": true,
		})
	}
}

func UpdatePost(c *gin.Context){
	post_id := strings.TrimSpace(c.PostForm("post_id"))
	title := strings.TrimSpace(c.PostForm("title"))
	content := strings.TrimSpace(c.PostForm("content"))
	hashtags, mentions := extractTags_Mentions(content)
	id, _ := UT.Get_Id_and_Username(c)
	if post_id == "" || id == "" || title == ""{
		panic("Wrong things happened before updating, double check your data")
	}else{
		db := UT.Conn_DB()
		defer db.Close()
		if len(hashtags) != 0{
			for _, eachHashTag := range hashtags{
				Create_Follow_HashTag(post_id, eachHashTag)
			}
		}
		if len(mentions) != 0{
			for _, eachMentionUser := range mentions{
				Create_Mention(post_id, eachMentionUser)
			}
		}
		_, err := db.Exec("UPDATE Posts SET title = ?, content = ? WHERE post_id = ? AND created_by = ?", title, content, post_id, id)
		UT.Err(err)
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Updated the post successfully",
			"success": true,
		})
	}
}

func LikePost(c *gin.Context){
	post_id := strings.TrimSpace(c.PostForm("post_id"))
	id, _ := UT.Get_Id_and_Username(c)
	if post_id == "" {
		panic("Please select a post to like")
	}else{
		var likeNum int
		db := UT.Conn_DB()
		defer db.Close()
		_, err := db.Exec("INSERT INTO Likes(post_id, like_by) VALUES(?, ?)", post_id, id)
		UT.Err(err)
		db.QueryRow("SELECT likes FROM Posts WHERE post_id = ?", post_id).Scan(&likeNum)
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Liked the post successfully",
			"success": true,
			"likes": likeNum,
		})
	}
}

func UnlikePost(c *gin.Context){
	post_id := strings.TrimSpace(c.PostForm("post_id"))
	id, _ := UT.Get_Id_and_Username(c)
	if post_id == "" {
		panic("Please select a post to unlike")
	}else{
		var likeNum int
		db := UT.Conn_DB()
		defer db.Close()
		_, err := db.Exec("DELETE FROM Likes WHERE post_id = ? AND like_by = ?", post_id, id)
		UT.Err(err)
		db.QueryRow("SELECT likes FROM Posts WHERE post_id = ?", post_id).Scan(&likeNum)
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Unliked the post successfully",
			"success": true,
			"likes": likeNum,
		})
	}
}

func DisplayProfile(target_id interface{}, my_id interface{}, c *gin.Context) map[string]interface{}{
	if target_id == ""{panic("Invalid target user")}
	db := UT.Conn_DB()
	defer db.Close()

	var(
		postID int
		title string
		content string
		likes int
		createdBy int
		allow_comments bool
	)

	var(
		userCount int
		userID int
		username string
		email string
		job string
		quote string
		views int
	)

	var(
		follower_num int
		following_num int
	)

	posts := []interface{}{}

	if target_id != my_id{
		stmt, _ := db.Prepare("UPDATE Profile SET views = views + 1 WHERE user_id = ?")
		_, err := stmt.Exec(target_id)
		UT.Err(err)
	}

	db.QueryRow("SELECT COUNT(user_id), user_id, username, email FROM Users WHERE user_id = ?", target_id).Scan(&userCount, &userID, &username, &email)
	db.QueryRow("SELECT job, quote, views FROM Profile WHERE user_id = ?", target_id).Scan(&job, &quote, &views)
	if userCount != 1 {panic("Invalid target user")}
	user := map[string]interface{}{
		"user_id": userID,
		"username": username, 
		"email": email,
		"job": job,
		"quote": quote,
		"views": views,
	}
	//goTo404(c, userCount)

	stmt, err := db.Prepare("SELECT post_id, title, content, likes, created_by, allow_comments FROM Posts WHERE created_by = ? ORDER BY created_date DESC")
	UT.Err(err)
	rows, err := stmt.Query(target_id)
	UT.Err(err)
	for rows.Next(){
		rows.Scan(&postID, &title, &content, &likes, &createdBy, &allow_comments)
		if allow_comments == true{
			post := map[string]interface{}{
				"post_id": postID,
				"title": title,
				"content": content,
				"created_by": createdBy,
				"likes": likes,
				"comments": ShowComments(c, postID),
			}
			posts = append(posts, post)
		}else{
			post := map[string]interface{}{
				"post_id": postID,
				"title": title,
				"content": content,
				"created_by": createdBy,
				"likes": likes,
				"comments": allow_comments,
			}
			posts = append(posts, post)
		}
	}

	db.QueryRow("SELECT COUNT(*) FROM Follow WHERE follow_by = ?", target_id).Scan(&following_num)
	db.QueryRow("SELECT COUNT(*) FROM Follow WHERE follow_to = ?", target_id).Scan(&follower_num)
	return map[string]interface{}{
		"message": "Found user posts",
		"success": true,
		"user": user,
		"posts": posts,
		"followers":  follower_num,
		"followings": following_num,
	}
}

func Profile (c *gin.Context){
	is_loggedin(c, "")
	target_id := c.Param("id") // id is part of url
	my_id, _ := UT.Get_Id_and_Username(c)
	blocked := is_Blacked(my_id, target_id)
	if blocked{
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Oops, you are blocked by this user",
			"success": true,
		})
	}else{
		open_for_unfollowers := open_for_Unfollowers(target_id)
		if open_for_unfollowers == false{ // not open for unfollowers
			is_following := is_Following(my_id, target_id)
			if is_following == false{
				c.JSON(http.StatusOK, map[string]interface{}{
					"message": "This user only allows follower views, please follow first",
					"success": true,
				})
			}else{
				c.JSON(http.StatusOK, DisplayProfile(target_id, my_id, c))
			}
		}else{
			c.JSON(http.StatusOK, DisplayProfile(target_id, my_id, c))
		}
	}
}

func CreateComments(c *gin.Context){
	var (
		allow_comments bool
		comments_num int
		posts_count int
	)
	post_id := strings.TrimSpace(c.PostForm("post_id"))
	content := strings.TrimSpace(c.PostForm("content"))
	if content == "" {panic("Comments content cannot be empty")}
	id, _ := UT.Get_Id_and_Username(c)
	db := UT.Conn_DB()
	defer db.Close()
	db.QueryRow("SELECT COUNT(*), comments_num, allow_comments FROM Posts WHERE post_id = ?", post_id).Scan(&posts_count, &comments_num, &allow_comments)
	if posts_count != 1 {panic("Invalid post id")}
	if allow_comments == true{
		comments_num = comments_num + 1
		stmt, err := db.Prepare("INSERT INTO Comments (post_id, user_id, comment_num, content) VALUES(?, ?, ?, ?)")
		UT.Err(err)
		stmt.Exec(post_id, id, comments_num, content)
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Your comments has been uploaded",
			"success": true,
		})
	}else{
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "The post does not allow comments",
			"success": true,
		})
	}
}

func ShowComments(c *gin.Context, post_id interface{}) []interface{}{
	var (
		user_id int
		comment_num int
		content string
	)
	comments := []interface{}{}
	db := UT.Conn_DB()
	defer db.Close()
	stmt, err := db.Prepare("SELECT user_id, comment_num, content from Comments where post_id = ? ORDER BY comment_num DESC")
	UT.Err(err)
	rows, err := stmt.Query(post_id)
	UT.Err(err)
	for rows.Next(){
		rows.Scan(&user_id, &comment_num, &content)
		comment := map[string]interface{}{
			"user_id": user_id,
			"comment_num": comment_num,
			"content": content,
		}
		comments = append(comments, comment)
	}
	return comments
}

func Explore(c *gin.Context){  // only show posts of people who you follow
	var (
		post_id int
		likes int
		created_by int
		comments_num int
		title string
		content string
		allow_comments bool
	)
	my_id, _ := UT.Get_Id_and_Username(c)
	db := UT.Conn_DB()
	defer db.Close()
	stmt, err := db.Prepare("select post_id, likes, created_by, comments_num, title, content, allow_comments from Posts where created_by in (select follow_to from Follow where follow_by = ? AND follow_to NOT IN (select black_by from Blacklist where black_to = ?)) ORDER BY created_date DESC LIMIT 10")
	UT.Err(err)
	rows, err := stmt.Query(my_id, my_id)
	UT.Err(err)
	posts := []interface{}{}
	for rows.Next(){
		rows.Scan(&post_id, &likes, &created_by, &comments_num, &title, &content, &allow_comments)
		if allow_comments == true{
			post := map[string]interface{}{
				"post_id": post_id,
				"likes": likes,
				"created_by": created_by,
				"comments_num": comments_num,
				"title": title,
				"content": content,
				"comments": ShowComments(c, post_id),
			}
			posts = append(posts, post)
		}else{
			post := map[string]interface{}{
				"post_id": post_id,
				"likes": likes,
				"created_by": created_by,
				"comments_num": comments_num,
				"title": title,
				"content": content,
				"comments": false,
			}
			posts = append(posts, post)
		}
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "View your friends posts",
		"success": true,
		"posts": posts,
	})
}

func ExploreHashtagPosts(c *gin.Context){
	var (
		hashtag_id int
		hashtag_count int
		post_id int
		likes int
		created_by int
		comments_num int
		title string
		content string
		allow_comments bool
	)
	hashtag_name := strings.TrimSpace(c.PostForm("hashtag_name"))
	if hashtag_name == "" {panic("Please enter a hashtag name")}
	my_id, _ := UT.Get_Id_and_Username(c)
	db := UT.Conn_DB()
	defer db.Close()
	db.QueryRow("SELECT COUNT(hashtag_id), hashtag_id FROM Hashtags WHERE hashtag_name = ?", hashtag_name).Scan(&hashtag_count, &hashtag_id)
	if hashtag_count != 1 {panic("Invalid Hashtag name")}
	stmt, err := db.Prepare("SELECT Posts.post_id, Posts.likes, Posts.created_by, Posts.comments_num, Posts.title, Posts.content, Posts.allow_comments FROM Posts INNER JOIN Posts_Hashtags using (post_id) WHERE hashtag_id = ? AND Posts.created_by IN (SELECT follow_to FROM Follow WHERE follow_by = ? AND follow_to NOT IN (SELECT black_by FROM Blacklist WHERE black_to = ?)) ORDER BY Posts.created_date DESC LIMIT 10")
	UT.Err(err)
	rows, err := stmt.Query(hashtag_id, my_id, my_id)
	UT.Err(err)
	posts := []interface{}{}
	for rows.Next(){
		rows.Scan(&post_id, &likes, &created_by, &comments_num, &title, &content, &allow_comments)
		if allow_comments == true{
			post := map[string]interface{}{
				"post_id": post_id,
				"likes": likes,
				"created_by": created_by,
				"comments_num": comments_num,
				"title": title,
				"content": content,
				"comments": ShowComments(c, post_id),
			}
			posts = append(posts, post)
		}else{
			post := map[string]interface{}{
				"post_id": post_id,
				"likes": likes,
				"created_by": created_by,
				"comments_num": comments_num,
				"title": title,
				"content": content,
				"comments": false,
			}
			posts = append(posts, post)
		}
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "View posts of Hashtags",
		"success": true,
		"posts": posts,
	})
}