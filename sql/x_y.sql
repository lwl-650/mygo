/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50558
Source Host           : localhost:3306
Source Database       : vueparent

Target Server Type    : MYSQL
Target Server Version : 50558
File Encoding         : 65001

Date: 2022-05-11 18:12:47
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `x_y`
-- ----------------------------
DROP TABLE IF EXISTS `x_y`;
CREATE TABLE `x_y` (
  `y_yid` bigint(20) NOT NULL DEFAULT '0',
  `x_xid` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`y_yid`,`x_xid`),
  KEY `fk_x_y_x` (`x_xid`),
  CONSTRAINT `fk_x_y_x` FOREIGN KEY (`x_xid`) REFERENCES `x` (`xid`),
  CONSTRAINT `fk_x_y_y` FOREIGN KEY (`y_yid`) REFERENCES `y` (`yid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of x_y
-- ----------------------------
INSERT INTO `x_y` VALUES ('1', '1');
INSERT INTO `x_y` VALUES ('4', '1');
INSERT INTO `x_y` VALUES ('2', '2');
INSERT INTO `x_y` VALUES ('3', '3');
INSERT INTO `x_y` VALUES ('5', '5');
INSERT INTO `x_y` VALUES ('2', '6');
