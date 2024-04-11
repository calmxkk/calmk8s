
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

DROP DATABASE IF EXISTS calmk8s;
create database calmk8s;
USE calmk8s;


--
-- 表的结构 `hg_admin_member`
--

CREATE TABLE IF NOT EXISTS `hg_admin_member` (
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
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb3 COMMENT='管理员_用户表';

INSERT INTO `hg_admin_member` (`id`, `dept_id`, `role_id`, `real_name`, `username`, `password_hash`, `salt`, `password_reset_token`, `integral`, `balance`, `avatar`, `sex`, `qq`, `email`, `mobile`, `birthday`, `city_id`, `address`, `pid`, `level`, `tree`, `invite_code`, `cash`, `last_active_at`, `remark`, `status`, `created_at`, `updated_at`) VALUES
(1, 100, 1, '孟帅', 'admin', 'a7c588fffeb2c1d99b29879d7fe97c78', '6541561', '', '89.00', '99290.78', 'http://bufanyun.cn-bj.ufileos.com/hotgo/attachment/2023-02-09/cqdq8er9nfkchdopav.png', 1, '133814250', '133814250@qq.com', '15303830571', '2016-04-16', 410172, '莲花街001号', 0, 1, '', '111', '{"name": "孟帅", "account": "15303830571", "payeeCode": "http://bufanyun.cn-bj.ufileos.com/hotgo/attachment/2023-02-09/cqdq8mqal5isvcb58g.jpg"}', '2023-11-25 11:21:38', NULL, 1, '2021-02-12 17:59:45', '2023-11-25 11:21:38'),
(3, 100, 2, '测试管理员', 'test', 'a7c588fffeb2c1d99b29879d7fe97c78', '6541561', '', '0.00', '4.00', 'http://alioss.qvnidaye.com//images/2021/03/12/image_1615529198_vMK4kwq2.jpg', 1, '', 'c@qq.cc', '15303888888', '2016-04-13', 371100, '大潮街道666号', 1, 2, 'tr_1 ', '222', NULL, '2023-08-11 17:42:10', '', 1, '2022-02-11 17:59:45', '2023-08-11 17:42:10'),
(8, 102, 200, '阿萌', 'ameng', '382df3b083a27886edb94e669a857c33', 'hfuUEb', '', '11.00', '4.22', '', 2, '', '', '', NULL, 0, '', 1, 2, 'tr_1 ', '333', NULL, '2023-08-02 16:31:59', '', 1, '2023-02-03 17:34:31', '2023-08-02 16:31:59'),
(9, 100, 206, '黄敏', 'test_finance', '151f5f6bb8b223fc7b589a32effb6f91', 'FhShzw', '', '0.00', '0.00', 'http://bufanyun.cn-bj.ufileos.com/hotgo/attachment/2023-02-09/cqdq8er9nfkchdopav.png', 1, '', '', '', NULL, 0, '', 1, 2, 'tr_1 ', '5jZUI3uWLfcj', NULL, NULL, '', 1, '2023-08-02 11:30:45', '2023-08-02 11:31:09'),
(11, 108, 201, '刘芳', 'abai', '5787c7a121190011fac8376b1d3e0396', 'puUFvx', '', '0.00', '0.00', 'http://bufanyun.cn-bj.ufileos.com/hotgo/attachment/2023-02-09/cqdq8er9nfkchdopav.png', 1, '', '', '', NULL, 0, '', 8, 3, 'tr_1 tr_8 ', 'SH5akjqInb2p', NULL, '2023-08-02 14:56:12', '', 1, '2023-08-02 14:24:50', '2023-08-02 14:56:12'),
(12, 109, 202, '李明', 'asong', '18d1e667e2e756c03186ff6d33b18fd4', 'ONYhgf', '', '0.00', '0.00', 'http://bufanyun.cn-bj.ufileos.com/hotgo/attachment/2023-02-09/cqdq8er9nfkchdopav.png', 1, '', '', '', NULL, 0, '', 11, 4, 'tr_1 tr_8 tr_11 ', 'pHffclXhgeg9', NULL, '2023-08-02 16:17:08', '', 1, '2023-08-02 14:50:49', '2023-08-02 16:17:08');