CREATE TABLE IF NOT EXISTS `contacts` (
    `id` SERIAL PRIMARY KEY,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) UNIQUE NOT NULL,
    `contact_number` VARCHAR(15) NOT NULL,
    description TEXT,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO `contacts` (`first_name`, `last_name`, `email`, `contact_number`, `description`)
VALUES
('John', 'Doe', 'john.doe@example.com', '1234567890', 'Hello! This is a test message.'),
('Jane', 'Smith', 'jane.smith@example.com', '0987654321', 'This is another test message.');