package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Label label对象
type Label struct {
	ID     int    `json:"id"`
	NodeID string `json:"node_id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
}

// Issue issue对象
type Issue struct {
	ID     int      `json:"id"`
	NodeID string   `json:"node_id"`
	Labels []*Label `json:"labels"`
	Title  string   `json:"title"`
	Body   string   `json:"body"`
}

// PostIssue 创建issue使用的结构
type PostIssue struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Labels []string `json:"labels"`
}

// GetIssues 获取当前repo的所有issue
func GetIssues(token, owner, repo string, labels []string) ([]*Issue, error) {
	labelStr := strings.Join(labels, ",")
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues?access_token=%s&labels=%s", owner, repo, token, labelStr)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	issues := make([]*Issue, 0, 0)
	err = json.Unmarshal(data, &issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}

// IssueToURLLabelName 从issues中取出url类型的label
func IssueToURLLabelName(issues []*Issue, commonLabels []string) []string {
	names := make([]string, 0, len(issues))
	commonMap := make(map[string]bool)
	for _, label := range commonLabels {
		commonMap[label] = true
	}

	for _, issue := range issues {
		if issue == nil {
			continue
		}
		for _, label := range issue.Labels {
			if label == nil {
				continue
			}
			names = append(names, label.Name)
		}
	}
	return names
}
