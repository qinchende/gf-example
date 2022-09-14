-- 公司员工
CREATE TABLE `sys_user`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `account`    varchar(32) NOT NULL DEFAULT '',
    `name`       varchar(64) NOT NULL DEFAULT '',
    `nickname`   varchar(64) NOT NULL DEFAULT '',
    `age`        int(11) NOT NULL DEFAULT 0,
    `email`      varchar(64) NULL,
    `status`     int(11) NOT NULL default 0,
    `created_at` datetime    NOT NULL,
    `updated_at` datetime    NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `sys_user_gm_info`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) NOT NULL,
    `is_open`    int(11) NOT NULL default 0,
    `open_time`  datetime NULL,
    `status`     int(11) NOT NULL default 0,
    `created_at` datetime    NOT NULL,
    `updated_at` datetime    NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
alter table `sys_user_gm_info` add unique index idx_sys_user_gm_info_user_id(user_id);

insert into sys_user_gm_info (user_id,is_open,open_time,status,created_at,updated_at) values(11,0,null,0,now(),now());