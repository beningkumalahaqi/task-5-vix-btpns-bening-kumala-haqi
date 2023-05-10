package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"task5-vix/configs"
	"task5-vix/helpers"
	"task5-vix/models"

	"gorm.io/gorm"
)

func GetPhoto(w http.ResponseWriter, r *http.Request) {

	userJWT := r.Context().Value("userinfo").(*helpers.MyCustomClaims)
	var photo models.Photo

	if err := configs.DB.First(&photo, "user_id = ?", userJWT.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(w, 404, "Profile Photo Not Found", nil)
			return
		}

		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, userJWT.Username+" Profile Photo", photo)
}

func UploadPhoto(w http.ResponseWriter, r *http.Request) {

	var uploadPhoto models.UploadPhoto

	if err := json.NewDecoder(r.Body).Decode(&uploadPhoto); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	userJWT := r.Context().Value("userinfo").(*helpers.MyCustomClaims)

	photo := models.Photo {
		Title: uploadPhoto.Title,
		Caption: uploadPhoto.Caption,
		PhotoUrl: uploadPhoto.PhotoUrl,
		UserID: userJWT.ID,
	}

	if err := configs.DB.Create(&photo).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Upload Photo Succesfully", nil)
}

func UpdatePhoto(w http.ResponseWriter, r *http.Request) {

	userJWT := r.Context().Value("userinfo").(*helpers.MyCustomClaims)
	var photoCheck models.Photo

	if err := configs.DB.First(&photoCheck, "user_id = ?", userJWT.ID).Error; err != nil {
		helpers.Response(w, 404, "Photo not found", nil)
		return
	}

	var update models.UpdatePhoto

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	photo := models.Photo {
		Title: update.Title,
		Caption: update.Caption,
		PhotoUrl: update.PhotoUrl,
	}

	if err := configs.DB.Where("user_id = ?", userJWT.ID).Updates(&photo).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	Response := &models.UpdatePhoto{
		Title: photo.Title,
		Caption: photo.Caption,
		PhotoUrl: photo.PhotoUrl,
	}

	helpers.Response(w, 201, "Profile Photo Updated", Response)
}

func DeletePhoto(w http.ResponseWriter, r *http.Request) {

	userJWT := r.Context().Value("userinfo").(*helpers.MyCustomClaims)

	var photo models.Photo

	res := configs.DB.Where("user_id = ?", userJWT.ID).Delete(&photo)

	if res.Error != nil {
		helpers.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helpers.Response(w, 404, "Photo not found", nil)
		return
	}

	helpers.Response(w, 200, "Photo Deleted: " + strconv.Itoa(int(photo.ID)), nil)
}
