-- 1--
SELECT * FROM transactions WHERE user_id = 1
UNION
SELECT * FROM transactions WHERE user_id = 2;

-- 2--
SELECT SUM(total_harga) AS total_harga
FROM transactions
WHERE user_id = 1;

-- 3--
SELECT SUM(td.total_harga) as total_transaksi
FROM transactions AS t
JOIN transaction_details AS td ON t.id = td.transaction_id
JOIN products AS p ON td.prod_id = p.id
WHERE products.prod_type_id = 2;

-- 4--
SELECT p.*, product_type.prod_type AS NAME
FROM products AS p
JOIN product_type ON p.prod_type_id = product_type.id;

-- 5--
SELECT transactions.*, products.*, users.*
FROM transactions
INNER JOIN users ON transactions.user_id = users.id
INNER JOIN transaction_details ON transactions.id = transaction_details.transaction_id
INNER JOIN products ON transaction_details.prod_id = products.id;

-- 6--
CREATE FUNCTION delete_transaction_details(transaction_id INT)
RETURNS BOOLEAN
BEGIN
    DELETE FROM transaction_details WHERE transaction_id = transaction_id;
    RETURN TRUE;
END

-- 7--
CREATE FUNCTION update_total_barang(transaction_id INT)
RETURNS INT
BEGIN
    DECLARE total_barang INT;
    SELECT SUM(total_barang) INTO total_barang FROM transaction_details WHERE transaction_id = transaction_id;
    UPDATE transactions SET total_barang = total_barang WHERE id = transaction_id;
    RETURN total_barang;
END

-- 8--