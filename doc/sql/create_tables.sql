CREATE TABLE `Tag` (
  `TagID` int unsigned NOT NULL AUTO_INCREMENT,
  `Name` varchar(100) DEFAULT '',
  `IsEnabled` tinyint unsigned DEFAULT '1',
  `InsertedTime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `UpdatedTime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`TagID`)
) ENGINE=InnoDB;

CREATE TABLE `User` (
  `UserID` int unsigned NOT NULL AUTO_INCREMENT,
  `Name` varchar(100) DEFAULT '',
  `Password` varchar(512) DEFAULT '',
  `InsertedTime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `UpdatedTime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`UserID`)
) ENGINE=InnoDB;

CREATE TABLE `Article` (
  `ArticleID` int unsigned NOT NULL AUTO_INCREMENT,
  `TagID` int unsigned DEFAULT '0',
  `Title` varchar(100) DEFAULT '',
  `Note` varchar(255) DEFAULT '',
  `Content` text,
  `IsEnabled` tinyint unsigned DEFAULT '1',
  `InsertedTime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `UpdatedTime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ArticleID`)
) ENGINE=InnoDB;


