package main

type Language struct {
	Code    string
	AceCode string
	Name    string
}

var Languages = []*Language{
	&Language{"c", "c_cpp", "C"},
	&Language{"css", "css", "CSS"},
	&Language{"go", "golang", "Go"},
	&Language{"html", "html", "HTML"},
}
