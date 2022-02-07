DROP TABLE IF EXISTS `user_profile`;
CREATE TABLE `nevad_db`.`user_profile` (
  `user_id` INT NOT NULL,
  `profile_id` INT NOT NULL,
  `status` INT(1) NULL DEFAULT '1',
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`, `profile_id`)
) ENGINE=InnoDB;