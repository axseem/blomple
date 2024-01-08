-- Create "article" table
CREATE TABLE `article` (`id` integer NULL, `title` text NOT NULL, `body` text NOT NULL, `created_at` datetime NULL DEFAULT (CURRENT_TIMESTAMP), `updated_at` datetime NULL DEFAULT (CURRENT_TIMESTAMP), PRIMARY KEY (`id`));
