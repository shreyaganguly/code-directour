package main

import (
	"strconv"
	"time"
)

//SnippetInfo contains information about snippets stored
type SnippetInfo struct {
	Key        string
	Time       string
	Title      string
	Language   string
	Code       string
	References string
}

//NewSnippet creates new snippet
func NewSnippet(title, language, code, references string) *SnippetInfo {
	return &SnippetInfo{
		Key:        strconv.Itoa(int(time.Now().In(location).Unix())),
		Time:       time.Now().In(location).Format("02-January-2006 15:04"),
		Title:      title,
		Language:   getLanguage(language),
		Code:       code,
		References: references,
	}
}

func findSnippetForUser(user, key string) (*SnippetInfo, error) {
	return find(user, key)
}

//Save saves the snippet
func (s *SnippetInfo) Save(user string) error {
	err := update(user, "manager", s)
	if err != nil {
		return err
	}
	return nil
}
