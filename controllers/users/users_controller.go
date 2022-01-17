package users

import (
	"github.com/aprilnurf/grocerystore_users-api/domain/users"
	"github.com/aprilnurf/grocerystore_users-api/services"
	"github.com/aprilnurf/grocerystore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(ctx *gin.Context) {
	userId, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user_id should be a number")
		ctx.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
		return
	}
	ctx.JSON(http.StatusOK, user)
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
		restError := errors.NewBadRequestError("invalid Json")
		ctx.JSON(restError.Status, restError)
		return
	}
	result, err := services.CreateUser(user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

//func SearchUser(ctx *gin.Context) {
//	ctx.String(http.StatusNotImplemented, "Implement me!")
//}
