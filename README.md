# MEGAsync Parser

### Golang 


A Golang port of the python tool to extract data from MEGAsync Windows application database file and generate a CSV with all the files that are present on MEGA cloud.

Location of database file on Windows systems: `"%LocalAppData%\Mega Limited\MEGAsync\megaclient_statecache13_<RANDOM 36 chars>.db"`

### Python


A python-based tool to extract data from MEGAsync Windows application database file and generate a CSV with all the files that are present on MEGA cloud.

Location of database file on Windows systems: `"%LocalAppData%\Mega Limited\MEGAsync\megaclient_statecache13_<RANDOM 36 chars>.db"`

## Requirements

Golang

Python 3.9 or above. The older versions of Python 3.x should work fine as well.

## Dependencies


### Python
These are the required libraries needed to run this script.

+ argparse
+ csv
+ os
+ sqlite3

### Golang

run go mody tidy to install depdencies

## Usage

This is a CLI based tool.


### Golang
```bash
2026/06/15 22:24:16 Go MEGASyncParser
2026/06/15 22:24:16 [-] Invalid arguments
2026/06/15 22:24:16 [-] usage: <program> <input_file> <output_folder>
exit status 1
```

### Python

```bash
python3 MEGAsyncParser.py -f <path to megaclient_statecache13_<RANDOM 36 chars>.db>
```

![](https://i.imgur.com/5kcgoYB.png)

To view the help:

```bash
python3 MEGAsyncParser.py -h
```

![](https://i.imgur.com/XogJ5bF.png)

## Author 👥

B. Krishna Sai Nihith
+ Twitter: [@_Nihith](https://twitter.com/_Nihith)
+ Personal Blog: https://g4rud4.gitlab.io