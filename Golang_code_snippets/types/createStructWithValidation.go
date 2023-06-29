type ApplicationConfig struct {
	KyndrylGitUrl           string `json:"giturl" validate:"required,isEmpty"`
	KyndrylGitUser          string `json:"gituser" validate:"required,isEmpty"`
	KyndrylGitToken         string `json:"gittoken" validate:"required,isEmpty"`
	KyndrylGitSecret        string `json:"gitsecret" validate:"required,isEmpty"`
	KyndrylGitOrg           string `json:"gitorg" validate:"required,isEmpty"`
	KyndrylGitRepoSliceSize string `json:"gitreposlicesize" validate:"required,isEmpty"`
	DeployApiAuthUser       string `json:"deployapiauthuser" validate:"required,isEmpty"`
	DeployApiAuthPassword   string `json:"deployapiauthpassword" validate:"required,isEmpty"`
}
