-- A--
INSERT INTO operator (op_name)
VALUES ('Operator A'), ('Operator B'), ('Operator C'), ('Operator D'), ('Operator E');

-- B--
INSERT INTO product_type (prod_type)
VALUES ('Type A'), ('Type B'), ('Type C');

-- C--
INSERT INTO products (prod_type_id, operator_id, prod_code, prod_name, prod_status)
VALUES 
(1, 3, 'P001', 'Product 1', '20'),
(1, 3, 'P002', 'Product 2', '15');

-- D--
INSERT INTO products (prod_type_id, operator_id, prod_code, prod_name, prod_status)
VALUES 
(2, 1, 'P003', 'Product 3', '20'),
(2, 1, 'P004', 'Product 4', '10'),
(2, 1, 'P005', 'Product 5', '15');

-- E--
INSERT INTO products (prod_type_id, operator_id, prod_code, prod_name, prod_status)
VALUES 
(3, 4, 'P006', 'Product 6', '20'),
(3, 4, 'P007', 'Product 7', '10'),
(3, 4, 'P008', 'Product 8', '15');

-- F--
INSERT INTO product_description (prod_desc, prod_id)
VALUES 
('Deskripsi Product 1', 1),
('Deskripsi Product 2', 2),
('Deskripsi Product 3', 3),
('Deskripsi Product 4', 4),
('Deskripsi Product 5', 5),
('Deskripsi Product 6', 6),
('Deskripsi Product 7', 7),
('Deskripsi Product 8', 8);

-- G--
INSERT INTO payment_methods (pay_method) 
VALUES
('Go-Pay'),
('OVO'),
('ShopeePay');

-- H--
INSERT INTO users (name, alamat, tanggal_lahir, status_user, gender) VALUES
('User 001', 'Jl. Sudirman No. 123, Jakarta', '1990-05-15', 'aktif', 'L'),
('User 002', 'Jl. Gatot Subroto No. 456, Jakarta', '1995-10-20', 'aktif', 'P'),
('User 003', 'Jl. Danau Sentani Tengah, Malang', '2002-04-08', 'aktif', 'L'),
('User 004', 'Jl. Danau Limboto, Malang', '2002-12-31', 'aktif', 'P'),
('User 005', 'Jl. Danau Sentai Raya, Malang', '2000-06-30', 'aktif', 'L');

-- I--
-- Transaksi user id = 1
INSERT INTO transactions (user_id, payment_method_id, trans_status, total_barang, total_harga) VALUES
(1, 1, 'sukses', 5, 200000),
(1, 2, 'gagal', 5, 300000),
(1, 3, 'sukses', 3, 50000);

-- Transaksi user id = 2
INSERT INTO transactions (user_id, payment_method_id, trans_status, total_barang, total_harga) VALUES
(2, 1, 'sukses', 7, 270000),
(2, 3, 'sukses', 4, 175000),
(2, 2, 'sukses', 3, 50000);

-- Transaksi user id = 3
INSERT INTO transactions (user_id, payment_method_id, trans_status, total_barang, total_harga) VALUES
(3, 2, 'gagal', 3, 150000),
(3, 1, 'sukses', 7, 235000),
(3, 3, 'sukses', 6, 115000);

-- Transaksi user id = 4
INSERT INTO transactions (user_id, payment_method_id, trans_status, total_barang, total_harga) VALUES
(4, 3, 'sukses', 4, 180000),
(4, 2, 'sukses', 4, 230000),
(4, 1, 'sukses', 12, 205000);

-- Transaksi user id = 5
INSERT INTO transactions (user_id, payment_method_id, trans_status, total_barang, total_harga) VALUES
(5, 1, 'gagal', 3, 150000),
(5, 2, 'sukses', 5, 235000),
(5, 3, 'sukses', 8, 215000);

-- J--
-- Insert 3 produk pada transaksi 1
INSERT INTO transaction_details (transaction_id, prod_id, detail_status, total_barang, total_harga)
VALUES (1, 1, 'sukses', 1, 20000),
       (1, 2, 'sukses', 2, 100000),
       (1, 3, 'sukses', 2, 80000);

-- Insert 3 produk pada transaksi 2
INSERT INTO transaction_details (transaction_id, prod_id, detail_status, total_barang, total_harga)
VALUES (2, 2, 'gagal', 1, 50000),
       (2, 3, 'gagal', 2, 80000),
       (2, 4, 'gagal', 2, 170000);

-- Insert 3 produk pada transaksi 3
INSERT INTO transaction_details (transaction_id, prod_id, detail_status, total_barang, total_harga)
VALUES (3, 5, 'sukses', 1, 10000),
       (3, 6, 'sukses', 1, 15000),
       (3, 7, 'sukses', 1, 25000);

-- Insert 3 produk pada transaksi 4
INSERT INTO transaction_details (transaction_id, prod_id, detail_status, total_barang, total_harga)
VALUES (4, 8, 'sukses', 5, 150000),
       (4, 1, 'sukses', 1, 20000),
       (4, 2, 'sukses', 1, 100000);

-- Insert 3 produk pada transaksi 5
INSERT INTO transaction_details (transaction_id, prod_id, detail_status, total_barang, total_harga)
VALUES (5, 3, 'sukses', 2, 80000),
       (5, 4, 'sukses', 1, 85000),
       (5, 5, 'sukses', 1, 10000);
    
-- Insert 3 produk pada transaksi 6
INSERT INTO transaction_details (transaction_id, prod_id, detail_status, total_barang, total_harga)
VALUES (6, 6, 'sukses', 1, 15000),
       (6, 7, 'sukses', 1, 25000),
       (6, 8, 'sukses', 1, 30000);

-- Insert 3 produk pada transaksi 7
INSERT INTO transaction_details (transaction_id, prod_id, detail_status, total_barang, total_harga)
VALUES (7, 8, 'gagal', 1, 30000),
       (7, 1, 'gagal', 1, 20000),
       (7, 2, 'gagal', 1, 100000);

-- Insert 3 produk pada transaksi 8
INSERT INTO transaction_details (transaction_id, prod_id, detail_status, total_barang, total_harga)
VALUES (8, 3, 'sukses', 3, 120000),
       (8, 4, 'sukses', 1, 85000),
       (8, 5, 'sukses', 3, 30000);
    
-- Insert 3 produk pada transaksi 9
INSERT INTO transaction_details (transaction_id, prod_id, detail_status, total_barang, total_harga)
VALUES (9, 6, 'sukses', 4, 60000),
       (9, 7, 'sukses', 1, 25000),
       (9, 8, 'sukses', 1, 30000);

-- Insert 3 produk pada transaksi 10
INSERT INTO transaction_details (transaction_id, prod_id, detail_status, total_barang, total_harga)
VALUES (10, 8, 'sukses', 2, 60000),
       (10, 1, 'sukses', 1, 20000),
       (10, 2, 'sukses', 1, 100000);

-- Insert 3 produk pada transaksi 11
INSERT INTO transaction_details (transaction_id, prod_id, detail_status, total_barang, total_harga)
VALUES (11, 3, 'sukses', 1, 40000),
       (11, 4, 'sukses', 2, 170000),
       (11, 5, 'sukses', 1, 10000);
    
-- Insert 3 produk pada transaksi 12
INSERT INTO transaction_details (transaction_id, prod_id, detail_status, total_barang, total_harga)
VALUES (12, 6, 'sukses', 10, 150000),
       (12, 7, 'sukses', 1, 25000),
       (12, 8, 'sukses', 1, 30000);

-- Insert 3 produk pada transaksi 13
INSERT INTO transaction_details (transaction_id, prod_id, detail_status, total_barang, total_harga)
VALUES (13, 8, 'sukses', 1, 30000),
       (13, 1, 'sukses', 1, 20000),
       (13, 2, 'sukses', 1, 100000);

-- Insert 3 produk pada transaksi 14
INSERT INTO transaction_details (transaction_id, prod_id, detail_status, total_barang, total_harga)
VALUES (14, 3, 'sukses', 3, 120000),
       (14, 4, 'sukses', 1, 85000),
       (14, 5, 'sukses', 1, 10000);
    
-- Insert 3 produk pada transaksi 15
INSERT INTO transaction_details (transaction_id, prod_id, detail_status, total_barang, total_harga)
VALUES (15, 6, 'sukses', 1, 15000),
       (15, 7, 'sukses', 6, 150000),
       (15, 8, 'sukses', 1, 30000);



