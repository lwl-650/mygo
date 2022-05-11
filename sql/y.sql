/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50558
Source Host           : localhost:3306
Source Database       : vueparent

Target Server Type    : MYSQL
Target Server Version : 50558
File Encoding         : 65001

Date: 2022-05-11 18:12:54
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `y`
-- ----------------------------
DROP TABLE IF EXISTS `y`;
CREATE TABLE `y` (
  `yid` bigint(20) NOT NULL AUTO_INCREMENT,
  `yname` longtext,
  PRIMARY KEY (`yid`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of y
-- ----------------------------
INSERT INTO `y` VALUES ('1', 'LS');
INSERT INTO `y` VALUES ('2', 'LS');
INSERT INTO `y` VALUES ('3', 'LS');
INSERT INTO `y` VALUES ('4', 'LS');
INSERT INTO `y` VALUES ('5', 'LS');
