package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

// NewPackCommand creates a new pack command which allows
// the user to package their Go modules
func NewPackAllCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "packall <go.mod file> <outputdirectory>",
		Short: "Package all your Go modules based on a go.mod file",
		Args:  cobra.MinimumNArgs(2),

		RunE: func(cmd *cobra.Command, args []string) error {
			return runPackAllCommand(args)
		},
	}

	return &cmd
}

func runPackAllCommand(args []string) error {
	goPath := os.Getenv("GOPATH")
	log.Printf("GOPATH is %s...", goPath)
	if goPath == "" {
		return fmt.Errorf("GOPATH environment variable not set")
	}

	goMod := args[0]
	log.Printf("go.mod path is %s...", goMod)
	if err := checkFileSyntax(goMod); err != nil {
		return fmt.Errorf("syntax of go.mod not correct (%w)", err)
	}

	outputDirectory, err := filepath.Abs(args[1])
	log.Printf("Packing module in path %s...", outputDirectory)
	if err != nil {
		return fmt.Errorf("get abs path of output directory: %w", err)
	}

	//if err := pack.Module(path, version, outputDirectory); err != nil {
	//	return fmt.Errorf("package module: %w", err)
	//}

	return nil
}

func checkFileSyntax(file string) error {

	//do some regex checks

	return nil
}

func packModules(goModFile string, outputDirectory string) error {
	goVendor(goModFile)
	return nil
}

func goVendor(goModFile string) error {
	tempDir := "temp"

	if err := os.Mkdir(tempDir, 0755); err != nil {
		return fmt.Errorf("could not create temporary directory (%w)", err)
	}

	cmd := exec.Command("COPY", "/Y", goModFile, tempDir)
	if err := cmd.Run(); err!=nil {
		return fmt.Errorf("could not copy go.mod into the temporary directory (%w)", err)
	}



	return nil
}

func writeMainGoFile(libraries string, tempDir string) error {

	d1 := []byte("package temp\n\nimport (\n"+ libraries +"\n)")
	if err := ioutil.WriteFile(tempDir+"temp.go", d1, 0644); err!=nil{
		return fmt.Errorf("could not write temp.go file (%w)", err)
	}


	return nil
}

func getLibraryList(goModFile string) error {
	// Regexp
	// \s*(\S*\/*)\s([vV]\d.\d.\d)

	re := regexp.MustCompile(`\s*(\S*\/*)\s([vV]\d.\d.\d)`)


	fileContent, _ := ioutil.ReadFile(goModFile)

	test := re.FindAll(fileContent, -1)

	fmt.Printf("%q\n", re.FindAll(fileContent, -1))
	println(test)

	return nil
}








func Test(){
	var goModFile string = "temp/go.mod"

	path, _ := filepath.Abs(goModFile)

	fmt.Println(path)

	if err := getLibraryList(goModFile); err != nil {
		fmt.Println("oops")
	}

}











