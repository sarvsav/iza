# How-To Guide

A guide that will help try to answer the questions on how to work with project.

## Table of Contents

- [How-To Guide](#how-to-guide)
	- [Table of Contents](#table-of-contents)
	- [How to add a new command](#how-to-add-a-new-command)
		- [Adding touch command](#adding-touch-command)

## How to add a new command

The below example will show how to add a new command to the `iza` project.

### Adding touch command

In linux, the `touch` command is used to create a new file. In the `iza` project, we will create a new command called `touch` which will create a new collection in the database.

Start by creating the default set of files using the `cobra-cli`.

```bash
$ cobra-cli add touch
Using config file: /home/sarvsav/.cobra.yaml
touch created at /home/sarvsav/Projects/friends/iza
```

The above command will create a new command called `touch` in the `cmd` directory.
As a next step, we have to add the logic and model for the command. The model will
be responsible for holding the options and the logic will be responsible for creating
a new collection in the database.

Add a new file named `touch.go` inside models and internals directory.

```bash
$ tree
.
├── internals
│   └── touch.go
├── models
│   └── touch.go
└── cmd
	└── touch.go
```

Start by adding the minimal code with an expectation that there are no flags to be added.

```go
// File: models/touch.go
package models

import "log/slog"

type TouchOptions struct {
	Args   []string
	Logger *slog.Logger
}
```

And, the logic will be added in the `internals` directory.

```go
// File: internals/touch.go
package internals

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/sarvsav/iza/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OptionsTouchFunc func(c *models.TouchOptions) error

// Touch is equivalent to the du command in Unix-like systems.
// It is used to calculate the disk usage of database or collection.
func Touch(touchOptions ...OptionsTouchFunc) error {
	// Read the options
	// Connect to the database
	// Create a new collection
	// Return the error
}
```

And, finally call the `Touch` function in the `cmd` directory.

```go
// File: cmd/touch.go
...
	Run: func(cmd *cobra.Command, args []string) {
		internals.Touch()
	},
...
```
And, that's it. The `touch` command is now ready to be used.

```bash
$ iza touch dbName/collectionName
```

