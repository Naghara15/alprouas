INSERT INTO products (name, description, price, stock, image)
VALUES
  ('ROG phone 8 pro, 24GB', 'High-end smartphone', 13899000, 25, 'ROG.png'),
  ('Laptop', 'High-end smartphone with OLED display', 20000000, 50, 'laptop2.png'),
  ('T-Shirt', 'Ergonomic wireless mouse', 135000, 200, 'baju.jpg'),
  ('Jeans', 'RGB backlit mechanical keyboard', 120000, 120, 'celana.jpg'),
  ('Gamepad', 'Full HD 24-inch monitor', 250000, 80, 'gamepad.png'),
  ('Bookshelf', '7-in-1 USB-C docking station', 120000, 150, 'rak.webp'),
  ('Paket Novel Tere Liye', 'Portable 1TB SSD drive', 480000, 60, 'tereliye.jpg'),
  ('Sepatu Arei Montreal', 'Adjustable gaming chair with lumbar support', 789000, 40, 'sepatu.webp');


INSERT INTO tags (name)
VALUES
('Electronics'),
('Fashion'),
('Home'),
('Books');


INSERT INTO product_tags (product_id, tag_id) VALUES
(1, 1),
(2, 1),
(3, 2),
(4, 2),
(5, 1),
(6, 3),
(7, 4),
(8, 2);
