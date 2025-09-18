SELECT EXISTS (
    SELECT * FROM payment_methods 
    WHERE user_id = 1 AND id = 1
);