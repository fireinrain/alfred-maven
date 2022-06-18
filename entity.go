package main

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

// GAVEntity for template generator
type GAVEntity struct {
	GroupId    string
	ArtifactId string
	Version    string
}

// Package for workflow item
type Package struct {
	//GroupId of package
	GroupId string
	//ArtifactId of package
	ArtifactId string
	// Version of package
	Version string
	// UpdateTime of package
	UpdateTime string
}
