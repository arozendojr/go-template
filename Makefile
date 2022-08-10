# Always run in project root
clean: 
	# Clean container and images not working
	docker system prune --all --force --volumes
cover: 
	# View test coverage
	go test -coverprofile=coverage.out ./... &&	go tool cover -html=coverage.out
commit: 
	# Semantic Git commits
	node ./node_modules/git-cz/bin/git-cz.js git cz
amend: 
	# commits with last comment
	git commit --amend --no-edit
squash: 
	# commits with last comment
	git rebase -i develop