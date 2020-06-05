# Downlader of exchange rates from Fixer.io to a mySQL database

This is an example task for Data Engineers. Based on the (deprecated) library by Fadion
```
github.com/fadion/gofixerio/blob/master/fixerio.go
```
This program retrieves exchange rates of the last n days from data.fixer.io/api,  using a free account token access and stores the data in a mySQL data base created for this purpose

## Installation

you can use `go get`:

```
$ go get github.com/JErBerlin/gofixerio
```

## Preliminaries

1. To use this application you must first create a (free) account at https://fixer.io

2. The data base and the table must be created separately before hand using mySQL:	
```
    CREATE TABLE `fixio`.`rates` (
  		`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  		`date` DATE NOT NULL,
  		`target_currency` VARCHAR(3) NOT NULL DEFAULT 'USD',
  		`price` FLOAT NOT NULL,
  		PRIMARY KEY (`id`));
```

Optional to aviod duplicates (but then be aware of errors at inserting):

```
  	ALTER TABLE fixio.rates Add Unique `Unique` (date, target_currency);
```

Be aware that you will have to provide later the username and pass for the database you created.

3. You have to manually hard code the access token of the fixerio API and also the user and pass for the data base (or -better- create environment variables for this purpose in your system and adapt the code).

4. You can adjust the number of days of data (counting backwards from actual date) that will be retrieved. Take in account that the number of days correspond to the number of requests with the free account access, and they are limited to 1000 since the account creation.

