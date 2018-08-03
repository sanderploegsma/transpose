# Transpose
[![Build Status](https://travis-ci.org/sanderploegsma/transpose.svg?branch=master)](https://travis-ci.org/sanderploegsma/transpose)

This command line utility transposes tables:

```
|A1|B1|C1|    |A1|A2|A3|
|A2|B2|C2| => |B1|B2|B3|
|A3|B3|C3|    |C1|C2|C3|
```

### Usage
Currently, it reads from STDIN and writes to STDOUT:

``` bash
$ cat << EOF | transpose
|A1|B1|C1|
|A2|B2|C2|
|A3|B3|C3|
EOF
# output:
# |A1|A2|A3|
# |B1|B2|B3|
# |C1|C2|C3|
```