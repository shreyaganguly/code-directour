package models

//Language contains structure for language
type Language struct {
	Code          string
	AceCode       string
	SlackFileType string
	Name          string
}

//Languages contain array of languages
var Languages = []*Language{
	&Language{"c", "c_cpp", "c", "C"},
	&Language{"css", "css", "css", "CSS"},
	&Language{"go", "golang", "go", "Go"},
	&Language{"html", "html", "html", "HTML"},
	&Language{"javascript", "javascript", "javascript", "Javascript"},
	&Language{"javascript", "json", "javascript", "JSON"},
}

func getLanguage(acecode string) string {
	for _, v := range Languages {
		if v.AceCode == acecode {
			return v.Name
		}
	}
	return "Plain Text"
}

func GetCode(language string) string {
	for _, v := range Languages {
		if v.Name == language {
			return v.Code
		}
	}
	return "bash"
}

func GetAceCode(language string) string {
	for _, v := range Languages {
		if v.Name == language {
			return v.AceCode
		}
	}
	return "bash"
}

func GetSlackFileType(language string) string {
	for _, v := range Languages {
		if v.Name == language {
			return v.SlackFileType
		}
	}
	return "auto"
}
