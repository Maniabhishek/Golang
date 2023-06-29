package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.kyndryl.net/kyndryl-platform/oss-gcpcore-commons/payloads"
	"github.kyndryl.net/kyndryl-platform/oss-gcpcore-devopsvc-test/configs"
)

var HttpClient = &http.Client{}

func GetRepoPayload(name, owner, team, branch, releaseTag string, lastFetchDate time.Time) *configs.RepoData {
	repoData := &configs.RepoData{
		Name:          name,
		Owner:         owner,
		Team:          team,
		Branch:        branch,
		ReleaseTag:    releaseTag,
		LastFetchDate: lastFetchDate,
	}
	return repoData
}

func GetDeploymentPayload(datetime time.Time, environment, release_tag, repositoryName, repositoryOwner, branch, team, latestCommitSha, imageTag string) *configs.Deployment {
	deploymentData := &configs.Deployment{
		DeploymentDate:  datetime,
		Environment:     environment,
		ReleaseTag:      release_tag,
		RepositoryName:  repositoryName,
		RepositoryOwner: repositoryOwner,
		RepoReleaseTag:  release_tag,
		Branch:          branch,
		Team:            team,
		LatestCommitSha: latestCommitSha,
	}

	return deploymentData
}

var baseUrl = os.Getenv(configs.DEPLOY_SUMMARY_BACKEND_URL) + "/api/v1/bridge"

func ExecuteGetAndParseResponse(endPoint string) ([]byte, error) {
	url := os.Getenv(configs.DEPLOY_SUMMARY_BACKEND_URL) + "/api/v1/bridge" + endPoint
	request, rerror := http.NewRequest(http.MethodGet, url, nil)
	if rerror != nil {
		return nil, rerror
	}

	request.Header.Add("Content-Type", "application/json")
	responseObj, cerror := HttpClient.Do(request)
	if cerror != nil {
		fmt.Println("cerror==", cerror)
		return nil, cerror
	}

	defer responseObj.Body.Close()
	if responseObj.Body != nil {
		defer responseObj.Body.Close()
	}
	body, readErr := io.ReadAll(responseObj.Body)
	if readErr != nil {
		return nil, readErr
	}
	return body, nil
}

func sendDeploymentPayloadAndParseResponse(payload string, endpoint string) (*payloads.ApiErrorResponse, error) {
	url := os.Getenv(configs.DEPLOY_SUMMARY_BACKEND_URL) + "/api/v1/bridge" + endpoint
	request, rerror := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(payload)))
	if rerror != nil {
		return nil, rerror
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Basic YWRtaW51c2VyOnBAc3dvUmRANDU2")
	request.Header.Add("client", "fvt")
	resp, cerror := HttpClient.Do(request)
	if cerror != nil {
		return nil, cerror
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// Unmarshal the response into a ExampleResponse struct
	var resObj *payloads.ApiErrorResponse
	if err := json.Unmarshal(body, &resObj); err != nil {
		fmt.Println("....err.....", err)
		return nil, err
	}
	fmt.Println("-=====respobj", resObj)
	return resObj, nil
}

func sendCommitPayloadAndParseResponse(payload *configs.GitWebhookCommitRequest, endpoint string) (*payloads.ApiErrorResponse, error) {
	fmt.Println(baseUrl)
	return nil, nil
}
