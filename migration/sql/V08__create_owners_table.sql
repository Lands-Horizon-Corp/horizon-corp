CREATE TABLE IF NOT EXISTS `owners` (
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
    `gender_id` BIGINT UNSIGNED NULL,
    PRIMARY KEY (`id`),
    INDEX `idx_owner_deleted_at` (`deleted_at`),
    FOREIGN KEY (`media_id`) REFERENCES `media`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`gender_id`) REFERENCES `genders`(`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;






INSERT INTO `owners` (`first_name`, `last_name`, `middle_name`, `permanent_address`, `description`, `birth_date`, `username`, `email`, `password`, `contact_number`, `is_email_verified`, `is_contact_verified`, `is_skip_verification`, `status`, `media_id`, `gender_id`)
VALUES
    ('John', 'Doe', 'Michael', '123 Main St, Los Angeles, CA', 'An experienced entrepreneur with a passion for technology.', '1985-06-15 00:00:00.000', 'johndoe', 'johndoe@example.com', 'JDJhJDEyJHBMNElVbzlpbjFaUWJPdVJvZVcuVHVSNExlRFc5ai9kYlJBRWJpbVd5SGs2Q2UyejlJUDRL', '1234567890', TRUE, TRUE, FALSE, 'Pending', NULL, 1),
    ('Jane', 'Smith', NULL, '456 Park Avenue, New York, NY', 'A creative business owner with expertise in design.', '1990-08-25 00:00:00.000', 'janesmith', 'janesmith@example.com', 'JDJhJDEyJHBMNElVbzlpbjFaUWJPdVJvZVcuVHVSNExlRFc5ai9kYlJBRWJpbVd5SGs2Q2UyejlJUDRL', '2345678901', TRUE, FALSE, FALSE, 'Pending', NULL, 2),
    ('Robert', 'Johnson', 'Lee', '789 Elm St, Chicago, IL', 'A professional with a background in real estate.', '1975-03-12 00:00:00.000', 'robertjohnson', 'robertj@example.com', 'JDJhJDEyJHBMNElVbzlpbjFaUWJPdVJvZVcuVHVSNExlRFc5ai9kYlJBRWJpbVd5SGs2Q2UyejlJUDRL', '3456789012', FALSE, TRUE, FALSE, 'Pending', NULL, 1),
    ('Emily', 'Davis', NULL, '321 Maple St, Seattle, WA', 'Business owner specializing in event planning.', '1995-11-03 00:00:00.000', 'emilydavis', 'emilyd@example.com', 'JDJhJDEyJHBMNElVbzlpbjFaUWJPdVJvZVcuVHVSNExlRFc5ai9kYlJBRWJpbVd5SGs2Q2UyejlJUDRL', '4567890123', FALSE, FALSE, TRUE, 'Pending', NULL, 2),
    ('Michael', 'Brown', 'James', '654 Pine St, Denver, CO', 'Entrepreneur with a focus on sustainable businesses.', '1980-01-01 00:00:00.000', 'michaelbrown', 'mbrown@example.com', 'JDJhJDEyJHBMNElVbzlpbjFaUWJPdVJvZVcuVHVSNExlRFc5ai9kYlJBRWJpbVd5SGs2Q2UyejlJUDRL', '5678901234', TRUE, TRUE, TRUE, 'Pending', NULL, 1);
