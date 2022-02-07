package main

import (
	"log"
	"nevad/component/appctx"
	"nevad/middleware"
	"nevad/module/profile/transport/ginprofile"
	"nevad/module/proxy/transport/ginproxy"
	"nevad/module/team/transport/ginteam"
	"nevad/module/user/transport/ginuser"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")

	secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	appContext := appctx.NewAppContext(db, secretKey)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(middleware.Recover(appContext))

	r.POST("/login", ginuser.Login(appContext))
	r.GET("/profile", middleware.RequiredAuth(appContext), ginuser.Profile(appContext))

	// User
	user := r.Group("/users", middleware.RequiredAuth(appContext))
	user.POST("", ginuser.Create(appContext))
	user.GET("", ginuser.ListUser(appContext))
	user.GET("/:id", ginuser.GetUser(appContext))
	user.DELETE("/:id", ginuser.DeleteUser(appContext))
	user.PUT("/:id", ginuser.UpdateUser(appContext))

	// Team
	team := r.Group("teams")
	team.POST("", ginteam.CreateTeam(appContext))
	team.GET("", ginteam.ListTeam(appContext))
	team.GET("/:id", ginteam.FindTeam(appContext))
	team.DELETE("/:id", ginteam.DeleteTeam(appContext))
	team.PATCH("/:id", ginteam.UpdateTeam(appContext))

	// Proxy
	proxy := r.Group("proxies")
	proxy.POST("", ginproxy.CreateProxy(appContext))
	proxy.GET("", ginproxy.ListProxy(appContext))
	proxy.GET("/:id", ginproxy.GetProxy(appContext))
	proxy.DELETE("/:id", ginproxy.DeleteProxy(appContext))
	proxy.PATCH("/:id", ginproxy.UpdateProxy(appContext))

	// Profile API Group
	profile := r.Group("profiles")
	profile.POST("", ginprofile.CreateProfile(appContext))
	profile.GET("", ginprofile.ListProfile(appContext))
	profile.GET("/:id", ginprofile.GetProfile(appContext))
	profile.DELETE("/:id", ginprofile.DeleteProfile(appContext))
	profile.PATCH("/:id", ginprofile.UpdateProfile(appContext))

	r.Run()

}
