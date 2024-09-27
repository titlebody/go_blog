-- MySQL dump 10.13  Distrib 8.0.39, for Win64 (x86_64)
--
-- Host: localhost    Database: go_blog
-- ------------------------------------------------------
-- Server version	8.0.39

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `advert_model`
--

DROP TABLE IF EXISTS `advert_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `advert_model` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` varchar(32) DEFAULT NULL,
  `href` longtext,
  `images` longtext,
  `is_show` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `advert_model`
--

LOCK TABLES `advert_model` WRITE;
/*!40000 ALTER TABLE `advert_model` DISABLE KEYS */;
/*!40000 ALTER TABLE `advert_model` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `article_model`
--

DROP TABLE IF EXISTS `article_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `article_model` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` varchar(32) DEFAULT NULL,
  `abstract` longtext,
  `content` longtext,
  `look_count` bigint DEFAULT NULL,
  `comment_count` bigint DEFAULT NULL,
  `digg_count` bigint DEFAULT NULL,
  `collects_count` bigint DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `category` varchar(32) DEFAULT NULL,
  `source` longtext,
  `link` longtext,
  `banner_id` bigint unsigned DEFAULT NULL,
  `nick_name` longtext,
  `banner_path` longtext,
  `tags` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_article_model_banner` (`banner_id`),
  KEY `fk_user_model_article_models` (`user_id`),
  CONSTRAINT `fk_article_model_banner` FOREIGN KEY (`banner_id`) REFERENCES `banner_model` (`id`),
  CONSTRAINT `fk_user_model_article_models` FOREIGN KEY (`user_id`) REFERENCES `user_model` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article_model`
--

LOCK TABLES `article_model` WRITE;
/*!40000 ALTER TABLE `article_model` DISABLE KEYS */;
/*!40000 ALTER TABLE `article_model` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `article_tag_models`
--

DROP TABLE IF EXISTS `article_tag_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `article_tag_models` (
  `article_model_id` bigint unsigned NOT NULL,
  `tag_model_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`article_model_id`,`tag_model_id`),
  KEY `fk_article_tag_models_tag_model` (`tag_model_id`),
  CONSTRAINT `fk_article_tag_models_article_model` FOREIGN KEY (`article_model_id`) REFERENCES `article_model` (`id`),
  CONSTRAINT `fk_article_tag_models_tag_model` FOREIGN KEY (`tag_model_id`) REFERENCES `tag_model` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article_tag_models`
--

LOCK TABLES `article_tag_models` WRITE;
/*!40000 ALTER TABLE `article_tag_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `article_tag_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `banner_model`
--

DROP TABLE IF EXISTS `banner_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `banner_model` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `path` longtext,
  `hash` longtext,
  `name` longtext,
  `type` bigint DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `banner_model`
--

LOCK TABLES `banner_model` WRITE;
/*!40000 ALTER TABLE `banner_model` DISABLE KEYS */;
/*!40000 ALTER TABLE `banner_model` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment_model`
--

DROP TABLE IF EXISTS `comment_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment_model` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `parent_comment_id` bigint unsigned DEFAULT NULL,
  `content` varchar(256) DEFAULT NULL,
  `digg_count` tinyint DEFAULT '0',
  `comment_count` tinyint DEFAULT '0',
  `article_id` bigint unsigned DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_comment_model_sub_comments` (`parent_comment_id`),
  KEY `fk_comment_model_user` (`user_id`),
  KEY `fk_article_model_comment_models` (`article_id`),
  CONSTRAINT `fk_article_model_comment_models` FOREIGN KEY (`article_id`) REFERENCES `article_model` (`id`),
  CONSTRAINT `fk_comment_model_sub_comments` FOREIGN KEY (`parent_comment_id`) REFERENCES `comment_model` (`id`),
  CONSTRAINT `fk_comment_model_user` FOREIGN KEY (`user_id`) REFERENCES `user_model` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment_model`
--

LOCK TABLES `comment_model` WRITE;
/*!40000 ALTER TABLE `comment_model` DISABLE KEYS */;
/*!40000 ALTER TABLE `comment_model` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `fade_back_model`
--

DROP TABLE IF EXISTS `fade_back_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `fade_back_model` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `content` varchar(128) DEFAULT NULL,
  `apply_content` varchar(128) DEFAULT NULL,
  `is_apply` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fade_back_model`
--

LOCK TABLES `fade_back_model` WRITE;
/*!40000 ALTER TABLE `fade_back_model` DISABLE KEYS */;
/*!40000 ALTER TABLE `fade_back_model` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `login_data_model`
--

DROP TABLE IF EXISTS `login_data_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `login_data_model` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `ip` varchar(20) DEFAULT NULL,
  `nick_name` varchar(42) DEFAULT NULL,
  `token` varchar(256) DEFAULT NULL,
  `device` varchar(256) DEFAULT NULL,
  `addr` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_login_data_model_user_model` (`user_id`),
  CONSTRAINT `fk_login_data_model_user_model` FOREIGN KEY (`user_id`) REFERENCES `user_model` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `login_data_model`
--

LOCK TABLES `login_data_model` WRITE;
/*!40000 ALTER TABLE `login_data_model` DISABLE KEYS */;
/*!40000 ALTER TABLE `login_data_model` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu_banner_model`
--

DROP TABLE IF EXISTS `menu_banner_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu_banner_model` (
  `menu_id` bigint unsigned DEFAULT NULL,
  `banner_id` bigint unsigned DEFAULT NULL,
  `sort` smallint DEFAULT NULL,
  KEY `fk_menu_banner_model_menu_model` (`menu_id`),
  KEY `fk_menu_banner_model_banner_model` (`banner_id`),
  CONSTRAINT `fk_menu_banner_model_banner_model` FOREIGN KEY (`banner_id`) REFERENCES `banner_model` (`id`),
  CONSTRAINT `fk_menu_banner_model_menu_model` FOREIGN KEY (`menu_id`) REFERENCES `menu_model` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu_banner_model`
--

LOCK TABLES `menu_banner_model` WRITE;
/*!40000 ALTER TABLE `menu_banner_model` DISABLE KEYS */;
/*!40000 ALTER TABLE `menu_banner_model` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu_model`
--

DROP TABLE IF EXISTS `menu_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu_model` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` varchar(32) DEFAULT NULL,
  `path` varchar(32) DEFAULT NULL,
  `slogan` varchar(64) DEFAULT NULL,
  `abstract` longtext,
  `abstract_time` bigint DEFAULT NULL,
  `banner_time` bigint DEFAULT NULL,
  `sort` smallint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu_model`
--

LOCK TABLES `menu_model` WRITE;
/*!40000 ALTER TABLE `menu_model` DISABLE KEYS */;
/*!40000 ALTER TABLE `menu_model` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message_model`
--

DROP TABLE IF EXISTS `message_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `message_model` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `send_user_id` bigint unsigned NOT NULL,
  `send_user_nick_name` varchar(42) DEFAULT NULL,
  `send_user_avatar` longtext,
  `rev_user_id` bigint unsigned NOT NULL,
  `rev_user_nick_name` varchar(42) DEFAULT NULL,
  `rev_user_avatar` longtext,
  `is_read` tinyint(1) DEFAULT '0',
  `content` longtext,
  PRIMARY KEY (`id`,`send_user_id`,`rev_user_id`),
  KEY `fk_message_model_send_user_model` (`send_user_id`),
  KEY `fk_message_model_rev_user_model` (`rev_user_id`),
  CONSTRAINT `fk_message_model_rev_user_model` FOREIGN KEY (`rev_user_id`) REFERENCES `user_model` (`id`),
  CONSTRAINT `fk_message_model_send_user_model` FOREIGN KEY (`send_user_id`) REFERENCES `user_model` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message_model`
--

LOCK TABLES `message_model` WRITE;
/*!40000 ALTER TABLE `message_model` DISABLE KEYS */;
/*!40000 ALTER TABLE `message_model` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tag_model`
--

DROP TABLE IF EXISTS `tag_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tag_model` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` varchar(16) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tag_model`
--

LOCK TABLES `tag_model` WRITE;
/*!40000 ALTER TABLE `tag_model` DISABLE KEYS */;
/*!40000 ALTER TABLE `tag_model` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_collect_model`
--

DROP TABLE IF EXISTS `user_collect_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_collect_model` (
  `user_id` bigint unsigned NOT NULL,
  `article_id` bigint unsigned NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`user_id`,`article_id`),
  KEY `fk_user_collect_model_article_model` (`article_id`),
  CONSTRAINT `fk_user_collect_model_article_model` FOREIGN KEY (`article_id`) REFERENCES `article_model` (`id`),
  CONSTRAINT `fk_user_collect_model_user_model` FOREIGN KEY (`user_id`) REFERENCES `user_model` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_collect_model`
--

LOCK TABLES `user_collect_model` WRITE;
/*!40000 ALTER TABLE `user_collect_model` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_collect_model` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_model`
--

DROP TABLE IF EXISTS `user_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_model` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `nick_name` varchar(36) DEFAULT NULL,
  `user_name` varchar(36) DEFAULT NULL,
  `password` varchar(128) DEFAULT NULL,
  `avatar` varchar(256) DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  `tel` varchar(18) DEFAULT NULL,
  `addr` varchar(64) DEFAULT NULL,
  `token` varchar(64) DEFAULT NULL,
  `ip` varchar(20) DEFAULT NULL,
  `role` tinyint DEFAULT '1',
  `sign_status` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_model`
--

LOCK TABLES `user_model` WRITE;
/*!40000 ALTER TABLE `user_model` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_model` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-09-27 15:58:26
