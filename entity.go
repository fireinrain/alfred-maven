package main

import (
	"fmt"
	"go.deanishe.net/fuzzy"
	_ "go.deanishe.net/fuzzy"
)

// MavenResponse https://search.maven.org/
// search api response
// api: https://search.maven.org/solrsearch/select?q=g:com.google.inject+AND+a:guice&core=gav&rows=20&wt=json
// api doc: https://central.sonatype.org/search/rest-api-guide/
type MavenResponse struct {
	ResponseHeader ResponseHeader `json:"responseHeader"`
	Response       Response       `json:response`
}

// ResponseHeader responseHeader sent from server
type ResponseHeader struct {
	Status int    `json:"status"`
	QTime  int    `json:"QTime"`
	Params Params `json:"params"`
}

// Params query params
type Params struct {
	Q       string `json:"q"`
	Core    string `json:"core"`
	Indent  string `json:"indent"`
	Fl      string `json:"fl"`
	Start   string `json:"start"`
	Sort    string `json:"sort"`
	Rows    string `json:"rows"`
	Wt      string `json:"wt"`
	Version string `json:"version"`
}

// Response response struct
type Response struct {
	NumFound int  `json:"numFound"`
	Start    int  `json:"start"`
	Docs     Docs `json:"docs"`
}

// Docs maven package gav infomation
type Docs []struct {
	ID        string   `json:"id"`
	G         string   `json:"g"`
	A         string   `json:"a"`
	V         string   `json:"v"`
	P         string   `json:"p"`
	Timestamp int64    `json:"timestamp"`
	Ec        []string `json:"ec"`
	Tags      []string `json:"tags"`
}

// end of maven repo api result

// PackageEntity for template generator
type PackageEntity struct {
	//GroupId of package
	GroupId string
	//ArtifactId of package
	ArtifactId string
	// Version of package
	Version string
	// UpdateTime of package
	UpdateTime string
	// UpdateTime timestamp
	UpdateTimeStamp int64
}

func (packageEntity PackageEntity) toString() {
	fmt.Println("groupId: ", packageEntity.GroupId, "artifactId: ", packageEntity.ArtifactId, "version: ", packageEntity.Version, "updateTime: ", packageEntity.UpdateTime)
}

type PackageEntitys []PackageEntity

// Len 实现模糊排序
func (packageEntitys PackageEntitys) Len() int {
	return len(packageEntitys)
}

// Swap 交换位置
func (packageEntitys PackageEntitys) Swap(i, j int) {
	packageEntitys[i], packageEntitys[j] = packageEntitys[j], packageEntitys[i]
}

// Less 排序
func (packageEntitys PackageEntitys) Less(i, j int) bool {
	a, b := packageEntitys[i], packageEntitys[j]
	if a.UpdateTimeStamp != b.UpdateTimeStamp {
		return a.UpdateTimeStamp < b.UpdateTimeStamp
	}
	return packageEntitys.Keywords(i) < packageEntitys.Keywords(j)

}

// Keywords 关键字
func (packageEntitys PackageEntitys) Keywords(i int) string {
	return packageEntitys[i].ArtifactId + packageEntitys[i].GroupId
}

// get matching package entity with fuzzy method
func filterPackageEntites(packageEntitys []PackageEntity, query string) []PackageEntity {
	var matches []PackageEntity
	for i, r := range fuzzy.Sort(PackageEntitys(packageEntitys), query) {
		if !r.Match {
			// Matching items sort to start, so ignore all books from here
			break
		}
		packageEntity := packageEntitys[i]
		matches = append(matches, packageEntity)
	}
	return matches
}
