package controller

import (
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/requests/order"
	"akimbaev/resources"
	"akimbaev/response"
	"akimbaev/service"
	"net/http"
	"net/url"
	"strconv"
)

type PostController struct {
	service service.PostService
}

func NewPostController(service service.PostService) *PostController {
	return &PostController{
		service: service,
	}
}

// r - page - count - sort
func (c *PostController) GetPosts(w http.ResponseWriter, r *http.Request) {
	userClaims, _ := helpers.ExctractUserFromToken(r)

	q := r.URL.Query()

	params := order.IndexRequest{
		Page:  getQueryInt(q, "page"),
		Count: getQueryInt(q, "count"),
		Sort:  q.Get("sort"),
	}

	msg, validErr := helpers.ValidateStruct(&params)

	if validErr != nil {
		response.Json(w, http.StatusBadRequest, msg)
		return
	}

	posts, err := c.service.GetPosts(int(userClaims.UserID), params)

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
		return
	}

	response.Json(w, http.StatusOK, resources.PostsResource(posts))
}

func (c *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	userClaims, _ := helpers.ExctractUserFromToken(r)

	var request models.Post

	e := helpers.ReadJson(r, w, &request)

	if e != nil {
		response.Json(w, e.HTTPStatus(), e.Details())
	}

	msg, validErr := helpers.ValidateStruct(request)

	if validErr != nil {
		response.Json(w, http.StatusBadRequest, msg)
		return
	}

	post, err := c.service.CreatePost(int(userClaims.UserID), request)

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
		return
	}

	response.Json(w, http.StatusOK, resources.PostResource(post))
}

func (c *PostController) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	err := c.service.Delete(id)

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
		return
	}

	response.Json(w, http.StatusOK, helpers.Envelope{"message": "Post deleted successfully"})
}

func getQueryInt(q url.Values, key string) int {
	if value := q.Get(key); value != "" {
		if num, err := strconv.Atoi(value); err == nil {
			return num
		}
	}
	return 1
}
