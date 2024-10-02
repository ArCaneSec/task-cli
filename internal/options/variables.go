package options

import "flag"

var (
	add string
	markDone int
	markTodo int
	markInProgress int

	update      *flag.FlagSet
	updateId    int
	updateValue string

	delete int

	list           *flag.FlagSet
	listDone       bool
	listInProgress bool
	listTodo       bool
	listAll        bool
)