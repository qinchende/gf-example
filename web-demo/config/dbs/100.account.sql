-- 公司员工
CREATE TABLE `sys_users`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `account`    varchar(32) NOT NULL DEFAULT '',
    `name`       varchar(64) NOT NULL DEFAULT '',
    `nickname`   varchar(64) NOT NULL DEFAULT '',
    `age`        int(11) NOT NULL DEFAULT '0',
    `created_at` datetime    NOT NULL,
    `updated_at` datetime    NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;