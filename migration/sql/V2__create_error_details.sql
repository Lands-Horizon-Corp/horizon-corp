-- V1__create_error_details_table.sql

-- Check if the table already exists
CREATE TABLE IF NOT EXISTS `error_details` (
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` DATETIME(3) DEFAULT NULL,
    `message` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `stack` TEXT DEFAULT NULL,
    `response` TEXT DEFAULT NULL,
    `status` INT DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX `idx_error_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert some seed data
INSERT INTO `error_details` (`message`, `name`, `stack`, `response`, `status`)
VALUES 
    ('Error connecting to the database', 'DatabaseError', 'Some stack trace here', 'Database unreachable', 500),
    ('Timeout error on request', 'TimeoutError', 'Request timed out after 30 seconds', 'No response from server', 408),
    ('Invalid input', 'ValidationError', 'Invalid format for email address', 'User input error', 400);
