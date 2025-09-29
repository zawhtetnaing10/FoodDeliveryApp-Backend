package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zawhtetnaing10/FoodDeliveryApp-Backend/internal/database"
)

// Request
type addDeliveryAddressAndPaymentMethodRequest struct {
	PaymentMethod   paymentMethodRequest   `json:"payment_method"`
	DeliveryAddress deliveryAddressRequest `json:"delivery_address"`
}

type deliveryAddressRequest struct {
	StreetAddress string `json:"street_address"`
}

type paymentMethodRequest struct {
	CardNumber string `json:"card_number"`
	ExpiryDate string `json:"expiry_date"`
	CVV        int32  `json:"cvv"`
	NameOnCard string `json:"name_on_card"`
}

// Responses
type deliveryAddressListAndPaymentMethodListForUserResponse struct {
	PaymentMethods    []paymentMethodResponse   `json:"payment_methods"`
	DeliveryAddresses []deliveryAddressResponse `json:"delivery_addresses"`
}

type deliveryAddressAndPaymentMethodResponse struct {
	PaymentMethod   paymentMethodResponse   `json:"payment_method"`
	DeliveryAddress deliveryAddressResponse `json:"delivery_address"`
}

type deliveryAddressResponse struct {
	Id            int64     `json:"id"`
	StreetAddress string    `json:"street_address"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type paymentMethodResponse struct {
	Id         int64     `json:"id"`
	CardNumber string    `json:"card_number"`
	ExpiryDate string    `json:"expiry_date"`
	CVV        int32     `json:"cvv"`
	NameOnCard string    `json:"name_on_card"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (cfg *ApiConfig) AddDeliveryAddressAndPaymentMethod(w http.ResponseWriter, r *http.Request) {
	// Check Token
	token, err := GetBearerToken(r.Header)
	if err != nil {
		cfg.LogError("Cannot read the authorization token", err)
		RespondWithError(w, http.StatusUnauthorized, "You are not authorized to add payment method and delivery addresses")
		return
	}

	user_id, jwtErr := ValidateJWT(token, cfg.TokenSecret)
	if jwtErr != nil {
		cfg.LogError("Validate jwt failed", jwtErr)
		RespondWithError(w, http.StatusUnauthorized, "You are not authorized to add payment method and delivery addresses")
		return
	}

	// Parse the request
	request := addDeliveryAddressAndPaymentMethodRequest{}
	decoder := json.NewDecoder(r.Body)
	if decodeErr := decoder.Decode(&request); decodeErr != nil {
		cfg.LogError(fmt.Sprintf("Error parsing request body: %v", request), decodeErr)
		RespondWithError(w, http.StatusBadRequest, "Failed to add delivery address and payment method. Please check if the data provided is correct.")
		return
	}

	// Add the payment method
	paymentMethodParams := database.CreatePaymentMethodParams{
		CardNumber: request.PaymentMethod.CardNumber,
		ExpiryDate: request.PaymentMethod.ExpiryDate,
		Cvv:        request.PaymentMethod.CVV,
		NameOnCard: request.PaymentMethod.NameOnCard,
		UserID: pgtype.Int8{
			Int64: user_id,
			Valid: true,
		},
	}
	paymentMethodDb, addPaymentMethodErr := cfg.Db.CreatePaymentMethod(r.Context(), paymentMethodParams)
	if addPaymentMethodErr != nil {
		cfg.LogError("Failed to add payment method", addPaymentMethodErr)
		RespondWithError(w, http.StatusInternalServerError, "Failed to add payment method")
		return
	}
	paymentMethodResponse := paymentMethodResponse{
		Id:         paymentMethodDb.ID,
		CardNumber: paymentMethodDb.CardNumber,
		ExpiryDate: paymentMethodDb.ExpiryDate,
		CVV:        paymentMethodDb.Cvv,
		NameOnCard: paymentMethodDb.NameOnCard,
		CreatedAt:  paymentMethodDb.CreatedAt.Time,
		UpdatedAt:  paymentMethodDb.UpdatedAt.Time,
	}

	// Add the delivery address
	deliveryAddressParams := database.CreateDeliveryAddressParams{
		StreetAddress: request.DeliveryAddress.StreetAddress,
		UserID: pgtype.Int8{
			Int64: user_id,
			Valid: true,
		},
	}
	deliveryAddressDb, deliveryAddressErr := cfg.Db.CreateDeliveryAddress(r.Context(), deliveryAddressParams)
	if deliveryAddressErr != nil {
		cfg.LogError("Failed to add delivery address", deliveryAddressErr)
		RespondWithError(w, http.StatusInternalServerError, "Failed to add  delivery address")
		return
	}
	deliveryAddressResponse := deliveryAddressResponse{
		Id:            deliveryAddressDb.ID,
		StreetAddress: deliveryAddressDb.StreetAddress,
		CreatedAt:     deliveryAddressDb.CreatedAt.Time,
		UpdatedAt:     deliveryAddressDb.UpdatedAt.Time,
	}

	// Build the response
	response := deliveryAddressAndPaymentMethodResponse{
		DeliveryAddress: deliveryAddressResponse,
		PaymentMethod:   paymentMethodResponse,
	}

	RespondWithJson(w, http.StatusCreated, response)
}

func (cfg *ApiConfig) GetDeliveryAddressAndPaymentMethodForUser(w http.ResponseWriter, r *http.Request) {
	// Check Token
	token, err := GetBearerToken(r.Header)
	if err != nil {
		cfg.LogError("Cannot read the authorization token", err)
		RespondWithError(w, http.StatusUnauthorized, "You are not authorized to add payment method and delivery addresses")
		return
	}

	user_id, jwtErr := ValidateJWT(token, cfg.TokenSecret)
	if jwtErr != nil {
		cfg.LogError("Validate jwt failed", jwtErr)
		RespondWithError(w, http.StatusUnauthorized, "You are not authorized to add payment method and delivery addresses")
		return
	}

	// Get Delivery Addresses
	deliveryAddressesDb, getDelAddErr := cfg.Db.GetDeliveryAddressesForUser(r.Context(), pgtype.Int8{Int64: user_id, Valid: true})
	if getDelAddErr != nil {
		cfg.LogError("Failed to get delivery addresses for user", getDelAddErr)
		RespondWithError(w, http.StatusInternalServerError, "Failed to get delivery addresses.")
		return
	}
	deliveryAddresses := []deliveryAddressResponse{}
	for _, addressDb := range deliveryAddressesDb {
		deliveryAddressResponse := deliveryAddressResponse{
			Id:            addressDb.ID,
			StreetAddress: addressDb.StreetAddress,
			CreatedAt:     addressDb.CreatedAt.Time,
			UpdatedAt:     addressDb.UpdatedAt.Time,
		}
		deliveryAddresses = append(deliveryAddresses, deliveryAddressResponse)
	}

	// Get Payment Methods
	paymentMethodsDb, getPaymentMethodsErr := cfg.Db.GetPaymentMethodsByUser(r.Context(), pgtype.Int8{Int64: user_id, Valid: true})
	if getPaymentMethodsErr != nil {
		cfg.LogError("Failed to get payment methods for user", getPaymentMethodsErr)
		RespondWithError(w, http.StatusInternalServerError, "Failed to get payment methods.")
		return
	}
	paymentMethods := []paymentMethodResponse{}
	for _, paymentMethodDb := range paymentMethodsDb {
		paymentMethodResponse := paymentMethodResponse{
			Id:         paymentMethodDb.ID,
			CardNumber: paymentMethodDb.CardNumber,
			ExpiryDate: paymentMethodDb.ExpiryDate,
			CVV:        paymentMethodDb.Cvv,
			NameOnCard: paymentMethodDb.NameOnCard,
			CreatedAt:  paymentMethodDb.CreatedAt.Time,
			UpdatedAt:  paymentMethodDb.UpdatedAt.Time,
		}
		paymentMethods = append(paymentMethods, paymentMethodResponse)
	}

	response := deliveryAddressListAndPaymentMethodListForUserResponse{
		DeliveryAddresses: deliveryAddresses,
		PaymentMethods:    paymentMethods,
	}

	RespondWithJson(w, http.StatusOK, response)
}
