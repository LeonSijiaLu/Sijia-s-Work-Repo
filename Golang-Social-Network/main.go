package main

import (
	route "Golang-Social-Network/routes"
	"github.com/gin-gonic/gin"
	"github.com/urfave/negroni"
)

func main(){
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/signup", route.ToSignUp)
		user.POST("/login", route.ToLogin)
		user.POST("/logout", route.ToLogout)
	}

	api := router.Group("/api")
	{
		api.GET("/explore", route.Explore)
		api.GET("/explore_hashtag_posts/:hashtagname", route.ExploreHashtagPosts)

		api.POST("/post", route.CreatePost)
		api.DELETE("/post/:postID", route.DeletePost)
		api.PUT("/post/:postID", route.UpdatePost)
		api.POST("/like_post", route.LikePost)
		api.POST("/unlike_post", route.UnlikePost)

		api.POST("/comments/:postID", route.CreateComments)
		api.PUT("/comments/:postID/:commentNum", route.EditComments)

		api.POST("/follow_user", route.FollowUser)
		api.POST("/unfollow_user", route.UnFollowUser)
		api.POST("/blacklist_user", route.BlockUser)
		api.POST("/unblacklist_user", route.UnBlockUser)

		api.GET("/profile/:id", route.Profile)
		api.POST("/profile", route.EditProfile)

//		api.POST("/follow_topic", route.FollowTopic)
//		api.POST("/unfollow_topic", route.UnFollowTopic)
		api.POST("/follow_hashtag", route.FollowHashTag)
		api.POST("/unfollow_hashtag", route.UnFollowHashTag)
	}
	server := negroni.Classic()
	server.UseHandler(router)
	server.Run(":8882")
}