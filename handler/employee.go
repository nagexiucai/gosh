package handler

import (
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/nagexiucai/gosh/model"
	"encoding/json"
	"github.com/gorilla/mux"
	)

func GetAllEmployees(db *gorm.DB, writer http.ResponseWriter, request *http.Request) {
	employees := []model.Employee{}
	db.Find(&employees)
	respondJSON(writer, http.StatusOK, employees)
}

func CreateEmployee(db *gorm.DB, writer http.ResponseWriter, request *http.Request) {
	employee := model.Employee{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondJSON(writer, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()

	if err := db.Save(&employee).Error; err != nil {
		respondError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(writer, http.StatusCreated, employee)
	return
}

func GetEmployeeOr404(db *gorm.DB, name string, writer http.ResponseWriter, request *http.Request) *model.Employee {
	employee := model.Employee{}
	if err := db.First(&employee, model.Employee{Name: name}).Error; err != nil {
		respondError(writer, http.StatusNotFound, err.Error())
		return nil
	}
	return &employee
}

func GetEmployee(db *gorm.DB, writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	name := vars["name"]
	employee := GetEmployeeOr404(db, name, writer, request)
	if employee == nil {
		return
	}
	respondJSON(writer, http.StatusOK, employee)
	return
}

func UpdateEmployee(db *gorm.DB, writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	name := vars["name"]
	employee := GetEmployeeOr404(db, name, writer, request)
	if employee == nil {
		return
	}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondError(writer, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()

	if err := db.Save(&employee).Error; err != nil {
		respondError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(writer, http.StatusOK, employee)
}

func DeleteEmployee(db *gorm.DB, writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	name := vars["name"]
	employee := GetEmployeeOr404(db, name, writer, request)
	if employee == nil {
		return
	}
	if err := db.Delete(&employee).Error; err != nil {
		respondError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(writer, http.StatusNoContent, nil)
	return
}

func DisableEmployee(db *gorm.DB, writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	name := vars["name"]
	employee := GetEmployeeOr404(db, name, writer, request)
	if employee == nil {
		return
	}
	employee.Disable()
	if err := db.Save(&employee).Error; err != nil {
		respondError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(writer, http.StatusOK, employee)
	return
}

func EnableEmployee(db *gorm.DB, writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	name := vars["name"]
	employee := GetEmployeeOr404(db, name, writer, request)
	if employee == nil {
		return
	}
	employee.Enable()
	if err := db.Save(&employee).Error; err != nil {
		respondError(writer, http.StatusOK, err.Error())
		return
	}
	respondJSON(writer, http.StatusOK, employee)
	return
}
