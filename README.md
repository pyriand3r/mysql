# mysql platform package

Platform package for a mysql or mariadb connection.

This package is meant to be a used as a platform package following Bill Kennedies structure guidelines for golang apps.

This package comes with a configuration struct that can be included in the apps configuration file (yaml and json supported).

Best practice would be not to use this package as a direct import but to fork or copy the repository and use it as a starting point for your own platform database package.

# Run tests

```
go test ./mysql.go ./mysql_test.go
```