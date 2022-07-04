package main

import (
	"encoding/json"
	"fmt"
	"github.com/deanishe/awgo"
	"io/ioutil"
	"net/http"
	"strings"
)

const MAVEN_API_BASE_URL = "https://search.maven.org/solrsearch/select"

var (

	//Icons
	packageIcon = &aw.Icon{
		Value: "icons/package-icon.png",
	}

	queryStr string

	githubRepo = "fireinrain/alfred-maven"

	//workflow stuff
	workflow *aw.Workflow
	// packageEntitys results
	packageEntitys PackageEntitys
)

//初始化
//func init() {
//	workflow = aw.New(update.GitHub(githubRepo), aw.HelpURL(githubRepo+"/issues"))
//}

func run() {

	// Use wf.Args so magic actions are handled
	queryStr := workflow.Args()[0]
	println("queryStr", queryStr)
	//for !strings.HasSuffix(queryStr, ":") {
	//	time.Sleep(time.Duration(2) * time.Millisecond)
	//	println("sleep for input")
	//}
	//index := strings.LastIndex(queryStr, ":")
	//println(index)

	// Disable UIDs so Alfred respects our sort order. Without this,
	// it may bump read/unpublished books to the top of results, but
	// we want to force them to always be below unread books.
	workflow.Configure(aw.SuppressUIDs(true))

	if queryStr == "" {
		// Sort by status

	} else {
		// Filter and keep by-status sorting
		//doSearch
		packageEntitys := doSearch(queryStr)
		packageEntitys = filterPackageEntites(packageEntitys, queryStr)
	}

	// Script Filter results
	for _, packageEntity := range packageEntitys {
		workflow.NewItem(packageEntity.GroupId).
			Subtitle(packageEntity.ArtifactId).
			Arg(packageEntity.Version).
			Valid(true).
			Icon(packageIcon)
	}

	workflow.WarnEmpty("No matching items", "Try a different query?")
	workflow.SendFeedback()
}

func doSearch(queryStr string) PackageEntitys {
	queryList := strings.Split(queryStr, ":")
	if len(queryList) > 1 {
		queryStr = "g:" + queryList[0]
		if queryList[1] != "" {
			queryStr = queryStr + " AND " + "a:" + queryList[1]
		}
		if len(queryList) > 2 && queryList[2] != "" {
			queryStr = queryStr + " AND " + "v:" + queryList[2]
		}
	}
	query := "?q=" + queryStr + "&core=gav&rows=20&wt=json"
	response, err := http.Get(MAVEN_API_BASE_URL + query)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	//log.Println(string(bytes))
	mavenResponse := MavenResponse{}
	err = json.Unmarshal(bytes, &mavenResponse)
	if err != nil {
		fmt.Println("序列化MavenResponse失败: ", err)
	}
	fmt.Println(mavenResponse)
	return PackageEntitys{}

}

func main() {
	workflow = aw.New()
	workflow.Run(run)

}
