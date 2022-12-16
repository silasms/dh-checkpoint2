
SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema dentistoffice
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `dentistoffice` DEFAULT CHARACTER SET utf8mb3 ;
USE `dentistoffice` ;


DROP TABLE IF EXISTS dentists;
DROP TABLE IF EXISTS patients;
DROP TABLE IF EXISTS appointments;

-- -----------------------------------------------------
-- Table `dentistoffice`.`dentists`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dentistoffice`.`dentists` (
  `id` INT(11) NOT NULL  AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `surname` VARCHAR(255) NOT NULL,
  `registry` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `registry_unique` (`registry` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`localities`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dentistoffice`.`patients` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,  
  `name` VARCHAR(255) NOT NULL,
  `surname` VARCHAR(255) NOT NULL,
  `rg` VARCHAR(255) NOT NULL,
  `registry_date` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `rg_unique` UNIQUE (`rg`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`carries`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dentistoffice`.`appointments` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `consult_date` DATETIME NOT NULL,
  `dentist_id` INT(11) NOT NULL,
  `patient_id` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_dentist1_idx` (`dentist_id` ASC) VISIBLE,
  CONSTRAINT `fk_dentist1`
    FOREIGN KEY (`dentist_id`)
    REFERENCES `dentistoffice`.`dentists` (`id`),
  INDEX `fk_patient1_idx` (`patient_id` ASC) VISIBLE,
  CONSTRAINT `fk_patient1`
    FOREIGN KEY (`patient_id`)
    REFERENCES `dentistoffice`.`patients` (`id`)

)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;

INSERT INTO `dentists` VALUES
(1,'Esther','Pimentel','COD-9878'),
(2, 'Silas', 'Medeiros', 'CPX-4312'),
(3,'Thays','Gama','CPX-6543');

INSERT INTO `patients` VALUES
(1,'Ronaldo', 'Nazário','12345678', '2010-12-05 23:18:44'),
(2,'Jair Messias', 'Bolsonaro', '22345678', '2015-12-09 12:17:04'),
(3,'Dilma', 'Roussef', '32345678', '2019-03-02 09:11:11');

INSERT INTO `appointments` VALUES
(1,'2019-09-10 14:42:55', 1, 1),
(2,'2018-06-09 12:33:11', 2, 2),
(3,'2021-04-15 12:00:00', 3, 3);
