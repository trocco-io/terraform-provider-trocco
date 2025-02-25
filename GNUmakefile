default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

# example)
# $ TROCCO_TEST_URL=https://localhost:4000 \
# 	TROCCO_API_KEY=＊＊＊＊ \
# 	make testacc TESTARGS="-run TestAccConnectionResource"
