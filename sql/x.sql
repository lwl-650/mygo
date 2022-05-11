/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50558
Source Host           : localhost:3306
Source Database       : vueparent

Target Server Type    : MYSQL
Target Server Version : 50558
File Encoding         : 65001

Date: 2022-05-11 18:12:41
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `x`
-- ----------------------------
DROP TABLE IF EXISTS `x`;
CREATE TABLE `x` (
  `xid` bigint(20) NOT NULL AUTO_INCREMENT,
  `xname` longtext,
  PRIMARY KEY (`xid`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of x
-- ----------------------------
INSERT INTO `x` VALUES ('1', 'zs');
INSERT INTO `x` VALUES ('2', 'zs');
INSERT INTO `x` VALUES ('3', 'zs');
INSERT INTO `x` VALUES ('4', 'zs');
INSERT INTO `x` VALUES ('5', 'zs');
INSERT INTO `x` VALUES ('6', 'zs');
