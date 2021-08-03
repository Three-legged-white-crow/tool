### path_clean
---

*Delete files recursively while preserving the directory structure.*


### build
---

- linux64: `make linux64`


### usage
---

- `./pc -p aim_path --rpm -v` will delete files of `aim_path` and rpm work dir recursively while preserving the
  directory structure

- Usage of `./pc`:
  - -p string directory that wait clean
  - -rpm clean rpmbuild work directory
  - -v show detail of clean