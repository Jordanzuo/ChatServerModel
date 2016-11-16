/*
SQLyog 企业版 - MySQL GUI v7.14 
MySQL - 5.5.27 : Database - chatserver_test
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

CREATE DATABASE /*!32312 IF NOT EXISTS*/`chatserver_test` /*!40100 DEFAULT CHARACTER SET utf8 */;

/*Table structure for table `config` */

DROP TABLE IF EXISTS `config`;

CREATE TABLE `config` (
  `AppId` varchar(32) NOT NULL COMMENT '应用Id',
  `AppName` varchar(32) NOT NULL COMMENT '应用名称',
  `AppKey` varchar(64) NOT NULL COMMENT '应用Key，用于加密',
  `ManageCenterAPI` varchar(128) NOT NULL COMMENT 'ManageCenterAPI',
  `PartnerListAPI` varchar(128) NOT NULL COMMENT '合作商API',
  `ServerListAPI` varchar(128) NOT NULL COMMENT '服务器API',
  `ServerGroupListAPI` varchar(128) NOT NULL COMMENT '服务器组API',
  `GroupType` varchar(16) NOT NULL COMMENT '服务器组类型',
  `ChatServerCenterRpcAddress` varchar(64) NOT NULL COMMENT 'ChatServerCenter监听的rpc地址（内网地址）',
  `ChatServerCenterWebAddress` varchar(64) NOT NULL COMMENT 'ChatServerCenter监听的web地址（内网地址）',
  `PlayerInfoAPI` varchar(128) NOT NULL COMMENT '获取玩家信息的API',
  `MaxMessageLength` int(11) NOT NULL COMMENT '消息最大长度',
  `MaxClientCount` int(11) NOT NULL COMMENT '最大的客户端数量',
  `IfRecordAPILog` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否记录API日志'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `config_ip` */

DROP TABLE IF EXISTS `config_ip`;

CREATE TABLE `config_ip` (
  `IP` varchar(128) NOT NULL COMMENT 'IP',
  PRIMARY KEY (`IP`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `config_word_forbid` */

DROP TABLE IF EXISTS `config_word_forbid`;

CREATE TABLE `config_word_forbid` (
  `Word` varchar(32) NOT NULL COMMENT '单词',
  PRIMARY KEY (`Word`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `config_word_sensitive` */

DROP TABLE IF EXISTS `config_word_sensitive`;

CREATE TABLE `config_word_sensitive` (
  `Word` varchar(32) NOT NULL COMMENT '单词',
  PRIMARY KEY (`Word`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `log_api` */

DROP TABLE IF EXISTS `log_api`;

CREATE TABLE `log_api` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键 ',
  `APIName` varchar(32) NOT NULL COMMENT 'API名称',
  `Content` varchar(1024) NOT NULL COMMENT '请求内容',
  `Crtime` datetime NOT NULL COMMENT '请求时间',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `log_message` */

DROP TABLE IF EXISTS `log_message`;

CREATE TABLE `log_message` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `PlayerId` varchar(36) NOT NULL COMMENT '玩家Id',
  `Name` varchar(32) NOT NULL COMMENT '玩家名称',
  `PartnerId` int(11) NOT NULL COMMENT '合作商Id',
  `ServerId` int(11) NOT NULL COMMENT '服务器Id',
  `ServerGroupId` int(11) NOT NULL COMMENT '服务器组Id',
  `Message` varchar(128) NOT NULL COMMENT '聊天内容',
  `ChannelType` int(11) NOT NULL COMMENT '聊天频道',
  `ToPlayerId` varchar(36) DEFAULT NULL COMMENT '如果是私聊，表示目标玩家',
  `Crtime` datetime NOT NULL COMMENT '发送时间',
  PRIMARY KEY (`Id`),
  KEY `IX_ServerGroupId_Crtime` (`ServerGroupId`,`Crtime`),
  KEY `IX_PlayerId_Crtime` (`PlayerId`,`Crtime`),
  KEY `IX_Name_Crtime` (`Name`,`Crtime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `log_online` */

DROP TABLE IF EXISTS `log_online`;

CREATE TABLE `log_online` (
  `OnlineTime` datetime NOT NULL COMMENT '在线时间',
  `Sid` int(11) NOT NULL COMMENT '序号',
  `ServerAddress` varchar(64) NOT NULL COMMENT '服务器Id',
  `ClientCount` int(11) NOT NULL COMMENT '客户端数量',
  `PlayerCount` int(11) NOT NULL COMMENT '玩家数量',
  `TotalCount` int(11) NOT NULL COMMENT '所有服务器的总数量',
  PRIMARY KEY (`OnlineTime`,`Sid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `player` */

DROP TABLE IF EXISTS `player`;

CREATE TABLE `player` (
  `Id` varchar(64) NOT NULL COMMENT '玩家Id',
  `Name` varchar(64) NOT NULL COMMENT '玩家名称',
  `PartnerId` int(11) NOT NULL COMMENT '合作商Id',
  `ServerId` int(11) NOT NULL COMMENT '服务器Id',
  `UnionId` varchar(64) DEFAULT NULL COMMENT '公会Id',
  `ExtraMsg` varchar(1024) DEFAULT NULL COMMENT '额外透传信息',
  `RegisterTime` datetime NOT NULL COMMENT '注册时间',
  `LoginTime` datetime NOT NULL COMMENT '登录时间',
  `IsForbidden` tinyint(1) NOT NULL COMMENT '是否封号',
  `SilentEndTime` datetime NOT NULL COMMENT '禁言结束时间',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
