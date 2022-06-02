/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50558
Source Host           : localhost:3306
Source Database       : vueparent

Target Server Type    : MYSQL
Target Server Version : 50558
File Encoding         : 65001

Date: 2022-06-02 18:19:06
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `xue`
-- ----------------------------
DROP TABLE IF EXISTS `xue`;
CREATE TABLE `xue` (
  `id` int(11) NOT NULL,
  `xue` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of xue
-- ----------------------------
