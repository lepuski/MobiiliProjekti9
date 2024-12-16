package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"mobiiliprojekti/internal/models"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// create an application struct for dependency injection
type application struct {
	users  *models.UserModel
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()


	app := &application{
		logger: logger,
		users:  &models.UserModel{DB: db},
	}

	r := gin.Default()

	store := cookie.NewStore([]byte("afuckinsecret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/register", app.registerGet)
	r.POST("/register", app.registerPost)
	r.GET("/login", app.loginGet)
	r.POST("/login", app.loginPost)
	r.POST("/logout", app.logout)

	auth := r.Group("/")

	auth.Use(app.authMiddleware())
	{
	auth.POST("/favorite-team", app.addFavoriteTeam)
	auth.DELETE("/favorite-team", app.removeFavoriteTeam)
	auth.GET("/favorite-team", app.getFavoriteTeam) 
	}


	r.Run(*addr) // listen and serve on port defined in addr

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
