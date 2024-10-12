CREATE TABLE IF NOT EXISTS `employees` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `middle_name` VARCHAR(255),
    `permanent_address` TEXT,
    `description` TEXT,
    `birthdate` DATE,
    `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` DATETIME(3) DEFAULT NULL,
    `username` VARCHAR(255) NOT NULL UNIQUE,
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `contact_number` VARCHAR(15) NOT NULL,
    `is_email_verified` TINYINT(1) DEFAULT 0,
    `is_contact_verified` TINYINT(1) DEFAULT 0,

   `media_id` BIGINT UNSIGNED NULL,
    PRIMARY KEY (`id`),
    INDEX `idx_employee_deleted_at` (`deleted_at`),
     FOREIGN KEY (`media_id`) REFERENCES `media`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `employees` 
    (`first_name`, `last_name`, `permanent_address`, `description`, `birthdate`, `username`, `email`, `password`, `media_id`, `created_at`, `updated_at`, `deleted_at`, `contact_number`)
VALUES 
    ('John', 'Doe', '123 Main St, Anytown, USA', 'A software engineer with 5 years of experience.', '1988-01-15', 'johndoe', 'johndoe@example.com', 'hashed_password_1', 1, NOW(), NOW(), NULL,'09194893088'),
    ('Jane', 'Smith', '456 Elm St, Anytown, USA', 'A project manager specializing in tech projects.', '1990-05-20', 'janesmith', 'janesmith@example.com', 'hashed_password_2', 2, NOW(), NOW(), NULL,'09194893088'),
    ('Michael', 'Johnson', '789 Oak St, Anytown, USA', 'A data analyst with a passion for statistics.', '1985-09-30', 'michaeljohnson', 'michaelj@example.com', 'hashed_password_3', 3, NOW(), NOW(), NULL,'09194893088');