package main

import (
	"time"

	"github.com/dchest/uniuri"
)

//SnippetInfo contains information about snippets stored
type SnippetInfo struct {
	Key             string
	Time            string
	Title           string
	Language        string
	Code            string
	References      string
	SharedBySomeone bool
	SharedBy        string
	SharedToSomeone bool
	SharedTo        string
}

//NewSnippet creates new snippet
func NewSnippet(title, language, code, references string, sharedBySomeone bool, sharedBy string, sharedToSomeone bool, sharedTo string) *SnippetInfo {
	return &SnippetInfo{
		Key:             uniuri.New(),
		Time:            time.Now().In(location).Format("02-January-2006 15:04"),
		Title:           title,
		Language:        getLanguage(language),
		Code:            code,
		References:      references,
		SharedBySomeone: sharedBySomeone,
		SharedBy:        sharedBy,
		SharedToSomeone: sharedToSomeone,
		SharedTo:        sharedTo,
	}
}

func findSnippetForUser(user, key string) (*SnippetInfo, error) {
	return find(user, key)
}

func findAndUpdateSnippet(user, key, sharedTo string) (*SnippetInfo, error) {
	return findAndUpdate(user, key, sharedTo)
}

func deleteSnippetForUser(user, key string) error {
	return delete(user, key)
}

//Save saves the snippet
func (s *SnippetInfo) Save(user string) error {
	err := update(user, "manager", s)
	if err != nil {
		return err
	}
	return nil
}

type Snippets []*SnippetInfo

func (s Snippets) own() []*SnippetInfo {
	var ownSnippets []*SnippetInfo
	for _, snippet := range s {
		if !snippet.SharedBySomeone {
			ownSnippets = append(ownSnippets, snippet)
		}
	}
	return ownSnippets
}

func (s Snippets) others() []*SnippetInfo {
	var otherSnippets []*SnippetInfo
	for _, snippet := range s {
		if snippet.SharedBySomeone {
			otherSnippets = append(otherSnippets, snippet)
		}
	}
	return otherSnippets
}
