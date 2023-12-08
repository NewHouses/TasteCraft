-- MariaDB dump 10.19-11.3.0-MariaDB, for Win64 (AMD64)
--
-- Host: 127.0.0.1    Database: tastecraft_test
-- ------------------------------------------------------
-- Server version	11.3.0-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `dish_types`
--

DROP TABLE IF EXISTS `dish_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `dish_types` (
  `dish_type_id` int(11) NOT NULL AUTO_INCREMENT,
  `dish_type` varchar(45) NOT NULL,
  PRIMARY KEY (`dish_type_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `foods`
--

DROP TABLE IF EXISTS `foods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `foods` (
  `food_id` int(11) NOT NULL AUTO_INCREMENT,
  `food_name` varchar(45) NOT NULL DEFAULT '',
  PRIMARY KEY (`food_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ingredients`
--

DROP TABLE IF EXISTS `ingredients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ingredients` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `recipe_id` int(11) NOT NULL,
  `food_id` int(11) NOT NULL,
  `quantity` int(10) unsigned NOT NULL DEFAULT 0,
  `measurement` varchar(45) NOT NULL DEFAULT 'grams',
  PRIMARY KEY (`id`),
  KEY `recipe_id_idx` (`recipe_id`),
  KEY `food_ingredient_id` (`food_id`),
  CONSTRAINT `food_ingredient_id` FOREIGN KEY (`food_id`) REFERENCES `foods` (`food_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `recipe_ingredient_id` FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`recipe_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `properties`
--

DROP TABLE IF EXISTS `properties`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `properties` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `food_id` int(11) NOT NULL,
  `kilocalories` int(10) unsigned NOT NULL DEFAULT 0,
  `proteins` int(10) unsigned NOT NULL DEFAULT 0,
  `carbohydrates` int(10) unsigned NOT NULL DEFAULT 0,
  `fat` int(10) unsigned NOT NULL DEFAULT 0,
  `fiber` int(10) unsigned NOT NULL DEFAULT 0,
  `calcium` int(10) unsigned NOT NULL DEFAULT 0,
  `iron` int(10) unsigned NOT NULL DEFAULT 0,
  `zinc` int(10) unsigned NOT NULL DEFAULT 0,
  `vitamin_a` int(10) unsigned NOT NULL DEFAULT 0,
  `vitamin_b` int(10) unsigned NOT NULL DEFAULT 0,
  `vitamin_b1` int(10) unsigned NOT NULL DEFAULT 0,
  `vitamin_b2` int(10) unsigned NOT NULL DEFAULT 0,
  `vitamin_b3` int(10) unsigned NOT NULL DEFAULT 0,
  `vitamin_b6` int(10) unsigned NOT NULL DEFAULT 0,
  `vitamin_b12` int(10) unsigned NOT NULL DEFAULT 0,
  `vitamin_c` int(10) unsigned NOT NULL DEFAULT 0,
  `vitamin_d` int(10) unsigned NOT NULL DEFAULT 0,
  `vitamin_e` int(10) unsigned NOT NULL DEFAULT 0,
  `vitamin_k` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `food_id_idx` (`food_id`),
  CONSTRAINT `food_property_id` FOREIGN KEY (`food_id`) REFERENCES `foods` (`food_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `recipes`
--

DROP TABLE IF EXISTS `recipes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `recipes` (
  `recipe_id` int(11) NOT NULL AUTO_INCREMENT,
  `recipe_name` varchar(45) NOT NULL,
  `dish_type` int(11) NOT NULL DEFAULT 1,
  PRIMARY KEY (`recipe_id`),
  KEY `dish_type_recipe_id_idx` (`dish_type`),
  CONSTRAINT `dish_type_recipe_id` FOREIGN KEY (`dish_type`) REFERENCES `dish_types` (`dish_type_id`) ON DELETE NO ACTION ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `schema_migration`
--

DROP TABLE IF EXISTS `schema_migration`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `schema_migration` (
  `version` varchar(14) NOT NULL,
  PRIMARY KEY (`version`),
  UNIQUE KEY `schema_migration_version_idx` (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `steps`
--

DROP TABLE IF EXISTS `steps`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `steps` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `recipe_id` int(11) NOT NULL,
  `step` varchar(120) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `recipe_step_id_idx` (`recipe_id`),
  CONSTRAINT `recipe_step_id` FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`recipe_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-12-08 21:43:51
