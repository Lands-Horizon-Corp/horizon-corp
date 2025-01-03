CREATE TABLE IF NOT EXISTS `genders` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (`id`),

    `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` DATETIME(3) DEFAULT NULL,
    `name` VARCHAR(255) NOT NULL UNIQUE,
    `description` TEXT,
    INDEX `idx_gender_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `genders` (`name`, `description`)
SELECT 'Male', 'A person who identifies as male, typically associated with masculinity.' 
WHERE NOT EXISTS (SELECT 1 FROM `genders` WHERE `name` = 'Male');

INSERT INTO `genders` (`name`, `description`)
SELECT 'Female', 'A person who identifies as female, typically associated with femininity.' 
WHERE NOT EXISTS (SELECT 1 FROM `genders` WHERE `name` = 'Female');

INSERT INTO `genders` (`name`, `description`)
SELECT 'Non-binary', 'A person who does not exclusively identify as male or female.' 
WHERE NOT EXISTS (SELECT 1 FROM `genders` WHERE `name` = 'Non-binary');

INSERT INTO `genders` (`name`, `description`)
SELECT 'Genderqueer', 'A person who identifies outside of the traditional gender binary.' 
WHERE NOT EXISTS (SELECT 1 FROM `genders` WHERE `name` = 'Genderqueer');

INSERT INTO `genders` (`name`, `description`)
SELECT 'Genderfluid', 'A person whose gender identity varies over time.' 
WHERE NOT EXISTS (SELECT 1 FROM `genders` WHERE `name` = 'Genderfluid');

INSERT INTO `genders` (`name`, `description`)
SELECT 'Agender', 'A person who identifies as having no gender or being gender-neutral.' 
WHERE NOT EXISTS (SELECT 1 FROM `genders` WHERE `name` = 'Agender');

INSERT INTO `genders` (`name`, `description`)
SELECT 'Bigender', 'A person who identifies as two genders, either simultaneously or at different times.' 
WHERE NOT EXISTS (SELECT 1 FROM `genders` WHERE `name` = 'Bigender');

INSERT INTO `genders` (`name`, `description`)
SELECT 'Two-Spirit', 'A term used by some Indigenous peoples to describe a person with both masculine and feminine spirits.' 
WHERE NOT EXISTS (SELECT 1 FROM `genders` WHERE `name` = 'Two-Spirit');

INSERT INTO `genders` (`name`, `description`)
SELECT 'Other', 'A gender identity that does not fall within traditional categories.' 
WHERE NOT EXISTS (SELECT 1 FROM `genders` WHERE `name` = 'Other');
