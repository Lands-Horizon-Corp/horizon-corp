CREATE TABLE IF NOT EXISTS `feedbacks` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` DATETIME(3) DEFAULT NULL,
    `email` VARCHAR(255) NOT NULL,
    `description` TEXT NOT NULL,
    `feedback_type` ENUM('bug', 'feature', 'general') NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `idx_feedback_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;