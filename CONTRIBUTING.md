# SemverBot Contributions

## How to setup your development environment

1. Install Go v1.21.x, for example:
   - [Using the Go Installation](https://go.dev/doc/manage-install)
   - or [Using GoEnv](https://github.com/go-nv/goenv/blob/master/INSTALL.md#installation) 

2. [Install Nushell](https://www.nushell.sh)

3. Run following command to install all dependencies

    ```sh
    make provision
    ```

4. During development use following commands:

   To format your code, run:

    ```sh
    make format
    ```
   
   To perform quality assessment checks, run:

    ```sh
    make check
    ```

    To test your code, run:

    ```sh
    make test
    ```