/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50558
Source Host           : localhost:3306
Source Database       : vueparent

Target Server Type    : MYSQL
Target Server Version : 50558
File Encoding         : 65001

Date: 2022-05-11 18:12:27
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `a_b`
-- ----------------------------
DROP TABLE IF EXISTS `a_b`;
CREATE TABLE `a_b` (
  `b_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `a_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`b_id`,`a_id`),
  KEY `fk_a` (`a_id`) USING BTREE,
  CONSTRAINT `fk_a_` FOREIGN KEY (`b_id`) REFERENCES `b` (`id`),
  CONSTRAINT `fk_a_b` FOREIGN KEY (`a_id`) REFERENCES `a` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of a_b
-- ----------------------------
INSERT INTO `a_b` VALUES ('2', '1');
INSERT INTO `a_b` VALUES ('3', '3');
