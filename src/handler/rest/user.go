package rest

import (
	"crud-user/src/business/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *rest) CreateUser(ctx *gin.Context) {
	var userParam entity.User
	if err := ctx.ShouldBindJSON(&userParam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "please fill all param correctly",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	user, err := r.uc.User.Create(userParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something wrong",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "successfully created new user",
		"error":   nil,
		"data":    user,
	})
}

func (r *rest) GetAllUsers(ctx *gin.Context) {
	users, err := r.uc.User.GetList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something wrong",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "successfully get all users",
		"error":   nil,
		"data":    users,
	})
}

func (r *rest) UpdateUsers(ctx *gin.Context) {
	var userSelect entity.UserSelectParam
	if err := ctx.ShouldBindUri(&userSelect); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "please fill all param correctly",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	var userParam entity.UserUpdateParam
	if err := ctx.ShouldBindJSON(&userParam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "please fill all param correctly",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	users, err := r.uc.User.Update(userSelect.Id, userParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something wrong",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "successfully update a user",
		"error":   nil,
		"data":    users,
	})
}

func (r *rest) DeleteUsers(ctx *gin.Context) {
	var userSelect entity.UserSelectParam
	if err := ctx.ShouldBindUri(&userSelect); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "please fill all param correctly",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	if err := r.uc.User.Delete(userSelect.Id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something wrong",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "successfully delete a user",
		"error":   nil,
		"data":    nil,
	})
}
