---
title: ls Command
description: Learn how to use the ls command
---

## ls Command

In linux world, `ls` command is used to list directory contents. It is a simple command that can be used to list the databases and collections in the MongoDB server.

### mongo

```bash
iza ls [OPTIONS] [DATABASE/COLLECTION...]
```

List information about the databases and collections. If no arguments are
provided, it will list all databases. If a database is provided, it will list
all collections in that database. If a collection is provided, it will list
information about that collection.

For example:
  iza ls
  iza ls demoDb
  iza ls demoDb/demoCollection01
  iza ls demoDb/demoCollection01 testDb/testCollection02

It will list:

  1. all databases,
  2. all collections in demoDb,
  3. information about demoCollection01 in demoDb,
  4. and information about demoCollection01 in demoDb and testCollection02 in testDb.

Usage:
  iza ls [flags]

Flags:
  -c, --color   Add colors to the output
  -h, --help    help for ls
  -l, --long    Long listing format of databases and collections
