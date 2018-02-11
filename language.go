package main

//Language contains structure for language
type Language struct {
	Code    string
	AceCode string
	Name    string
}

//Languages contain array of language
var Languages = []*Language{
	&Language{"c", "c_cpp", "C"},
	&Language{"css", "css", "CSS"},
	&Language{"go", "golang", "Go"},
	&Language{"html", "html", "HTML"},
	&Language{"javascript", "javascript", "Javascript"},
	&Language{"javascript", "json", "JSON"},
}

func getLanguage(acecode string) string {
	for _, v := range Languages {
		if v.AceCode == acecode {
			return v.Name
		}
	}
	return "Plain Text"
}

func getCode(language string) string {
	for _, v := range Languages {
		if v.Name == language {
			return v.Code
		}
	}
	return "bash"
}
