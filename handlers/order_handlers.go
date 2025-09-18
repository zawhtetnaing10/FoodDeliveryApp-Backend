package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zawhtetnaing10/FoodDeliveryApp-Backend/internal/database"
)

type orderRequest struct {
	PaymentMethodId   int64             `json:"payment_method_id"`
	DeliveryAddressId int64             `json:"delivery_address_id"`
	FoodItems         []foodItemRequest `json:"food_items"`
}

type foodItemRequest struct {
	Id       int64 `json:"id"`
	Quantity int   `json:"quantity"`
}

// TODO: -
type orderResponse struct {
	Id              int64                          `json:"id"`
	OrderNumber     string                         `json:"order_number"`
	TotalCost       float64                        `json:"total_cost"`
	CreatedAt       time.Time                      `json:"created_at"`
	UpdatedAt       time.Time                      `json:"updated_at"`
	DeliveryAddress deliveryAddressResponse        `json:"delivery_address"`
	PaymentMethod   paymentMethodResponse          `json:"payment_method"`
	FoodItems       []foodItemWithQuantityResponse `json:"food_items"`
}

type foodItemWithQuantityResponse struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	ImageUrl    string  `json:"image_url"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
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
	request := orderRequest{}
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

	// Does DA Exists
	daExistsParams := database.DoesDAExistsForUserParams{
		ID: request.DeliveryAddressId,
		UserID: pgtype.Int8{
			Int64: userId,
			Valid: true,
		},
	}
	daExists, daErr := cfg.Db.DoesDAExistsForUser(r.Context(), daExistsParams)
	if daErr != nil {
		cfg.LogError("DA Exists error in order", daErr)
		RespondWithError(w, http.StatusInternalServerError, daErr.Error())
		return
	}
	if !daExists {
		RespondWithError(w, http.StatusBadRequest, "Delivery address does not exists for the user")
		return
	}

	// Does PM Exists
	pmExistsParams := database.DoesPMExistsForUserParams{
		ID: request.PaymentMethodId,
		UserID: pgtype.Int8{
			Int64: userId,
			Valid: true,
		},
	}
	pmExists, pmErr := cfg.Db.DoesPMExistsForUser(r.Context(), pmExistsParams)
	if pmErr != nil {
		cfg.LogError("PM Exists error in order", pmErr)
		RespondWithError(w, http.StatusInternalServerError, pmErr.Error())
		return
	}
	if !pmExists {
		RespondWithError(w, http.StatusBadRequest, "Payment Method does not exist for the user")
		return
	}

	// Insert
	insertedOrderId := int64(0)

	// The whole insert transaction
	transactionErr := cfg.WithTransaction(r.Context(), func(qtx *database.Queries) error {

		foodItemsByte, foodItemsErr := json.Marshal(request.FoodItems)
		if foodItemsErr != nil {
			return foodItemsErr
		}

		totalCost, totalCostErr := cfg.Db.CalculateTotalCost(r.Context(), foodItemsByte)
		if totalCostErr != nil {
			return totalCostErr
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
		insertedOrderId = orderInDb.ID
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

	// Check the inserted order id
	if insertedOrderId == 0 {
		cfg.LogError("Order id not updated", errors.New("order id not updated"))
		RespondWithError(w, http.StatusInternalServerError, "There was something wrong while submitting order")
		return
	}

	// Get the inserted order and then return
	insertedOrderRows, getOrderErr := cfg.Db.GetOrderById(r.Context(), insertedOrderId)
	if getOrderErr != nil || len(insertedOrderRows) == 0 {
		cfg.LogError("Cannot get the inserted order", getOrderErr)
		RespondWithError(w, http.StatusInternalServerError, "There was something wrong while returning the inserted order")
		return
	}

	// Get the inserted order
	firstItem := insertedOrderRows[0]

	// Parse total cost
	totalCost, totalCostErr := firstItem.OrderTotalCost.Float64Value()
	if totalCostErr != nil {
		cfg.LogError("Cannot parse total cost", totalCostErr)
		RespondWithError(w, http.StatusInternalServerError, totalCostErr.Error())
		return
	}

	// Populate order data, delivery address and payment methods
	response := orderResponse{
		Id:          firstItem.OrderID,
		OrderNumber: firstItem.OrderNumber,
		TotalCost:   totalCost.Float64,
		CreatedAt:   firstItem.OrderCreatedAt.Time,
		UpdatedAt:   firstItem.OrderUpdatedAt.Time,
		DeliveryAddress: deliveryAddressResponse{
			Id:            firstItem.DeliveryAddressID,
			StreetAddress: firstItem.DeliveryAddress,
			CreatedAt:     firstItem.DeliveryAddressCreatedAt.Time,
			UpdatedAt:     firstItem.DeliveryAddressUpdatedAt.Time,
		},
		PaymentMethod: paymentMethodResponse{
			Id:         firstItem.PaymentMethodID,
			CardNumber: firstItem.PaymentMethodCardNumber,
			ExpiryDate: firstItem.PaymentMethodExpiryDate,
			CVV:        int(firstItem.PaymentMethodCvv),
			NameOnCard: firstItem.PaymentMethodNameOnCard,
			CreatedAt:  firstItem.PaymentMethodCreatedAt.Time,
			UpdatedAt:  firstItem.PaymentMethodUpdatedAt.Time,
		},
		FoodItems: []foodItemWithQuantityResponse{},
	}

	// Populate food items
	for _, orderRow := range insertedOrderRows {

		price, priceErr := orderRow.FoodItemPrice.Float64Value()
		if priceErr != nil {
			cfg.LogError("Cannot parse price", priceErr)
			RespondWithError(w, http.StatusInternalServerError, priceErr.Error())
			return
		}

		foodItemWithQuantityResponse := foodItemWithQuantityResponse{
			Id:          orderRow.FoodItemID,
			Name:        orderRow.FoodItemName,
			ImageUrl:    orderRow.FoodItemImageUrl,
			Description: orderRow.FoodItemDescription,
			Price:       price.Float64,
			Quantity:    int(orderRow.FoodItemQuantity),
		}

		response.FoodItems = append(response.FoodItems, foodItemWithQuantityResponse)
	}

	RespondWithJson(w, http.StatusCreated, response)
}

func (cfg *ApiConfig) GetOrdersForUser(w http.ResponseWriter, r *http.Request) {
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

	orderRows, orderErr := cfg.Db.GetOrdersForUser(r.Context(), pgtype.Int8{Int64: userId, Valid: true})
	if orderErr != nil {
		cfg.LogError("Fetch orders failed", orderErr)
		RespondWithError(w, http.StatusInternalServerError, "There was something wrong while fetching orders")
		return
	}

	// Data structures to map the response
	ordersMap := map[int64]*orderResponse{}
	response := []*orderResponse{}

	// Create response
	for _, orderRow := range orderRows {

		totalCost, totalCostErr := orderRow.OrderTotalCost.Float64Value()
		if totalCostErr != nil {
			cfg.LogError("Cannot parse total cost", totalCostErr)
			RespondWithError(w, http.StatusInternalServerError, totalCostErr.Error())
			return
		}

		price, priceErr := orderRow.FoodItemPrice.Float64Value()
		if priceErr != nil {
			cfg.LogError("Cannot parse price", priceErr)
			RespondWithError(w, http.StatusInternalServerError, priceErr.Error())
			return
		}

		order, ok := ordersMap[orderRow.OrderID]
		if !ok {
			// No order found in map. Craft new
			orderResponse := orderResponse{
				Id:          orderRow.OrderID,
				OrderNumber: orderRow.OrderNumber,
				TotalCost:   totalCost.Float64,
				CreatedAt:   orderRow.OrderCreatedAt.Time,
				UpdatedAt:   orderRow.OrderUpdatedAt.Time,
				DeliveryAddress: deliveryAddressResponse{
					Id:            orderRow.DeliveryAddressID,
					StreetAddress: orderRow.DeliveryAddress,
					CreatedAt:     orderRow.DeliveryAddressCreatedAt.Time,
					UpdatedAt:     orderRow.DeliveryAddressUpdatedAt.Time,
				},
				PaymentMethod: paymentMethodResponse{
					Id:         orderRow.PaymentMethodID,
					CardNumber: orderRow.PaymentMethodCardNumber,
					ExpiryDate: orderRow.PaymentMethodExpiryDate,
					CVV:        int(orderRow.PaymentMethodCvv),
					NameOnCard: orderRow.PaymentMethodNameOnCard,
					CreatedAt:  orderRow.PaymentMethodCreatedAt.Time,
					UpdatedAt:  orderRow.PaymentMethodUpdatedAt.Time,
				},
				FoodItems: []foodItemWithQuantityResponse{
					{
						Id:          orderRow.FoodItemID,
						Name:        orderRow.FoodItemName,
						ImageUrl:    orderRow.FoodItemImageUrl,
						Description: orderRow.FoodItemDescription,
						Price:       price.Float64,
						Quantity:    int(orderRow.FoodItemQuantity),
					},
				},
			}

			ordersMap[orderRow.OrderID] = &orderResponse
			response = append(response, &orderResponse)
		} else {
			// Order already found. Update food item
			foodItemWithQuantityResponse := foodItemWithQuantityResponse{
				Id:          orderRow.FoodItemID,
				Name:        orderRow.FoodItemName,
				ImageUrl:    orderRow.FoodItemImageUrl,
				Description: orderRow.FoodItemDescription,
				Price:       price.Float64,
				Quantity:    int(orderRow.FoodItemQuantity),
			}
			order.FoodItems = append(order.FoodItems, foodItemWithQuantityResponse)
		}
	}

	// Return the responses
	RespondWithJson(w, http.StatusOK, response)
}

// Validate order request
func validateOrderRequest(request orderRequest) error {
	if request.PaymentMethodId == 0 {
		return errors.New("no payment method in request")
	}

	if request.DeliveryAddressId == 0 {
		return errors.New("no delivery address in request")
	}

	if len(request.FoodItems) == 0 {
		return errors.New("there are no food items in request")
	}

	// Check if food items in request have duplicate ids
	idHash := map[int64]bool{}
	for _, foodItem := range request.FoodItems {
		_, ok := idHash[foodItem.Id]
		if !ok {
			idHash[foodItem.Id] = true
		} else {
			return errors.New("there must not be duplicate food item ids in request")
		}
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
