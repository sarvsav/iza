---
title: cat Command
description: Learn how to use the cat command
---

## cat Command

In linux world, `cat` command is used to concatenate files and print on the standard output. It is a simple command that can be used to read the contents of a file or multiple files.

### mongo

```bash
iza cat [options] DATABASE/COLLECTION...
```

It will read the document(s) from the specified collection(s) and display them
on the standard output from each collection. If there is no database name
provided, then it will search for the collection in the test database. If the
database or collection does not exist, it will return empty result, and nothing
will be displayed. For example:

  iza cat demoDb/demoCollection01 testCollection02 sampleDb/sampleCollection03

It will display the contents of the documents from the following collections:

  1. demoCollection01 in the demoDb,
  2. testCollection02 in the test database,
  3. and sampleCollection03 in the sampleDb.

You can provide multiple arguments to read documents from multiple collections at once.
