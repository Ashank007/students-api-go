package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Ashank007/students-api-go/internal/storage"
	"github.com/Ashank007/students-api-go/internal/types"
	"github.com/Ashank007/students-api-go/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New (storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Creating a Student")

		var student types.Student 

		err := json.NewDecoder(r.Body).Decode(&student)	

		if errors.Is(err,io.EOF) {
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(fmt.Errorf("Empty Body")))
			return
		}

		if err != nil {
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w,http.StatusBadRequest,response.ValidationError(validateErrs))
			return
		}

		lastId, err := storage.CreateStudent(student.Name,student.Email,student.Age)

		slog.Info("User Created Successfully",slog.String("UserId",fmt.Sprint(lastId)))

		if err != nil {
      response.WriteJson(w,http.StatusInternalServerError,err)
			return
		}

		response.WriteJson(w,http.StatusCreated,map[string]int64{"id":lastId})
	}
}
