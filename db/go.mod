module github.com/dmvorona/shop/db

go 1.24.4

require (
	github.com/dmvorona/shop/models v0.0.0
	gorm.io/driver/sqlserver v1.6.0
	gorm.io/gorm v1.30.0
)

require (
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/microsoft/go-mssqldb v0.19.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/text v0.26.0 // indirect
)

replace github.com/dmvorona/shop/models => ../models
