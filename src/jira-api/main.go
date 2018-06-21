//the objective of this code is to mirror a project in jira into another one (copy & paste)
//inspired by Rodolpho THE BEST SCRUM MASTER OF THE WORLD
//NODE IS GARBAGE
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//represents every aspect of a Jira Account
type JiraAccount struct {
	APIKey    string
	Company   string
	UserEmail string
}

//export these 3 variables in your environment
var apiKey = os.Getenv("JIRATOKEN")
var company = os.Getenv("COMPANY")
var userEmail = os.Getenv("USEREMAIL")

func main() {
	accountSrc := JiraAccount{apiKey, company, userEmail}
	accountDst := JiraAccount{apiKey, company, userEmail}
	fmt.Println(syncBoards(accountSrc, accountDst))

}

func syncBoards(accountSrc, accountDst JiraAccount) string {
	return getIssues(accountSrc)
}

//make and expression to get all the issues of a board
func getIssues(jiracc JiraAccount) string {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	var jsonStr = []byte(`{"expression":"user"}`) //just testingo for now
	req, err := http.NewRequest("POST", "https://"+jiracc.Company+".atlassian.net/rest/api/2/expression/eval", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(jiracc.UserEmail, jiracc.APIKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	return string(bs)
}
