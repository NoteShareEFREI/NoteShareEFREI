CREATE DATABASE if not exists NoteShareEFREI;
USE NoteShareEFREI;

CREATE TABLE if not exists Account(
   Id_Account INT auto_increment,
   Pseudo VARCHAR(50),
   Email VARCHAR(50),
   HashPassword VARCHAR(257),
   Phonenumber VARCHAR(15),
   salt Int,
   Role BOOLEAN, -- 0 for admin, 1 for user
   PRIMARY KEY(Id_Account),
   INDEX idx_pseudo (Pseudo),
   INDEX idx_email (Email)
);

CREATE TABLE if not exists Category(
    Id_Category INT AUTO_INCREMENT,
    Name VARCHAR(50),
    PRIMARY KEY(Id_Category)
);

CREATE TABLE if not exists SubCategory(
    Id_SubCategory INT AUTO_INCREMENT,
    Name VARCHAR(50),
    Id_Category INT NOT NULL,
    PRIMARY KEY(Id_SubCategory),
    FOREIGN KEY(Id_Category) REFERENCES Category(Id_Category)
);

CREATE TABLE if not exists StudySheet(
    Id_Sheet INT AUTO_INCREMENT,
    Hash VARCHAR(257),
    Name VARCHAR(50),
    Id_SubCategory INT NOT NULL,
    Id_Account INT NOT NULL,
    PRIMARY KEY(Id_Sheet),
    FOREIGN KEY(Id_SubCategory) REFERENCES SubCategory(Id_SubCategory),
    FOREIGN KEY(Id_Account) REFERENCES Account(Id_Account),
    INDEX idx_hash (Hash),
    INDEX idx_account (Id_Account),
    INDEX idx_subcategory (Id_SubCategory)
);

CREATE TABLE if not exists Comment(
    Id_Comment INT AUTO_INCREMENT,
    Content VARCHAR(1000),
    Id_Sheet INT NOT NULL,
    Id_Account INT NOT NULL,
    PRIMARY KEY(Id_Comment),
    FOREIGN KEY(Id_Sheet) REFERENCES StudySheet(Id_Sheet),
    FOREIGN KEY(Id_Account) REFERENCES Account(Id_Account),
    INDEX idx_sheet (Id_Sheet),
    INDEX idx_account (Id_Account)
);
