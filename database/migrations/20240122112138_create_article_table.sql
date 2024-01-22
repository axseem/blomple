-- Create "article" table
CREATE TABLE `article` (`id` integer NULL, `title` text NOT NULL, `content` text NOT NULL, `created_at` datetime NULL DEFAULT (CURRENT_TIMESTAMP), `updated_at` datetime NULL DEFAULT (CURRENT_TIMESTAMP), PRIMARY KEY (`id`));
