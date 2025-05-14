default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test -cover ./... -v $(TESTARGS) \
		-count=1 \
		-timeout 120m \
		-coverprofile=tmp/cover.out

# example)
# $ TROCCO_TEST_URL=https://localhost:4000 \
# 	TROCCO_API_KEY=＊＊＊＊ \
# 	make testacc TESTARGS="-run TestAccConnectionResource"


.PHONY: testacc-html
cover-html:
	go tool cover -html=tmp/cover.out -o tmp/cover.html
