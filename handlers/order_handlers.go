package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zawhtetnaing10/FoodDeliveryApp-Backend/internal/database"
)

// TODO: - Calculate total cost manually
type OrderRequest struct {
	TotalCost         float64           `json:"total_cost"`
	PaymentMethodId   int64             `json:"payment_method_id"`
	DeliveryAddressId int64             `json:"delivery_address_id"`
	FoodItems         []FoodItemRequest `json:"food_items"`
}

type FoodItemRequest struct {
	Id       int64 `json:"id"`
	Quantity int   `json:"quantity"`
}

func (cfg *ApiConfig) SubmitOrder(w http.ResponseWriter, r *http.Request) {
	// Validate JWT
	token, err := GetBearerToken(r.Header)
	if err != nil {
		cfg.LogError("Cannot read the authorization token", err)
		RespondWithError(w, http.StatusUnauthorized, "You are not authorized to add orders")
		return
	}
	userId, jwtErr := ValidateJWT(token, cfg.TokenSecret)
	if jwtErr != nil {
		cfg.LogError("Validate jwt failed", jwtErr)
		RespondWithError(w, http.StatusUnauthorized, "You are not authorized to add orders")
		return
	}

	// Decode request
	decoder := json.NewDecoder(r.Body)
	request := OrderRequest{}
	if decodeErr := decoder.Decode(&request); decodeErr != nil {
		cfg.LogError("Error decoding request", decodeErr)
		RespondWithError(w, http.StatusBadRequest, "Invalid order request")
		return
	}

	// Validate request
	if validateErr := validateOrderRequest(request); validateErr != nil {
		cfg.LogError("Invalid request", validateErr)
		RespondWithError(w, http.StatusBadRequest, validateErr.Error())
		return
	}

	// The whole insert transaction
	transactionErr := cfg.WithTransaction(r.Context(), func(qtx *database.Queries) error {

		// Convert total cost
		totalCost, convertErr := convertFloatToPgtypeNumeric(request.TotalCost)
		if convertErr != nil {
			return convertErr
		}

		// Generate order number
		orderNum, orderNumErr := GenerateOrderNumber()
		if orderNumErr != nil {
			return orderNumErr
		}

		// Insert Order First
		orderParams := database.CreateOrderParams{
			UserID:            pgtype.Int8{Int64: userId, Valid: true},
			DeliveryAddressID: pgtype.Int8{Int64: request.DeliveryAddressId, Valid: true},
			PaymentMethodID:   pgtype.Int8{Int64: request.PaymentMethodId, Valid: true},
			TotalCost:         totalCost,
			OrderNumber:       orderNum,
		}

		// Insert order first
		orderInDb, orderInsertErr := qtx.CreateOrder(r.Context(), orderParams)
		if orderInsertErr != nil {
			return orderInsertErr
		}

		// Insert order with food items
		// Prepare request
		orderFoodItemRequests := []database.BulkInsertFoodItemsOrdersParams{}
		for _, foodItem := range request.FoodItems {
			request := database.BulkInsertFoodItemsOrdersParams{
				FoodItemID: foodItem.Id,
				OrderID:    orderInDb.ID,
				Quantity:   int32(foodItem.Quantity),
			}
			orderFoodItemRequests = append(orderFoodItemRequests, request)
		}

		// Bulk Insert
		_, orderAndFoodItemsErr := qtx.BulkInsertFoodItemsOrders(r.Context(), orderFoodItemRequests)
		if orderAndFoodItemsErr != nil {
			return orderAndFoodItemsErr
		}

		return nil
	})

	if transactionErr != nil {
		cfg.LogError("Transaction error", transactionErr)
		RespondWithError(w, http.StatusInternalServerError, transactionErr.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Order successfully submitted"))
}

// Validate order request
func validateOrderRequest(request OrderRequest) error {
	if request.TotalCost == 0.0 {
		return errors.New("total cost must not be null and must be greater than zero")
	}

	if request.PaymentMethodId == 0 {
		return errors.New("no payment method in request")
	}

	if request.DeliveryAddressId == 0 {
		return errors.New("no delivery address in request")
	}

	if len(request.FoodItems) == 0 {
		return errors.New("there are no food items in request")
	}

	return nil
}

// Generate random order number
func GenerateOrderNumber() (string, error) {
	bytes := make([]byte, 12)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", errors.New("cannot generate order number")
	}

	result := hex.EncodeToString(bytes)
	return result, nil
}
