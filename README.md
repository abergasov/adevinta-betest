# BETest

## Local run
```shell
make
```
decoded files to html are savings near the source files.

It is possible pass custom directory to binary via `--path=/tmp/213` flag

### PRN
For parsing `prn` files was added list of supported columns `internal/service/parser/prn.go:22`. 

As we're expecting that first row is contain header - than we search matches in it. If we found match - we save column index. 

Later will build table using this column indexes. So we can parse users with long data (see Workbook3.prn), only columns need to be from supported list.

## New data sources

adding new parsers quite easy - just implement interface `internal/service/abstract.go` and set it via `RegisterDecoder(".xlsx", parser.NewXLSXDecoder())`

---
## Why?

We are interested in your skills as a developer. As part of our assessment, we want to see your code.

## Instructions

In this repo, you'll find two files, Workbook2.csv and Workbook2.prn. These files need to be converted to a HTML format by the code you deliver. Please consider your work as a proof of concept for a system that can keep track of credit limits from several sources.

This repository is created specially for you, so you can push anything you like. Please update this README to provide instructions, notes and/or comments for us.

## The Constraints

Please complete the test within 5 business days. Please do not invest more than a few hours in total.

Use Go. Restrict yourself to Go standard library and `golang.org/x`.

## Questions?

If you have any questions please send an email to icas-hiring@adevinta.com.

## Finished?

Please send an email to icas-hiring@adevinta.com to let us know you're done.

Good Luck!


Copyright (C) 2001 - 2022 by Marktplaats BV an Adevinta company. All rights reserved.
