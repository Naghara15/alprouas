INSERT INTO products (name, description, price, stock)
VALUES
  ('Laptop X200', 'A powerful business laptop', 1499.99, 25),
  ('Smartphone Z10', 'High-end smartphone with OLED display', 899.99, 50),
  ('Wireless Mouse', 'Ergonomic wireless mouse', 29.99, 200),
  ('Mechanical Keyboard', 'RGB backlit mechanical keyboard', 79.99, 120),
  ('Monitor 24"', 'Full HD 24-inch monitor', 139.99, 80),
  ('USB-C Hub', '7-in-1 USB-C docking station', 49.99, 150),
  ('External SSD 1TB', 'Portable 1TB SSD drive', 129.99, 60),
  ('Gaming Chair', 'Adjustable gaming chair with lumbar support', 249.99, 40),
  ('Webcam 1080p', 'HD webcam for streaming and video calls', 59.99, 90),
  ('Bluetooth Speaker', 'Portable waterproof speaker', 39.99, 100),
  ('Graphics Card RTX 4060', 'Mid-range GPU for gaming and productivity', 399.99, 35),
  ('Power Bank 20,000mAh', 'Fast-charging portable power bank', 44.99, 180),
  ('Smartwatch Pro', 'Advanced smartwatch with fitness tracking', 199.99, 75),
  ('Desk Lamp LED', 'Adjustable LED desk lamp with touch control', 24.99, 110),
  ('Noise Cancelling Headphones', 'Over-ear headphones with ANC', 149.99, 55);


INSERT INTO tags (name)
VALUES
  ('Electronics'),
  ('Laptop'),
  ('Smartphone'),
  ('Accessories'),
  ('Input Devices'),
  ('Monitor'),
  ('Storage'),
  ('Furniture'),
  ('Audio'),
  ('Gaming'),
  ('Power'),
  ('Wearable'),
  ('Lighting'),
  ('Video'),
  ('Networking');


INSERT INTO product_tag (product_id, tag_id)
VALUES
  (1, 1), (1, 2),
  (2, 1), (2, 3),
  (3, 1), (3, 4), (3, 5),
  (4, 1), (4, 4), (4, 5),
  (5, 1), (5, 6),
  (6, 1), (6, 4),
  (7, 1), (7, 4), (7, 7),
  (8, 8), (8, 10),
  (9, 1), (9, 4), (9, 14),
  (10, 1), (10, 4), (10, 9),
  (11, 1), (11, 10),
  (12, 1), (12, 4), (12, 11),
  (13, 1), (13, 3), (13, 12),
  (14, 4), (14, 13),
  (15, 1), (15, 9);
