package main

type LanguageMap struct {
	Code    string
	AceCode string
	Name    string
}

var LanguageMaps = []*LanguageMap{
	&LanguageMap{"c", "c_cpp", "C"},
	&LanguageMap{"css", "css", "CSS"},
	&LanguageMap{"go", "golang", "Go"},
	&LanguageMap{"html", "html", "HTML"},
}
