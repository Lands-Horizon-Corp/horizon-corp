CREATE TABLE IF NOT EXISTS `media` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` DATETIME(3) DEFAULT NULL,
    `file_name` VARCHAR(255) NOT NULL,
    `file_size` BIGINT NOT NULL,
    `file_type` VARCHAR(50) NOT NULL,
    `storage_key` VARCHAR(255) NOT NULL UNIQUE,
    `url` VARCHAR(255) NOT NULL,
    `key` VARCHAR(255),
    `bucket_name` VARCHAR(255),
    PRIMARY KEY (`id`),
    INDEX `idx_media_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;