package main

import (
	route "Golang-Social-Network/routes"
	"github.com/gin-gonic/gin"
	"github.com/urfave/negroni"
	"net/http"
)

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("web/*.html")
	router.StaticFS("/assets", http.Dir("web/assets"))
	router.StaticFS("/users", http.Dir("web/users"))

	user := router.Group("/user")
	{
		user.POST("/signup", route.ToSignUp)
		user.POST("/login", route.ToLogin)
		user.POST("/logout", route.ToLogout)
	}

	router.GET("/basics", route.Basics)
	router.GET("/signup", route.Signup)
	router.GET("/login", route.Login)
	router.GET("/home", route.Home) // direct to page that shows the posts of followings order by created_date only
	router.GET("/upload", route.Upload) // direct to page that can upload posts
	router.GET("/exploring", route.Exploring) // direct to page that shows the most popular posts and images of all users (not limited to your followings)
	router.GET("/stories", route.Stories) // direct to page that 
	router.GET("/view_post", route.ViewPost) // direct to page that display your profile info
	router.GET("/view_users", route.ViewUserProfile) // direct to page that display your profile info
	router.GET("/view_profile", route.ViewProfile) // direct to page that display your profile info
	router.GET("/followers", route.Followers)
	router.GET("/followings", route.Followings)

	api := router.Group("/api")
	{
		api.GET("/explore", route.Explore)
		api.GET("/explore/hashtag/posts/all/:hashtagname", route.ExploreAllHashtagPosts)
		api.GET("/explore/hashtag/posts/following/:hashtagname", route.ExploreFriendsHashtagPosts)

		api.GET("/post/popular", route.ShowHottestPosts)
		api.POST("/post/add", route.CreatePost)
		api.POST("/post/delete/:postID", route.DeletePost)
		api.POST("/post/edit/:postID", route.UpdatePost)
		api.POST("/post/like/:postID", route.LikePost)
		api.POST("/post/unlike/:postID", route.UnlikePost)

		api.GET("/images", route.ShowImages)
		api.GET("/images/popular", route.GetHottestImages)

		api.POST("/comments/add/:postID", route.CreateComments)
		api.POST("/comments/edit/:commentID", route.EditComments)
		api.POST("/comments/like/:commentID", route.LikeComments)
		api.POST("/comments/unlike/:commentID", route.UnlikeComments)
		api.POST("/comments/delete/:commentID", route.DeleteComments)

		api.GET("/user/popular", route.ShowHottestUsers)
		api.GET("/user/followers", route.GetFollowers)
		api.GET("/user/followings", route.GetFollowings)
		api.GET("/user/hashtags", route.GetHashtags)
		api.GET("/user/followers/:userName", route.GetFollowers)
		api.GET("/user/followings/:userName", route.GetFollowings)
		api.GET("/user/hashtags/:userName", route.GetHashtags)
		api.POST("/user/follow/:userName", route.FollowUser)
		api.POST("/user/unfollow/:userName", route.UnFollowUser)
		api.POST("/user/blacklist/:userName", route.BlockUser)
		api.POST("/user/unblacklist/:userName", route.UnBlockUser)
		api.POST("/user/ID/:userName", route.GetUserID)

		api.GET("/profile/:id", route.Profile)
		api.POST("/profile", route.EditProfile)

//		api.POST("/follow_topic", route.FollowTopic)
//		api.POST("/unfollow_topic", route.UnFollowTopic)
		api.GET("/hashtag/popular", route.ShowHottestHashtags)
		api.POST("/hashtag/follow/:hashtagName", route.FollowHashTag)
		api.POST("/hashtag/unfollow/:hashtagName", route.UnFollowHashTag)
	}
	server := negroni.Classic()
	server.UseHandler(router)
	server.Run(":8882")
}