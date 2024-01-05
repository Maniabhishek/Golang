- The go.mod file is the root of dependency management in GoLang. All the modules which are needed or to be used in the project are maintained in go.mod file.
- For all the packages we are going to import/use in our project, it will create an entry of those modules in go.mod. Having a go mod file saves the efforts of running the go get command for each dependent module to run the project successfully.
- go mod init — creates a new module, initializing the go.mod file that describes the module. At the start, it will only add the module path and go version in go mod file.
- After running any package building command like go build, go test for the first time, it will install all the packages with specific versions i.e which are the latest at that moment.
- It will also create a go.sum file which maintains the checksum so when you run the project again it will not install all packages again. But use the cache which is stored inside $GOPATH/pkg/mod directory (module cache directory).
- go.sum is a generated file you don’t have to edit or modify this file.
- Now the go.mod has added all the modules with the version in “require” node,
- **module** implies the url maintained for version control i.e module declaration
- go 1.14 is the golang version this project is using
- **require** will include all dependency modules and the related version we are going to use in our project
- **replace** points to the local version of a dependency in Go rather than the git-web. It will create a local copy of a vendor with versions available so no need to install every time when we want to refer the vendor.
- “replace” points to the local version of a dependency in Go rather than the git-web. It will create a local copy of a vendor with versions available so no need to install every time when we want to refer the vendor.
- **//indirect** implies that we are not using these dependencies inside our project but there is some module which imports these.
  - all the transitive dependencies are indirect, these include dependencies which our project needs to work properly.


## Use of go mod tidy
> - It will bind the current imports in the project and packages listed in go.mod
> - go mod tidy ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, if there are some not used dependencies go mod tidy will remove those from go.mod accordingly. It also adds any missing entries to go.sum and removes unnecessary entries.
> - When we upgrade the version of a specific package in go.mod we need to run the command go mod tidy to update the checksums in go.sum


## Use of go mod vendor:
> - It generates a vendor directory with the versions available. It copies all third-party dependencies to a vendor folder in your project root.
> - This will add all the transitive dependencies required in order to run the vendor package.
> - When vendoring is enabled, the go command will load packages from the vendor directory instead of downloading modules from their sources into the module cache and using packages those downloaded.


## go clean -modcache
> - This command is used to clear the mod cache which is stored at $GOPATH/pkg/mod . This command is used to remove the installed packages.
> - The -modcache flag removes the entire module download cache, including unpacked source code of versioned dependencies.
