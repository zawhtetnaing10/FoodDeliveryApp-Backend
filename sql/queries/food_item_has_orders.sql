-- name: BulkInsertFoodItemsOrders :copyfrom
INSERT INTO food_items_has_orders (food_item_id, order_id, quantity)
VALUES($1, $2, $3);