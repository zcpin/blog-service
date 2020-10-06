## Go blog-service

### sql
```sql
CREATE DATABASE IF NOT EXISTS `blog_service` DEFAULT CHARACTER SET UTF8MB4 DEFAULT COLLATE UTF8MB4_GENERAL_CI ;

USE blog_service;
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE IF NOT EXISTS  `blog_tag`(
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT ,
	`name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '标签名称',
	`created_on` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
	`created_by` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '创建人',
	`modified_on` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
	`modified_by` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '修改人',
	`delete_on` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
	`is_delete` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除，0-未删除；1-已删除',
	`state` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态，0-禁用；1-启用',
	PRIMARY KEY (`id`)
)ENGINE = INNODB DEFAULT CHARSET =UTF8MB4 COMMENT '标签管理';

DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE IF NOT EXISTS `blog_article`(
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '文章标题',
	`desc` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '文章简介',
	`cover_image_url` VARCHAR(255) DEFAULT '' COMMENT '文章封面图片地址',
	`content` LONGTEXT COMMENT '文章内容',
	`created_on` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
	`created_by` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '创建人',
	`modified_on` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
	`modified_by` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '修改人',
	`delete_on` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
	`is_delete` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除，0-未删除；1-已删除',
	`state` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态，0-禁用；1-启用',
	PRIMARY KEY (`id`)
) ENGINE = INNODB DEFAULT CHARSET = UTF8MB4 COMMENT ='文章管理';

DROP TABLE IF EXISTS `blog_article_tag`;
CREATE TABLE IF NOT EXISTS `blog_article_tag`(
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT ,
	`article_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章ID',
	`tag_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '标签ID',
	`created_on` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
	`created_by` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '创建人',
	`modified_on` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
	`modified_by` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '修改人',
	`delete_on` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
	`is_delete` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除，0-未删除；1-已删除',
	PRIMARY KEY (`id`)
)ENGINE = INNODB DEFAULT CHARSET = UTF8MB4 COMMENT ='文章标签关联表';
```
