CREATE TABLE IF NOT EXISTS `companies` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT,
    `address` VARCHAR(500),
    `longitude` DECIMAL(10,7),
    `latitude` DECIMAL(10,7),
    `contact_number` VARCHAR(255) NOT NULL UNIQUE,
    `owner_id` BIGINT UNSIGNED NULL,
    `media_id` BIGINT UNSIGNED NULL,
    `is_admin_verified` BOOLEAN DEFAULT FALSE,
    `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` DATETIME(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`owner_id`) REFERENCES `owners`(`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (`media_id`) REFERENCES `media`(`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


INSERT INTO `companies` (`name`, `description`, `address`, `longitude`, `latitude`, `contact_number`, `owner_id`, `media_id`, `is_admin_verified`)
VALUES
    ('Tech Solutions Inc.', 'A leading tech solutions provider.', '123 Tech Park, Silicon Valley, CA', -121.895, 37.339, '1234567890', 1, null, TRUE),
    ('Green Earth Landscaping', 'Professional landscaping services for residential and commercial spaces.', '45 Elm Street, Denver, CO', -104.9903, 39.7392, '2345678901', 2, null, FALSE),
    ('Urban Eats', 'Trendy urban restaurant offering a fusion menu.', '78 Downtown Ave, New York, NY', -74.006, 40.7128, '3456789012', 3, null, TRUE),
    ('Mountain Trek Adventures', 'Adventure tours and expeditions for outdoor enthusiasts.', '56 Trailblazer Lane, Boulder, CO', -105.2705, 40.015, '4567890123', NULL, NULL, FALSE),
    ('Creative Studios', 'Graphic design and marketing agency.', '89 Creative Road, Austin, TX', -97.7431, 30.2672, '5678901234', 4, null, TRUE);
