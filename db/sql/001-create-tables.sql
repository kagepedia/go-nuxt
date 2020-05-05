DROP DATABASE IF EXISTS sample_db;
CREATE DATABASE sample_db;
USE sample_db;

---- drop ----
DROP TABLE IF EXISTS `t_user`;

---- create ----
create table IF not exists `t_user`
(
 `id`               INT(11) AUTO_INCREMENT,
 `name`             VARCHAR(64) NOT NULL,
 `email`            VARCHAR(256) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
INSERT INTO t_user VALUES (1,'tanaka','tanaka@gmail.com');
INSERT INTO t_user VALUES (2,'kage','kage@gmail.com');
INSERT INTO t_user VALUES (3,'yamada','yamada@gmail.com');

---- drop ----
DROP TABLE IF EXISTS `t_task`;

---- create ----
create table IF not exists `t_task`
(
 `id`               INT(11) AUTO_INCREMENT,
 `task`             VARCHAR(256) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
INSERT INTO t_task VALUES (1,'tanaka');
INSERT INTO t_task VALUES (2,'kage');
INSERT INTO t_task VALUES (3,'yamada');
