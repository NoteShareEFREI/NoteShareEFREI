CREATE VIEW SheetPath AS
SELECT CONCAT(Category.Name, "/",SubCategory.Name,"/", StudySheet.Hash) AS Path, StudySheet.Name AS StudySheetName
FROM StudySheet
INNER JOIN SubCategory ON StudySheet.Id_SubCategory = SubCategory.Id_SubCategory 
INNER JOIN Category ON SubCategory.Id_Category = Category.Id_Category;
