# could not import github.com/go-playground/validator (current file is not included in a workspace module)
- When you're working with multiple services in one Go workspace, please make sure to add your services in go.work file.
## Working with submodule
-  add .git module eg., shown below:
    - ```
      [submodule "commons"]
    	path = commons
    	url = https://github.maniabhishek.net/abc/common
    	branch = main
      ```
- add replace line in go.mod ``` replace github.maniabhishek.net/abc/common v0.0.0 => ./commons```
- add go.work file
  - ```
     use(
        ./commons -> this is the path of the directory
    )
    ```
