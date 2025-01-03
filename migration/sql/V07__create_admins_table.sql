CREATE TABLE IF NOT EXISTS `admins` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `middle_name` VARCHAR(255),
    `permanent_address` TEXT,
    `description` TEXT,
    `birth_date` DATETIME(3),
    `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` DATETIME(3) DEFAULT NULL,
    `username` VARCHAR(255) NOT NULL UNIQUE,
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `contact_number` VARCHAR(15) NOT NULL UNIQUE,
    `is_email_verified` BOOLEAN DEFAULT FALSE,
    `is_contact_verified` BOOLEAN DEFAULT FALSE,
    `is_skip_verification` BOOLEAN DEFAULT FALSE,
    `status` VARCHAR(11) DEFAULT 'Pending',
    `media_id` BIGINT UNSIGNED NULL,
    `role_id` BIGINT UNSIGNED NULL,
    `gender_id` BIGINT UNSIGNED NULL,
    PRIMARY KEY (`id`),
    INDEX `idx_admin_deleted_at` (`deleted_at`),
    FOREIGN KEY (`media_id`) REFERENCES `media`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (`gender_id`) REFERENCES `genders`(`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
