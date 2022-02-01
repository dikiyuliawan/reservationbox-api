-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.4.17-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for reservationbox-api
CREATE DATABASE IF NOT EXISTS `reservationbox-api` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `reservationbox-api`;

-- Dumping structure for table reservationbox-api.hotels
CREATE TABLE IF NOT EXISTS `hotels` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `hotel_name` varchar(100) DEFAULT NULL,
  `address` varchar(150) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table reservationbox-api.hotels: ~5 rows (approximately)
/*!40000 ALTER TABLE `hotels` DISABLE KEYS */;
INSERT INTO `hotels` (`id`, `hotel_name`, `address`) VALUES
	(1, 'Bobobox Dago', 'Jalan Sultan Tirtayasa No 11, Dago, Bandung'),
	(2, 'Bobobox Paskal', 'Jalan Pasir Kaliki No 76A, Pasirkaliki, Bandung'),
	(3, 'Bobobox Alun Alun', 'Jalan Kepatihan No 8, Balonggede, Bandung'),
	(6, 'Bobobox Cipaganti', 'Cipaganti, Bandung'),
	(7, '', '');
/*!40000 ALTER TABLE `hotels` ENABLE KEYS */;

-- Dumping structure for table reservationbox-api.prices
CREATE TABLE IF NOT EXISTS `prices` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `date` datetime DEFAULT NULL,
  `room_type_id` int(11) DEFAULT NULL,
  `price` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table reservationbox-api.prices: ~2 rows (approximately)
/*!40000 ALTER TABLE `prices` DISABLE KEYS */;
INSERT INTO `prices` (`id`, `date`, `room_type_id`, `price`) VALUES
	(1, '2022-02-02 00:00:00', 1, 150000),
	(2, '2022-02-04 00:00:00', 2, 200000);
/*!40000 ALTER TABLE `prices` ENABLE KEYS */;

-- Dumping structure for table reservationbox-api.promos
CREATE TABLE IF NOT EXISTS `promos` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `room_type_id` int(11) NOT NULL,
  `promo_name` varchar(150) NOT NULL,
  `promo` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table reservationbox-api.promos: ~4 rows (approximately)
/*!40000 ALTER TABLE `promos` DISABLE KEYS */;
INSERT INTO `promos` (`id`, `room_type_id`, `promo_name`, `promo`) VALUES
	(1, 1, 'Promo Lunar New Year', 25000),
	(2, 2, 'Promo New Year', 35000),
	(3, 3, 'Promo3', 30000),
	(4, 0, '', 0);
/*!40000 ALTER TABLE `promos` ENABLE KEYS */;

-- Dumping structure for table reservationbox-api.reservations
CREATE TABLE IF NOT EXISTS `reservations` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` int(11) DEFAULT NULL,
  `customer_name` varchar(100) DEFAULT NULL,
  `book_room_count` int(11) DEFAULT NULL,
  `checkin_date` date DEFAULT NULL,
  `checkout_date` date DEFAULT NULL,
  `hotel_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Dumping data for table reservationbox-api.reservations: ~0 rows (approximately)
/*!40000 ALTER TABLE `reservations` DISABLE KEYS */;
/*!40000 ALTER TABLE `reservations` ENABLE KEYS */;

-- Dumping structure for table reservationbox-api.rooms
CREATE TABLE IF NOT EXISTS `rooms` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `hotel_id` int(11) DEFAULT NULL,
  `room_type_id` int(11) DEFAULT NULL,
  `room_number` int(11) DEFAULT NULL,
  `room_status` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table reservationbox-api.rooms: ~1 rows (approximately)
/*!40000 ALTER TABLE `rooms` DISABLE KEYS */;
INSERT INTO `rooms` (`id`, `hotel_id`, `room_type_id`, `room_number`, `room_status`) VALUES
	(1, 1, 1, 101, 'out of service');
/*!40000 ALTER TABLE `rooms` ENABLE KEYS */;

-- Dumping structure for table reservationbox-api.roomtypes
CREATE TABLE IF NOT EXISTS `roomtypes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table reservationbox-api.roomtypes: ~4 rows (approximately)
/*!40000 ALTER TABLE `roomtypes` DISABLE KEYS */;
INSERT INTO `roomtypes` (`id`, `name`) VALUES
	(1, 'Sky Single'),
	(2, 'Earth Single'),
	(3, 'Sky Double'),
	(4, 'Earth Double');
/*!40000 ALTER TABLE `roomtypes` ENABLE KEYS */;

-- Dumping structure for table reservationbox-api.stayrooms
CREATE TABLE IF NOT EXISTS `stayrooms` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `stay_id` int(11) DEFAULT NULL,
  `room_id` int(11) DEFAULT NULL,
  `date` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Dumping data for table reservationbox-api.stayrooms: ~0 rows (approximately)
/*!40000 ALTER TABLE `stayrooms` DISABLE KEYS */;
/*!40000 ALTER TABLE `stayrooms` ENABLE KEYS */;

-- Dumping structure for table reservationbox-api.stays
CREATE TABLE IF NOT EXISTS `stays` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `reservation_id` int(11) DEFAULT NULL,
  `guest_name` varchar(100) DEFAULT NULL,
  `room_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Dumping data for table reservationbox-api.stays: ~0 rows (approximately)
/*!40000 ALTER TABLE `stays` DISABLE KEYS */;
/*!40000 ALTER TABLE `stays` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
