package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/nlopes/slack"
	"github.com/shreyaganguly/code-directour/db"
	"github.com/shreyaganguly/code-directour/models"
)

var (
	slackUsersMap map[string]string
	slackClient   *slack.Client
	SlackEnabled  bool
)

func makeSlackUserMap() {
	users, err := slackClient.GetUsers()
	if err != nil {
		log.Fatalln("Obtaining users from slack", err)
	}
	slackUsersMap = make(map[string]string)
	for _, v := range users {
		slackUsersMap[v.Name] = v.ID
	}
}

func UploadSnippet(snippet *models.SnippetInfo, receiverName string) error {
	user, err := db.LookupinUser(snippet.Owner)
	if err != nil {
		return err
	}
	if user.Slack != nil && user.Slack.Token != "" {
		slackClient = slack.New(user.Slack.Token)
		makeSlackUserMap()
	}
	id, ok := slackUsersMap[receiverName]
	if !ok {
		return fmt.Errorf("User %s does not exist  ", receiverName)
	}
	_, err = slackClient.UploadFile(slack.FileUploadParameters{
		Content:        snippet.Code,
		Title:          snippet.Title,
		Filetype:       models.GetSlackFileType(snippet.Language),
		Channels:       []string{id},
		InitialComment: fmt.Sprintf("%s via Code Directour", strings.Title(snippet.Owner)),
	})
	if err != nil {
		return err
	}
	return nil
}
