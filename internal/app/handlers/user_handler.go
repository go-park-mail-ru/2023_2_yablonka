package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	apperrors "server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	dto "server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/service"

	"github.com/asaskevich/govalidator"
)

type UserHandler struct {
	us service.IUserService
}

// @Summary Поменять пароль
// @Description Получает старый и новый пароли, а также id пользователя
// @Tags user
//
// @Accept  json
// @Produce  json
//
// @Param passwords body dto.PasswordChangeInfo true "id, старый и новый пароли пользователя"
//
// @Success 200  {string} string "no content"
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /user/edit/change_password/ [post]
func (uh UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var passwords dto.PasswordChangeInfo

	err := json.NewDecoder(r.Body).Decode(&passwords)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(passwords)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}

	err = uh.us.UpdatePassword(rCtx, passwords)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	r.Body.Close()
}

// @Summary Поменять данные профиля
// @Description В ответ ничего не шлёт
// @Tags auth
//
// @Accept  json
// @Produce  json
//
// @Param newProfileInfo body dto.UserProfileInfo true "id пользователя, имя, фамилия, описание пользователя"
//
// @Success 200  {string} string "no content"
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /user/edit/ [post]
func (uh UserHandler) ChangeProfile(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var newProfileInfo dto.UserProfileInfo

	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&newProfileInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	log.Println("request struct decoded")

	// _, err = govalidator.ValidateStruct(newProfileInfo)
	// if err != nil {
	// 	*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
	// 	return
	// }
	// log.Println("request struct validated")

	newProfileInfo.UserID = user.ID

	err = uh.us.UpdateProfile(rCtx, newProfileInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("user profile updated")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response generated")

	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("response written")

	r.Body.Close()
	log.Println("response closed")
}

// @Summary Поменять аватарку
// @Description В ответ шлёт ссылку на файл
// @Tags auth
//
// @Accept  json
// @Produce  json
//
// @Param avatarChangeInfo body dto.AvatarChangeInfo true "id пользователя, изображение"
//
// @Success 200  {object}  doc_structs.AvatarUploadResponse "Объект пользователя"
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /user/edit/change_avatar/ [post]
// @Router /user/edit/change_avatar/ [post]
func (uh UserHandler) ChangeAvatar(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var avatarChangeInfo dto.AvatarChangeInfo

	err := json.NewDecoder(r.Body).Decode(&avatarChangeInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	log.Println("request struct decoded")

	_, err = govalidator.ValidateStruct(avatarChangeInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}
	log.Println("request struct validated")

	url, err := uh.us.UpdateAvatar(rCtx, avatarChangeInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("avatar uploaded")

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"avatar_url": url,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response generated")

	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("response written")

	r.Body.Close()
	log.Println("response closed")
}
