-- A--
UPDATE products
SET prod_name = 'product dummy'
WHERE id = 1;

-- B--
UPDATE transaction_details
SET total_barang = 3
WHERE prod_id = 1;