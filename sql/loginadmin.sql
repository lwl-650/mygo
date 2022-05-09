/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50558
Source Host           : localhost:3306
Source Database       : vueparent

Target Server Type    : MYSQL
Target Server Version : 50558
File Encoding         : 65001

Date: 2022-05-09 18:07:15
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `loginadmin`
-- ----------------------------
DROP TABLE IF EXISTS `loginadmin`;
CREATE TABLE `loginadmin` (
  `lid` int(11) NOT NULL AUTO_INCREMENT,
  `lname` varchar(255) DEFAULT NULL,
  `ltime` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`lid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of loginadmin
-- ----------------------------
INSERT INTO `loginadmin` VALUES ('1', 'admin', '2022-05-07 10:01:58');
INSERT INTO `loginadmin` VALUES ('2', 'admin', '2022-05-07 10:03:25');
INSERT INTO `loginadmin` VALUES ('3', 'zs', '2022-05-07 10:05:19');
