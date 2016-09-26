/*
SQLyog Ultimate v11.11 (64 bit)
MySQL - 5.5.5-10.1.12-MariaDB : Database - ms
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`ms` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `ms`;

/*Table structure for table `tags` */

DROP TABLE IF EXISTS `tags`;

CREATE TABLE `tags` (
  `Id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `Name` varchar(100) NOT NULL,
  `Count` int(11) NOT NULL,
  `IsBlocked` tinyint(1) NOT NULL,
  `CreatedTime` int(11) NOT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Name` (`Name`)
) ENGINE=Aria AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1 DELAY_KEY_WRITE=1;

/*Data for the table `tags` */

insert  into `tags`(`Id`,`Name`,`Count`,`IsBlocked`,`CreatedTime`) values (1,'آتش‌سوزی',210,0,1471759996),(2,'انسان',356,0,1471759996),(3,'شده',380,0,1471759996),(4,'MicroTugs',390,0,1471759996),(5,'تکنیک‌های',217,0,1471759996),(6,'شوند.',191,0,1471759996),(7,'این',374,0,1471759996),(8,'خود',201,0,1471759996),(9,'بیشتر',185,0,1471759996),(10,'کنفرانس',193,0,1471759996),(11,'مهندسین',202,0,1471759996),(12,'ربات‌ها',387,0,1471759996),(13,'آزمایشگاه',356,0,1471759996),(14,'اشیایی',207,0,1471759996),(15,'راز',205,0,1471759996),(16,'دانشگاه',357,0,1471759996),(17,'نجات',187,0,1471759996),(18,'دوباره',207,0,1471759996),(19,'هنگامی',211,0,1471759996),(20,'استنفورد',181,0,1471759996),(21,'هنگام',205,0,1471759996),(22,'سازندگان',191,0,1471759996),(23,'عمودی',192,0,1471759996),(24,'به',187,0,1471759996),(25,'ولی',182,0,1471759996),(26,'هزار',158,0,1471759996),(27,'انسان‌های',215,0,1471759996),(28,'کشیدن',200,0,1471759996),(29,'سطح',187,0,1471759996),(30,'قدرت',207,0,1471759996),(31,'کوچک',388,0,1471759996),(32,'قوی‌ترین',205,0,1471759996),(33,'کوچک‌ترین',192,0,1471759997);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
