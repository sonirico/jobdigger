DROP TABLE IF EXISTS `site`;
CREATE TABLE `site` (
  `id` INTEGER AUTO_INCREMENT PRIMARY KEY,
  `name` TEXT NOT NULL,
  `domain` VARCHAR(255) NOT NULL,
  `link` VARCHAR(255) NOT NULL,
  `description` TEXT NOT NULL,
  `last_checked_at` INTEGER NULL
);

DROP TABLE IF EXISTS `offer`;
CREATE TABLE `offer` (
  `id` INTEGER AUTO_INCREMENT PRIMARY KEY,
  `title` TEXT NOT NULL,
  `permalink` VARCHAR(255) NOT NULL,
  `description` TEXT NOT NULL,
  `published_at`  INTEGER,
  `site_id` INTEGER NULL,
  FOREIGN KEY (`site_id`) REFERENCES `site`(`id`)
);

INSERT INTO `site` (`name`, `domain`, `link`, `description`)
VALUES (
  "Insituto para la formación y el empleo de Ponferrada",
  "https://empleo.ponferrada.org",
  "https://empleo.ponferrada.org/rss",
  "Insituto para la formación y el empleo de Ponferrada"
)