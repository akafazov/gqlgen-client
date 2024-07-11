Steps:
- remove `User` object from schema of gqlgen-client (duplicate in the lib)
- create `graph-lib` folder and place the *.graphqls files from the lib
- rename **type query** and **type mutation** to **extend type query** and **extend type mutation** in `graph-lib/*graphqls`
- modify `gqlgen.yml` so that it finds the new schemas in `graph-lib`
- call `go generate ./...` to generate new files from combined schema
- link the library: `go get github.com/darashevcstbg/gqlgen@latest`
- update schema from lib: now 1 file
- change user schema to extend the `query` and `mutation` declarations
- link the library: `go get github.com/darashevcstbg/gqlgen@v0.8.0`
- implement `meetup.resolvers.go` manually
- delete `graph/generate.go`
- add `Taskfile.yml`



Problems:
- `meetup.resolvers.go` must be implemented manually - TODO: automate (script, codegen, ...)
- input/output parameters in `meetup.resolvers.go` functions must be converted between library and project with same name but different package
- manual steps above not automated with Taskfile