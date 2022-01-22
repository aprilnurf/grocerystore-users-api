package users

import (
	"github.com/aprilnurf/grocerystore_users-api/domain/users"
	"github.com/aprilnurf/grocerystore_users-api/services"
	"github.com/aprilnurf/grocerystore_users-api/utils/errors_utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getUserId(UserIdParam string) (int64, *errors_utils.RestError) {
	userId, userErr := strconv.ParseInt(UserIdParam, 10, 64)
	if userErr != nil {
		return 0, errors_utils.NewBadRequestError("user_id should be a number")
	}
	return userId, nil
}

func GetUser(ctx *gin.Context) {
	userId, userErr := getUserId(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
		return
	}
	ctx.JSON(http.StatusOK, user.Marshall(ctx.GetHeader("X-Public") == "true"))
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
		restError := errors_utils.NewBadRequestError("invalid Json")
		ctx.JSON(restError.Status, restError)
		return
	}
	result, err := services.CreateUser(user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusCreated, result.Marshall(ctx.GetHeader("X-Public") == "true"))
}

//func SearchUser(ctx *gin.Context) {
//	ctx.String(http.StatusNotImplemented, "Implement me!")
//}

func UpdateUser(ctx *gin.Context) {
	userId, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors_utils.NewBadRequestError("user_id should be a number")
		ctx.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		//return bad request
		restError := errors_utils.NewBadRequestError("invalid Json")
		ctx.JSON(restError.Status, restError)
		return
	}

	user.Id = userId

	isPartial := ctx.Request.Method == http.MethodPatch
	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusCreated, result.Marshall(ctx.GetHeader("X-Public") == "true"))
}

func Search(ctx *gin.Context) {
	status := ctx.Query("status")
	isActive, err := strconv.ParseBool(status)
	if err != nil {
		restError := errors_utils.NewBadRequestError("invalid Json")
		ctx.JSON(restError.Status, restError)
		return
	}
	users, errSearch := services.Search(isActive)

	if errSearch != nil {
		ctx.JSON(errSearch.Status, errSearch)
	}
	//users.Marshall(ctx.GetHeader("X-Public") == "true")
	ctx.JSON(http.StatusOK, users)
}

func Delete(ctx *gin.Context) {
	userId, userErr := getUserId(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}