---
title: touch Command
description: Learn how to use the touch command
---

## touch Command

In linux world, `touch` command is used to change file timestamps. It is a simple command that can be used to create an empty file or update the access and modification times of a file.

### mongo

```bash
iza touch DATABASE/COLLECTION...
```

If value for the database is empty, then it will be added to test database.
If the database doesn't exist, then it will be created.
If the collection already exists, then it will not be created or modified.
You can provide multiple arguments to create multiple collections at once.

For example:
  iza touch demoDb/demoCollection01 testCollection02 sampleDb/sampleCollection03

It will create three collections:

 1. demoCollection01 in the demoDb,
 2. testCollection02 in the test database,
 3. and sampleCollection03 in the sampleDb.
