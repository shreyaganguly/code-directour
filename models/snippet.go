package models

import (
	"time"

	"github.com/dchest/uniuri"
)

var location *time.Location

func SetLocation(loc *time.Location) {
	location = loc
}

//SnippetInfo contains information about snippets stored
type SnippetInfo struct {
	Owner           string
	Key             string
	Time            string
	Title           string
	Language        string
	Code            string
	References      string
	SharedBySomeone bool
	SharedBy        string
	SharedToSomeone bool
	SharedTo        []*ShareInfo
	CreatedAt       int64
	ModifiedAt      int64
	DeletedAt       int64
}

type ShareInfo struct {
	Method   string
	SharedTo string
}

// Snippets is the type for array of snippets
type Snippets []*SnippetInfo

//NewSnippet creates new snippet
func NewSnippet(owner, title, language, code, references string, sharedBySomeone bool, sharedBy string, sharedToSomeone bool, sharedTo []*ShareInfo) *SnippetInfo {
	return &SnippetInfo{
		Owner:           owner,
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
		CreatedAt:       time.Now().Unix(),
		ModifiedAt:      time.Now().Unix(),
	}
}

func (s Snippets) Own() Snippets {
	var ownSnippets []*SnippetInfo
	for _, snippet := range s {
		if !snippet.SharedBySomeone {
			ownSnippets = append(ownSnippets, snippet)
		}
	}
	return ownSnippets
}

func (s Snippets) Others() Snippets {
	var otherSnippets []*SnippetInfo
	for _, snippet := range s {
		if snippet.SharedBySomeone {
			otherSnippets = append(otherSnippets, snippet)
		}
	}
	return otherSnippets
}

func (s Snippets) SharedTo() Snippets {
	var otherSnippets []*SnippetInfo
	for _, snippet := range s {
		if snippet.SharedToSomeone {
			otherSnippets = append(otherSnippets, snippet)
		}
	}
	return otherSnippets
}

func (s Snippets) Reverse() Snippets {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (s *SnippetInfo) BucketName() string {
	return "manager"
}

func (s *SnippetInfo) ID() string {
	return s.Owner
}

func (s *SnippetInfo) Value() interface{} {
	return s
}
