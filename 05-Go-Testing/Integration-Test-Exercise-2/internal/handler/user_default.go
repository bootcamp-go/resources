package handler

import (
	"app/internal"
	"app/platform/web/request"
	"app/platform/web/response"
	"errors"
	"net/http"
)

// NewUserDefault returns a new instance of UserDefault.
func NewUserDefault(rp internal.UserRepository) *UserDefault {
	return &UserDefault{rp}
}

// UserDefault is an struct that contain the handler for user
type UserDefault struct {
	// rp is the repository.
	rp internal.UserRepository
}

// UserJSON is an struct that represent the user in JSON format.
type UserJSON struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// GetAll returns all the users.
func (h *UserDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		u, err := h.rp.FindAll()
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error getting users")
			return
		}

		// response
		// - serialize
		data := make(map[int]UserJSON)
		for k, v := range u {
			data[k] = UserJSON{
				Id:    v.Id,
				Name:  v.Name,
				Age:   v.Age,
				Email: v.Email,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// RequestBodyCreateUser is an struct that represent the request body for create user.
type RequestBodyCreateUser struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// Create creates a new user.
func (h *UserDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body
		var body RequestBodyCreateUser
		err := request.JSON(r, &body)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body request")
			return
		}

		// process
		u := internal.User{
			Id: body.Id,
			UserAttributes: internal.UserAttributes{
				Name:  body.Name,
				Age:   body.Age,
				Email: body.Email,
			},
		}
		err = h.rp.Save(&u)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrFieldDuplicated):
				response.Error(w, http.StatusConflict, "user has duplicated field")
			default:
				response.Error(w, http.StatusInternalServerError, "error creating user")
			}
			return
		}

		// response
		// - serialize
		data := UserJSON{
			Id:    u.Id,
			Name:  u.Name,
			Age:   u.Age,
			Email: u.Email,
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
