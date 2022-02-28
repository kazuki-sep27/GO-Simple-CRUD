CREATE TABLE `accounts` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `owner` varchar(255) NOT NULL,
  `balance` bigint NOT NULL,
  `currency` varchar(3) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `entries` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `account_id` bigint NOT NULL,
  `amount` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `transfers` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `from_account_id` bigint NOT NULL,
  `to_account_id` bigint NOT NULL,
  `amount` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE `entries` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`from_account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`to_account_id`) REFERENCES `accounts` (`id`);

CREATE INDEX `accounts_index_0` ON `accounts` (`owner`);

CREATE INDEX `transfers_index_1` ON `transfers` (`from_account_id`);

CREATE INDEX `transfers_index_2` ON `transfers` (`to_account_id`);

CREATE INDEX `transfers_index_3` ON `transfers` (`from_account_id`, `to_account_id`);
