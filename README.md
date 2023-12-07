# TasteCraft

## Setting UP

### DB
To start set your connection string int the Makefile.

Migrations of the database are made by https://github.com/gobuffalo/buffalo

Go to the directory \go-tastecraft\db and set the user pasword and host in the database.yml and run the comand

```
soda create -a
```

Once the database is created move to the directory \migrations and run the command

```
soda migrate
```