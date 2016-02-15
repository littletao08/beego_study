
# ************************************************************
# Sequel Pro SQL dump
# Version 4096
#
# http://www.sequelpro.com/
# http://code.google.com/p/sequel-pro/
#
# Host: 127.0.0.1 (MySQL 5.1.63)
# Database: beego_study
# Generation Time: 2015-11-13 07:23:06 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

drop database beego_study;

create database beego_study default character set utf8 collate utf8_general_ci;


# Dump of table article
# ------------------------------------------------------------

DROP TABLE IF EXISTS `article`;

CREATE TABLE `article` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户编号',
  `title` varchar(200) DEFAULT NULL COMMENT '标题',
  `tag` varchar(200) DEFAULT NULL,
  `category_id` int(11) DEFAULT NULL COMMENT '类别编号',
  `content` varchar(5000) DEFAULT NULL COMMENT '内容',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `article` WRITE;
/*!40000 ALTER TABLE `article` DISABLE KEYS */;

INSERT INTO `article` (`id`, `user_id`, `title`, `tag`, `category_id`, `content`, `created_at`, `updated_at`)
VALUES
	(5,1,'go study','',0,'<p>  近日小生迷上了golang,用了一个礼拜的闲暇时间，学习了一下golang的数据结构及并发goroutine。贴一下学习成果，不要见笑，并上传了附件</p>\n<p>1：环境搭建\n    下载go sdk 并安装\n    下载地址：<a href=\"http://dl.iteye.com/topics/download/522115e7-d053-3267-8d3e-efce8fb21ce9\">http://dl.iteye.com/topics/download/522115e7-d053-3267-8d3e-efce8fb21ce9</a>\n    安装后，配置好环境变量和GOPATH\n    我的mac 环境变量设置供参考：</p>\n<pre><code>export GO_HOME=/usr/local/go  \nexport GO_ROOT=$GO_HOME  \nexport GOPATH=/Users/xiaosheng/go_workspace  \nexport PATH=&quot;$GO_HOME/bin:$PATH&quot;  \n</code></pre><p>2：贴一下工程结构，比较传统的，呵呵。      </p>\n<p>   bin：存放编译后的可执行文件\n   pkg：编译生成的文件\n   src:   源代码\n3:开发环境\n   1) 我用的intellij idea,贴一下idea上golang的环境配置\n    安装 go-lang-pugin-org ,安装方式如下\n    plugins&gt;browse repositories&gt;manage repositories<br>    点击+ 号 输入 <a href=\"https://plugins.jetbrains.com/plugins/alpha/5047\">https://plugins.jetbrains.com/plugins/alpha/5047</a>\n    点击check，check通过后 点击ok，回到plugins search go，\n    勾选go 并安装，重启\n   2) 设置go sdk\n     file&gt;project struct&gt;sdks  添加sdk</p>\n<p>4:go 工程创建\n    因为GOPATH已经指定了go 工程目录，idea中直接import 那个目录就ok了</p>\n<p>5:工程目录结构</p>\n<p>6:上nb的代码hello world</p>\n<pre><code>    package demos\n\n    import &quot;fmt&quot;\n\n    func SayHello() {\n    fmt.Println(&quot;hello world&quot;)\n    }\n\n       测试用例\n    package main\n\n    import (\n    &quot;xiaosheng/demos&quot;\n    )\n    func main() {\n    demos.SayHello()\n    }\n</code></pre><p>7:学习demos\n  <a href=\"http://dl.iteye.com/topics/download/57600994-5c4c-3ade-ab38-3892d318f6d3\">http://dl.iteye.com/topics/download/57600994-5c4c-3ade-ab38-3892d318f6d3</a></p>\n','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(7,1,'go study','',0,'<p> 近日小生迷上了golang,用了一个礼拜的闲暇时间，学习了一下golang的数据结构及并发goroutine。贴一下学习成果，不要见笑，并上传了附件</p> <p>1：环境搭建 下载go sdk 并安装 下载地址：<a href=\"http://dl.iteye.com/topics/download/522115e7-d053-3267-8d3e-efce8fb21ce9\">http://dl.iteye.com/topics/download/522115e7-d053-3267-8d3e-efce8fb21ce9</a> 安装后，配置好环境变量和GOPATH 我的mac 环境变量设置供参考：</p> <pre><code>export GO_HOME=/usr/local/go export GO_ROOT=$GO_HOME export GOPATH=/Users/xiaosheng/go_workspace export PATH=&quot;$GO_HOME/bin:$PATH&quot; </code></pre><p>2：贴一下工程结构，比较传统的，呵呵。 </p> <p> bin：存放编译后的可执行文件 pkg：编译生成的文件 src: 源代码 3:开发环境 1) 我用的intellij idea,贴一下idea上golang的环境配置 安装 go-lang-pugin-org ,安装方式如下 plugins&gt;browse repositories&gt;manage repositories<br> 点击+ 号 输入 <a href=\"https://plugins.jetbrains.com/plugins/alpha/5047\">https://plugins.jetbrains.com/plugins/alpha/5047</a> 点击check，check通过后 点击ok，回到plugins search go， 勾选go 并安装，重启 2) 设置go sdk file&gt;project struct&gt;sdks 添加sdk</p> <p>4:go 工程创建 因为GOPATH已经指定了go 工程目录，idea中直接import 那个目录就ok了</p> <p>5:工程目录结构</p> <p>6:上nb的代码hello world</p> <pre><code> package demos import &quot;fmt&quot; func SayHello() { fmt.Println(&quot;hello world&quot;) } 测试用例 package main import ( &quot;xiaosheng/demos&quot; ) func main() { demos.SayHello() } </code></pre><p>7:学习demos <a href=\"http://dl.iteye.com/topics/download/57600994-5c4c-3ade-ab38-3892d318f6d3\">http://dl.iteye.com/topics/download/57600994-5c4c-3ade-ab38-3892d318f6d3</a></p>','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(8,1,'go study','',0,'<p>  近日小生迷上了golang,用了一个礼拜的闲暇时间，学习了一下golang的数据结构及并发goroutine。贴一下学习成果，不要见笑，并上传了附件</p>\n<p>1：环境搭建\n    下载go sdk 并安装\n    下载地址：<a href=\"http://dl.iteye.com/topics/download/522115e7-d053-3267-8d3e-efce8fb21ce9\">http://dl.iteye.com/topics/download/522115e7-d053-3267-8d3e-efce8fb21ce9</a>\n    安装后，配置好环境变量和GOPATH\n    我的mac 环境变量设置供参考：</p>\n<pre><code>export GO_HOME=/usr/local/go  \nexport GO_ROOT=$GO_HOME  \nexport GOPATH=/Users/xiaosheng/go_workspace  \nexport PATH=&quot;$GO_HOME/bin:$PATH&quot;  \n</code></pre><p>2：贴一下工程结构，比较传统的，呵呵。      </p>\n<p>   bin：存放编译后的可执行文件\n   pkg：编译生成的文件\n   src:   源代码\n3:开发环境\n   1) 我用的intellij idea,贴一下idea上golang的环境配置\n    安装 go-lang-pugin-org ,安装方式如下\n    plugins&gt;browse repositories&gt;manage repositories<br>    点击+ 号 输入 <a href=\"https://plugins.jetbrains.com/plugins/alpha/5047\">https://plugins.jetbrains.com/plugins/alpha/5047</a>\n    点击check，check通过后 点击ok，回到plugins search go，\n    勾选go 并安装，重启\n   2) 设置go sdk\n     file&gt;project struct&gt;sdks  添加sdk</p>\n<p>4:go 工程创建\n    因为GOPATH已经指定了go 工程目录，idea中直接import 那个目录就ok了</p>\n<p>5:工程目录结构</p>\n<p>6:上nb的代码hello world</p>\n<pre><code>    package demos\n\n    import &quot;fmt&quot;\n\n    func SayHello() {\n    fmt.Println(&quot;hello world&quot;)\n    }\n\n       测试用例\n    package main\n\n    import (\n    &quot;xiaosheng/demos&quot;\n    )\n    func main() {\n    demos.SayHello()\n    }\n</code></pre><p>7:学习demos\n  <a href=\"http://dl.iteye.com/topics/download/57600994-5c4c-3ade-ab38-3892d318f6d3\">http://dl.iteye.com/topics/download/57600994-5c4c-3ade-ab38-3892d318f6d3</a></p>\n','0000-00-00 00:00:00','0000-00-00 00:00:00');

/*!40000 ALTER TABLE `article` ENABLE KEYS */;
UNLOCK TABLES;


alter table `user` add column head varchar(200) comment '头像地址' after nick;

alter table `user` add column province varchar(20) comment '省份' after qq;

alter table `user` add column city varchar(200) comment '头像地址' after province;

alter table `user` add column open_id varchar(20) comment '第三方openid' after city;


# Dump of table category
# ------------------------------------------------------------

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` bigint(22) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT '0',
  `name` varchar(200) DEFAULT NULL COMMENT '名称',
  `order` int(11) DEFAULT '0' COMMENT '排序编号',
  `article_count` int(11) DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;

INSERT INTO `category` (`id`, `name`, `order`, `created_at`, `updated_at`)
VALUES
	(1,'go',1,'2015-11-08 16:53:47',NULL),
	(2,'java',2,'2015-11-08 16:53:47',NULL),
	(3,'redis',3,'2015-11-08 16:53:47',NULL),
	(4,'mysql',4,'2015-11-08 16:53:47',NULL),
	(5,'git',5,'2015-11-08 16:53:47',NULL),
	(6,'maven',6,'2015-11-08 16:53:47',NULL),
	(7,'js',7,'2015-11-08 16:53:47',NULL),
	(8,'tool',8,'2015-11-08 16:53:47',NULL);

/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL COMMENT '用户名',
  `nick` varchar(100) DEFAULT NULL COMMENT '昵称',
  `password` varchar(50) DEFAULT NULL COMMENT '密码',
  `age` int(11) DEFAULT NULL COMMENT '年龄',
  `sex` int(11) DEFAULT NULL COMMENT '1：男 ；2：女',
  `cell` varchar(20) DEFAULT NULL COMMENT '手机号',
  `mail` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `qq` varchar(50) DEFAULT NULL COMMENT 'qq号',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;

INSERT INTO `user` (`id`, `name`, `nick`, `password`, `age`, `sex`, `cell`, `mail`, `qq`, `created_at`, `updated_at`)
VALUES
	(1,'406504302@qq.com','java小生','xxxxx',26,1,'xxxxxx',NULL,NULL,'2015-08-28 12:03:50',NULL);

/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;



DROP TABLE IF EXISTS `parameter`;

CREATE TABLE `parameter` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(50) DEFAULT NULL COMMENT '参数键',
  `value` varchar(50) DEFAULT NULL COMMENT '参数值',
  `desc` varchar(100) DEFAULT NULL COMMENT '参数描述',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `parameter` WRITE;
/*!40000 ALTER TABLE `parameter` DISABLE KEYS */;


DROP TABLE IF EXISTS `article_view`;

CREATE TABLE `article_view` (
  `id` bigint(22) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(22) DEFAULT NULL COMMENT '用户编号',
  `article_id` bigint(22) DEFAULT NULL COMMENT '文章编号',
  `ip` varchar(50) DEFAULT NULL COMMENT '访问ip',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `article_like`;

CREATE TABLE `article_like` (
  `id` bigint(22) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(22) DEFAULT '0' COMMENT '用户编号',
  `article_id` bigint(22) DEFAULT NULL COMMENT '文章编号',
  `valid` tinyint(1) DEFAULT '1' COMMENT '1:有效;0:无效',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_userid_articleid` (`user_id`,`article_id`)
) ENGINE=InnoDB AUTO_INCREMENT=102 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
  `id` bigint(22) unsigned NOT NULL AUTO_INCREMENT,
  `comment_user_id` bigint(22) DEFAULT NULL COMMENT '评论者编号',
  `be_commented_user_id` bigint(22) DEFAULT NULL COMMENT '被评论者编号',
  `article_id` bigint(22) DEFAULT NULL COMMENT '文章编号',
  `parent_id` bigint(22) DEFAULT NULL COMMENT '操作者ip',
  `valid` tinyint(1) DEFAULT 1 COMMENT '1:有效;0:无效',
  `content` varchar(500) DEFAULT null COMMENT '评论内容',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#一个用户指定的类别是唯一的
create  unique index unique_userid_name on category(user_id,name);

#文章内容长度限制
insert into parameter(`key`,`value`,`desc`,created_at)
values('article_content_max_length',10000,'文章内容最大长度',now());
insert into parameter(`key`,`value`,`desc`,created_at)
values('article_title_max_length',100,'文章标题最大长度',now());
insert into parameter(`key`,`value`,`desc`,created_at)
values('article_categories_max_length',100,'文章分类最大长度',now());
insert into parameter(`key`,`value`,`desc`,created_at)
values('article_tags_max_length',100,'文章标签最大长度',now());
insert into parameter(`key`,`value`,`desc`,created_at)
values('max_create_article_count_one_day',20,'每日创建文章最大个数',now());
insert into parameter(`key`,`value`,`desc`,created_at)
values('max_create_comment_count_one_day',100,'每日创建评论最大个数',now());
insert into parameter(`key`,`value`,`desc`,created_at)
values('user_name_min_length',4,'用户名最小字符数',now());
insert into parameter(`key`,`value`,`desc`,created_at)
values('user_name_max_length',20,'用户名最大字符数',now());
insert into parameter(`key`,`value`,`desc`,created_at)
values('user_pass_min_length',6,'用户密码最小字符数',now());
insert into parameter(`key`,`value`,`desc`,created_at)
values('user_pass_max_length',12,'用户密码最大字符数',now());


CREATE TABLE `open_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `open_id` varchar(100) DEFAULT NULL COMMENT '第三方openid',
  `user_id` bigint(22) DEFAULT '0' COMMENT '绑定的用户编号',
  `type` int(11) DEFAULT NULL COMMENT '1：qq;2：新浪微博',
  `nick` varchar(50) DEFAULT NULL COMMENT '第三方昵称',
  `head` varchar(200) DEFAULT NULL COMMENT '头像',
  `sex` tinyint(1) DEFAULT '1' COMMENT '1：男；2：女',
  `age` int(10) DEFAULT '0' COMMENT '年龄',
  `province` varchar(20) DEFAULT NULL COMMENT '省份',
  `city` varchar(20) DEFAULT NULL COMMENT '城市',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_openid_type` (`open_id`,`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#20160216
alter table `user` add column view_count int(10) default 0 comment '文章浏览总数' after city ;
alter table `user` add column like_count int(10) default 0 comment '文章点赞总数' after city ;
insert into parameter(`key`,`value`,`desc`,created_at)
values('like_article_day_range',30,'精华文章筛选天数范围',now());


