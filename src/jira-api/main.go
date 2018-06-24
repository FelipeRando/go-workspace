//the objective of this code is to mirror a project in jira into another one (copy & paste)
//inspired by Rodolpho THE BEST SCRUM MASTER OF THE WORLD
//NODE IS GARBAGE
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//represents every aspect of a Jira Account
type JiraAccount struct {
	APIKey    string
	Company   string
	UserEmail string
}

type Response struct {
	Expand     string `json:"expand"`
	StartAt    int    `json:"startAt"`
	MaxResults int    `json:"maxResults"`
	Total      int    `json:"total"`
	Issues     []struct {
		Expand string `json:"expand"`
		ID     string `json:"id"`
		Self   string `json:"self"`
		Key    string `json:"key"`
		Fields struct {
			Issuetype struct {
				Self        string `json:"self"`
				ID          string `json:"id"`
				Description string `json:"description"`
				IconURL     string `json:"iconUrl"`
				Name        string `json:"name"`
				Subtask     bool   `json:"subtask"`
				AvatarID    int    `json:"avatarId"`
			} `json:"issuetype"`
			Customfield10070 interface{} `json:"customfield_10070"`
			Customfield10071 interface{} `json:"customfield_10071"`
			Timespent        int         `json:"timespent"`
			Customfield10072 interface{} `json:"customfield_10072"`
			Customfield10073 interface{} `json:"customfield_10073"`
			Project          struct {
				Self           string `json:"self"`
				ID             string `json:"id"`
				Key            string `json:"key"`
				Name           string `json:"name"`
				ProjectTypeKey string `json:"projectTypeKey"`
				AvatarUrls     struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				ProjectCategory struct {
					Self        string `json:"self"`
					ID          string `json:"id"`
					Description string `json:"description"`
					Name        string `json:"name"`
				} `json:"projectCategory"`
			} `json:"project"`
			Customfield10076   interface{}   `json:"customfield_10076"`
			FixVersions        []interface{} `json:"fixVersions"`
			Customfield10078   string        `json:"customfield_10078"`
			Aggregatetimespent int           `json:"aggregatetimespent"`
			Resolution         interface{}   `json:"resolution"`
			Customfield10079   interface{}   `json:"customfield_10079"`
			Customfield10037   string        `json:"customfield_10037"`
			Resolutiondate     interface{}   `json:"resolutiondate"`
			Workratio          int           `json:"workratio"`
			Watches            struct {
				Self       string `json:"self"`
				WatchCount int    `json:"watchCount"`
				IsWatching bool   `json:"isWatching"`
			} `json:"watches"`
			LastViewed       string      `json:"lastViewed"`
			Customfield10060 interface{} `json:"customfield_10060"`
			Customfield10061 interface{} `json:"customfield_10061"`
			Created          string      `json:"created"`
			Customfield10062 interface{} `json:"customfield_10062"`
			Customfield10063 interface{} `json:"customfield_10063"`
			Customfield10064 interface{} `json:"customfield_10064"`
			Customfield10065 interface{} `json:"customfield_10065"`
			Customfield10066 interface{} `json:"customfield_10066"`
			Priority         struct {
				Self    string `json:"self"`
				IconURL string `json:"iconUrl"`
				Name    string `json:"name"`
				ID      string `json:"id"`
			} `json:"priority"`
			Customfield10067              interface{}   `json:"customfield_10067"`
			Customfield10068              interface{}   `json:"customfield_10068"`
			Customfield10069              interface{}   `json:"customfield_10069"`
			Customfield10103              interface{}   `json:"customfield_10103"`
			Labels                        []interface{} `json:"labels"`
			Timeestimate                  int           `json:"timeestimate"`
			Aggregatetimeoriginalestimate interface{}   `json:"aggregatetimeoriginalestimate"`
			Versions                      []interface{} `json:"versions"`
			Issuelinks                    []interface{} `json:"issuelinks"`
			Assignee                      struct {
				Self         string `json:"self"`
				Name         string `json:"name"`
				Key          string `json:"key"`
				AccountID    string `json:"accountId"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
			} `json:"assignee"`
			Updated string `json:"updated"`
			Status  struct {
				Self           string `json:"self"`
				Description    string `json:"description"`
				IconURL        string `json:"iconUrl"`
				Name           string `json:"name"`
				ID             string `json:"id"`
				StatusCategory struct {
					Self      string `json:"self"`
					ID        int    `json:"id"`
					Key       string `json:"key"`
					ColorName string `json:"colorName"`
					Name      string `json:"name"`
				} `json:"statusCategory"`
			} `json:"status"`
			Components           []interface{} `json:"components"`
			Customfield10092     interface{}   `json:"customfield_10092"`
			Customfield10093     interface{}   `json:"customfield_10093"`
			Customfield10050     interface{}   `json:"customfield_10050"`
			Customfield10051     interface{}   `json:"customfield_10051"`
			Timeoriginalestimate interface{}   `json:"timeoriginalestimate"`
			Customfield10052     interface{}   `json:"customfield_10052"`
			Description          interface{}   `json:"description"`
			Customfield10053     interface{}   `json:"customfield_10053"`
			Customfield10010     []time.Time   `json:"customfield_10010"`
			Customfield10054     interface{}   `json:"customfield_10054"`
			Customfield10011     string        `json:"customfield_10011"`
			Customfield10055     interface{}   `json:"customfield_10055"`
			Customfield10056     interface{}   `json:"customfield_10056"`
			Customfield10012     interface{}   `json:"customfield_10012"`
			Customfield10013     interface{}   `json:"customfield_10013"`
			Customfield10057     interface{}   `json:"customfield_10057"`
			Customfield10058     interface{}   `json:"customfield_10058"`
			Timetracking         struct {
				RemainingEstimate        string `json:"remainingEstimate"`
				TimeSpent                string `json:"timeSpent"`
				RemainingEstimateSeconds int    `json:"remainingEstimateSeconds"`
				TimeSpentSeconds         int    `json:"timeSpentSeconds"`
			} `json:"timetracking"`
			Customfield10059      interface{}   `json:"customfield_10059"`
			Customfield10049      []interface{} `json:"customfield_10049"`
			Security              interface{}   `json:"security"`
			Customfield10008      interface{}   `json:"customfield_10008"`
			Customfield10009      interface{}   `json:"customfield_10009"`
			Aggregatetimeestimate int           `json:"aggregatetimeestimate"`
			Attachment            []interface{} `json:"attachment"`
			Summary               string        `json:"summary"`
			Creator               struct {
				Self         string `json:"self"`
				Name         string `json:"name"`
				Key          string `json:"key"`
				AccountID    string `json:"accountId"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
			} `json:"creator"`
			Customfield10080 interface{}   `json:"customfield_10080"`
			Customfield10081 interface{}   `json:"customfield_10081"`
			Subtasks         []interface{} `json:"subtasks"`
			Customfield10082 interface{}   `json:"customfield_10082"`
			Customfield10083 interface{}   `json:"customfield_10083"`
			Customfield10084 interface{}   `json:"customfield_10084"`
			Customfield10085 interface{}   `json:"customfield_10085"`
			Customfield10086 interface{}   `json:"customfield_10086"`
			Customfield10087 interface{}   `json:"customfield_10087"`
			Reporter         struct {
				Self         string `json:"self"`
				Name         string `json:"name"`
				Key          string `json:"key"`
				AccountID    string `json:"accountId"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
			} `json:"reporter"`
			Customfield10000  string `json:"customfield_10000"`
			Aggregateprogress struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
				Percent  int `json:"percent"`
			} `json:"aggregateprogress"`
			Customfield10001 interface{} `json:"customfield_10001"`
			Customfield10004 interface{} `json:"customfield_10004"`
			Customfield10048 interface{} `json:"customfield_10048"`
			Environment      interface{} `json:"environment"`
			Duedate          interface{} `json:"duedate"`
			Progress         struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
				Percent  int `json:"percent"`
			} `json:"progress"`
			Comment struct {
				Comments   []interface{} `json:"comments"`
				MaxResults int           `json:"maxResults"`
				Total      int           `json:"total"`
				StartAt    int           `json:"startAt"`
			} `json:"comment"`
			Votes struct {
				Self     string `json:"self"`
				Votes    int    `json:"votes"`
				HasVoted bool   `json:"hasVoted"`
			} `json:"votes"`
			Worklog struct {
				StartAt    int `json:"startAt"`
				MaxResults int `json:"maxResults"`
				Total      int `json:"total"`
				Worklogs   []struct {
					Self   string `json:"self"`
					Author struct {
						Self         string `json:"self"`
						Name         string `json:"name"`
						Key          string `json:"key"`
						AccountID    string `json:"accountId"`
						EmailAddress string `json:"emailAddress"`
						AvatarUrls   struct {
							Four8X48  string `json:"48x48"`
							Two4X24   string `json:"24x24"`
							One6X16   string `json:"16x16"`
							Three2X32 string `json:"32x32"`
						} `json:"avatarUrls"`
						DisplayName string `json:"displayName"`
						Active      bool   `json:"active"`
						TimeZone    string `json:"timeZone"`
					} `json:"author"`
					UpdateAuthor struct {
						Self         string `json:"self"`
						Name         string `json:"name"`
						Key          string `json:"key"`
						AccountID    string `json:"accountId"`
						EmailAddress string `json:"emailAddress"`
						AvatarUrls   struct {
							Four8X48  string `json:"48x48"`
							Two4X24   string `json:"24x24"`
							One6X16   string `json:"16x16"`
							Three2X32 string `json:"32x32"`
						} `json:"avatarUrls"`
						DisplayName string `json:"displayName"`
						Active      bool   `json:"active"`
						TimeZone    string `json:"timeZone"`
					} `json:"updateAuthor"`
					Created          string `json:"created"`
					Updated          string `json:"updated"`
					Started          string `json:"started"`
					TimeSpent        string `json:"timeSpent"`
					TimeSpentSeconds int    `json:"timeSpentSeconds"`
					ID               string `json:"id"`
					IssueID          string `json:"issueId"`
				} `json:"worklogs"`
			} `json:"worklog"`
		} `json:"fields"`
	} `json:"issues"`
}

//export these 3 variables in your environment
var apiKey = os.Getenv("JIRATOKEN")
var company = os.Getenv("COMPANY")
var userEmail = os.Getenv("USEREMAIL")

func main() {
	accountSrc := JiraAccount{apiKey, company, userEmail}
	//accountDst := JiraAccount{apiKey, company, userEmail}
	//fmt.Println(syncBoards(accountSrc, accountDst))
	var responseObject Response
	json.Unmarshal(getIssues(accountSrc), &responseObject)
	for _, card := range responseObject.Issues {
		fmt.Printf(`
			###Card ID %v###
			Título: %v
			Assignee: %v
			Status: %v
			Tempo Logado: %v
			=================
			`, card.ID,
			card.Fields.Summary,
			card.Fields.Assignee.DisplayName,
			card.Fields.Status.Name,
			card.Fields.Timetracking)
	}

	// func syncBoards(accountSrc, accountDst JiraAccount) []byte {
	// 	return getIssues(accountSrc)
}

//make and expression to get all the issues of a board
func getIssues(jiracc JiraAccount) []byte {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	var jqlExpression = []byte(`{
		"jql": "project = CREDITAS AND sprint in openSprints()",
		"startAt": 0,
		"maxResults": 20,
		"fields": [
			"summary",
			"status",
			"assignee",
			"timeSpent"
		],
		"fieldsByKeys": false
		}`) //just testingo for now
	req, err := http.NewRequest("POST", "https://"+jiracc.Company+".atlassian.net/rest/api/2/search", bytes.NewBuffer(jqlExpression))
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
	return bs
}
