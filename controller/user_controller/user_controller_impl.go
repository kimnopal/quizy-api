package user_controller

import (
	"net/http"
	"quizy-api/helper"
	"quizy-api/model/web"
	"quizy-api/model/web/user_dto"
	"quizy-api/service/user_service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	service user_service.UserService
}

func NewUserController(service user_service.UserService) UserController {
	return &UserControllerImpl{
		service: service,
	}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := user_dto.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.service.Create(request.Context(), userCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := user_dto.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	id := params.ByName("userId")
	userId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	userUpdateRequest.Id = int64(userId)

	userResponse := controller.service.Update(request.Context(), userUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("userId")
	userId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	controller.service.Delete(request.Context(), int64(userId))

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
