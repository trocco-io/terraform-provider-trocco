.PHONY: testacc
testacc:
	TF_ACC=1 go test -cover ./... -v $(TESTARGS) \
		-timeout 120m \
		-count=1 \
		-coverprofile=tmp/cover.out

.PHONY: cover-html
cover-html:
	go tool cover -html=tmp/cover.out -o tmp/cover.html
