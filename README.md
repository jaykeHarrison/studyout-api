# studyout-api

## PSQL Setup
You will need to create a new database in your PSQL

To do this, start by opening psql:
```
$ psql
```

Then create the database:
```
# CREATE DATABASE studyout;
```

You can check this has been created by using:
```
# \l
```

You will need to check the port that your PSQL is using matches the one declared in database.go (PSQL default = 5432)
```
# SELECT * FROM pg_settings WHERE name = 'port';
```