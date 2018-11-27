CREATE TABLE `blog_menu` (
  `id` varchar(40) NOT NULL COMMENT '主键',
  `name` varchar(20) NOT NULL COMMENT '菜单名称',
  `url` varchar(50) NOT NULL COMMENT '菜单url',
  `sort` int(11) NOT NULL COMMENT '菜单排序',
  `status` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'Y' COMMENT '是否启动，Y-启用，N-停用',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `blog_article` (
  `id` varchar(40) NOT NULL,
  `title` varchar(40) NOT NULL COMMENT '标题',
  `author` varchar(50) NOT NULL COMMENT '作者',
  `cover_img` varchar(255) DEFAULT NULL COMMENT '封面图片',
  `category_id` varchar(40) NOT NULL COMMENT '分类ID',
  `label` varchar(10) DEFAULT NULL COMMENT '标签',
  `Introduction` varchar(255) DEFAULT NULL COMMENT '简介',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章';

CREATE TABLE `blog_category` (
  `id` varchar(40) NOT NULL COMMENT '文章分类名',
  `name` varchar(40) NOT NULL COMMENT '文章分类名',
  `sort` int(255) NOT NULL DEFAULT '0' COMMENT '排序',
  `show_main` varchar(1) NOT NULL DEFAULT 'Y',
  `count` int(11) NOT NULL DEFAULT '0' COMMENT '文章数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `blog_comment` (
  `id` varchar(40) NOT NULL,
  `article_id` varchar(40) NOT NULL COMMENT '文章ID',
  `Commentator` varchar(40) NOT NULL COMMENT '评价人',
  `content` varchar(255) NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评价';

CREATE TABLE `blog_extend` (
  `id` varchar(40) NOT NULL,
  `article_id` varchar(40) NOT NULL COMMENT '文章ID',
  `read_count` int(255) NOT NULL DEFAULT '0' COMMENT '阅读数量',
  `like_count` int(255) NOT NULL DEFAULT '0',
  `pre_article` varchar(40) DEFAULT NULL,
  `nex_article` varchar(40) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='扩展';

CREATE TABLE `blog_picture` (
  `id` varchar(40) NOT NULL,
  `url` varchar(255) NOT NULL,
  `hot` int(255) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='图片';

CREATE TABLE `blog_my_desc` (
  `id` varchar(40) NOT NULL,
  `name` varchar(10) NOT NULL COMMENT '姓名',
  `my_desc` varchar(255) NOT NULL COMMENT '自我描述',
  `blog_desc` varchar(255) NOT NULL COMMENT '网站描述',
  `icp` varchar(255) NOT NULL COMMENT '网站备案号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `blog_leave_msg` (
  `id` varchar(40) NOT NULL,
  `name` varchar(50) NOT NULL COMMENT '姓名',
  `mail` varchar(255) NOT NULL COMMENT '邮箱',
  `content` varchar(255) NOT NULL COMMENT '留言内容',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `reply_content` varchar(255) DEFAULT NULL COMMENT '回复内容',
  `reply_time` datetime DEFAULT NULL COMMENT '回复时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='留言';

INSERT INTO `blog`.`blog_menu` (`id`, `name`, `url`, `sort`, `status`, `create_time`) VALUES (UUID(), '网站首页', 'index.html', 1, 'Y', NOW());
INSERT INTO `blog`.`blog_menu` (`id`, `name`, `url`, `sort`, `status`, `create_time`) VALUES (UUID(), '我的相册', 'share.html', 2, 'Y', NOW());
INSERT INTO `blog`.`blog_menu` (`id`, `name`, `url`, `sort`, `status`, `create_time`) VALUES (UUID(), '我的日记', 'list.html', 3, 'Y', NOW());
INSERT INTO `blog`.`blog_menu` (`id`, `name`, `url`, `sort`, `status`, `create_time`) VALUES (UUID(), '关于我', 'about.html', 4, 'Y', NOW());
INSERT INTO `blog`.`blog_menu` (`id`, `name`, `url`, `sort`, `status`, `create_time`) VALUES (UUID(), '留言', 'gbook.html', 5, 'Y', NOW());
INSERT INTO `blog`.`blog_menu` (`id`, `name`, `url`, `sort`, `status`, `create_time`) VALUES (UUID(), '内容页', 'info.html', 6, 'Y', NOW());
INSERT INTO `blog`.`blog_menu` (`id`, `name`, `url`, `sort`, `status`, `create_time`) VALUES (UUID(), '内容页', 'infopic.html', 7, 'Y', NOW());

INSERT INTO `blog`.`blog_my_desc` (`id`, `name`, `my_desc`, `blog_desc`, `icp`) VALUES (''a854d281-ed66-11e8-b460-00163e086cfc'', ''楚林少'', ''一个80后草根程序员，熟悉java相关、go语言等'', ''楚林少的博客'', ''湘ICP备18001708号-1'');

INSERT INTO `blog`.`blog_picture` (`id`, `url`, `hot`) VALUES (UUID(), 'https://img.126134.com/blog/img/1.jpg!1', 1);
INSERT INTO `blog`.`blog_picture` (`id`, `url`, `hot`) VALUES (UUID(), 'https://img.126134.com/blog/img/2.jpg!1', 2);
INSERT INTO `blog`.`blog_picture` (`id`, `url`, `hot`) VALUES (UUID(), 'https://img.126134.com/blog/img/3.jpg!1', 3);
INSERT INTO `blog`.`blog_picture` (`id`, `url`, `hot`) VALUES (UUID(), 'https://img.126134.com/blog/img/4.jpg!1', 4);
INSERT INTO `blog`.`blog_picture` (`id`, `url`, `hot`) VALUES (UUID(), 'https://img.126134.com/blog/img/5.jpg!1', 5);
INSERT INTO `blog`.`blog_picture` (`id`, `url`, `hot`) VALUES (UUID(), 'https://img.126134.com/blog/img/6.jpg!1', 6);
INSERT INTO `blog`.`blog_picture` (`id`, `url`, `hot`) VALUES (UUID(), 'https://img.126134.com/blog/img/7.jpg!1', 7);
INSERT INTO `blog`.`blog_picture` (`id`, `url`, `hot`) VALUES (UUID(), 'https://img.126134.com/blog/img/8.jpg!1', 8);
INSERT INTO `blog`.`blog_picture` (`id`, `url`, `hot`) VALUES (UUID(), 'https://img.126134.com/blog/img/1.jpg!1', 9);
INSERT INTO `blog`.`blog_picture` (`id`, `url`, `hot`) VALUES (UUID(), 'https://img.126134.com/blog/img/2.jpg!1', 10);
INSERT INTO `blog`.`blog_picture` (`id`, `url`, `hot`) VALUES (UUID(), 'https://img.126134.com/blog/img/3.jpg!1', 11);
INSERT INTO `blog`.`blog_picture` (`id`, `url`, `hot`) VALUES (UUID(), 'https://img.126134.com/blog/img/4.jpg!1', 12);
INSERT INTO `blog`.`blog_picture` (`id`, `url`, `hot`) VALUES (UUID(), 'https://img.126134.com/blog/img/5.jpg!1', 13);
INSERT INTO `blog`.`blog_picture` (`id`, `url`, `hot`) VALUES (UUID(), 'https://img.126134.com/blog/img/6.jpg!1', 14);
INSERT INTO `blog`.`blog_picture` (`id`, `url`, `hot`) VALUES (UUID(), 'https://img.126134.com/blog/img/7.jpg!1', 15);

INSERT INTO `blog`.`blog_category` (`id`, `name`, `sort`, `show_main`, `count`) VALUES ('123456', 'java基础', 1, 'Y', 1);
INSERT INTO `blog`.`blog_category` (`id`, `name`, `sort`, `show_main`, `count`) VALUES ('1234567', 'java-web', 2, 'Y', 1);
INSERT INTO `blog`.`blog_category` (`id`, `name`, `sort`, `show_main`, `count`) VALUES (UUID(), 'java并发', 3, 'Y', 2);
INSERT INTO `blog`.`blog_category` (`id`, `name`, `sort`, `show_main`, `count`) VALUES (UUID(), '分页式', 4, 'Y', 1);
INSERT INTO `blog`.`blog_category` (`id`, `name`, `sort`, `show_main`, `count`) VALUES (UUID(), 'mq', 5, 'Y', 1);
INSERT INTO `blog`.`blog_category` (`id`, `name`, `sort`, `show_main`, `count`) VALUES (UUID(), '生活', 6, 'Y', 1);

INSERT INTO `blog`.`blog_article` (
	`id`,
	`title`,
	`author`,
	`cover_img`,
	`category_id`,
	`label`,
	`Introduction`,
	`content`,
	`create_time`
)
VALUES
	(
		'123456',
		'大家好',
		'李四',
		null,
		'123456',
		'第一次见面',
		'大家好呀',
		'内容很少，给大家打个招呼',
		NOW()
	);
INSERT INTO `blog`.`blog_extend` (`id`, `article_id`, `read_count`, `like_count`, `pre_article`, `nex_article`) VALUES ('123456', '123456', 10, 20, NULL, NULL);

xorm reverse mysql root:123456@/blog?charset=utf8 templates/goxorm