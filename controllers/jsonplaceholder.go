package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanhaiya15/gopf/lib/jsend"
	"github.com/kanhaiya15/gopf/services/apiclient"
)

// GetPosts posts
func GetPosts(c *gin.Context) {
	resultjson, err := apiclient.GetService("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		jsend.Wrap(c.Writer).
			Status(http.StatusInternalServerError).
			Message(http.StatusText(http.StatusInternalServerError)).
			Send()
		return
	}
	jsend.Wrap(c.Writer).
		Status(http.StatusOK).
		Data(resultjson).
		Message(http.StatusText(http.StatusOK)).
		Send()
}

// AddPost AddPost
func AddPost(c *gin.Context) {
	payload := map[string]interface{}{
		"title":  "foo",
		"body":   "bar",
		"userId": 1,
	}
	resultjson, err := apiclient.PostService("https://jsonplaceholder.typicode.com/posts", payload)
	if err != nil {
		jsend.Wrap(c.Writer).
			Status(http.StatusInternalServerError).
			Message(http.StatusText(http.StatusInternalServerError)).
			Send()
		return
	}
	jsend.Wrap(c.Writer).
		Status(http.StatusOK).
		Data(resultjson).
		Message(http.StatusText(http.StatusOK)).
		Send()
}

// UpdatePost UpdatePost
func UpdatePost(c *gin.Context) {
	payload := map[string]interface{}{
		"title":  "fooss",
		"body":   "barss",
		"userId": 1,
	}
	resultjson, err := apiclient.PutService("https://jsonplaceholder.typicode.com/posts/1", payload)
	if err != nil {
		jsend.Wrap(c.Writer).
			Status(http.StatusInternalServerError).
			Message(http.StatusText(http.StatusInternalServerError)).
			Send()
		return
	}
	jsend.Wrap(c.Writer).
		Status(http.StatusOK).
		Data(resultjson).
		Message(http.StatusText(http.StatusOK)).
		Send()
}

// DeletePost DeletePost
func DeletePost(c *gin.Context) {
	resultjson, err := apiclient.DeleteService("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		jsend.Wrap(c.Writer).
			Status(http.StatusInternalServerError).
			Message(http.StatusText(http.StatusInternalServerError)).
			Send()
		return
	}
	jsend.Wrap(c.Writer).
		Status(http.StatusOK).
		Data(resultjson).
		Message(http.StatusText(http.StatusOK)).
		Send()
}
