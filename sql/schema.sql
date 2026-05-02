-- CREATE DATABASE NoteShareEFREI;
USE NoteShareEFREI;

CREATE TABLE Account(
   Id_Account INT,
   Pseudo VARCHAR(50),
   Email VARCHAR(50),
   HashPassword VARCHAR(257),
   Phonenumber VARCHAR(15),
   Role BOOLEAN, -- 0 for admin, 1 for user
   PRIMARY KEY(Id_Account)
);

CREATE TABLE Category(
   Id_Category INT,
   Name VARCHAR(50),
   PRIMARY KEY(Id_Category)
);

CREATE TABLE SubCategory(
   Id_SubCategory INT,
   Name VARCHAR(50),
   Id_Category INT NOT NULL,
   PRIMARY KEY(Id_SubCategory),
   FOREIGN KEY(Id_Category) REFERENCES Category(Id_Category)
);

CREATE TABLE StudySheet(
   Id_Sheet INT,
   Hash VARCHAR(257),
   Name VARCHAR(50),
   Id_SubCategory INT NOT NULL,
   Id_Account INT NOT NULL,
   PRIMARY KEY(Id_Sheet),
   FOREIGN KEY(Id_SubCategory) REFERENCES SubCategory(Id_SubCategory),
   FOREIGN KEY(Id_Account) REFERENCES Account(Id_Account)
);

CREATE TABLE Comment(
   Id_Comment INT,
   Content VARCHAR(1000),
   Id_Sheet INT NOT NULL,
   Id_Account INT NOT NULL,
   PRIMARY KEY(Id_Comment),
   FOREIGN KEY(Id_Sheet) REFERENCES StudySheet(Id_Sheet),
   FOREIGN KEY(Id_Account) REFERENCES Account(Id_Account)
);

