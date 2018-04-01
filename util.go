package main

import "github.com/shreyaganguly/code-directour/models"

func reverse(s []*models.SnippetInfo) []*models.SnippetInfo {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	return s
}
