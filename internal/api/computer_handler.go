package api

import (
	"net/http"
	"strconv"

	"github.com/eniworoeva/sample-company/internal/models"
	"github.com/eniworoeva/sample-company/internal/util"
	"github.com/eniworoeva/sample-company/services"
	"github.com/gin-gonic/gin"
)

func (u *HTTPHandler) CreateComputer(c *gin.Context) {
	var computer models.Computer
	if err := c.ShouldBindJSON(&computer); err != nil {
		util.Response(c, "Invalid request", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	if len(computer.EmployeeAbbr) != 3 {
		util.Response(c, "Invalid request", http.StatusBadRequest, nil, []string{"Employee abbreviation must be 3 characters"})
		return
	}

	if err := u.Repository.CreateComputer(&computer); err != nil {
		util.Response(c, "Failed to create computer", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	// Check if the employee has 3 or more computers
	computers, err := u.Repository.GetComputersByEmployee(computer.EmployeeAbbr)
	if err == nil && len(computers) >= 3 {
		notification := services.Notification{
			Level:                "warning",
			EmployeeAbbreviation: computer.EmployeeAbbr,
			Message:              "Employee has 3 or more computers assigned",
		}
		services.Notify(notification)
	}

	util.Response(c, "Computer created", http.StatusCreated, computer, nil)
}

func (u *HTTPHandler) GetAllComputers(c *gin.Context) {
	computers, err := u.Repository.GetAllComputers()
	if err != nil {
		util.Response(c, "Failed to get computers", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}
	util.Response(c, "Success", http.StatusOK, computers, nil)
}

func (u *HTTPHandler) GetComputersByEmployee(c *gin.Context) {
	employeeAbbr := c.Param("abbr")
	computers, err := u.Repository.GetComputersByEmployee(employeeAbbr)
	if err != nil {
		util.Response(c, "Failed to get computers", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}
	util.Response(c, "Success", http.StatusOK, computers, nil)
}

func (u *HTTPHandler) GetComputerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.Response(c, "Invalid ID", http.StatusBadRequest, nil, []string{"Invalid ID"})
		return
	}

	computer, err := u.Repository.GetComputerByID(uint(id))
	if err != nil {
		util.Response(c, "Failed to get computer", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}
	util.Response(c, "Success", http.StatusOK, computer, nil)
}

func (u *HTTPHandler) UpdateComputer(c *gin.Context) {
	var computer models.Computer
	if err := c.ShouldBindJSON(&computer); err != nil {
		util.Response(c, "Invalid request", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	if err := u.Repository.UpdateComputer(&computer); err != nil {
		util.Response(c, "Failed to update computer", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}
	util.Response(c, "Computer updated", http.StatusOK, computer, nil)
}

func (u *HTTPHandler) DeleteComputer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.Response(c, "Invalid ID", http.StatusBadRequest, nil, []string{"Invalid ID"})
		return
	}

	if err := u.Repository.DeleteComputer(uint(id)); err != nil {
		util.Response(c, "Failed to delete computer", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}
	util.Response(c, "Computer deleted", http.StatusOK, nil, nil)
}

func (u *HTTPHandler) AssignComputer(c *gin.Context) {
	var input struct {
		EmployeeAbbr string `json:"employee_abbr"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		util.Response(c, "Invalid request", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.Response(c, "Invalid ID", http.StatusBadRequest, nil, []string{"Invalid ID"})
		return
	}

	computer, err := u.Repository.GetComputerByID(uint(id))
	if err != nil {
		util.Response(c, "Failed to get computer", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	computer.EmployeeAbbr = input.EmployeeAbbr

	if err := u.Repository.UpdateComputer(&computer); err != nil {
		util.Response(c, "Failed to assign computer", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	// Check if the employee has 3 or more computers
	computers, err := u.Repository.GetComputersByEmployee(computer.EmployeeAbbr)
	if err == nil && len(computers) >= 3 {
		notification := services.Notification{
			Level:                "warning",
			EmployeeAbbreviation: computer.EmployeeAbbr,
			Message:              "Employee has 3 or more computers assigned",
		}
		services.Notify(notification)
	}

	util.Response(c, "Computer assigned", http.StatusOK, computer, nil)
}
