DROP TABLE IF EXISTS `team_profile`;
CREATE TABLE `nevad_db`.`team_profile` (
  `team_id` INT NOT NULL,
  `profile_id` INT NOT NULL,
  `status` INT(1) NULL DEFAULT '1',
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`team_id`, `profile_id`)
) ENGINE=InnoDB;