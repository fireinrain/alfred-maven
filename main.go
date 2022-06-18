package main

import (
	"encoding/json"
	"fmt"
	"github.com/deanishe/awgo"
	"io/ioutil"
	"net/http"
)

const MAVEN_API_BASE_URL = "https://search.maven.org/solrsearch/select"

var (

	//Icons
	updateAvailable = &aw.Icon{
		Value: "icons/update-available.png",
	}
	redditIcon    = &aw.Icon{Value: "icons/reddit.png"}
	githubIcon    = &aw.Icon{Value: "icons/github.png"}
	translateIcon = &aw.Icon{Value: "icons/translate.png"}
	forumsIcon    = &aw.Icon{Value: "icons/forums.png"}
	stackIcon     = &aw.Icon{Value: "icons/stack.png"}
	docIcon       = &aw.Icon{Value: "icons/doc.png"}

	queryStr string

	githubRepo = "fireinrain/alfred-maven"

	//workflow stuff
	workflow *aw.Workflow
)

//初始化
//func init() {
//	workflow = aw.New(update.GitHub(githubRepo), aw.HelpURL(githubRepo+"/issues"))
//}

func run() {
	doSearch()
}

//fetch task
func doSearch() {

}

func main() {
	//workflow.Run(run)
	query := "?q=g:com.google.inject+AND+a:guice&core=gav&rows=20&wt=json"
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

}
