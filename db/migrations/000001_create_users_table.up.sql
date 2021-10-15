CREATE TABLE IF NOT EXISTS users(
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
);
