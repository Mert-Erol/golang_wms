-- --------------------------------------------------------
-- Sunucu:                       127.0.0.1
-- Sunucu sürümü:                5.7.31 - MySQL Community Server (GPL)
-- Sunucu İşletim Sistemi:       Win64
-- HeidiSQL Sürüm:               11.0.0.5919
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- wms için veritabanı yapısı dökülüyor
CREATE DATABASE IF NOT EXISTS `wms` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `wms`;

-- tablo yapısı dökülüyor wms.products
CREATE TABLE IF NOT EXISTS `products` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_name` varchar(50) DEFAULT NULL,
  `stock_code` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- wms.products: 0 rows tablosu için veriler indiriliyor
DELETE FROM `products`;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` (`id`, `product_name`, `stock_code`) VALUES
	(1, 'Telefon', '222');
/*!40000 ALTER TABLE `products` ENABLE KEYS */;

-- tablo yapısı dökülüyor wms.shelfs
CREATE TABLE IF NOT EXISTS `shelfs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `capacity` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- wms.shelfs: 2 rows tablosu için veriler indiriliyor
DELETE FROM `shelfs`;
/*!40000 ALTER TABLE `shelfs` DISABLE KEYS */;
INSERT INTO `shelfs` (`id`, `name`, `capacity`) VALUES
	(1, 'Raf1', 100),
	(2, 'Raf2', 50);
/*!40000 ALTER TABLE `shelfs` ENABLE KEYS */;

-- tablo yapısı dökülüyor wms.stocks
CREATE TABLE IF NOT EXISTS `stocks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_code` int(11) DEFAULT NULL,
  `quantity` int(11) DEFAULT NULL,
  `shelf_id` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- wms.stocks: 0 rows tablosu için veriler indiriliyor
DELETE FROM `stocks`;
/*!40000 ALTER TABLE `stocks` DISABLE KEYS */;
/*!40000 ALTER TABLE `stocks` ENABLE KEYS */;

-- tablo yapısı dökülüyor wms.transactions
CREATE TABLE IF NOT EXISTS `transactions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `stock_code` varchar(50) DEFAULT NULL,
  `quantity` int(11) DEFAULT NULL,
  `type` tinyint(4) DEFAULT NULL COMMENT '1: Mal Kabul, 2: Sipariş',
  `statu` int(11) DEFAULT '0' COMMENT '0: Beklemede, 1:İşlem Tamamlandı',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- wms.transactions: 2 rows tablosu için veriler indiriliyor
DELETE FROM `transactions`;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` (`id`, `stock_code`, `quantity`, `type`, `statu`) VALUES
	(1, '111', 3, 0, 0),
	(2, '222', 5, 0, 0);
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;

-- tablo yapısı dökülüyor wms.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  `role` tinyint(4) DEFAULT NULL COMMENT '1:İdari Personel, 2: Depo Personeli',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- wms.users: 2 rows tablosu için veriler indiriliyor
DELETE FROM `users`;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` (`id`, `username`, `password`, `role`) VALUES
	(1, 'merol', '$2a$10$V8NKlnHuz8amCIMaHTFDN.lxsSOD1SYEUfof.UP/Bo/VV15/E8ste', 1),
	(2, 'mertt', '$2a$10$VtxZNi4yxi3n6MNlTwM.wOhxMI2GlTeYAIQ/bR/Y93/h4Rs2fl9mW', NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
