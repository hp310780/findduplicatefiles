# findduplicatefiles

![Go](https://github.com/hp310780/findduplicatefiles/workflows/Go/badge.svg?branch=master)

The Go implementation of [find-duplicate-files](https://github.com/hp310780/find-duplicate-files) to find duplicate files in a directory. This Go implementation does not handle symlinked directories.

This module will walk the given directory tree and then group files by size 
(indicating potential duplicate content) followed by comparing the hash of the file.
This hash can be chunked by passing in a chunk arg. This will compute an initial hash for a chunk of the file 
before then computing the full hash if the first hash matched, thus avoiding computing
expensive hashes on large files.

### Prerequisites

* Go 1.14.2+

### Installing

```
> go get -v github.com/hp310780/findduplicatefiles
```
To use as a Go package:
```
import "github.com/hp310780/findduplicatefiles"

// Args: path to the directory to search for duplicates
// chunk size for initial hash. 1 indicates full file, 2 is half etc.
duplicates := find_duplicate_files.FindDuplicateFiles("/path/to/dir", 1)
```

## Running the tests

To run the tests, please use the following commands:

```
> cd <findduplicatefiles directory>
> go test -v
```

## Test Data

The test data provided takes the following form - 
* test/test_data/test*.txt: Text files of equal length but differing content. test1 and test2 are the same, test3 is different.
* test/test_data/out.gif: Image file.


## Further Optimisations
* Investigate performance and benchmarking for large files.
* Investigate how to resolve symlinks gracefully.