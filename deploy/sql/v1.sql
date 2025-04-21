CREATE DATABASE `go-zero-box` DEFAULT CHARACTER SET = `utf8mb4`;

CREATE TABLE `user` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `uuid` varchar(100) NOT NULL DEFAULT '',
    `account` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
    `name` varchar(100) NOT NULL DEFAULT '' COMMENT '姓名',
    `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
    `email` varchar(150) NOT NULL DEFAULT '' COMMENT '邮箱',
    `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
    `gender` tinyint(2) NOT NULL COMMENT '性别',
    `note` varchar(500) NOT NULL DEFAULT '' COMMENT '备注',
    `status` tinyint(3) unsigned NOT NULL DEFAULT '2' COMMENT '状态 1启动 2禁用',
    `is_delete` tinyint(3) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 1是 2否',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `username` (`account`),
    KEY `email` (`email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户管理/用户表';

INSERT INTO `user` (`id`, `uuid`, `account`, `name`, `mobile`, `email`, `password`, `gender`, `note`, `status`, `is_delete`, `created_at`, `updated_at`)
VALUES
    (1, '', 'admin', '测试管理员', '', 'demo@qq.com', '$2a$04$nnU0o/uNUjkWr9ZfqvqzAusjR5AaiJ.KrJz6lfyUgZQbgF7slYF6a', 1, '1', 1, 2, '2025-02-12 18:28:42', '2025-02-12 18:28:42');

CREATE TABLE `message` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` int(10) unsigned NOT NULL,
    `content` varchar(255) NOT NULL DEFAULT '',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='消息表';