DROP TABLE IF EXISTS `profiles`;
CREATE TABLE `nevad_db`.`profiles` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `profile_name` VARCHAR(50) NOT NULL,
  `user_agent` VARCHAR(250) NULL,
  `proxy_id` INT NULL,
  `status` INT(1) NULL DEFAULT '1',
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `profile_name_UNIQUE` (`profile_name` ASC)
) ENGINE=InnoDB;