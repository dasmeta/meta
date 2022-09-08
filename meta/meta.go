package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/thatisuday/commando"
)


var version = "v0.1.1";

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createConfigFile() {
	f, err := os.Create(".pre-commit-config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	words := []string{"repos:", "- repo: https://github.com/pre-commit/pre-commit-hooks", "  rev: v4.3.0", "  hooks:", "    - id: check-added-large-files", "    - id: check-merge-conflict", "    - id: check-vcs-permalinks", "    - id: end-of-file-fixer", "    - id: trailing-whitespace", "      args: [--markdown-linebreak-ext=md]", "      exclude: CHANGELOG.md", "    - id: check-yaml", "    - id: check-merge-conflict", "    - id: check-executables-have-shebangs", "    - id: check-case-conflict", "    - id: mixed-line-ending", "      args: [--fix=lf]", "    - id: detect-aws-credentials", "      args: ['--allow-missing-credentials']", "    - id: detect-private-key", "- repo: https://github.com/antonbabenko/pre-commit-terraform", "  rev: v1.74.1", "  hooks:", "    - id: terraform_fmt", "    - id: terraform_docs", "      args:", "        - --hook-config=--path-to-file=README.md", "        - --hook-config=--add-to-existing-file=true", "        - --hook-config=--create-file-if-not-exist=true"}

	for _, word := range words {

		_, err := f.WriteString(word + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("done")
}

func createFIle() {
	f, err := os.Create("pre-commit")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	words := []string{"#!/bin/bash", "set -e", "for file in $(git status --short | grep '^[MARCD]')","do", "  git show ':$file'", "  pre-commit run --show-diff-on-failure --color=always --all-files",  "  if [ $? -ne 0 ]; then", "    exit 1", "  fi", "done"}

	for _, word := range words {

		_, err := f.WriteString(word + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(".pre-commit file creation is done")
}

func createFolder() {
	if err := os.Mkdir("githooks", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	os.Chdir("githooks")
	createFIle()
	os.Chmod("pre-commit", 0777)
	os.Chdir("../")
}

func installPreCommit() {
	if runtime.GOOS == "linux" {
		app := "pip"

		arg0 := "install"
		arg1 := "pre-commit"

		cmd := exec.Command(app, arg0, arg1)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(stdout))
	} else if runtime.GOOS == "windows" {
		app := "conda"

		arg0 := "install"
		arg1 := "-c"
		arg2 := "conda-forge"
		arg3 := "pre-commit"

		cmd := exec.Command(app, arg0, arg1, arg2, arg3)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(stdout))
	} else {
		app := "brew"

		arg0 := "install"
		arg1 := "pre-commit"

		cmd := exec.Command(app, arg0, arg1)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(stdout))
	}
}

func configPreCommitGlobally() {
	app := "git"

	arg0 := "config"
	arg1 := "--get"
	arg2 := "core.hooksPath"
	arg3 := "githooks"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	fmt.Println(cmd)
}

func createPackageJsonFile() {
	file, err := os.Create("package.json")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	content, err := http.Get("https://github.com/dasmeta/meta/releases/download/" + version + "/package.json")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Body.Close()

	io.Copy(file, content.Body)
	fmt.Println("package.json file creation is done")
}

func createCommitLintConfigFile() {

	file, err := os.Create("commitlint.config.js")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	content, err := http.Get("https://github.com/dasmeta/meta/releases/download/" + version +  "/commitlint.config.js")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Body.Close()

	io.Copy(file, content.Body)
	fmt.Println("commitlint.config.js file creation is done")
}

func createGithubFolder() {
	if err := os.Mkdir(".github", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	os.Chdir(".github")
	fmt.Println("create .github folder")
	createWorkflowFolder()
	createCommitLintYamlFile()
	createGitHubCI()
	os.Chdir("../../")

}

func createWorkflowFolder() {
	if err := os.Mkdir("workflows", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	os.Chdir("workflows")
	fmt.Println("create workflows folder")
}

func createCommitLintYamlFile() {
	file, err := os.Create("commitlint.yaml")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	content, err := http.Get("https://github.com/dasmeta/meta/releases/download/" + version + "/commitlint.yaml")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Body.Close()

	io.Copy(file, content.Body)
	fmt.Println("commitlint.yaml file creation is done")

}

func createGitHubCI() {
	os.Chdir("workflows")
	file, err := os.Create("pre-commit.yaml")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	content, err := http.Get("https://github.com/dasmeta/meta/releases/download/" + version + "/pre-commit.yml")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Body.Close()

	io.Copy(file, content.Body)
	fmt.Println("pre-commit.yaml file creation is done")

}

func createHuskyFolder() {
	if err := os.Mkdir(".husky", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	os.Chdir(".husky")
	fmt.Println("create .husky folder")
	createCommitMSGFile()
	create_Folder()
	os.Chdir("_")
}

func create_Folder() {
	if err := os.Mkdir("_", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	fmt.Println("create _ folder")
	createHuskyshFile()
	createGitIgnoreFile()
}

func createHuskyshFile() {
	file, err := os.Create("husky.sh")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	content, err := http.Get("https://github.com/dasmeta/meta/releases/download/" + version +  "/husky.sh")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Body.Close()

	io.Copy(file, content.Body)
	fmt.Println("husky.sh file creation is done")

}

func createCommitMSGFile() {
	file, err := os.Create("commit-msg")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	content, err := http.Get("https://github.com/dasmeta/meta/releases/download/" + version +  "/commit-msg")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Body.Close()

	io.Copy(file, content.Body)
	fmt.Println("commit-msg file creation is done")

}

func createGitIgnoreFile() {
	file, err := os.Create(".gitignore")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	content, err := http.Get("https://github.com/dasmeta/meta/releases/download/" + version +  "/default.gitignore")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Body.Close()

	io.Copy(file, content.Body)
	fmt.Println(".gitignore file creation is done")

}

func main() {

	commando.
		SetExecutableName("meta").
		SetVersion(version).
		SetDescription("Meta is a command-line tool to generate semantic-release and pre-commit hooks in your projects.\nIt helps you create needed files and run it and much more.").
		SetEventListener(func(eventName string) {
			//fmt.Println("event-name: ", eventName)
		})

	commando.
		Register("pre-commit").
		SetDescription("This command creates .pre-commit.yaml files and after running scripts for needed pre-commit checks and outputs that files in the project directory.").
		SetShortDescription("This command creates .pre-commit.yaml files and after running scripts for needed pre-commit checks and outputs that files in the project directory.").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			createConfigFile()
			createFolder()
			configPreCommitGlobally()
			installPreCommit()
		})

	commando.
		Register("semantic-release").
		SetDescription("This command creates files and github actions CI for semantic release also for pre-commit checks and outputs that files in the project directory.").
		SetShortDescription("This command creates files and github actions CI for semantic release also for pre-commit checks and outputs that files in the project directory.").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			createPackageJsonFile()
			createCommitLintConfigFile()
			createGithubFolder()
			createGitHubCI()
			createHuskyFolder()
		})

	commando.Parse(nil)
}
