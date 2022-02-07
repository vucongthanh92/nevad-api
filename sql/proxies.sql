DROP TABLE IF EXISTS `proxies`;
CREATE TABLE `proxies` (
  `id` int NOT NULL AUTO_INCREMENT,
  `type` enum('http','socks') NOT NULL DEFAULT 'http',
  `ip` varchar(100),
  `port` varchar(100),
  `username` varchar(100),
  `password` varchar(100),
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE (`ip`)
) ENGINE=InnoDB;