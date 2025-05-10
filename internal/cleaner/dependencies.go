package cleaner

import (
    "encoding/xml"
    "fmt"
    "os"
)

type Dependency struct {
    Group    string `xml:"groupId"`
    Artifact string `xml:"artifactId"`
    Version  string `xml:"version"`
}

type Project struct {
    Dependencies []Dependency `xml:"dependencies>dependency"`
}

func ParsePom(path string) *Project {
    content, err := os.ReadFile(path)
    if err != nil {
        fmt.Println("Error reading POM file:", err)
        return nil
    }

    var project Project
    parseErr := xml.Unmarshal(content, &project)
    if parseErr != nil {
        fmt.Println("Error parsing POM XML:", parseErr)
        return nil
    }

    return &project
}
