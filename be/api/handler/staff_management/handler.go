package staffmanagement

import (
	"BE_Hospital_Management/constant"
	"BE_Hospital_Management/internal/domain/dto"
	service "BE_Hospital_Management/internal/service/staff_management"
	"BE_Hospital_Management/pkg"
	"BE_Hospital_Management/pkg/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type StaffManagementHandler struct {
	service service.StaffManagementService
}

func NewStaffManagementHandler(service service.StaffManagementService) *StaffManagementHandler {
	return &StaffManagementHandler{service: service}
}

// StaffManagement godoc
// @Summary      Assign task to staff
// @Description  Assign task to staff
// @Tags         StaffManagement
// @Accept 		json
// @Produce      json
// @Param 		 uid path int true "Staff ID"
// @Param		request	 	body		dto.TaskInfoRequest		true	"Task information"
// @param Authorization header string true "Authorization"
// @Router       /api/staff_management/staffs/{uid}/tasks [POST]
// @Success      201   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *StaffManagementHandler) AssignTaskToStaff(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when assigning task to staff. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	staffUIDStr := c.Param("uid")
	staffUID, err := strconv.ParseInt(staffUIDStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	var request dto.TaskInfoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	newTask, err := h.service.CreateTask(*authUserId, staffUID, &request)
	if err != nil {
		log.Error("Happened error when assigning task to staff. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrInvalidTimeRange):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		case errors.Is(err, service.ErrExistsOverlapTask):
			pkg.PanicExeption(constant.Conflict, err.Error())
		case errors.Is(err, service.ErrExistsOverlapAppointment):
			pkg.PanicExeption(constant.Conflict, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when assigning task to staff.")
		}
	}
	c.JSON(http.StatusCreated, pkg.BuildResponseSuccess(constant.Success, newTask))
}

// StaffManagement godoc
// @Summary      Get tasks for current staff
// @Description  Get tasks for current staff
// @Tags         StaffManagement
// @Accept 		 json
// @Produce      json
// @Router       /api/staff_management/me/tasks [get]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @param Authorization header string true "User Authorization"
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *StaffManagementHandler) GetMyTasks(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting current staff's tasks. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	taskInfo, err := h.service.GetTasksByStaffUID(*authUserId)
	if err != nil {
		log.Error("Happened error when getting current staff's tasks. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting current staff's tasks.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, taskInfo))
}

// StaffManagement godoc
// @Summary      Get tasks assigned by current manager
// @Description  Get tasks assigned by current manager
// @Tags         StaffManagement
// @Accept 		 json
// @Produce      json
// @Router       /api/staff_management/me/assigned-tasks [get]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @param Authorization header string true "User Authorization"
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *StaffManagementHandler) GetMyAssignedTasks(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting tasks assigned by current manager. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	taskInfo, err := h.service.GetTasksByManagerUID(*authUserId)
	if err != nil {
		log.Error("Happened error when getting tasks assigned by current manager. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting tasks assigned by current manager.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, taskInfo))
}

// StaffManagement godoc
// @Summary      Get tasks assigned by current manager to a specific staff
// @Description  Get tasks assigned by current manager to a specific staff
// @Tags         StaffManagement
// @Accept 		 json
// @Produce      json
// @Router       /api/staff_management/staffs/{uid}/tasks [get]
// @Param 		 uid path int true "Staff UID"
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @param Authorization header string true "User Authorization"
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *StaffManagementHandler) GetMyAssignedTasksToAStaff(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting tasks assigned by current manager to a staff. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	staffUIDStr := c.Param("uid")
	staffUID, err := strconv.ParseInt(staffUIDStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	taskInfo, err := h.service.GetTasksByMangerUIDAndStaffUID(*authUserId, staffUID)
	if err != nil {
		log.Error("Happened error when getting tasks assigned by current manager to a staff. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting tasks assigned by current manager to a staff.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, taskInfo))
}

// StaffManagement godoc
// @Summary      Get task by Id
// @Description  Get task by Id
// @Tags         StaffManagement
// @Accept 		 json
// @Produce      json
// @Router       /api/staff_management/tasks/{id} [get]
// @Param 		 id path int true "Task ID"
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @param Authorization header string true "User Authorization"
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *StaffManagementHandler) GetTaskById(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting task by id. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	authUserRole := utils.GetAuthUserRole(c)
	if authUserRole == nil {
		log.Error("Happened error when getting task by id. Error: ", "Missing user role in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user role in context")
		return
	}
	taskIdStr := c.Param("id")
	taskId, err := strconv.ParseInt(taskIdStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	taskInfo, err := h.service.GetTaskById(*authUserId, *authUserRole, taskId)
	if err != nil {
		log.Error("Happened error when getting task by id. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrTaskNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrOutOfWorkingHours):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting task by id.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, taskInfo))
}

// //StaffManagement godoc
// //@Summary      Delete task by Id
// //@Description  Delete task by Id
// //@Tags         StaffManagement
// //@Accept 		 json
// //@Produce      json
// //@Router       /api/staff_management/tasks/{id} [delete]
// //@Param 		 id path int true "Task ID"
// //@Success      200   {object}  dto.ApiResponseSuccessStruct
// //@param Authorization header string true "User Authorization"
// //@securityDefinitions.apiKey token
// //@in header
// //@name Authorization
// //@Security JWT
func (h *StaffManagementHandler) DeleteTaskById(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when deleting task by ID. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	taskIdStr := c.Param("id")
	taskId, err := strconv.ParseInt(taskIdStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	err = h.service.DeleteTaskById(*authUserId, taskId)
	if err != nil {
		log.Error("Happened error when deleting task by ID. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrTaskNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when deleting task by ID.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccessNoData())
}

// StaffManagement godoc
// @Summary      Update task by Id
// @Description  Update task by Id
// @Tags         StaffManagement
// @Accept 		 json
// @Produce      json
// @Router       /api/staff-management/tasks/{id} [PATCH]
// @Param 		 id path int true "Task ID"
// @Param		request	 	body		dto.UpdateTaskInfoRequest		true	"Updated task information"
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @param Authorization header string true "User Authorization"
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *StaffManagementHandler) UpdateTaskById(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when updating task. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	taskIdStr := c.Param("id")
	taskId, err := strconv.ParseInt(taskIdStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	var request dto.UpdateTaskInfoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	updatedTask, err := h.service.UpdateTaskById(*authUserId, taskId, &request)
	if err != nil {
		log.Error("Happened error when updating task. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrInvalidTimeRange):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		case errors.Is(err, service.ErrExistsOverlapTask):
			pkg.PanicExeption(constant.Conflict, err.Error())
		case errors.Is(err, service.ErrExistsOverlapAppointment):
			pkg.PanicExeption(constant.Conflict, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when updating task.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, updatedTask))
}
