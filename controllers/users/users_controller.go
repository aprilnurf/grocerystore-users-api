package users

import (
	"github.com/aprilnurf/grocerystore_users-api/domain/users"
	"github.com/aprilnurf/grocerystore_users-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "Implement me!")
}

func CreateUser(ctx *gin.Context) {
	var user users.User
	//bytes, err := ioutil.ReadAll(ctx.Request.Body)
	//if err != nil {
	//	// TO DO Error handler
	//	return
	//}
	//
	//err = json.Unmarshal(bytes, &user)
	//if err != nil {
	//	// TO DO Error handler
	//	return
	//}
	//OR you can use ShouldBindJSON
	if err := ctx.ShouldBindJSON(&user); err != nil {
		//return bad request
	}
	result, err := services.CreateUser(user)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

//func SearchUser(ctx *gin.Context) {
//	ctx.String(http.StatusNotImplemented, "Implement me!")
//}
