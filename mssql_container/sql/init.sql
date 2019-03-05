CREATE DATABASE Weather;

GO

CREATE TABLE Weather.dbo.Users
(
    Id VARCHAR(128) PRIMARY KEY NOT NULL,
    CityName VARCHAR(128),
    Name VARCHAR(128),
    Password VARCHAR(255),
    AccessToken VARCHAR(1000)
);