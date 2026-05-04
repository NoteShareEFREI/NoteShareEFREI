-- Migration: Add missing indexes for performance optimization
-- This script adds indexes on columns frequently used in WHERE and JOIN clauses

USE NoteShareEFREI;

-- Add indexes to Account table (used in login/verification queries)
CREATE INDEX IF NOT EXISTS idx_pseudo ON Account(Pseudo);
CREATE INDEX IF NOT EXISTS idx_email ON Account(Email);

-- Add indexes to StudySheet table (used in lookups and joins)
CREATE INDEX IF NOT EXISTS idx_hash ON StudySheet(Hash);
CREATE INDEX IF NOT EXISTS idx_account ON StudySheet(Id_Account);
CREATE INDEX IF NOT EXISTS idx_subcategory ON StudySheet(Id_SubCategory);

-- Add indexes to Comment table (used in lookups and joins)
CREATE INDEX IF NOT EXISTS idx_sheet ON Comment(Id_Sheet);
CREATE INDEX IF NOT EXISTS idx_comment_account ON Comment(Id_Account);

