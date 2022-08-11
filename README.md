# Go Template <br/><br/> Summary



- [About](#about)
- [Download and run project](#download-and-run-project)
- [Debug](#debug)
- [Local development recommend](#local-development-recommend)
- [Dependencies](#dependencies)
- [Optional tools](#optional-tools)
- [Setting the local variables](#setting-the-local-variables)
- [Loading local variables](#loading-local-variables)
- [Dependencies for debug](#dependencies-for-debug)
- [Generation backup plugins vscode](#generation-backup-plugins-vscode)
- [Install plugins](#install-plugins)
- [Test in kubernets local](#test-in-kubernets-local)
- [Modifying the project](#modifying-the-project)

## About
- [Inspirado em golang-standards project-layout](https://github.com/golang-standards/project-layout)
- Purpose of being a base project
- Have necessary tools
- Have debug .json configured
- Hello World Rest

## Download and run project
```bash

    $ git clone https://github.com/arozendojr/go-template

    $ cd go-template
    $ go run internal/cmd/main.go
    
    # Test
    $ curl localhost:8080

    # Response
    {"Text":"Hello World","Code":"200"}

```

## Debug
- Open project with vscode
- [Loading local variables](#loading-local-variables)
- See picture


## Local development recommend 
- [Asdf](https://asdf-vm.com/guide/getting-started.html)
- [Install Dependencies with asdf](https://github.com/asdf-vm/asdf-plugins)

## Dependencies
- [Vscode](https://code.visualstudio.com/download)
- [Go](https://go.dev/doc/install)

## Optional tools
- [Docker](https://docs.docker.com/engine/install/)
- [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [Skaffold](https://skaffold.dev/docs/install/)
- [Nodejs](https://nodejs.org/en/download/package-manager/)

## Setting the local variables
- nano ~/.bashrc
- export GOPATH=~/go
- export GOROOT=$(asdf where golang)/go
- #export GOROOT=~/golang/1.18.4/go
- export PATH=$PATH:$GOROOT/bin

## Loading local variables
source ~/.bashrc

## Dependencies for debug
- go install github.com/go-delve/delve/cmd/dlv@master
- go install github.com/cweill/gotests/gotests@latest 
- go install github.com/fatih/gomodifytags@latest 
- go install github.com/josharian/impl@latest 
- go install github.com/haya14busa/goplay/cmd/goplay@latest 
- go install honnef.co/go/tools/cmd/staticcheck@latest 
- go install golang.org/x/tools/gopls@latest 
- go install golang.org/x/tools/cmd/goimports@latest

## Plugins

## Generation backup plugins vscode
- code --list-extensions

## Install plugins
- code --install-extension alphabotsec.vscode-eclipse-keybindings
- code --install-extension bierner.markdown-preview-github-styles
- code --install-extension golang.go
- code --install-extension hediet.vscode-drawio
- code --install-extension jebbs.plantuml
- code --install-extension ms-azuretools.vscode-docker
- code --install-extension ms-vscode-remote.remote-containers
- code --install-extension ms-vscode.makefile-tools
- code --install-extension neonxp.gotools
- code --install-extension PKief.material-icon-theme
- code --install-extension roonie007.hide-files
- code --install-extension trietho.file-size-explorer
- #code --install-extension WakaTime.vscode-wakatime

# Test in kubernets local

## kind create cluster

```
Using Skaffold takes a long time to build
Close all programs you don't need, before starting the command below
```

- skaffold dev --port-forward
- or
- skaffold debug --port-forward

## Install commit tools
- npm install 

## Modifying the project

```
git checkout develop
```
or
```
git switch develop
```

- git checkout -b feature/my-feature

### Commit 1
- git add .
- make commit

![commit](https://user-images.githubusercontent.com/36340691/184033851-bbb7a419-a8dd-44ea-b6f9-d48cb69efa8f.png)

### Commit 2
- git add .
- make commit

![commit](https://user-images.githubusercontent.com/36340691/184033851-bbb7a419-a8dd-44ea-b6f9-d48cb69efa8f.png)

### Running squash
- make squash

![squash](https://user-images.githubusercontent.com/36340691/184034961-dacba4e1-1597-423b-bd9e-113ef2636b3c.png)

- Next pass
- Edit commit for squash
- switch from pick to squash, save and quit

![squash_2](https://user-images.githubusercontent.com/36340691/184035047-57b3792c-73c3-43ef-88c4-4f5710e03f54.png)

- Next pass
- Edit long message

![squash_3](https://user-images.githubusercontent.com/36340691/184035116-2007e596-c295-43dc-9e54-f22ee60d1817.png)

- save and quit

### commit with last commet
- make amend
