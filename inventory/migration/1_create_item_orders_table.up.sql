CREATE TABLE IF NOT EXISTS item_orders(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    item_id INT NOT NULL,
    order_id INT NOT NULL,
    order_quantity INT,
    order_fulfilled BOOLEAN,
    UNIQUE(item_id, order_id),
    
    CONSTRAINT fk_item
    FOREIGN KEY (item_id)
    REFERENCES items(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);