/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50558
Source Host           : localhost:3306
Source Database       : vueparent

Target Server Type    : MYSQL
Target Server Version : 50558
File Encoding         : 65001

Date: 2022-04-27 16:27:44
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `r_http`
-- ----------------------------
DROP TABLE IF EXISTS `r_http`;
CREATE TABLE `r_http` (
  `r_id` int(11) NOT NULL AUTO_INCREMENT,
  `r_name` varchar(255) DEFAULT NULL,
  `r_method` varchar(255) DEFAULT NULL,
  `r_url` varchar(255) DEFAULT NULL,
  `r_actualurl` varchar(255) DEFAULT NULL,
  `r_status` int(11) DEFAULT NULL,
  `r_time` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`r_id`)
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of r_http
-- ----------------------------
INSERT INTO `r_http` VALUES ('2', 'admin', 'GET', '/add', null, '200', '2022-04-27 14:10:49');

