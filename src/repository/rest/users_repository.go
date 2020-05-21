package rest

import (
	"encoding/json"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"

	"github.com/f4nt0md3v/bookstore_oauth-api/src/domain/users"
	"github.com/f4nt0md3v/bookstore_oauth-api/src/utils/errors"
)

var (
	usersRestClient = rest.RequestBuilder{
		Timeout:        100 * time.Millisecond,
		BaseURL:        "https://api.bookstore.com",
	}
)

type UsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}

type usersRepository struct {}

func NewRepository() UsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestError) {
	req := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	res := usersRestClient.Post("/users/login", req)
	if res == nil || res.Response == nil {
		return nil, errors.NewInternalServerError("invalid rest client response when trying to login user")
	}
	if res.StatusCode > 299 {
		var restErr *errors.RestError
		err := json.Unmarshal(res.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, restErr
	}
	var u users.User
	if err := json.Unmarshal(res.Bytes(), &u); err != nil {
		return nil, errors.NewInternalServerError("invalid user interface when trying to login user")
	}
	return &u, nil
}
