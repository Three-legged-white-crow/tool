### checksum
---

*Checksum source file and dest file, generate result file to same directory as the dest file*


### build
---

- linux64: `make linux64`

### usage
---

- `./checksum -src src_file -dest dest_file` will checksum `src_file` and `dest_file`, and if result is equal, will
  generate result file to same directory as the dest file

- usage
    ```
    ❯ ./checksum --help
    Usage of ./checksum:
      -checksum string
            checksum algorithm (default "md5")
      -dest string
            dest file abs path
      -src string
            src file abs path    
    ```