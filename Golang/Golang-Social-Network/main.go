package main

import (
	route "Golang-Social-Network/routes"
	"github.com/gin-gonic/gin"
	"github.com/urfave/negroni"
)

func main(){
	router := gin.Default()
	//router.LoadHTMLGlob("views/*.html")

	user := router.Group("/user")
	{
		user.POST("/signup", route.ToSignUp)
		user.POST("/login", route.ToLogin)
		user.POST("/logout", route.ToLogout)
	}

	//router.GET("/signup", route.Signup)

	api := router.Group("/api")
	{
		api.GET("/explore", route.Explore)
		api.GET("/explore/hashtag/posts/all/:hashtagname", route.ExploreAllHashtagPosts)
		api.GET("/explore/hashtag/posts/following/:hashtagname", route.ExploreFriendsHashtagPosts)

		api.POST("/post/add", route.CreatePost)
		api.POST("/post/delete/:postID", route.DeletePost)
		api.POST("/post/edit/:postID", route.UpdatePost)
		api.POST("/post/like/:postID", route.LikePost)
		api.POST("/post/unlike/:postID", route.UnlikePost)

		api.POST("/comments/add/:postID", route.CreateComments)
		api.POST("/comments/edit/:commentID", route.EditComments)
		api.POST("/comments/like/:commentID", route.LikeComments)
		api.POST("/comments/unlike/:commentID", route.UnlikeComments)
		api.POST("/comments/delete/:commentID", route.DeleteComments)

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

		api.GET("/profile/:id", route.Profile)
		api.POST("/profile", route.EditProfile)

//		api.POST("/follow_topic", route.FollowTopic)
//		api.POST("/unfollow_topic", route.UnFollowTopic)
		api.POST("/hashtag/follow/:hashtagName", route.FollowHashTag)
		api.POST("/hashtag/unfollow/:hashtagName", route.UnFollowHashTag)
	}
	server := negroni.Classic()
	server.UseHandler(router)
	server.Run(":8882")
}