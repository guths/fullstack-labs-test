CREATE TABLE `battles` (
  `id` integer,
  `monster_a` integer,
  `monster_b` integer,
  `winner` integer,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_battles_monster_a` FOREIGN KEY (`monster_a`) REFERENCES `monsters`(`id`),
  CONSTRAINT `fk_battles_monster_b` FOREIGN KEY (`monster_b`) REFERENCES `monsters`(`id`),
  CONSTRAINT `fk_battles_winner` FOREIGN KEY (`winner`) REFERENCES `monsters`(`id`)
);

CREATE TABLE `monsters` (
  `id` integer,
  `attack` integer NOT NULL,
  `defense` integer NOT NULL,
  `hp` integer NOT NULL,
  `speed` integer NOT NULL,
  `image_url` text NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime,
  PRIMARY KEY (`id`)
);

CREATE INDEX `idx_battles_deleted_at` ON `battles`(`deleted_at`);
CREATE INDEX `idx_monsters_deleted_at` ON `monsters`(`deleted_at`);