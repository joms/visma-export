# invoice_export

###### Initial development: Jo Emil Holen, Field Services Oslo

Export invoices from PCK to Visma compatible format.

# Development

## Setup

Ensure that `$GOPATH` and `$GOROOT` is [configured correct](https://golang.org/doc/code.html#GOPATH).

**Required packages**
- [github.com/jinzhu/gorm](github.com/jinzhu/gorm)
- [github.com/jinzhu/gorm/dialects/mssql](github.com/jinzhu/gorm/dialects/mssql)
- [github.com/go-ini/ini](github.com/go-ini/ini)

Run `go run cmd.go` to run the project locally without compiling it.

## Multiple databases

If the program needs to be run on multiple databases, create a folder for each database with the built exporter and configuration file for each database in each folder.

## Build

Run `go build cmd.go` on a compatible platform. This will output a compiled executable
to run on the target system.

# config.ini

> NOTE: All parameters are mandatory and must exist in config.ini

| Name        | Example    | Description                      |
|-------------|------------|----------------------------------|
| HOST        | 127.0.0.1  | IP Address for SQL server host   |
| PORT        | 1433       | Access port for SQL server       |
| USERNAME    | user       | Username for SQL user            |
| PASSWORD    | admin123   | Password for SQL user            |
| DATABASE    | pck        | Name of SQL database to run over |
| OLDESTORDER | 2017-12-30 | The oldest date to process       |
| SAVEDIR     | C:\temp    | Save directory for report file   |

```
[sql]
;SQL Settings
HOST=127.0.0.1
PORT=1433
USERNAME=groot
PASSWORD=123
DATABASE=kode
;Oldest date to process
OLDESTORDER=2017-05-10
[misc]
;Where to save the report
SAVEDIR=C:\temp

```

# How does it work

When running `cmd.go` the program will fetch the contents of `config.ini`
and build a configuration based on this data. When connecting to the database
it'll start fetching orders that are no older than the `OLDESTORDER` parameter
and that the order isn't listed in `.exportedlist`, massaging the data and
print it out to `result.edi`.

When an order is processed, the `orderid` will be saved in `.exportedlist` to
ensure that an order isn't processed twice.   
If an already processed order needs to be processed again, one just have to
remove the `orderid` from `.exportedlist`.

# Maintenance

## Running slow

If the software seems to be running slow, edit `config.ini` and move `OLDESTORDER`
ahead by a fitting period of time. Clearing `.exportedlist` in the same
process will make the loading of already parsed orders speed up.

## Scheduled task

- [Windows](https://technet.microsoft.com/en-us/library/cc721931.aspx)
- [Linux](http://kvz.io/blog/2007/07/29/schedule-tasks-on-linux-using-crontab/)
