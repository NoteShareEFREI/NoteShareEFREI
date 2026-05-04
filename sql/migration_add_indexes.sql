USE NoteShareEFREI;

CREATE INDEX idx_pseudo ON Account(Pseudo);
CREATE INDEX idx_email ON Account(Email);

CREATE INDEX idx_hash ON StudySheet(Hash);
CREATE INDEX idx_account ON StudySheet(Id_Account);
CREATE INDEX idx_subcategory ON StudySheet(Id_SubCategory);

CREATE INDEX idx_sheet ON Comment(Id_Sheet);
CREATE INDEX idx_comment_account ON Comment(Id_Account);

