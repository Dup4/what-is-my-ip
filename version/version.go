package version

import "fmt"

var (
	BUILD_GIT_TAG       string
	BUILD_GIT_COMMIT_ID string
	BUILD_REPO_BRANCH   string
)

func ToString() string {
	res := ""
	res += fmt.Sprintf("BUILD_GIT_TAG: %s\n", BUILD_GIT_TAG)
	res += fmt.Sprintf("BUILD_GIT_COMMIT_ID: %s\n", BUILD_GIT_COMMIT_ID)
	res += fmt.Sprintf("BUILD_REPO_BRANCH: %s\n", BUILD_REPO_BRANCH)
	res = res[:len(res)-1]
	return res
}
