### path_clean
---

*Delete files recursively while preserving the directory structure.*


### build
---

- linux64: `make linux64`


### usage
---

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