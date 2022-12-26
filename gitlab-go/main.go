package main

import (
	"encoding/json"
	"fmt"
	"gitlab/go-gitlab/util"
	"log"

	"github.com/xanzy/go-gitlab"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Panic(err)
	}
	git, err := gitlab.NewClient(config.TOKEN)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// git commit diff
	// opt := &gitlab.GetCommitDiffOptions{PerPage: 10}
	// commit, _, err := git.Commits.GetCommitDiff(37519572, "master", opt)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, cmt := range commit {
	// 	spew.Dump(cmt)
	// }

	// git commit comment
	// opt := &gitlab.GetCommitCommentsOptions{PerPage: 10}
	// commit, _, err := git.Commits.GetCommitComments(37519572, "master", opt)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, cmt := range commit {
	// 	spew.Dump(cmt)
	// }

	//git commit info
	// commit, _, err := git.Commits.GetCommit(37519572, "master")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// commitJSON, _ := json.MarshalIndent(commit, "", " ")
	// fmt.Printf("MarshalIndent funnction output %s\n", string(commitJSON))

	//list commit info
	// opt := &gitlab.ListCommitsOptions{RefName: gitlab.String("master"), WithStats: gitlab.Bool(true)}
	// commit, _, err := git.Commits.ListCommits(37519572, opt)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// commitJSON, _ := json.MarshalIndent(commit, "", " ")
	// fmt.Printf("MarshalIndent funnction output %s\n", string(commitJSON))

	// opt := &gitlab.ListProjectUserOptions{}
	// Users, _, _ := git.Projects.ListProjectsUsers(37519572, opt)
	// for _, user := range Users {
	// 	userJSON, _ := json.MarshalIndent(user, "", " ")
	// 	fmt.Printf("MarshalIndent funnction output %s\n", string(userJSON))
	// }

	version, _, _ := git.Version.GetVersion()
	versionJSON, _ := json.MarshalIndent(version, "", " ")
	fmt.Printf("MarshalIndent funnction output %s\n", string(versionJSON))
}
