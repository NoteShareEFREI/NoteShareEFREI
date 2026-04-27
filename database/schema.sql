DROP DATABASE IF EXISTS noteShare;
CREATE DATABASE noteShare;
USE noteShare;

DROP TABLE IF EXISTS Category;
DROP TABLE IF EXISTS Sub_category;
DROP TABLE IF EXISTS Account;

DROP TABLE IF EXISTS Path;
DROP TABLE IF EXISTS file;
DROP TABLE IF EXISTS Comment;

DROP TABLE IF EXISTS Session;



CREATE TABLE Category(
   Cat_ID INT auto_increment,
   Name VARCHAR(50) NOT NULL,
   PRIMARY KEY(Cat_ID)
);

CREATE TABLE Sub_category(
   SCat_ID INT auto_increment,
   name VARCHAR(50),
   Cat_ID INT NOT NULL,
   PRIMARY KEY(SCat_ID),
   FOREIGN KEY(Cat_ID) REFERENCES Category(Cat_ID)
);

CREATE TABLE Account(
   User_ID INT auto_increment,
   name VARCHAR(50),
   hash VARCHAR(50),
   email VARCHAR(50),
   phone VARCHAR(50),
   PRIMARY KEY(User_ID)
);

CREATE TABLE Path(
   fID INT auto_increment,
   name VARCHAR(50),
   Prev INT NOT NULL,
   SCat_ID INT NOT NULL,
   PRIMARY KEY(fID),
   FOREIGN KEY(SCat_ID) REFERENCES Sub_category(SCat_ID),
   FOREIGN KEY(Prev) REFERENCES Path(fID)
);

CREATE TABLE file(
   Sheet_ID INT auto_increment,
   Sheet blob,
   User_ID INT NOT NULL,
   fID INT NOT NULL,
   PRIMARY KEY(Sheet_ID),
   FOREIGN KEY(User_ID) REFERENCES Account(User_ID),
   FOREIGN KEY(fID) REFERENCES Path(fID)
);

CREATE TABLE Comment(
   Comment_ID INT auto_increment,
   Text VARCHAR(100),
   User_ID INT NOT NULL,
   Sheet_ID INT NOT NULL,
   PRIMARY KEY(Comment_ID),
   FOREIGN KEY(User_ID) REFERENCES Account(User_ID),
   FOREIGN KEY(Sheet_ID) REFERENCES file(Sheet_ID)
);

CREATE TABLE Session(
    Session_ID INT auto_increment PRIMARY KEY,
    User_ID INT NOT NULL,
    Token_Hash VARCHAR(256) NOT NULL UNIQUE,
    Created_At TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    Expires_At TIMESTAMP NOT NULL,
    FOREIGN KEY(User_ID) REFERENCES Account(User_ID)
);