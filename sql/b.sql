/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50558
Source Host           : localhost:3306
Source Database       : vueparent

Target Server Type    : MYSQL
Target Server Version : 50558
File Encoding         : 65001

Date: 2022-05-11 18:12:34
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `b`
-- ----------------------------
DROP TABLE IF EXISTS `b`;
CREATE TABLE `b` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `bid` bigint(20) DEFAULT NULL,
  `bname` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of b
-- ----------------------------
INSERT INTO `b` VALUES ('1', '3', '发生的');
INSERT INTO `b` VALUES ('2', '5', '吧v部分');
INSERT INTO `b` VALUES ('3', '5', '66');
