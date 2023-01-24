package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	gin "github.com/gin-gonic/gin"
	metacall "github.com/metacall/core/source/ports/go_port/source"
)

func remove_stopwords_py(text string) (string, error) {
	ret, err := metacall.Call("remove_stopwords", text)

	if err != nil {
		return "", err
	}

	if ret, ok := ret.(string); ok {
		return ret, nil
	} else {
		return "", errors.New("An error ocurred after executing the call when casting the result.")
	}
}

func main() {
	// Initialize MetaCall
	if err := metacall.Initialize(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Defer MetaCall destruction
	defer metacall.Destroy()

	scripts := []string{"nlp_script.py"}

	if err := metacall.LoadFromFile("py", scripts); err != nil {
		fmt.Println(err)
		return
	}

	r := gin.Default()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	r.GET("/remove_stopwords", func(c *gin.Context) {
		// result, err := DeployTransaction(30, 50)

		result, err := remove_stopwords_py("This is a dummy text to test the deployed model to remove stop words.")

		if err != nil {
			c.JSON(400, gin.H{
				"Error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"Deployment Status": result,
			})
		}
	})

	r.GET("/close", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
	})

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Listen:", err)
	}
}