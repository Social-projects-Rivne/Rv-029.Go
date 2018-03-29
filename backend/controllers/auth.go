package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/jwt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/mail"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
	"bytes"
	"strings"
	"io"
	"encoding/csv"
)

type errorResponse struct {
	Status  bool
	Message string
}

type loginResponse struct {
	Status  bool
	Message string
	Token   string
	User    models.User
}

type registerResponse struct {
	Status  bool
	Message string
	User    models.User
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequestData validator.LoginRequestData

	err := decodeAndValidate(r, &loginRequestData)
	if err != nil {
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/auth.go error: %+v", err), StatusCode: http.StatusBadRequest}
		response.Failed(w)
		return
	}

	user := models.User{}
	user.Email = loginRequestData.Email
	user, err = models.UserDB.CheckUserPassword(user)

	if err != nil {
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/auth.go error: %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	if user.Password != password.EncodePassword(loginRequestData.Password, user.Salt) {
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/auth.go There is no such user with email and password combination. error: %+v", err), StatusCode: http.StatusUnauthorized}
		response.Failed(w)
		return
	}

	// generate jwt token from user claims
	token, err := jwt.GenerateToken(&user)
	if err != nil {
		log.Printf("Error in controllers/auth error: %+v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	//TODO change on helper, than will be not Token: token but Data: token, and change it in frontend
	jsonResponse, err := json.Marshal(loginResponse{
		Status:  true,
		Message: "You was successfully authenticated.",
		Token:   token,
		User:    user,
	})
	if err != nil {
		log.Printf("Error in controllers/auth error: %+v", err)
		return
	}

	w.Write(jsonResponse)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var registerRequestData validator.RegisterRequestData
	err := decodeAndValidate(r, &registerRequestData)
	if err != nil {
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/auth.go error: %+v", err)}
		response.Failed(w)
		return
	}

	salt := password.GenerateSalt(8)
	user := models.User{
		UUID:      gocql.TimeUUID(),
		Email:     registerRequestData.Email,
		FirstName: registerRequestData.FirstName,
		LastName:  registerRequestData.LastName,
		Salt:      salt,
		Password:  password.EncodePassword(registerRequestData.Password, salt),
		Role:      models.ROLE_USER,
		Status:    0,
		Photo:     "../static/nigga.png",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := models.UserDB.Insert(&user); err != nil {
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/auth.go error: %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	message := fmt.Sprintf("Hello %s,\nYou was successfully registered in \"Task manager\".\n To activate your account please go to the <a href=\"http://localhost/authorization/login/?token=%s&uuid=%s\">LINK</a>\n Your ID: %s\n Regards\n", user.FirstName, user.Password, user.UUID, user.UUID)
	mail.Mailer.Send(user.Email, user.FirstName, "Successfully Registered", message)

	jsonResponse, err := json.Marshal(registerResponse{
		Status:  true,
		Message: "You was successfully registered",
		User:    user,
	})
	if err != nil {
		log.Printf("Error in controllers/auth error: %+v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	return
}

func ConfirmRegistration(w http.ResponseWriter, r *http.Request) {
	var confirmRegistrationRequestData validator.ConfirmRegistrationRequestData

	err := decodeAndValidate(r, &confirmRegistrationRequestData)
	if err != nil {
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/auth.go error: %+v", err), StatusCode: http.StatusBadRequest}
		response.Failed(w)
		return
	}
	user := models.User{}
	user.Status = 1
	if err = models.UserDB.Update(&user); err != nil {
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/auth.go error: %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	jsonResponse, err := json.Marshal(struct {
		Status  bool
		Message string
	}{
		Status:  true,
		Message: "Your account was successfully activated.",
	})
	if err != nil {
		log.Printf("Error in controllers/auth error: %+v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	return
}

//ForgotPassword ..
func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var forgotRequestData validator.ForgotPasswordRequestData

	err := decodeAndValidate(r, &forgotRequestData)
	if err != nil {
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/auth.go error: %+v", err), StatusCode: http.StatusBadRequest}
		response.Failed(w)
		return
	}

	user := models.User{}
	user.Email = forgotRequestData.Email
	// TODO err handler
	if err = models.UserDB.FindByEmail(&user); err != nil {
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/auth.go error: %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	message := fmt.Sprintf("Hello %s,\nIt is your link to restore password <a href=\"http://localhost/authorization/new-password/%s\">LINK</a>\n", user.FirstName, user.Password)
	mail.Mailer.Send(user.Email, user.FirstName, "Successfully Registered", message)

	jsonResponse, err := json.Marshal(registerResponse{
		Status:  true,
		Message: "Your link sent",
	})
	if err != nil {
		log.Printf("Error in controllers/auth error: %+v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	return

}

//ResetPassword ..
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var resetRequestData validator.ResetPasswordRequestData
	err := decodeAndValidate(r, &resetRequestData)

	if err != nil {
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/auth.go error: %+v", err), StatusCode: http.StatusBadRequest}
		response.Failed(w)
		return
	}
	user := models.User{}
	user.Email = resetRequestData.Email
	user, err = models.UserDB.CheckUserPassword(user)
	if err != nil {
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/auth.go error: %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	// FIXME doesn't work properly
	if user.Password != resetRequestData.Token {
		jsonResponse, err := json.Marshal(errorResponse{
			Status:  false,
			Message: "Invalid reset token",
		})
		if err != nil {
			log.Printf("Error in controllers/auth error: %+v", err)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	user.Password = password.EncodePassword(resetRequestData.Password, user.Salt)

	models.UserDB.Update(&user)

	jsonResponse, err := json.Marshal(registerResponse{
		Status:  true,
		Message: "Your password restored",
	})
	if err != nil {
		log.Printf("Error in controllers/auth error: %+v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	return

}

//GetUserInfo gives frontend information about user
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := gocql.ParseUUID(vars["user_id"])
	user := models.User{}
	if vars["user_id"] == "own" {
		user = r.Context().Value("user").(models.User)
	} else {
		user = models.User{
			UUID: userID,
		}
		models.UserDB.FindByID(&user)		
	}

	var b gocql.UUID
	a := make([]gocql.UUID, 0)
	for k := range user.Projects {
		b, _ = gocql.ParseUUID(fmt.Sprintf("%s", k))
		a = append(a, b)
	}
	projects, err := models.ProjectDB.GetProjectsNamesList(a)
	if err != nil {
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/auth.go error: %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}
	flag := true
	if len(projects) > 0 {
		for k := range user.Projects {
			for i := 0; flag; i++ {
				if k == projects[i].ID {
					user.Projects[k] = projects[i].Name					
					flag = false
				}
			}
			flag = true
		}
	}
	user.Photo = "static/nigga.jpeg"

	response := helpers.Response{Message: "Done", Data: user, StatusCode: http.StatusOK}
	response.Success(w)
	return

}

//UpdateUserInfo gives frontend information about user
func UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)
	var updateRequestData validator.UpdateUserRequestData
	err := decodeAndValidate(r, &updateRequestData)
	if err != nil {
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/auth.go error: %+v", err), StatusCode: http.StatusBadRequest}
		response.Failed(w)
		return
	}
	user.FirstName = updateRequestData.FirstName
	user.LastName = updateRequestData.LastName
	user.UpdatedAt = time.Now()
	if err = models.UserDB.UpdateFirstAndLastName(&user); err != nil {
		response := helpers.Response{Status: false, Message: fmt.Sprintf("Error occured in controllers/auth.go error: %+v", err), StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Done", StatusCode: http.StatusOK}
	response.Success(w)
	return

}

// Import users from csv file
func Import(w http.ResponseWriter, r *http.Request)  {
	var respLogs []string
	var Buf bytes.Buffer
	// get file from request
	file, _, err := r.FormFile("import")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// copy the file data to the buffer
	io.Copy(&Buf, file)
	contents := Buf.String()

	// create reader
	reader := csv.NewReader(strings.NewReader(contents))

	// read file content line by line
	for i := 0 ; ; i++  {
		line, err := reader.Read()
		// check if not the end of file
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			response := helpers.Response{Status: false, Message: fmt.Sprintf("error during reading file content. error: %+v", err), StatusCode: http.StatusInternalServerError}
			response.Failed(w)
			return
		}
		if i == 0 { // skip headers
			continue
		}

		user := models.User{
			Email: line[0],
		}
		// check if user with such email already exists
		err = models.UserDB.FindByEmail(&user)
		if err != nil { // user not exists
			user.UUID = gocql.TimeUUID()
			user.FirstName = line[1]
			user.LastName = line[2]
			user.Salt = password.GenerateSalt(8)
			user.Password = password.EncodePassword(password.EncodeMD5(line[3]), user.Salt)
			user.Role = line[3]
			user.Status = 1
			user.CreatedAt = time.Now()
			user.UpdatedAt = time.Now()

			err = models.UserDB.Insert(&user)
			if err != nil {
				respLogs = append(respLogs, fmt.Sprintf("Error occurred: %v", err))
				response := helpers.Response{Status: false, Message: fmt.Sprintf("error during importing users. error: %+v", err), StatusCode: http.StatusInternalServerError}
				response.Failed(w)
				return
			}

			respLogs = append(respLogs, fmt.Sprintf("User with email %s successfully imported", line[0]))
		} else {
			respLogs = append(respLogs, fmt.Sprintf("User with email %s already exists", line[0]))
		}
	}

	Buf.Reset()

	response := helpers.Response{Message: "Done", Data: respLogs, StatusCode: http.StatusOK}
	response.Success(w)
	return
}