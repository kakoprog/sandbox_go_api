DROP TABLE IF EXISTS `task`;

CREATE TABLE `task` (
	`id`	INTEGER PRIMARY KEY,
	`name`	TEXT NOT NULL,
	`comment`	TEXT,
	`status`	INTEGER NOT NULL,
	`create_at`	TIMESTAMP DEFAULT (DATETIME('now','localtime')),
	`update_at`	TIMESTAMP DEFAULT (DATETIME('now','localtime'))
);

CREATE INDEX `idx_task_name` ON `task` (`name` );
CREATE INDEX `idx_task_create_at` ON `task` (`create_at` );
CREATE INDEX `idx_task_update_at` ON `task` (`update_at` );

DROP TABLE IF EXISTS `status`;

CREATE TABLE `status` (
	`id`	INTEGER PRIMARY KEY,
	`code`	TEXT NOT NULL UNIQUE,
	`name`	TEXT NOT NULL,
	`create_at`	TIMESTAMP DEFAULT (DATETIME('now','localtime')),
	`update_at`	TIMESTAMP DEFAULT (DATETIME('now','localtime'))
);

INSERT INTO `status` (code, name) VALUES('1', 'todo');
INSERT INTO `status` (code, name) VALUES('2', 'progress');
INSERT INTO `status` (code, name) VALUES('3', 'done');
