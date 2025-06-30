package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Ashank007/students-api-go/internal/storage"
	"github.com/Ashank007/students-api-go/internal/types"
	"github.com/Ashank007/students-api-go/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handleCreateStudent(storage, w, r)
		case http.MethodGet:
			handleGetStudents(storage, w)
		case http.MethodPut:
			handleUpdateStudent(storage, w, r)
		case http.MethodDelete:
			handleDeleteStudent(storage, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func handleCreateStudent(storage storage.Storage, w http.ResponseWriter, r *http.Request) {
	slog.Info("Creating a Student")
	var student types.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if errors.Is(err, io.EOF) {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Empty Body")))
		return
	}
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

	if err := validator.New().Struct(student); err != nil {
		validateErrs := err.(validator.ValidationErrors)
		response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
		return
	}

	lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)
	if err != nil {
		response.WriteJson(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("User Created Successfully", slog.String("UserId", fmt.Sprint(lastId)))
	response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
}

func handleGetStudents(storage storage.Storage, w http.ResponseWriter) {
	slog.Info("Fetching Students List")

	students, err := storage.GetAllStudents()
	if err != nil {
		response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
		return
	}
	response.WriteJson(w, http.StatusOK, students)
}

func handleUpdateStudent(storage storage.Storage, w http.ResponseWriter, r *http.Request) {
	slog.Info("Updating Student")

	var student types.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil || student.Id == 0 {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Invalid input")))
		return
	}

	if err := storage.UpdateStudent(student); err != nil {
		response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
		return
	}

	response.WriteJson(w, http.StatusOK, map[string]string{"message": "Student updated"})
}

func handleDeleteStudent(storage storage.Storage, w http.ResponseWriter, r *http.Request) {
	slog.Info("Deleting Student")

	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id == 0 {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Invalid ID")))
		return
	}

	if err := storage.DeleteStudent(id); err != nil {
		response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
		return
	}

	response.WriteJson(w, http.StatusOK, map[string]string{"message": "Student deleted"})
}


