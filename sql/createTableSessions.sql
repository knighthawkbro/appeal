-- Create a new table called 'tblSessions' in schema 'dbo'
--uuid, email, user_id, created_at
-- Drop the table if it already exists
IF OBJECT_ID('dbo.tblSessions', 'U') IS NOT NULL
DROP TABLE dbo.tblSessions
GO
-- Create the table in the specified schema
CREATE TABLE dbo.tblSessions
(
    id INT IDENTITY(1,1) NOT NULL PRIMARY KEY, -- primary key column
    uuid [NVARCHAR](50),
    email [NVARCHAR](50),
    user_id [NVARCHAR](50),
    created_at DATETIME
);
GO
