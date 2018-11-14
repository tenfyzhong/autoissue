package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
)

var root = flag.String("root", ".", "root path")

// 出现在lhs而不出现在rhs的数据
func subtraction(lhs, rhs []string) []string {
	rhsExist := make(map[string]bool)
	for _, r := range rhs {
		rhsExist[r] = true
	}
	results := make([]string, 0, len(lhs))
	for _, l := range lhs {
		if !rhsExist[l] {
			results = append(results, l)
		}
	}
	return results
}

func postIssue(owner, repo, token, domain, root, labelName string, labels []string) error {
	indexPath := path.Join(root, "public", labelName, "index.html")
	title, err := PostPathToTitle(indexPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(title)
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues?access_token=%s", owner, repo, token)

	req := &PostIssue{
		Title:  title,
		Body:   domain + labelName,
		Labels: append(labels, labelName),
	}
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	res, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("status code: %d", res.StatusCode)
	}

	return nil
}

func main() {
	flag.Parse()
	configPath := path.Join(*root, "_config.yml")
	config, err := NewConfig(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "new config %v\n", err)
		os.Exit(1)
	}

	token := os.Getenv("AUTH_TOKEN")
	if token == "" {
		fmt.Fprintf(os.Stderr, "AUTH_TOKEN is empty\n")
		os.Exit(2)
	}

	issues, err := GetIssues(token, config.Owner, config.CommentRepo, config.Labels)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get issues %v\n", err)
		os.Exit(3)
	}
	issueLabelNames := IssueToURLLabelName(issues, config.Labels)

	sitemapPath := path.Join(*root, "public", "sitemap.xml")
	sitemap, err := ParseSitemap(sitemapPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse sitemap %v\n", err)
		os.Exit(4)
	}

	sitemapLabelNames := SitemapToURLLabelName(sitemap, config.URL)

	needInitLabelNames := subtraction(sitemapLabelNames, issueLabelNames)
	for _, name := range needInitLabelNames {
		postIssue(config.Owner, config.CommentRepo, token, config.URL, *root, name, config.Labels)
	}
}
