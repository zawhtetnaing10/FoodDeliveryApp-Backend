package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/zawhtetnaing10/FoodDeliveryApp-Backend/internal/database"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type forgetPasswordCheckEmailRequest struct {
	Email string `json:"email"`
}

type forgetPasswordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerRequest struct {
	Email    string `json:"email"`
	FullName string `json:"fullname"`
	Password string `json:"password"`
}

type userWithTokenResponse struct {
	Id          int64     `json:"id"`
	Email       string    `json:"email"`
	Fullname    string    `json:"fullname"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AccessToken string    `json:"access_token"`
}

type forgetPasswordCheckEmailResponse struct {
	User               userWithoutTokenResponse `json:"user"`
	ResetPasswordToken string                   `json:"reset_password_token"`
}

type userWithoutTokenResponse struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Register
func (cfg *ApiConfig) Register(w http.ResponseWriter, request *http.Request) {
	// Decode request
	decoder := json.NewDecoder(request.Body)
	registerRequest := registerRequest{}
	if decodeErr := decoder.Decode(&registerRequest); decodeErr != nil {
		cfg.LogError(fmt.Sprintf("failed to decode register request: %v", registerRequest), decodeErr)
		RespondWithError(w, http.StatusBadRequest, "Please check if all fields are present in register request")
		return
	}

	// Check the request
	if registerRequest.Email == "" || registerRequest.FullName == "" || registerRequest.Password == "" {
		RespondWithError(w, http.StatusBadRequest, "Please check if all fields are present in register request")
		return
	}

	// Hash the password
	hashedPass, hashErr := HashPassword(registerRequest.Password)
	if hashErr != nil {
		cfg.LogError("failed to hash password", hashErr)
		RespondWithError(w, http.StatusInternalServerError, "Please check if all fields are present in register request")
		return
	}

	// Register params
	params := database.CreateUserParams{
		Fullname:       registerRequest.FullName,
		Email:          registerRequest.Email,
		Hashedpassword: hashedPass,
	}

	// Create the user
	createdUser, createUserErr := cfg.Db.CreateUser(request.Context(), params)
	if createUserErr != nil {
		cfg.LogError("Failed to create user", createUserErr)
		RespondWithError(w, http.StatusInternalServerError, "Please check if all fields are present in register request")
		return
	}

	jwtToken, jwtErr := MakeJWT(createdUser.ID, cfg.TokenSecret, 1*time.Hour)
	if jwtErr != nil {
		cfg.LogError(fmt.Sprintf("Error creating jwt token %v", jwtErr), jwtErr)
		RespondWithError(w, http.StatusInternalServerError, "Please check if all fields are present in register request")
		return
	}

	response := userWithTokenResponse{
		Id:          createdUser.ID,
		Email:       createdUser.Email,
		Fullname:    createdUser.Fullname,
		CreatedAt:   createdUser.CreatedAt,
		UpdatedAt:   createdUser.UpdatedAt,
		AccessToken: jwtToken,
	}

	RespondWithJson(w, http.StatusCreated, response)
}

// Login
func (cfg *ApiConfig) Login(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	loginRequest := loginRequest{}
	// Decode request
	if decodeErr := decoder.Decode(&loginRequest); decodeErr != nil {
		cfg.LogError(fmt.Sprintf("failed to decode login request %v", loginRequest), decodeErr)
		RespondWithError(w, http.StatusBadRequest, "Incorrect email or password.")
		return
	}

	// Check request
	if loginRequest.Email == "" || loginRequest.Password == "" {
		RespondWithError(w, http.StatusBadRequest, "Incorrect email or password.")
		return
	}

	// Find user by email
	userFromDb, err := cfg.Db.GetUserByEmail(request.Context(), loginRequest.Email)
	if err != nil {
		cfg.LogError(fmt.Sprintf("error getting user by email %v", loginRequest), err)
		RespondWithError(w, http.StatusBadRequest, "Incorrect email or password.")
		return
	}

	// Check password
	if checkPassErr := CheckPasswordHash(userFromDb.Hashedpassword, loginRequest.Password); checkPassErr != nil {
		cfg.LogError(fmt.Sprintf("Password check failed : %v", checkPassErr), checkPassErr)
		RespondWithError(w, http.StatusBadRequest, "Incorrect email or password.")
		return
	}

	jwtToken, jwtErr := MakeJWT(userFromDb.ID, cfg.TokenSecret, 1*time.Hour)
	if jwtErr != nil {
		cfg.LogError(fmt.Sprintf("Error creating jwt token %v", jwtErr), jwtErr)
		RespondWithError(w, http.StatusBadRequest, "Incorrect email or password.")
		return
	}

	// Create request
	response := userWithTokenResponse{
		Id:          userFromDb.ID,
		Email:       userFromDb.Email,
		Fullname:    userFromDb.Fullname,
		CreatedAt:   userFromDb.CreatedAt,
		UpdatedAt:   userFromDb.UpdatedAt,
		AccessToken: jwtToken,
	}

	RespondWithJson(w, http.StatusOK, response)
}

// Check Email
func (cfg *ApiConfig) CheckEmail(w http.ResponseWriter, request *http.Request) {
	// Parse request
	decoder := json.NewDecoder(request.Body)
	req := forgetPasswordCheckEmailRequest{}
	if decodeErr := decoder.Decode(&req); decodeErr != nil {
		cfg.LogError(fmt.Sprintf("Failed to decode request: %v", req), decodeErr)
		RespondWithError(w, http.StatusBadRequest, "Please confirm if there's already an account with this email.")
		return
	}

	// Check request
	if req.Email == "" {
		RespondWithError(w, http.StatusBadRequest, "Email cannot be empty")
		return
	}

	// Get User By Email
	userFromDb, err := cfg.Db.GetUserByEmail(request.Context(), req.Email)
	if err != nil {
		cfg.LogError("Failed to find user", err)
		RespondWithError(w, http.StatusNotFound, "Please confirm if there's already an account with this email.")
		return
	}

	// Generate Token
	token, jwtErr := MakeJWT(userFromDb.ID, cfg.TokenSecret, 1*time.Hour)
	if jwtErr != nil {
		cfg.LogError("Error creating jwt", jwtErr)
		RespondWithError(w, http.StatusInternalServerError, "Error creating password reset token.")
		return
	}

	userWithoutToken := userWithoutTokenResponse{
		Id:        userFromDb.ID,
		Email:     userFromDb.Email,
		Fullname:  userFromDb.Fullname,
		CreatedAt: userFromDb.CreatedAt,
		UpdatedAt: userFromDb.UpdatedAt,
	}

	response := forgetPasswordCheckEmailResponse{
		User:               userWithoutToken,
		ResetPasswordToken: token,
	}

	RespondWithJson(w, http.StatusOK, response)
}

// Forget Password
func (cfg *ApiConfig) ForgetPassword(w http.ResponseWriter, request *http.Request) {

	token, getTokenErr := GetBearerToken(request.Header)
	if getTokenErr != nil {
		cfg.LogError("Error getting bearer token", getTokenErr)
		RespondWithError(w, http.StatusUnauthorized, "You are not authorized to reset your password")
		return
	}

	_, jwtErr := ValidateJWT(token, cfg.TokenSecret)
	if jwtErr != nil {
		cfg.LogError("Validate jwt failed", jwtErr)
		RespondWithError(w, http.StatusUnauthorized, "You are not authorized to reset your password")
		return
	}

	// Parse request
	decoder := json.NewDecoder(request.Body)
	req := forgetPasswordRequest{}
	if decodeErr := decoder.Decode(&req); decodeErr != nil {
		cfg.LogError(fmt.Sprintf("Failed to decode request: %v", req), decodeErr)
		RespondWithError(w, http.StatusBadRequest, "Reset password failed. Please check if data being sent is correct.")
		return
	}

	// Check Request
	if req.Email == "" || req.Password == "" {
		RespondWithError(w, http.StatusBadRequest, "Reset password failed. Please check if data being sent is correct.")
		return
	}

	newHashedPass, hashErr := HashPassword(req.Password)
	if hashErr != nil {
		cfg.LogError("Password hash failed", hashErr)
		RespondWithError(w, http.StatusInternalServerError, "There was something wrong while resetting the password.")
		return
	}
	params := database.UpdatePasswordParams{
		Email:          req.Email,
		Hashedpassword: newHashedPass,
	}
	_, updateErr := cfg.Db.UpdatePassword(request.Context(), params)
	if updateErr != nil {
		cfg.LogError(fmt.Sprintf("Update password failed: %v", req), updateErr)
		RespondWithError(w, http.StatusInternalServerError, "There was something wrong while resetting the password.")
		return
	}

	w.WriteHeader(http.StatusOK)
}
