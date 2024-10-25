CREATE TABLE IF NOT EXISTS `roles` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` DATETIME(3) DEFAULT NULL,
    `name` VARCHAR(255) UNIQUE NOT NULL,
    `description` TEXT,
    `api_key` VARCHAR(255) UNIQUE NOT NULL,
    `read_role` BOOLEAN DEFAULT FALSE,
    `write_role` BOOLEAN DEFAULT FALSE,
    `update_role` BOOLEAN DEFAULT FALSE,
    `delete_role` BOOLEAN DEFAULT FALSE,
    `read_error_details` BOOLEAN DEFAULT FALSE,
    `write_error_details` BOOLEAN DEFAULT FALSE,
    `update_error_details` BOOLEAN DEFAULT FALSE,
    `delete_error_details` BOOLEAN DEFAULT FALSE,
    `read_gender` BOOLEAN DEFAULT FALSE,
    `write_gender` BOOLEAN DEFAULT FALSE,
    `update_gender` BOOLEAN DEFAULT FALSE,
    `delete_gender` BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (`id`),
    INDEX `idx_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `roles` (
    `name`, `description`, `api_key`, 
    `read_role`, `write_role`, `update_role`, `delete_role`,
    `read_error_details`, `write_error_details`, `update_error_details`, `delete_error_details`,
    `read_gender`, `write_gender`, `update_gender`, `delete_gender`
) VALUES 
    (
        'Administrator', 
        'Full access to all resources and permissions.', 
        'ADMIN_API_KEY_123456',
        TRUE, TRUE, TRUE, TRUE,
        TRUE, TRUE, TRUE, TRUE,
        TRUE, TRUE, TRUE, TRUE
    ),
    (
        'Editor', 
        'Can read and modify most resources but has limited administrative privileges.', 
        'EDITOR_API_KEY_abcdef',
        TRUE, TRUE, TRUE, FALSE,
        TRUE, TRUE, TRUE, FALSE,
        TRUE, TRUE, FALSE, FALSE
    ),
    (
        'Viewer', 
        'Can only read resources without any modification permissions.', 
        'VIEWER_API_KEY_654321',
        TRUE, FALSE, FALSE, FALSE,
        TRUE, FALSE, FALSE, FALSE,
        TRUE, FALSE, FALSE, FALSE
    ),
    (
        'ErrorManager', 
        'Specialized role for managing error details.', 
        'ERRMGR_API_KEY_zxy123',
        FALSE, FALSE, FALSE, FALSE,
        TRUE, TRUE, TRUE, TRUE,
        FALSE, FALSE, FALSE, FALSE
    );
