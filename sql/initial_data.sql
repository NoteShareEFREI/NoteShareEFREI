-- Initial data for categories and subcategories based on files/ directory structure

USE NoteShareEFREI;

-- Insert categories
INSERT INTO Category (Id_Category, Name) VALUES
(1, 'Maths'),
(2, 'Physics'),
(3, 'Programming'),
(4, 'Formation_Generale');

-- Insert subcategories
INSERT INTO SubCategory (Id_SubCategory, Name, Id_Category) VALUES
-- Maths
(1, 'Calculus', 1),
(2, 'LinearAlgebra', 1),
-- Physics
(3, 'DigitalSystems', 2),
(4, 'SignalProcessing', 2),
(5, 'TransmissionSystems', 2),
-- Programming
(6, 'C', 3),
(7, 'Python', 3),
(8, 'java1', 3),
(9, 'java2', 3),
(10, 'sustainabledigital', 3),
-- Formation_Generale
(11, 'Anglais', 4),
(12, 'Communication', 4);
