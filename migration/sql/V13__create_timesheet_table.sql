CREATE TABLE IF NOT EXISTS `timesheets` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `employee_id` BIGINT UNSIGNED NOT NULL,
    `time_in` DATETIME(3) NOT NULL,
    `time_out` DATETIME(3) NOT NULL,
    `media_in_id` BIGINT UNSIGNED NULL,
    `media_out_id` BIGINT UNSIGNED NULL,
    `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`),
    INDEX `idx_timesheet_employee_id` (`employee_id`),
    FOREIGN KEY (`employee_id`) REFERENCES `employees`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`media_in_id`) REFERENCES `media`(`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (`media_out_id`) REFERENCES `media`(`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
