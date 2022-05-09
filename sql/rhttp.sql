/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50558
Source Host           : localhost:3306
Source Database       : vueparent

Target Server Type    : MYSQL
Target Server Version : 50558
File Encoding         : 65001

Date: 2022-05-09 18:07:23
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `rhttp`
-- ----------------------------
DROP TABLE IF EXISTS `rhttp`;
CREATE TABLE `rhttp` (
  `r_id` int(11) NOT NULL AUTO_INCREMENT,
  `r_name` varchar(255) DEFAULT NULL,
  `r_method` varchar(255) DEFAULT NULL,
  `r_url` varchar(255) DEFAULT NULL,
  `r_actualurl` varchar(255) DEFAULT NULL,
  `r_status` varchar(255) DEFAULT NULL,
  `r_time` varchar(255) DEFAULT NULL,
  `adminid` int(11) NOT NULL,
  PRIMARY KEY (`r_id`),
  KEY `adminid` (`adminid`),
  CONSTRAINT `adminid` FOREIGN KEY (`adminid`) REFERENCES `admin` (`a_id`)
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of rhttp
-- ----------------------------
INSERT INTO `rhttp` VALUES ('1', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:21:34', '1');
INSERT INTO `rhttp` VALUES ('2', 'admin', 'DELETE', '/findRhttp', '/findRhttp', '500', '2022-04-28 14:21:36', '2');
INSERT INTO `rhttp` VALUES ('3', 'admin', 'GET', '/findRhttp', '/findRhttp', '600', '2022-04-28 14:21:38', '2');
INSERT INTO `rhttp` VALUES ('4', 'admin', 'POST', '/findRhttp', '/findRhttp', '100', '2022-04-28 14:21:39', '1');
INSERT INTO `rhttp` VALUES ('5', 'admin', 'PUT', '/findRhttp', '/findRhttp', '300', '2022-04-28 14:21:41', '1');
INSERT INTO `rhttp` VALUES ('6', 'admin', 'PATCH', '/findRhttp', '/findRhttp', '400', '2022-04-28 14:21:43', '1');
INSERT INTO `rhttp` VALUES ('7', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:21:46', '1');
INSERT INTO `rhttp` VALUES ('8', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:21:49', '1');
INSERT INTO `rhttp` VALUES ('9', 'admin', 'POST', '/getArr', '/getArr', '200', '2022-04-28 14:21:54', '1');
INSERT INTO `rhttp` VALUES ('10', 'admin', 'POST', '/getMap', '/getMap', '200', '2022-04-28 14:21:57', '1');
INSERT INTO `rhttp` VALUES ('11', 'admin', 'POST', '/getMapArr', '/getMapArr', '200', '2022-04-28 14:22:00', '1');
INSERT INTO `rhttp` VALUES ('12', 'admin', 'POST', '/getArrMap', '/getArrMap', '200', '2022-04-28 14:22:03', '1');
INSERT INTO `rhttp` VALUES ('13', 'admin', 'POST', '/getssq', '/getssq', '200', '2022-04-28 14:22:06', '1');
INSERT INTO `rhttp` VALUES ('14', 'admin', 'POST', 'da', '/ssq', '404', '2022-04-28 14:22:11', '1');
INSERT INTO `rhttp` VALUES ('15', 'admin', 'POST', '/delById', '/delById?id=22', '200', '2022-04-28 14:22:13', '1');
INSERT INTO `rhttp` VALUES ('16', 'admin', 'POST', '/login', '/login?n', '200', '2022-04-28 14:22:15', '1');
INSERT INTO `rhttp` VALUES ('17', 'admin', 'POST', '/addUser', '11', '200', '2022-04-28 14:22:17', '1');
INSERT INTO `rhttp` VALUES ('18', 'admin', 'POST', '/findId', '/findId?id=1', '200', '2022-04-28 14:22:19', '1');
INSERT INTO `rhttp` VALUES ('19', 'admin', 'POST', '/find', '/find', '200', '2022-04-28 14:22:22', '1');
INSERT INTO `rhttp` VALUES ('20', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:22:25', '1');
INSERT INTO `rhttp` VALUES ('21', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:22:27', '1');
INSERT INTO `rhttp` VALUES ('22', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:22:31', '1');
INSERT INTO `rhttp` VALUES ('23', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:22:32', '1');
INSERT INTO `rhttp` VALUES ('24', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:22:38', '1');
INSERT INTO `rhttp` VALUES ('25', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:22:41', '1');
INSERT INTO `rhttp` VALUES ('26', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:22:45', '1');
INSERT INTO `rhttp` VALUES ('27', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:22:47', '1');
INSERT INTO `rhttp` VALUES ('28', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:22:49', '1');
INSERT INTO `rhttp` VALUES ('29', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:22:58', '1');
INSERT INTO `rhttp` VALUES ('30', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:23:15', '1');
INSERT INTO `rhttp` VALUES ('31', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:23:17', '1');
INSERT INTO `rhttp` VALUES ('32', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:23:23', '1');
INSERT INTO `rhttp` VALUES ('33', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:23:24', '1');
INSERT INTO `rhttp` VALUES ('34', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:23:26', '1');
INSERT INTO `rhttp` VALUES ('35', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:23:31', '1');
INSERT INTO `rhttp` VALUES ('36', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:23:43', '1');
INSERT INTO `rhttp` VALUES ('37', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:23:46', '1');
INSERT INTO `rhttp` VALUES ('38', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:23:57', '1');
INSERT INTO `rhttp` VALUES ('39', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:24:01', '1');
INSERT INTO `rhttp` VALUES ('40', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:24:09', '1');
INSERT INTO `rhttp` VALUES ('41', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:24:11', '1');
INSERT INTO `rhttp` VALUES ('42', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:24:12', '1');
INSERT INTO `rhttp` VALUES ('43', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:24:43', '1');
INSERT INTO `rhttp` VALUES ('44', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:24:58', '1');
INSERT INTO `rhttp` VALUES ('45', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:25:00', '1');
INSERT INTO `rhttp` VALUES ('46', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:25:02', '1');
INSERT INTO `rhttp` VALUES ('47', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:25:03', '1');
INSERT INTO `rhttp` VALUES ('48', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:25:04', '1');
INSERT INTO `rhttp` VALUES ('49', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:25:45', '1');
INSERT INTO `rhttp` VALUES ('50', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:25:46', '1');
INSERT INTO `rhttp` VALUES ('51', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:25:47', '1');
INSERT INTO `rhttp` VALUES ('52', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:25:49', '1');
INSERT INTO `rhttp` VALUES ('53', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:25:53', '1');
INSERT INTO `rhttp` VALUES ('54', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:25:58', '1');
INSERT INTO `rhttp` VALUES ('55', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:26:00', '1');
INSERT INTO `rhttp` VALUES ('56', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:26:04', '1');
INSERT INTO `rhttp` VALUES ('57', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:26:12', '1');
INSERT INTO `rhttp` VALUES ('58', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:26:13', '1');
INSERT INTO `rhttp` VALUES ('59', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:26:17', '1');
INSERT INTO `rhttp` VALUES ('60', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:26:20', '1');
INSERT INTO `rhttp` VALUES ('61', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:26:23', '3');
INSERT INTO `rhttp` VALUES ('62', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:26:26', '1');
INSERT INTO `rhttp` VALUES ('63', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:26:30', '1');
INSERT INTO `rhttp` VALUES ('64', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 14:27:01', '1');
INSERT INTO `rhttp` VALUES ('65', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 16:54:38', '1');
INSERT INTO `rhttp` VALUES ('66', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 16:54:43', '1');
INSERT INTO `rhttp` VALUES ('67', 'admin', 'POST', '/findRhttp', '/findRhttp', '200', '2022-04-28 16:54:50', '1');
