package rest

import (
	"context"
	"crud-user/src/business/usecase"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type REST interface {
	Run()
}

var once = &sync.Once{}

type rest struct {
	http *gin.Engine
	uc   *usecase.Usecase
}

func Init(uc *usecase.Usecase) REST {
	r := &rest{}

	once.Do(func() {
		httpServ := gin.New()

		r = &rest{
			http: httpServ,
			uc:   uc,
		}

		r.http.Use(cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowHeaders:    []string{"*"},
			AllowMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
		}))

		r.http.Use(gin.Recovery())

		r.Register()
	})

	return r
}

func (r *rest) Run() {
	port := ":8080"

	server := &http.Server{
		Addr:    port,
		Handler: r.http,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Serving HTTP error: %s\n", err.Error())
		}
	}()
	fmt.Printf("Listening and Serving HTTP on %s\n", server.Addr)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
	}

	fmt.Println("Server exiting")
}

func (r *rest) Register() {
	publicApi := r.http.Group("/public")

	publicApi.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "hello world",
		})
	})

	user := r.http.Group("/user")
	user.POST("/", r.CreateUser)
	user.GET("/", r.GetAllUsers)
	user.PUT("/:user_id", r.UpdateUsers)
	user.DELETE("/:user_id", r.DeleteUsers)
}
