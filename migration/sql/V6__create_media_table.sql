-- V6__create_media_table.sql

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

-- Insert seed data into media table
INSERT INTO `media` (`file_name`, `file_size`, `file_type`, `storage_key`, `url`, `key`, `bucket_name`)
VALUES 
    ('example_image.jpg', 1048576, 'image/jpeg', 'unique_key_123', 'http://example.com/images/example_image.jpg', 'img_key_123', 'images_bucket'),
    ('example_video.mp4', 20971520, 'video/mp4', 'unique_key_456', 'http://example.com/videos/example_video.mp4', 'vid_key_456', 'videos_bucket'),
    ('example_document.pdf', 5242880, 'application/pdf', 'unique_key_789', 'http://example.com/docs/example_document.pdf', 'doc_key_789', 'docs_bucket');

