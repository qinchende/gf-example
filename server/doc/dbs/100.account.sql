-- 公司员工
CREATE TABLE `sys_user`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `mobile`     varchar(11) NOT NULL DEFAULT '',
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
insert into sys_user(mobile,account,name,nickname,age,email,status,created_at,updated_at) values('13466663333','admin','chende','qinchende',38,'cd@tl50.com',0,now(),now());
insert into sys_user(mobile,account,name,nickname,age,email,status,created_at,updated_at) values('13466663333','adm22','bmccde','qinchende',38,'cd@tl50.com',0,now(),now());
insert into sys_user(mobile,account,name,nickname,age,email,status,created_at,updated_at) values('13466663333','adm33','bmccde','qinchende',38,'cd@tl50.com',0,now(),now());

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

CREATE TABLE `sys_user_pass`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) NOT NULL,
    `hash_pass`  varchar(256) NOT NULL DEFAULT '',
    `status`     int(11) NOT NULL default 0,
    `created_at` datetime    NOT NULL,
    `updated_at` datetime    NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
alter table `sys_user_pass` add unique index idx_sys_user_pass_user_id(user_id);

insert into sys_user_pass(user_id,hash_pass,status,created_at,updated_at) values(1,'a877a79727be7e5461da5cc81e329701a40c84db',0,now(),now());
insert into sys_user_pass(user_id,hash_pass,status,created_at,updated_at) values(2,'32af6f66051df24d4bba08a36c9225dc0a2e36ad',0,now(),now());
insert into sys_user_pass(user_id,hash_pass,status,created_at,updated_at) values(3,'62a4dc06fe39e2723b6211d00320fd3260cd65fe',0,now(),now());
