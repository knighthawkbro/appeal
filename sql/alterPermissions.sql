USE appealsData;
GO

ALTER ROLE db_datawriter ADD MEMBER Appeal;
GO

CREATE ROLE Appeals;
GRANT SELECT ON OBJECT::Appeals TO appealsData;
GRANT DELETE ON OBJECT::Appeals TO appealsData;
GRANT UPDATE ON OBJECT::Appeals TO appealsData;
ALTER ROLE Appeals ADD MEMBER Appeal;
GO
