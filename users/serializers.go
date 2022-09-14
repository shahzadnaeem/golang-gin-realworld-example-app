package users

import (
	"github.com/gin-gonic/gin"

	"github.com/gothinkster/golang-gin-realworld-example-app/common"
)

type ProfileSerializer struct {
	C *gin.Context
	UserModel
}

// Declare your response schema here
type ProfileResponse struct {
	ID        uint    `json:"-"`
	Username  string  `json:"username"`
	Bio       string  `json:"bio"`
	Image     *string `json:"image"`
	Following bool    `json:"following"`
}

// Put your response logic including wrap the userModel here.
func (self *ProfileSerializer) Response() ProfileResponse {
	myUserModel := self.C.MustGet("my_user_model").(UserModel)
	profile := ProfileResponse{
		ID:        self.ID,
		Username:  self.Username,
		Bio:       self.Bio,
		Image:     self.Image,
		Following: myUserModel.isFollowing(self.UserModel),
	}
	return profile
}

type UserSerializer struct {
	c *gin.Context
	// TODO: Check this update is the accepted practice
	//       see related changes to 'routers.go'
	UserModel
}

type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
}

func (self *UserSerializer) Response() UserResponse {
	// myUserModel := self.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		Username: self.Username,
		Email:    self.Email,
		Bio:      self.Bio,
		Image:    self.Image,
		Token:    common.GenToken(self.ID),
	}
	return user
}

type UsersSerializer struct {
	c     *gin.Context
	Users []UserModel
}

func (s *UsersSerializer) Response() []UserResponse {
	response := []UserResponse{}
	for _, user := range s.Users {
		serializer := UserSerializer{s.c, user}
		response = append(response, serializer.Response())
	}
	return response
}
