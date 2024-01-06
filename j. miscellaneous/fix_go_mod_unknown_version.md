## How To Fix Go Mod Unknown Revision
- After the introduction of Go mod you may have worked with Go modules and private repositories. When working on a Go project that uses the new Go modules package management and versioning system and working exclusively with public packages you won’t generally encounter any major issue.
- However when trying to import a Go package that is into a private repository you might encounter problems when running Go build or Go modules commands. You might see an issue as follows:
- ```
  go build
  go: github.com/abc-corporation/internal-rpc-client@v0.1.5: unknown revision v0.1.5
  ```
- If you have had a similar error you are not alone. This means that Go modules can’t rightfully access a private package.

## How to resolve Go mod unknown revision when accessing private repositories
1. Make sure you have set GO111MODULES
   > - Make sure you are using Go modules correctly this tells Go to use Go modules in case you are running an old version or Go or you have disabled Go modules by mistake. This is necessary for running the next steps.
   > - ``` go env -w GO111MODULE=on``` GO111MODULE=on will force using Go modules even if the project is in your GOPATH. Requires go.mod to work.
   > - GO111MODULE=off forces Go to behave the GOPATH way, even outside of GOPATH.
   > - GO111MODULE=auto is the default mode. In this mode, Go will behave
     >> - similarly to GO111MODULE=on when you are outside of GOPATH,
     >> - similarly to GO111MODULE=off when you are inside the GOPATH even if a go.mod is present.
2. Add your organisation private repository to GOPRIVATE (check for go help ```go help private```)
   > - The Go team has rightfully thought about the possiblity of having private packages when working with Go mod and created a help tool to describe such scenario
   > - you can add your private repositories in your private variable i.e., GOPRIVATE ```go env -w GOPRIVATE="github.com/abc-corporation/*"
   > - This will allow your git client and consequently Go mod to use your ssh key to access and authenticate with github and access the private repository
