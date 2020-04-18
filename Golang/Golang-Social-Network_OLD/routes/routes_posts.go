package routes

import (
	UT "Golang-Social-Network/utils"
	"strings"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePost(c *gin.Context) {
	is_loggedin(c, "")
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
	is_loggedin(c, "")
	post_id := c.Param("postID")
	my_id, _ := UT.Get_Id_and_Username(c)
	if post_id == "" {
		panic("Please select a post to delete")
	}else{
		db := UT.Conn_DB()
		defer db.Close()
		var verifyPost int
		db.QueryRow("SELECT COUNT(post_id) FROM Posts WHERE post_id = ? AND created_by = ?", post_id, my_id).Scan(&verifyPost)
		if verifyPost != 1 {panic("Invalid Post Ownership, cannot delete")}
		_, err := db.Exec("DELETE FROM Posts WHERE post_id = ? AND created_by = ?", post_id, my_id)
		UT.Err(err)
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Deleted the post successfully",
			"success": true,
		})
	}
}

func UpdatePost(c *gin.Context){
	is_loggedin(c, "")
	post_id := c.Param("postID")
	title := strings.TrimSpace(c.PostForm("title"))
	content := strings.TrimSpace(c.PostForm("content"))
	allow_comments := c.PostForm("allow_comments")
	hashtags, mentions := extractTags_Mentions(content)
	id, _ := UT.Get_Id_and_Username(c)
	if post_id == "" || id == "" || title == "" || allow_comments == ""{
		panic("Wrong things happened before updating, double check your data")
	}else{
		db := UT.Conn_DB()
		defer db.Close()
		var verifyPost int
		db.QueryRow("SELECT COUNT(post_id) FROM Posts WHERE post_id = ? AND created_by = ?", post_id, id).Scan(&verifyPost)
		if verifyPost != 1 {panic("Invalid Post Ownership, cannot delete")}
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
		_, err := db.Exec("UPDATE Posts SET title = ?, content = ?, allow_comments = ? WHERE post_id = ? AND created_by = ?", title, content, allow_comments, post_id, id)
		UT.Err(err)
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Updated the post successfully",
			"success": true,
			"allow_comments": allow_comments,
			"title": title,
			"content": content,
		})
	}
}

func LikePost(c *gin.Context){
	is_loggedin(c, "")
	post_id := c.Param("postID")
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
	is_loggedin(c, "")
	post_id := c.Param("postID")
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

func EditProfile(c *gin.Context){
	is_loggedin(c, "")
	my_id, _ := UT.Get_Id_and_Username(c)
	job := strings.TrimSpace(c.PostForm("job"))
	quote := strings.TrimSpace(c.PostForm("quote"))
	allow_unfollowed_views := c.PostForm("allow_unfollowed_views")
	db := UT.Conn_DB()
	defer db.Close()
	if allow_unfollowed_views == ""{
		_, err := db.Exec("UPDATE Profile SET job = ?, quote = ? WHERE user_id = ?", job, quote, my_id)
		UT.Err(err)
	}else{
		_, err := db.Exec("UPDATE Profile SET job = ?, quote = ?, allow_unfollowed_views = ? WHERE user_id = ?", job, quote, allow_unfollowed_views, my_id)
		UT.Err(err)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Edited your profile successfully",
		"success": true,
	})
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
	is_loggedin(c, "")
	var (
		allow_comments bool
		posts_count int
	)
	post_id := c.Param("postID")
	content := strings.TrimSpace(c.PostForm("content"))
	if content == "" {panic("Comments content cannot be empty")}
	id, _ := UT.Get_Id_and_Username(c)
	db := UT.Conn_DB()
	defer db.Close()
	db.QueryRow("SELECT COUNT(*), allow_comments FROM Posts WHERE post_id = ?", post_id).Scan(&posts_count, &allow_comments)
	if posts_count != 1 {panic("Invalid post id")}
	if allow_comments == true{
		stmt, err := db.Prepare("INSERT INTO Comments (post_id, user_id, content) VALUES(?, ?, ?)")
		UT.Err(err)
		stmt.Exec(post_id, id, content)
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

func EditComments(c *gin.Context){
	is_loggedin(c, "")
	comment_id := c.Param("commentID")
	my_id, _ := UT.Get_Id_and_Username(c)
	if comment_id == "" || my_id == "" {panic("Invalid value")}
	db := UT.Conn_DB()
	defer db.Close()
	content := strings.TrimSpace(c.PostForm("content"))
	if content != ""{
		stmt, err := db.Prepare("UPDATE Comments SET content = ? WHERE AND comment_id = ? AND user_id = ?")
		UT.Err(err)
		stmt.Exec(content, comment_id, my_id)
	}else{
		panic("Comments cannot be empty")
	}
}

func LikeComments(c *gin.Context){
	is_loggedin(c, "")
	comment_id := c.Param("commentID")
	my_id, _ := UT.Get_Id_and_Username(c)
	if comment_id == "" || my_id == "" {panic("Invalid value")}
	db := UT.Conn_DB()
	defer db.Close()
	stmt, err := db.Prepare("UPDATE Comments SET likes = likes + 1 WHERE comment_id = ? AND user_id = ?")
	UT.Err(err)
	stmt.Exec(comment_id, my_id)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Liked the comment successfully",
		"success": true,
	})
}

func UnlikeComments(c *gin.Context){
	is_loggedin(c, "")
	comment_id := c.Param("commentID")
	my_id, _ := UT.Get_Id_and_Username(c)
	if comment_id == "" || my_id == "" {panic("Invalid value")}
	db := UT.Conn_DB()
	defer db.Close()
	stmt, err := db.Prepare("UPDATE Comments SET likes = likes - 1 WHERE comment_id = ? AND user_id = ?")
	UT.Err(err)
	stmt.Exec(comment_id, my_id)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Unliked the comment successfully",
		"success": true,
	})
}

func DeleteComments(c *gin.Context){
	is_loggedin(c, "")
	comment_id := c.Param("commentID")
	my_id, _ := UT.Get_Id_and_Username(c)
	var (
		post_count int
		post_id int
	)
	if comment_id == "" || my_id == "" {panic("Invalid value")}
	db := UT.Conn_DB()
	defer db.Close()
	db.QueryRow("SELECT COUNT(*), post_id FROM Comments WHERE comment_id = ? AND user_id = ?", comment_id, my_id).Scan(&post_count, &post_id)
	if post_count != 1 {panic("Incorrect comment ID")}
	db.Exec("DELETE FROM Comments WHERE post_id = ? AND comment_id = ? AND user_id = ?", post_id, comment_id, my_id)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully deleted comments",
		"success": true,
	})
}

func ShowComments(c *gin.Context, post_id interface{}) []interface{}{
	var (
		comment_id int
		user_id int
		content string
		likes int
	)
	comments := []interface{}{}
	db := UT.Conn_DB()
	defer db.Close()
	stmt, err := db.Prepare("SELECT comment_id, user_id, content, likes from Comments where post_id = ? ORDER BY likes, created_date DESC")
	UT.Err(err)
	rows, err := stmt.Query(post_id)
	UT.Err(err)
	for rows.Next(){
		rows.Scan(&comment_id, &user_id, &content, &likes)
		comment := map[string]interface{}{
			"comment_id": comment_id,
			"user_id": user_id,
			"post_id": post_id,
			"content": content,
			"likes": likes,
		}
		comments = append(comments, comment)
	}
	return comments
}

func Explore(c *gin.Context){  // only show posts of people who you follow
	is_loggedin(c, "")
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

func ExploreFriendsHashtagPosts(c *gin.Context){
	is_loggedin(c, "")
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
	hashtag_name := c.Param("hashtagname")
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

func ExploreAllHashtagPosts (c *gin.Context){
	is_loggedin(c, "")
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
	hashtag_name := c.Param("hashtagname")
	if hashtag_name == "" {panic("Please enter a hashtag name")}
	my_id, _ := UT.Get_Id_and_Username(c)
	db := UT.Conn_DB()
	defer db.Close()
	db.QueryRow("SELECT COUNT(hashtag_id), hashtag_id FROM Hashtags WHERE hashtag_name = ?", hashtag_name).Scan(&hashtag_count, &hashtag_id)
	if hashtag_count != 1 {panic("Invalid Hashtag name")}
	stmt, err := db.Prepare("SELECT Posts.post_id, Posts.likes, Posts.created_by, Posts.comments_num, Posts.title, Posts.content, Posts.allow_comments FROM Posts INNER JOIN Posts_Hashtags using (post_id) WHERE hashtag_id = ? AND Posts.created_by NOT IN (SELECT black_by FROM Blacklist WHERE black_to = ?) ORDER BY Posts.created_date DESC LIMIT 10")
	UT.Err(err)
	rows, err := stmt.Query(hashtag_id, my_id)
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