package main

import "fmt"

type PlatFormType string

const (
	PolicyMaven     PlatFormType = "Apache Maven"
	PolicyGradle    PlatFormType = "Gradle Groovy"
	PolicyKotlin    PlatFormType = "Gradle Kotlin"
	PolicySbt       PlatFormType = "Scala SBT"
	PolicyIvy       PlatFormType = "Apache Ivy"
	PolicyGrape     PlatFormType = "Groovy Grape"
	PolicyLeiningen PlatFormType = "Leiningen"
	PolicyBuildr    PlatFormType = "Apache Buildr"
	PolicyBadge     PlatFormType = "Maven Central Badge"
	PolicyPurl      PlatFormType = "PURL"
	PolicyBazel     PlatFormType = "Bazel"
)

// GAVEntity
type GAVEntity struct {
	GroupId    string
	ArtifactId string
	Version    string
}

func (gav GAVEntity) toString() {
	fmt.Println("groupId: ", gav.GroupId, "artifactId: ", gav.ArtifactId, "version: ", gav.Version)
}

func GenDependencyTemplate(platFormType PlatFormType, gAVEntity GAVEntity) string {

	switch platFormType {
	case PolicyMaven:
		return fmt.Sprintf("<dependency>\n  <groupId>%s</groupId>\n  <artifactId>%s</artifactId>\n  <version>%s</version>\n  <type>pom</type>\n</dependency>", gAVEntity.GroupId, gAVEntity.ArtifactId, gAVEntity.Version)

	case PolicyGradle:
		return fmt.Sprintf("implementation '%s:%s:%s'", gAVEntity.GroupId, gAVEntity.ArtifactId, gAVEntity.Version)

	case PolicyBazel:
		return fmt.Sprintf("maven_jar(\n    name = \"%s\",\n    artifact = \"%s:%s\",\n    sha1 = \"calculating...\",\n)", gAVEntity.GroupId, gAVEntity.ArtifactId, gAVEntity.Version)
	case PolicyPurl:
		return fmt.Sprintf("pkg:maven/%s/%s@%s", gAVEntity.GroupId, gAVEntity.ArtifactId, gAVEntity.Version)
	case PolicySbt:
		return "libraryDependencies +=" + "\"" + gAVEntity.GroupId + "\"" + "%" + "\"" + gAVEntity.ArtifactId + "\"" + "%" + "\"" + gAVEntity.Version + "\""
	case PolicyIvy:
		return fmt.Sprintf("<dependency org=\"%s\" name=\"%s\" rev=\"%s\" />", gAVEntity.GroupId, gAVEntity.ArtifactId, gAVEntity.Version)
	case PolicyGrape:
		return fmt.Sprintf("@Grapes(\n  @Grab(group='%s', module='%s', version='%s')\n)", gAVEntity.GroupId, gAVEntity.ArtifactId, gAVEntity.Version)
	case PolicyBadge:
		return fmt.Sprintf("[![Maven Central](https://img.shields.io/maven-central/v/%s/%s.svg?label=Maven Central)](https://search.maven.org/search?q=g:\"%s\" AND a:\"%s\")", gAVEntity.GroupId, gAVEntity.ArtifactId, gAVEntity.GroupId, gAVEntity.ArtifactId)
	case PolicyBuildr:
		return fmt.Sprintf("'%s:%s:jar:%s'", gAVEntity.GroupId, gAVEntity.ArtifactId, gAVEntity.Version)
	case PolicyKotlin:
		return fmt.Sprintf("implementation(\"%s:%s:%s\")", gAVEntity.GroupId, gAVEntity.ArtifactId, gAVEntity.Version)
	case PolicyLeiningen:
		return fmt.Sprintf("[%s/%s \"%s\"]", gAVEntity.GroupId, gAVEntity.ArtifactId, gAVEntity.Version)
	}

	return "unsupported platform"
}
