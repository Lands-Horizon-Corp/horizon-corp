CREATE TABLE IF NOT EXISTS `footsteps` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `admin_id` BIGINT UNSIGNED NULL,
    `employee_id` BIGINT UNSIGNED NULL,
    `owner_id` BIGINT UNSIGNED NULL,
    `member_id` BIGINT UNSIGNED NULL,
    `account_type` VARCHAR(11) NOT NULL,
    `description` VARCHAR(1000),
    `activity` VARCHAR(255) NOT NULL,
    `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` DATETIME(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`admin_id`) REFERENCES `admins`(`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (`employee_id`) REFERENCES `employees`(`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (`owner_id`) REFERENCES `owners`(`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (`member_id`) REFERENCES `members`(`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;