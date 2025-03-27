-- +goose Up
CREATE TABLE products(
  ID UUID PRIMARY KEY NOT NULL,
  SKU VARCHAR(32) UNIQUE NOT NULL,
  name VARCHAR(64) NOT NULL,
  description TEXT NOT NULL,
  price float(10) NOT NULL,
  stock_qty INTEGER NOT NULL,
  category VARCHAR(32) NOT NULL REFERENCES categories(name)
  ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
INSERT INTO products(
 ID, SKU, name,
 description, price,
 stock_qty, category,
 created_at, updated_at
 ) VALUES (
 '002b856b-dcdf-4e4e-93ac-fd751c6836e0', 'black-forest-cake', 'Black Forest Cake',
 'A delicious black forest cake.', 100.0, 
 6, 'cake', 
 CURRENT_TIMESTAMP AT TIME ZONE 'UTC', CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
 );

 INSERT INTO products(
 ID, SKU, name,
 description, price,
 stock_qty, category,
 created_at, updated_at
 ) VALUES (
 '5fdc9d1a-1428-4cbd-b552-15eef4c36429', 'strawberry-vanilla-cake', 'Vanila Cake',
 'A fresh Vanila Cake.', 100.0, 
 2, 'cake', 
 CURRENT_TIMESTAMP AT TIME ZONE 'UTC', CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
 );

 INSERT INTO products(
 ID, SKU, name,
 description, price,
 stock_qty, category,
 created_at, updated_at
 ) VALUES (
 'e1ce6d90-ea0c-4811-8458-70d755715e4a', 'opera-coffee-cake', 'Opera Cake',
 'A fresh Opera cake.', 100.0, 
 3, 'cake', 
 CURRENT_TIMESTAMP AT TIME ZONE 'UTC', CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
 );

 INSERT INTO products(
 ID, SKU, name,
 description, price,
 stock_qty, category,
 created_at, updated_at
 ) VALUES (
 'd763cb38-c780-4a91-915a-592c7ba2143e', 'cheese-cake-oreo', 'Cheesecake Cake',
 'A fresh Cheesecake oreo.', 100.0, 
 5, 'cake', 
 CURRENT_TIMESTAMP AT TIME ZONE 'UTC', CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
 );

 INSERT INTO products(
 ID, SKU, name,
 description, price,
 stock_qty, category,
 created_at, updated_at
 ) VALUES (
 'f41f03db-9db6-463c-941d-f22b910b9bad', 'red-velvet-cake', 'Red Velvet Cake',
 'A rich, creamy red velvet cake.', 80.0, 
 5, 'cake', 
 CURRENT_TIMESTAMP AT TIME ZONE 'UTC', CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
 );

 INSERT INTO products(
 ID, SKU, name,
 description, price,
 stock_qty, category,
 created_at, updated_at
 ) VALUES (
 '0fd94f84-3d20-4362-b785-b45f01772cda', 'choco-late-cake', 'Chocolate Cake',
 'A Delicious Chocolate cake.', 110.0, 
 2, 'cake', 
 CURRENT_TIMESTAMP AT TIME ZONE 'UTC', CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
 );

INSERT INTO products(
 ID, SKU, name,
 description, price,
 stock_qty, category,
 created_at, updated_at
 ) VALUES (
 '13828a45-9831-4ea3-ab62-a918e379983e', 'southern-strawberry-cake', 'Southern Strawberry Cake',
 'A delicious Southern Strawberry Cake.', 50.0, 
 3, 'cake', 
 CURRENT_TIMESTAMP AT TIME ZONE 'UTC', CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
 ) RETURNING *;

 INSERT INTO products(
 ID, SKU, name,
 description, price,
 stock_qty, category,
 created_at, updated_at
 ) VALUES (
 'f28e45a2-7701-415a-9656-9db85f46f65f', 'straw-berry-cake', 'Strawberry Cake',
 'A fresh strawberry cake.', 70.0, 
 3, 'cake', 
 CURRENT_TIMESTAMP AT TIME ZONE 'UTC', CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
 );

-- +goose Down
DROP TABLE products;