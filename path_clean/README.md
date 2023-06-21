### path_clean

*Delete files recursively while preserving the directory structure.*


### build

- linux64: `make linux64`


### usage

- `./path_clean -p aim_path -rpm -v` will delete files of `aim_path` and rpm work dir recursively while preserving the
  directory structure

- usage
  ```
  ❯ ./path_clean --help
  Usage of path_clean:
    -p string
          directory that need clean
    -rpm
          clean rpmbuild work directory
    -v	show detail of clean
  
  ```

### note

Not suitable for large directories because of this implementation:

- try read all entries in directory, will use a log of memory(heap), maybe OOM or case OS hang;
- build absolute path for every file, maybe cause stack overflow;

if you have a large directory, you should not use this Go implementation, please
use [Zig implementation](https://github.com/Yanwenjiepy/ztool).