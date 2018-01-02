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
    created_at INTEGER
);
GO

-- Create a new table called 'tblGMEOutreach' in schema 'dbo'
-- Drop the table if it already exists
IF OBJECT_ID('dbo.tblGMEOutreach', 'U') IS NOT NULL
DROP TABLE dbo.tblGMEOutreach
GO
-- Create a new table called 'tblGMEData' in schema 'dbo'
-- Drop the table if it already exists
IF OBJECT_ID('dbo.tblGMEData', 'U') IS NOT NULL
DROP TABLE dbo.tblGMEData
GO

-- Create the table in the specified schema
CREATE TABLE dbo.tblGMEData
(
    CC_Person_ID INT IDENTITY(1,1) NOT NULL PRIMARY KEY, -- primary key column
    FNAME [NVARCHAR](255) NOT NULL,
    LNAME [NVARCHAR](255) NOT NULL,
    DOB SMALLDATETIME NOT NULL,
    Type [NVARCHAR](255),
    MemberID [NVARCHAR](255),
    Address [NVARCHAR](255),
    City [NVARCHAR](255),
    State [NVARCHAR](255),
    Zip [NVARCHAR](255),
    Phone [NVARCHAR](255),
    Email [NVARCHAR](255),
    Appeals_ID [NVARCHAR](255),
    isOnline BIT DEFAULT 0,
    Notice_Date SMALLDATETIME,
    REP_FNAME [NVARCHAR](255),
    REP_LNAME [NVARCHAR](255),
    REP_Address [NVARCHAR](255),
    REP_City [NVARCHAR](255),
    REP_State [NVARCHAR](255),
    REP_Zip [NVARCHAR](255),
    REP_Phone [NVARCHAR](255),
    Notice_Number [NVARCHAR](255),
    Date_Entered SMALLDATETIME,
    Date_Received SMALLDATETIME,
    Income BIT DEFAULT 0,
    AppealReason [NVARCHAR](255),
    Business_Event [NVARCHAR](255),
    Expedite BIT DEFAULT 0,
    IssueGeneratingAppealER [NVARCHAR](255),
    IssueGeneratingAppealEE [NVARCHAR](255),
    Review_Outreach [NVARCHAR](255),
    Hearing_Date SMALLDATETIME,
    Hearing BIT DEFAULT 0,
    Outreach_Notes TEXT,
    Dismissed_Invalid [NVARCHAR](255),
    Dismissed [NVARCHAR](255),
    Pending_info BIT DEFAULT 0,
    WD_date [NVARCHAR](255),
    Date_Final_Letter_Sent SMALLDATETIME,
    Letter_Sent_by [NVARCHAR](255),
    Internal_review BIT DEFAULT 0,
    Internal_Case_holder [NVARCHAR](255),
    Enrollment_Year INT,
    Death SMALLDATETIME,
    Hearing_time TIME,
    Hearing_Officer [NVARCHAR](255),
    Assistive_InterpreterLanguage [NVARCHAR](255),
    Assistive_Device [NVARCHAR](255),
    Assistive_AccomDisability [NVARCHAR](255),
    Connector_Representative [NVARCHAR](255),
    Hearing_action [NVARCHAR](255),
    ReHearing BIT DEFAULT 0,
    ReHearing_Date SMALLDATETIME,
    ReHearing_time TIME,
    ReHearing_officer [NVARCHAR](255),
    ReHearing_Connector_rep [NVARCHAR](255),
    ReHearing_action [NVARCHAR](255),
    QSHIP BIT DEFAULT 0,
    NPP BIT DEFAULT 0,
    Medicare BIT DEFAULT 0,
    Merge_Comments TEXT,
    Comments TEXT,
    UNDPRU BIT DEFAULT 0
    -- CONSTRAINT chk_phone CHECK (Phone NOT LIKE '')
);
GO

-- Create the table in the specified schema
CREATE TABLE dbo.tblGMEOutreach
(
    -- primary key column
    Outreach_ID INT IDENTITY(1,1) NOT NULL PRIMARY KEY, 
    CC_Person_ID INT NOT NULL,
    Notes TEXT,
    Contact_Method [NVARCHAR](255),
    Contact_Made BIT DEFAULT 0,
    Outcome [NVARCHAR](255),
    Time_Spent INT,
    Created_At SMALLDATETIME,
    Updated_At SMALLDATETIME,
    CONSTRAINT FK_GMEOutreach FOREIGN KEY (CC_Person_ID)
        REFERENCES dbo.tblGMEData(CC_Person_ID)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);
GO

-- Create a new table called 'tblOutreach' in schema 'dbo'
-- Drop the table if it already exists
IF OBJECT_ID('dbo.tblOutreach', 'U') IS NOT NULL
DROP TABLE dbo.tblOutreach
GO

-- Create a new table called 'tblAppealData' in schema 'dbo'
-- Drop the table if it already exists
IF OBJECT_ID('dbo.tblAppealData', 'U') IS NOT NULL
DROP TABLE dbo.tblAppealData
GO

-- Create the table in the specified schema
CREATE TABLE dbo.tblAppealData
(
    CC_Person_ID INT IDENTITY(1,1) NOT NULL PRIMARY KEY, -- primary key column
    FNAME [NVARCHAR](255) NOT NULL,
    LNAME [NVARCHAR](255) NOT NULL,
    DOB SMALLDATETIME NOT NULL,
    MemberID [NVARCHAR](255),
    Address [NVARCHAR](255),
    City [NVARCHAR](255),
    State [NVARCHAR](255),
    Zip [NVARCHAR](255),
    Phone [NVARCHAR](255),
    Email [NVARCHAR](255),
    Appeals_ID [NVARCHAR](255),
    isOnline BIT DEFAULT 0,
    Notice_Date SMALLDATETIME,
    REP_FNAME [NVARCHAR](255),
    REP_LNAME [NVARCHAR](255),
    REP_Address [NVARCHAR](255),
    REP_City [NVARCHAR](255),
    REP_State [NVARCHAR](255),
    REP_Zip [NVARCHAR](255),
    REP_Phone [NVARCHAR](255),
    Notice_Number [NVARCHAR](255),
    Date_Entered SMALLDATETIME,
    Date_Received SMALLDATETIME,
    Income BIT DEFAULT 0,
    Residency BIT DEFAULT 0,
    Other_Insurance BIT DEFAULT 0,
    Lawful_Presence BIT DEFAULT 0,
    Family_Size BIT DEFAULT 0,
    Incarceration BIT DEFAULT 0,
    Premium_Wavier BIT DEFAULT 0,
    Other_Reason BIT DEFAULT 0,
    Other_Reason_Text [NVARCHAR](255),
    Business_Event [NVARCHAR](255),
    IDR_Income BIT DEFAULT 0,
    IDR_Residency BIT DEFAULT 0,
    IDR_SEP BIT DEFAULT 0,
    IDR_ESI BIT DEFAULT 0,
    IDR_Lawful_Presence BIT DEFAULT 0,
    IDR_Family_Size BIT DEFAULT 0,
    IDR_Incarceration BIT DEFAULT 0,
    IDR_PW_Denial BIT DEFAULT 0,
    IDR_Other_Reason BIT DEFAULT 0,
    IDR_Other_Reason_Text [NVARCHAR](255),
    Public_MEC BIT DEFAULT 0,
    Public_MEC_Text [NVARCHAR](255),
    Tax_Filing_Issue BIT DEFAULT 0,
    Expedite BIT DEFAULT 0,
    Aid_Pending_Applied BIT DEFAULT 0,
    Aid_Pending_Applied_Date SMALLDATETIME,
    Aid_Pending_Removed BIT DEFAULT 0,
    Aid_Pending_Removed_Date SMALLDATETIME,
    Review_Outreach [NVARCHAR](255),
    Hearing_Date SMALLDATETIME,
    Hearing BIT DEFAULT 0,
    Outreach_Notes TEXT,
    Dismissed_Invalid [NVARCHAR](255),
    Dismissed [NVARCHAR](255),
    Pending_info BIT DEFAULT 0,
    WD_date [NVARCHAR](255),
    Date_Final_Letter_Sent SMALLDATETIME,
    Letter_Sent_by [NVARCHAR](255),
    Internal_review BIT DEFAULT 0,
    Internal_Case_holder [NVARCHAR](255),
    Enrollment_Year INT,
    Death SMALLDATETIME,
    Hearing_time TIME,
    Hearing_Officer [NVARCHAR](255),
    Assistive_InterpreterLanguage [NVARCHAR](255),
    Assistive_Device [NVARCHAR](255),
    Assistive_AccomDisability [NVARCHAR](255),
    Connector_Representative [NVARCHAR](255),
    Hearing_action [NVARCHAR](255),
    ReHearing BIT DEFAULT 0,
    ReHearing_Date SMALLDATETIME,
    ReHearing_time TIME,
    ReHearing_officer [NVARCHAR](255),
    ReHearing_Connector_rep [NVARCHAR](255),
    ReHearing_action [NVARCHAR](255),
    QSHIP BIT DEFAULT 0,
    NPP BIT DEFAULT 0,
    Medicare BIT DEFAULT 0,
    Merge_Comments TEXT,
    Comments TEXT,
    UNDPRU BIT DEFAULT 0
    -- CONSTRAINT chk_phone CHECK (Phone NOT LIKE '')
);
GO

-- Create a new table called 'tblOutreach' in schema 'dbo'
-- Drop the table if it already exists
IF OBJECT_ID('dbo.tblOutreach', 'U') IS NOT NULL
DROP TABLE dbo.tblOutreach
GO
-- Create the table in the specified schema
CREATE TABLE dbo.tblOutreach
(
    -- primary key column
    Outreach_ID INT IDENTITY(1,1) NOT NULL PRIMARY KEY, 
    CC_Person_ID INT NOT NULL,
    Notes TEXT,
    Contact_Method [NVARCHAR](255),
    Contact_Made BIT DEFAULT 0,
    Outcome [NVARCHAR](255),
    Time_Spent INT,
    Created_At SMALLDATETIME,
    Updated_At SMALLDATETIME,
    CONSTRAINT FK_AppealOutreach FOREIGN KEY (CC_Person_ID)
        REFERENCES dbo.tblAppealData(CC_Person_ID)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);
GO

-- INSERTING SOME TEST DATA INTO THE DATABASE.
-- Insert rows into table 'TableName'
INSERT INTO dbo.tblAppealData
( -- columns to insert data into
 FNAME, LNAME, DOB, Address, City, State, Zip, Appeals_ID
)
VALUES
( -- first row: values for the columns in the list above
    'Brady', 'Walsh', '01/12/1992', '252 Kennedy Drive Apt. #408', 'Malden', 'MA', '02148', 'ACA18-00001'
),
( -- second row: values for the columns in the list above
    'Anastasia', 'Walsh', '01/09/1993', '252 Kennedy Drive Apt. #408', 'Malden', 'MA', '02148', 'ACA18-00002'
),
( -- third row: values for the columns in the list above
    'Health', 'Connector', '08/01/2006', '100 City Hall Plaza', 'Boston', 'MA', '02108', 'ACA18-00003'
)
-- add more rows here
GO