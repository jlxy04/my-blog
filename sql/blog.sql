CREATE TABLE `blog_menu` (
  `id` varchar(40) NOT NULL COMMENT '主键',
  `name` varchar(20) NOT NULL COMMENT '菜单名称',
  `url` varchar(50) NOT NULL COMMENT '菜单url',
  `sort` int(11) NOT NULL COMMENT '菜单排序',
  `status` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'Y' COMMENT '是否启动，Y-启用，N-停用',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

