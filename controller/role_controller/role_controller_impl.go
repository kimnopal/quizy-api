package role_controller

import (
	"encoding/json"
	"net/http"
	"quizy-api/helper"
	"quizy-api/model/web"
	"quizy-api/model/web/role_dto"
	"quizy-api/service/role_service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type RoleControllerImpl struct {
	service role_service.RoleService
}

func NewServiceController(service role_service.RoleService) RoleController {
	return &RoleControllerImpl{
		service: service,
	}
}

func (controller *RoleControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	roleCreateRequest := role_dto.RoleCreateRequest{}
	err := decoder.Decode(&roleCreateRequest)
	helper.PanicIfError(err)

	roleResponse := controller.service.Create(request.Context(), roleCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "success create role",
		Data:   roleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoleControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleUpdateRequest := role_dto.RoleUpdateRequest{}
	helper.ReadFromRequestBody(request, &roleUpdateRequest)

	roleId := params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	helper.PanicIfError(err)
	roleUpdateRequest.Id = int64(id)

	roleResponse := controller.service.Update(request.Context(), roleUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "success update role",
		Data:   roleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoleControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleId := params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	helper.PanicIfError(err)

	controller.service.Delete(request.Context(), int64(id))

	webResponse := web.WebResponse{
		Code:   200,
		Status: "success delete role",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
