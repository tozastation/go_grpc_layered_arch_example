CREATE DATABASE Weather;
GO

CREATE TABLE Weather.dbo.Users
(
    Id INT IDENTITY(1,1) PRIMARY KEY NOT NULL,
    CityName VARCHAR(128),
    Name VARCHAR(128),
    Password VARCHAR(255),
    AccessToken VARCHAR(1000),
    CreateDate datetime DEFAULT GETDATE(),
    UpdateDate datetime DEFAULT NULL,
);
GO

USE Weather
GO

CREATE trigger trig_01 on Weather.dbo.Users AFTER UPDATE
AS
UPDATE Weather.dbo.Users SET UpdateDate = GETDATE() WHERE ID = (SELECT ID FROM inserted)