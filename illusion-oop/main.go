package main
//in go abstraction is acheived using interfaces, as that interface is set of collection of methods and not
//the actual implementation of them.

import "fmt"

type File interface {
	OpenFile() string
	ReadFile() string
	WriteFile(content string) string
}

type MyFile struct {
	name string
	content string
}// har ek file ke pass uska name and content hoga thats how we defined struct here for File

// now writing actual implmentation of these above ops like readfile and all

func (f *MyFile) ReadFile() string {
	return f.content
}

func (f *MyFile) OpenFile() string {
	return fmt.Sprintf("File opened: %s", f.name)
}

func (f *MyFile) WriteFile(content string) string {
	f.content = content
	return fmt.Sprintf("Written to a file: %s", f.name)
}

func main() {
	file := &MyFile{
		name: "example.txt",
	}
	fmt.Println(file.OpenFile())
	fmt.Println(file.WriteFile("This is atharva here"))
	fmt.Println(file.ReadFile())
}

//In Go, an identifier is exported if it begins with an uppercase letter, making it accessible from other packages. {Encapsulation}
//Now Polymorphism

