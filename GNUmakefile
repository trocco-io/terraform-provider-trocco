default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test -cover ./... -v $(TESTARGS) \
		-timeout 120m \
		-count=1 \
		-coverprofile=tmp/cover.out

# example)
# $ TROCCO_TEST_URL=https://localhost:4000 \
# 	TROCCO_API_KEY=＊＊＊＊ \
# 	make testacc TESTARGS="-run TestAccConnectionResource"


.PHONY: cover-html
cover-html:
	go tool cover -html=tmp/cover.out -o tmp/cover.html
