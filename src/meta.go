package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

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

	words := []string{"repos:", "- repo: https://github.com/pre-commit/pre-commit-hooks", "  rev: v4.2.0", "  hooks:", "    - id: check-added-large-files", "    - id: check-merge-conflict", "    - id: check-vcs-permalinks", "    - id: end-of-file-fixer", "    - id: trailing-whitespace", "      args: [--markdown-linebreak-ext=md]", "      exclude: CHANGELOG.md", "    - id: check-yaml", "    - id: check-merge-conflict", "    - id: check-executables-have-shebangs", "    - id: check-case-conflict", "    - id: mixed-line-ending", "      args: [--fix=lf]", "    - id: detect-aws-credentials", "      args: ['--allow-missing-credentials']", "    - id: detect-private-key", "- repo: https://github.com/antonbabenko/pre-commit-terraform", "  rev: v1.62.3", "  hooks:", "    - id: terraform_fmt", "    - id: terraform_docs", "  args:", "      - '--args=--lockfile=false'"}

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

	words := []string{"#!/bin/bash", "echo 'Start git pre-commit hooks and checks... ';", "pre-commit run -a;", "echo 'End git pre-commit hooks and checks... ';"}

	for _, word := range words {

		_, err := f.WriteString(word + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("done")
}

func createFolder() {
	if err := os.Mkdir("githooks", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	file := os.Chdir("githooks")
	fmt.Println(file)
	createFIle()
	cmd := os.Chmod("pre-commit", 0777)
	fmt.Println(cmd)
	file1 := os.Chdir("../")
	fmt.Println(file1)
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

func main() {
	createConfigFile()
	createFolder()
	configPreCommitGlobally()
	installPreCommit()
}
