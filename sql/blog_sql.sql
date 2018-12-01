/*
Navicat MySQL Data Transfer

Source Server         : 120.76.157.34_3306
Source Server Version : 50636
Source Host           : 120.76.157.34:3306
Source Database       : blog

Target Server Type    : MYSQL
Target Server Version : 50636
File Encoding         : 65001

Date: 2018-12-01 17:41:27
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
  `id` varchar(40) NOT NULL,
  `title` varchar(40) NOT NULL COMMENT '标题',
  `author` varchar(50) NOT NULL COMMENT '作者',
  `cover_img_url` varchar(255) DEFAULT NULL COMMENT '封面图片',
  `category_id` varchar(40) NOT NULL COMMENT '分类ID',
  `label` varchar(50) DEFAULT NULL COMMENT '标签',
  `Introduction` varchar(256) DEFAULT NULL COMMENT '简介',
  `content` text NOT NULL COMMENT '内容',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章';

-- ----------------------------
-- Records of blog_article
-- ----------------------------
INSERT INTO `blog_article` VALUES ('1121212', '个人博客，属于我的小世界！', '李五', 'https://img.126134.com/blog/img/java.jpg!1', '123456', null, '个人博客，用来做什么？我刚开始就把它当做一个我吐槽心情的地方，也就相当于一个网络记事本，写上一些关于自己生活工作中的小情小事，也会放上一些照片，音乐。每天工作回家后就能访问自己的网站，一边听着音乐，一边写写文章。', '本文很长，记录了我博客建站初到现在的过程，还有我从毕业到现在的一个状态，感谢您的阅读，如果你还是学生，也许你能从此文中，找到我们曾经相似的地方。如果你已经工作，有自己的博客，我想，你并没有忘记当初建立个人博客的初衷吧！\r\n\r\n我的个人博客已经建站有8年的时间了，对它的热爱，一直都是只增未减。回想大学读书的那几年，那会儿非常流行QQ空间，我们寝室的室友还经常邀约去学校的网吧做自己的空间。系里有个男生，空间做得非常漂亮，什么悬浮，开场动画，音乐播放器，我们女生羡慕得不得了。还邀约他跟我们一起去通宵弄空间，网上可以找到很多免费的flash资源，还有音乐，那也是第一次接触js，知道在浏览器输入一个地址，修改一下数据，就能调用一些背景出来。然后把自己QQ空间弄得漂漂亮亮的，经常邀约室友来互踩。我记得08年地震，第二天晚上，我们寝室的几个人还淡定的在寝室装扮空间呢！\r\n\r\n\r\n\r\n后来空间收费项目也多了，官方漏洞也修复了，加上临近毕业，又要忙着做毕业设计，就没再打理QQ空间。我知道现在的九零后，零零后，你们肯定没看过《一帘幽梦》，那会儿我也是疯狂追剧，喜欢上紫菱，喜欢上她的网站。想看看她的小世界，而我更想学着做一个她那样的网站。那会儿还天真的以为网上真的有她的网站，百度搜了好些天也没有。\r\n\r\n要毕业的时候，要交作业了，感觉自己什么都没学会。室友拉着我们去看了她同学做的网站，我们一个个佩服得五体投地，甚至觉得太不可思议了。有难度，又怕自己不会。老师教我们怎么布局，怎么做，并没有教我们右键保存网页。不知道是谁先会了这绝技，然后我们一个个又跟打了鸡血似的，疯狂在网上找网页，右键另存为。然后一个个修改文字，图片。仿佛又回到了那会儿做QQ空间那个时候。拿着copy来，并且精心修改的作品，递交了毕业设计，顺利结业。那会儿还是很蒙，一种云里雾里的感觉，竟有种不知道自己到底是会还是不会的感觉，也就是大家常说的毕业迷茫期。\r\n\r\n\r\n\r\n工作后进入社会，出去谈业务，遇到一个对网页设计超级感兴趣的人，聊了一下午都还不够，他是完全自学的，做了一个首页宣传他们的产品。他眼里的我就是专业的，总是请教我一些问题。其实我内心特羡慕人家，每次问我，我也似懂非懂的跟人家解决问题，但我还是经常靠百度来搜索他要的答案。他身上那种好学好问的那股劲儿，也成为我迫切想拥有自己的个人博客的一个重要原因。\r\n\r\n做博客不是说做就做的，很多东西我都不懂，也不知道要购买域名还有空间。前期要做的工作还是很多。幸好张园同学，也是我实习期的同事，他会这些，教我网上找免费的虚拟空间，然后就是把自己做的页面上传进去，还给了有一个地址，然后就能访问了。那是第一次接触，也了解了整个网站的制作过程，只可惜买域名还有空间需要费用，还在实习期的我，想想也就算了。虚拟空间毕竟是免费的，没多长时间，做过的网页就不能访问了，又得重新注册，重新上传。\r\n\r\n等自己有一些资金和技术后，我开始买域名和空间。从一开始，我就没想过只是练练手，或者用一段时间就行了。我会一直用下去，所以精心挑选了域名和空间。这些年除了域名没有更换外，后台程序由asp换成了php，空间从西部数据换成万网，也就是现在的阿里云。一步步升级，就想把最好的一面呈现给大家。很多人问我网站速度怎么访问那么快，其实一是网站程序，页面最好是静态页面。每次我写的代码的时候，我都在琢磨怎么减少代码，减少使用div和图片，让html结构简单化，而又不失美观。所以，网站也改版了有好几次。二是空间还有带宽。这个其实很重要，现在备案跟以前比，快很多了，快的话一星期，慢的话顶多二十天。所以，不是因为特着急的话，还是用国内空间。关于国外空间，其实现在阿里云的香港虚拟主机也不错，访问还是上传都比以前好很多了。用它的小伙伴也挺多。延伸阅读 《我的个人博客之——阿里云空间选择》\r\n\r\n\r\n\r\n个人博客，用来做什么？我刚开始就把它当做一个我吐槽心情的地方，也就相当于一个网络记事本，写上一些关于自己生活工作中的小情小事，也会放上一些照片，音乐。每天工作回家后就能访问自己的网站，一边听着音乐，一边写写文章。虽然访问人数少，我也不在乎这些，个人博客就好像我自己的一个日记本，来窥探的人多了，反而不自在了。因为博客，我还是交了不少朋友。我网站之前的logo就是网友给我设计的。\r\n\r\n后来，也是因为同学说百度搜不着我的个人博客，我才开始琢磨怎么把自己的博客排名提上去，放上去一些自己做过的模板，提供免费下载。页面以前也只有一个首页，当初我就是想着设计一个首页就行了，真正想学的人，一个页面足够了，其他页面，都靠自己设计。现在不同啦，近期版本我提供的模板都很全，实在是架不住人多都来问我怎么只有一个页面，虽然我已经在下载说明里面说得很清楚了。也就是提供了这样一些资源，我的个人博客也不再是一个只关乎自己的一个平台。曝光率越来越高，光天化日之下，哪敢放自己照片，写自己心情啊，于是乎，继续这样的模式吧，把自己写的模板放到网站上去。乐于分享，是能结交到很多朋友的。\r\n\r\n我的个人博客，在今年三月份之前是没有提供后台安装，调试这些服务的。其实一直都有人找我做个人博客，因为调试时间还有修改的问题，从上传到安装到修改调试，一般要四五个小时，我怕收费高，所以一直都是拒绝的。自从推出来《心蓝时间轴》后，主动找我做网站的也多了，他们都说是我粉丝，喜欢我的模板，所以我在模板下面写了付费说明，我的个人博客也开始有了付费项目。\r\n\r\n\r\n\r\n我接到的第一单，是一位父亲，想给自家的宝贝建立一个博客。他说他想买一个域名，谈了好几天，最终花了8000多购买到，我挺敬佩他的，对孩子的爱，毫不吝啬。域名固然重要，其实我想说坚持用博客来记录宝贝成长点滴，才是最有价值的。第二单是一个阅读网站，从加他到跟他聊，看他朋友圈，很有才气的一个人，也相当有爱心，他的网站大部分是分享给一些爱好阅读的长者。还有些就是技术类的站长，做资讯，做旅游，做推广的。通常，从选的模板就能看出来网站类型。《心蓝时间轴》偏个人，《绅士》，《格调》偏技术，资讯。《清雅》，《水墨古典》偏文艺。但也有对我防备心的，比如说付款方式，能不能走淘宝，这些我也能理解，但是有些吧，从一开始聊，感觉就不太好，理应就认为不应该收钱，收钱干嘛，到处都是免费的，听到这些，我也慢慢解释。其实他们并不了解我，了解我的网站，只要是从我博客来的，关注过一段时间的，都知道青姐的人品是咋样的。用“心塞”这个词，毫不为过吧。\r\n\r\n我做过的网站，每过一段时间，我都会一个个点击看看进展怎么样，个人博客，不像真实的面对面聊，更多的是文字的交流，不得不说有时候文字传达的信息更能了解一个人，甚至有心灵共鸣。我想我还会再重新做一个仅仅属于自己的个人博客，如果你也想要做一个博客，用来记录自己的家庭，工作，生活，或者讨女友欢心，不妨现在就开始吧！', '2018-11-25 19:05:22');
INSERT INTO `blog_article` VALUES ('123456', '大家好', '李四', 'https://img.126134.com/blog/img/java.jpg!1', '123456', '第一次见面;小世界;大舞台', '大家好呀11111', '内容很少，给大家打个招呼', '2018-11-21 16:56:19');
INSERT INTO `blog_article` VALUES ('a8e5e98e-ed66-11e8-b460-00163e086cfc', '大家好11', '李四', 'https://img.126134.com/blog/img/golang.jpg!1', '123456', '第一次见面;小世界;大舞台', '大家好呀，这是一个简介', '内容很少，给大家打个招呼', '2018-11-21 16:23:01');

-- ----------------------------
-- Table structure for blog_category
-- ----------------------------
DROP TABLE IF EXISTS `blog_category`;
CREATE TABLE `blog_category` (
  `id` varchar(40) NOT NULL COMMENT '文章分类名',
  `name` varchar(40) NOT NULL COMMENT '文章分类名',
  `sort` int(255) NOT NULL DEFAULT '0' COMMENT '排序',
  `show_main` varchar(1) NOT NULL DEFAULT 'Y',
  `count` int(11) NOT NULL DEFAULT '0' COMMENT '文章数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog_category
-- ----------------------------
INSERT INTO `blog_category` VALUES ('123456', 'java基础', '1', 'Y', '1');
INSERT INTO `blog_category` VALUES ('1234567', 'java-web', '2', 'Y', '1');
INSERT INTO `blog_category` VALUES ('a8cb68b5-ed66-11e8-b460-00163e086cfc', 'java并发', '3', 'Y', '2');
INSERT INTO `blog_category` VALUES ('a8d1e3fe-ed66-11e8-b460-00163e086cfc', '分页式', '4', 'Y', '1');
INSERT INTO `blog_category` VALUES ('a8d82cf8-ed66-11e8-b460-00163e086cfc', 'mq', '5', 'Y', '1');
INSERT INTO `blog_category` VALUES ('a8de69ef-ed66-11e8-b460-00163e086cfc', '生活', '6', 'Y', '1');

-- ----------------------------
-- Table structure for blog_comment
-- ----------------------------
DROP TABLE IF EXISTS `blog_comment`;
CREATE TABLE `blog_comment` (
  `id` varchar(40) NOT NULL,
  `article_id` varchar(40) NOT NULL COMMENT '文章ID',
  `commentator` varchar(40) NOT NULL COMMENT '评价人',
  `content` varchar(255) NOT NULL COMMENT '评价内容',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评价';

-- ----------------------------
-- Records of blog_comment
-- ----------------------------
INSERT INTO `blog_comment` VALUES ('dfsafsadfsa', 'a8e5e98e-ed66-11e8-b460-00163e086cfc', '张三大', '非常不错的文章', '2018-12-01 15:39:10');
INSERT INTO `blog_comment` VALUES ('dfsafsadfsa1', 'a8e5e98e-ed66-11e8-b460-00163e086cfc', '李四', '很好，很喜欢', '2018-12-01 15:39:10');
INSERT INTO `blog_comment` VALUES ('dfsafsadfsa11', 'a8e5e98e-ed66-11e8-b460-00163e086cfc', '李四', '很好，很喜欢', '2018-12-01 15:39:10');

-- ----------------------------
-- Table structure for blog_extend
-- ----------------------------
DROP TABLE IF EXISTS `blog_extend`;
CREATE TABLE `blog_extend` (
  `id` varchar(40) NOT NULL,
  `article_id` varchar(40) NOT NULL COMMENT '文章ID',
  `read_count` int(255) NOT NULL DEFAULT '0' COMMENT '阅读数量',
  `like_count` int(255) NOT NULL DEFAULT '0',
  `pre_article` varchar(40) DEFAULT NULL,
  `nex_article` varchar(40) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uq_article_id` (`article_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='扩展';

-- ----------------------------
-- Records of blog_extend
-- ----------------------------
INSERT INTO `blog_extend` VALUES ('123456', 'a8e5e98e-ed66-11e8-b460-00163e086cfc', '123', '6', '1121212', '123456');
INSERT INTO `blog_extend` VALUES ('1234561', '123456', '71', '6', 'a8e5e98e-ed66-11e8-b460-00163e086cfc', '');
INSERT INTO `blog_extend` VALUES ('12345612', '1121212', '81', '6', '', 'a8e5e98e-ed66-11e8-b460-00163e086cfc');

-- ----------------------------
-- Table structure for blog_leave_msg
-- ----------------------------
DROP TABLE IF EXISTS `blog_leave_msg`;
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

-- ----------------------------
-- Records of blog_leave_msg
-- ----------------------------
INSERT INTO `blog_leave_msg` VALUES ('1', '夜月归途', 'xxx', '从青姐朋友圈分享的文章《我为什么要做个人网站》过来的，自习看了下你的网站非常不错，看的出来你一直在坚持!', '2018-11-21 17:32:40', '感谢捧场啊！看了你的网站，有两个月没更新了哦~', null);
INSERT INTO `blog_leave_msg` VALUES ('2', '周宏', 'xxx', '读《从今日起，我永久卸载今日头条》有感。正如作者所说，这个APP抓住了很多人性的特点，在简单、重复、爽这三点上做到了极致。但是我认为永久卸载这个想法比较荒诞，任何东西你只要有自控力，就能将它为我所用。曾经一度我也是刷头条根本停不下来', '2018-11-25 17:03:15', '嗯，我也是自制力有限，删除头条就是矫枉过正而已，这个因人而异，不强求他人，也不想标题党。', null);
INSERT INTO `blog_leave_msg` VALUES ('4db625494c6c4d21a0ba99b9b6925996', '楚林少', 'fdafasfsa@mail.com', '挺不错的，我挺喜欢的。。。。。。。。。。', '2018-11-25 18:36:34', '', null);
INSERT INTO `blog_leave_msg` VALUES ('6ffe7ee46af64849b469004c53d141fc', '大V1', 'fdafsadf.@.fadsf', '关注很久了，希望能学习下', '2018-11-25 18:36:58', '', null);

-- ----------------------------
-- Table structure for blog_menu
-- ----------------------------
DROP TABLE IF EXISTS `blog_menu`;
CREATE TABLE `blog_menu` (
  `id` varchar(40) NOT NULL COMMENT '主键',
  `name` varchar(20) NOT NULL COMMENT '菜单名称',
  `url` varchar(50) NOT NULL COMMENT '菜单url',
  `sort` int(11) NOT NULL COMMENT '菜单排序',
  `status` varchar(1) NOT NULL DEFAULT 'Y' COMMENT '是否启动，Y-启用，N-停用',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog_menu
-- ----------------------------
INSERT INTO `blog_menu` VALUES ('a8284cd7-ed66-11e8-b460-00163e086cfc', '网站首页', 'index.html', '1', 'Y', '2018-11-21 16:22:59');
INSERT INTO `blog_menu` VALUES ('a82eb83e-ed66-11e8-b460-00163e086cfc', '我的相册', 'share.html', '2', 'Y', '2018-11-21 16:22:59');
INSERT INTO `blog_menu` VALUES ('a834e08e-ed66-11e8-b460-00163e086cfc', '我的日记', 'list.html', '3', 'Y', '2018-11-21 16:22:59');
INSERT INTO `blog_menu` VALUES ('a83b6645-ed66-11e8-b460-00163e086cfc', '关于我', 'about.html', '4', 'Y', '2018-11-21 16:22:59');
INSERT INTO `blog_menu` VALUES ('a8419f92-ed66-11e8-b460-00163e086cfc', '留言', 'gbook.html', '5', 'Y', '2018-11-21 16:23:00');
INSERT INTO `blog_menu` VALUES ('a8482252-ed66-11e8-b460-00163e086cfc', '内容页', 'info.html', '6', 'Y', '2018-11-21 16:23:00');
INSERT INTO `blog_menu` VALUES ('a84e7040-ed66-11e8-b460-00163e086cfc', '内容页', 'infopic.html', '7', 'Y', '2018-11-21 16:23:00');

-- ----------------------------
-- Table structure for blog_my_desc
-- ----------------------------
DROP TABLE IF EXISTS `blog_my_desc`;
CREATE TABLE `blog_my_desc` (
  `id` varchar(40) NOT NULL,
  `name` varchar(10) NOT NULL COMMENT '姓名',
  `my_desc` varchar(255) NOT NULL COMMENT '自我描述',
  `blog_desc` varchar(100) NOT NULL COMMENT '网站描述',
  `icp` varchar(50) NOT NULL COMMENT '网站备案号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog_my_desc
-- ----------------------------
INSERT INTO `blog_my_desc` VALUES ('a854d281-ed66-11e8-b460-00163e086cfc', '楚林少', '一个80后草根程序员，熟悉java相关、go语言等', '楚林少的博客', '湘ICP备18001708号-1 ©2018-lijun');

-- ----------------------------
-- Table structure for blog_picture
-- ----------------------------
DROP TABLE IF EXISTS `blog_picture`;
CREATE TABLE `blog_picture` (
  `id` varchar(40) NOT NULL,
  `url` varchar(255) NOT NULL,
  `hot` int(255) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='图片';

-- ----------------------------
-- Records of blog_picture
-- ----------------------------
INSERT INTO `blog_picture` VALUES ('14c9fb9e-ed69-11e8-b460-00163e086cfc', 'https://img.126134.com/blog/img/1.jpg!1', '1');
INSERT INTO `blog_picture` VALUES ('14d04215-ed69-11e8-b460-00163e086cfc', 'https://img.126134.com/blog/img/2.jpg!1', '2');
INSERT INTO `blog_picture` VALUES ('14d661f8-ed69-11e8-b460-00163e086cfc', 'https://img.126134.com/blog/img/3.jpg!1', '3');
INSERT INTO `blog_picture` VALUES ('14deb575-ed69-11e8-b460-00163e086cfc', 'https://img.126134.com/blog/img/4.jpg!1', '4');
INSERT INTO `blog_picture` VALUES ('14e5674d-ed69-11e8-b460-00163e086cfc', 'https://img.126134.com/blog/img/5.jpg!1', '5');
INSERT INTO `blog_picture` VALUES ('14ed6a59-ed69-11e8-b460-00163e086cfc', 'https://img.126134.com/blog/img/6.jpg!1', '6');
INSERT INTO `blog_picture` VALUES ('14f4f135-ed69-11e8-b460-00163e086cfc', 'https://img.126134.com/blog/img/7.jpg!1', '7');
INSERT INTO `blog_picture` VALUES ('14fca428-ed69-11e8-b460-00163e086cfc', 'https://img.126134.com/blog/img/8.jpg!1', '8');
INSERT INTO `blog_picture` VALUES ('1504af94-ed69-11e8-b460-00163e086cfc', 'https://img.126134.com/blog/img/1.jpg!1', '9');
INSERT INTO `blog_picture` VALUES ('150ce260-ed69-11e8-b460-00163e086cfc', 'https://img.126134.com/blog/img/2.jpg!1', '10');
INSERT INTO `blog_picture` VALUES ('15166887-ed69-11e8-b460-00163e086cfc', 'https://img.126134.com/blog/img/3.jpg!1', '11');
INSERT INTO `blog_picture` VALUES ('152091ca-ed69-11e8-b460-00163e086cfc', 'https://img.126134.com/blog/img/4.jpg!1', '12');
INSERT INTO `blog_picture` VALUES ('15286288-ed69-11e8-b460-00163e086cfc', 'https://img.126134.com/blog/img/5.jpg!1', '13');
INSERT INTO `blog_picture` VALUES ('152f429a-ed69-11e8-b460-00163e086cfc', 'https://img.126134.com/blog/img/6.jpg!1', '14');
INSERT INTO `blog_picture` VALUES ('1536cbe1-ed69-11e8-b460-00163e086cfc', 'https://img.126134.com/blog/img/7.jpg!1', '15');
