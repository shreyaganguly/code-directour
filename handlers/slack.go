package handlers

import (
	"fmt"
	"log"

	"github.com/nlopes/slack"
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

func SetSlackClient(token string) {
	slackClient = nil
	if token != "" {
		slackClient = slack.New(token)
		makeSlackUserMap()
	}

}

func UploadSnippet(snippet *models.SnippetInfo, receiverName string) error {
	id, ok := slackUsersMap[receiverName]
	if !ok {
		return fmt.Errorf("User %s does not exist  ", receiverName)
	}
	_, err := slackClient.UploadFile(slack.FileUploadParameters{
		Content:        snippet.Code,
		Title:          snippet.Title,
		Filetype:       models.GetSlackFileType(snippet.Language),
		Channels:       []string{id},
		InitialComment: fmt.Sprintf("%s via Code Directour", snippet.Owner),
	})
	if err != nil {
		return err
	}
	return nil
}
