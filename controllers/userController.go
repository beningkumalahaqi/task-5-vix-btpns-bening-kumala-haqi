package controllers

import (
	"encoding/json"
	"net/http"
	"task5-vix/configs"
	"task5-vix/helpers"
	"task5-vix/models"
)

func GetAllProfile(w http.ResponseWriter, r *http.Request) {

	var user []models.User

	if err := configs.DB.Find(&user).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Users List", user)
}

//Users

func Me(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)

	userResponse := &models.MyProfile{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	helpers.Response(w, 200, "My Profile", userResponse)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	userJWT := r.Context().Value("userinfo").(*helpers.MyCustomClaims)
	var userCheck models.User

	if err := configs.DB.First(&userCheck, "id = ?", userJWT.ID).Error; err != nil {
		helpers.Response(w, 404, "ID not found", nil)
		return
	}

	var update models.UpdateProfile

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	passwordHash, err := helpers.HashPassword(update.Password)
	if err != nil {
		helpers.Response(w, 500, err.Error(),nil)
		return
	}

	user := models.User {
		Username: update.Username,
		Email: update.Email,
		Password: passwordHash,
	}

	if err := configs.DB.Where("id = ?", userJWT.ID).Updates(&user).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	
	userResponse := &models.UpdateProfile{
		Username: user.Username,
		Email:    user.Email,
		Password: "Password Changed",
	}

	helpers.Response(w, 201, "Succes Edit Profile", userResponse)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userJWT := r.Context().Value("userinfo").(*helpers.MyCustomClaims)

	var user models.User

	res := configs.DB.Where("id = ?", userJWT.ID).Delete(&user)

	if res.Error != nil {
		helpers.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helpers.Response(w, 404, "User not found", nil)
		return
	}

	helpers.Response(w, 200, "User Deleted, Username : " + userJWT.Username , nil)	
}
