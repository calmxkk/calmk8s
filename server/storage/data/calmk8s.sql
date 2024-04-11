
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

DROP DATABASE IF EXISTS calmk8s;
create database calmk8s;
USE calmk8s;


--
-- 表的结构 `ck8s_user`
--

CREATE TABLE IF NOT EXISTS `ck8s_user` (
  `id` bigint NOT NULL COMMENT '主键',
  `role` bigint NOT NULL COMMENT '权限等级',
  `real_name` varchar(32) DEFAULT '' COMMENT '真实姓名',
  `username` varchar(20) NOT NULL DEFAULT '' COMMENT '帐号',
  `password_hash` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` char(16) NOT NULL COMMENT '密码盐',
  `password_reset_token` varchar(150) DEFAULT '' COMMENT '密码重置令牌',
  `avatar` char(150) DEFAULT '' COMMENT '头像',
  `sex` tinyint(1) DEFAULT '1' COMMENT '性别',
  `qq` varchar(20) DEFAULT '' COMMENT 'qq',
  `email` varchar(60) DEFAULT '' COMMENT '邮箱',
  `mobile` varchar(20) DEFAULT '' COMMENT '手机号码',
  `birthday` date DEFAULT NULL COMMENT '生日',
  `address` varchar(100) DEFAULT '' COMMENT '联系地址',
  `last_active_at` datetime DEFAULT NULL COMMENT '最后活跃时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间'
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb3 COMMENT='用户表';

INSERT INTO `ck8s_user` (`id`, `role`, `real_name`, `username`, `password_hash`, `salt`, `password_reset_token`, `avatar`, `sex`, `qq`, `email`, `mobile`, `birthday`, `address`, `last_active_at`, `remark`, `status`, `created_at`, `updated_at`) VALUES  
(1, 100, '张三', 'zhangsan', '5f4dcc3b5aa765d61d8327deb882cf99', '1a2b3c4d5e6f7g8h', '', 'http://example.com/avatar1.jpg', 1, '12345678', 'zhangsan@example.com', '13800138000', '1990-01-01', '北京市朝阳区', NOW(), '这是张三', 1, NOW(), NOW()),  
(2, 90, '李四', 'lisi', 'e10adc3949ba59abbe56e057f20f883e', '9h8i7g6f5e4d3c2b', '', 'http://example.com/avatar2.jpg', 0, '87654321', 'lisi@example.com', '13900139000', '1995-05-05', '上海市浦东新区', NOW(), '这是李四', 1, NOW(), NOW()),  
(3, 80, '王五', 'wangwu', 'd033e22ae348aeb5660fc2140aec3585', '7c6b5a49382716fa', '', 'http://example.com/avatar3.jpg', 1, '55555555', 'wangwu@example.com', '13700137000', '2000-10-10', '深圳市南山区', NOW(), '这是王五', 1, NOW(), NOW());


--
-- 表的结构 `ck8s_cluster`
--

CREATE TABLE IF NOT EXISTS `ck8s_cluster` (
  `id` bigint NOT NULL COMMENT '主键',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '集群名称',
  `configstr` varchar(65535) DEFAULT NULL COMMENT '配置文件',
  
  `salt` char(16) NOT NULL COMMENT '密码盐',
  `password_reset_token` varchar(150) DEFAULT '' COMMENT '密码重置令牌',
  `avatar` char(150) DEFAULT '' COMMENT '头像',
  `sex` tinyint(1) DEFAULT '1' COMMENT '性别',
  `qq` varchar(20) DEFAULT '' COMMENT 'qq',
  `email` varchar(60) DEFAULT '' COMMENT '邮箱',
  `mobile` varchar(20) DEFAULT '' COMMENT '手机号码',
  `birthday` date DEFAULT NULL COMMENT '生日',
  `address` varchar(100) DEFAULT '' COMMENT '联系地址',
  `last_active_at` datetime DEFAULT NULL COMMENT '最后活跃时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `created_by` datetime DEFAULT NULL COMMENT '创建用户',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间'
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb3 COMMENT='集群表';