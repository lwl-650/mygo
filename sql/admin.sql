/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50558
Source Host           : localhost:3306
Source Database       : vueparent

Target Server Type    : MYSQL
Target Server Version : 50558
File Encoding         : 65001

Date: 2022-05-09 18:07:05
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `admin`
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `a_id` int(11) NOT NULL AUTO_INCREMENT,
  `a_name` varchar(255) DEFAULT NULL,
  `a_pass` varchar(255) DEFAULT NULL,
  `a_portrait` varchar(255) DEFAULT NULL,
  `a_grade` int(11) DEFAULT NULL,
  PRIMARY KEY (`a_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES ('1', 'admin', 'admin', 'http://q.vuetitle.lwlsl.top/admin.jpg', '10');
INSERT INTO `admin` VALUES ('2', 'zs', '123', 'http://q.vuetitle.lwlsl.top/title.png', '1');
INSERT INTO `admin` VALUES ('3', 'hello', '18', 'false', '2');
INSERT INTO `admin` VALUES ('4', 'world', '18', '5165', '2');
