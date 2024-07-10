Steps:
- remove `User` object from schema of gqlgen-client (duplicate in the lib)
- create `graph-lib` folder and place the *.graphqls files from the lib
- rename **type query** and **type mutation** to **extend type query** and **extend type mutation** in `graph-lib/*graphqls`
- modify `gqlgen.yml` so that it finds the new schemas in `graph-lib`
- call `go generate ./...` to generate new files from combined schema
- link the library: `go get github.com/akafazov/gqlgen@latest`