-- A--
SELECT name
FROM users
WHERE gender = 'L';

-- B--
SELECT *
FROM products
WHERE id = 3;

-- C--
SELECT * 
FROM users 
WHERE created_at >= CURRENT_DATE - INTERVAL 7 DAY 
AND 
name LIKE '%a%';

-- C--
SELECT COUNT(*) as total
FROM users
WHERE gender = 'P';

-- D--
SELECT *
FROM users
ORDER BY name ASC;

-- E--
SELECT *
FROM products
LIMIT 5;