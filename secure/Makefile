proto: ## Generate up to date protobuf code based on protobuf defintions in api/proto/defintions.
	protoc -I=./api/v1/proto --go_out=plugins=grpc:./api/v1/proto/generated api/v1/proto/definitions/*.proto
	mv api/v1/proto/generated/definitions/** api/v1/proto/generated
	rmdir api/v1/proto/generated/definitions

help:
	$(info Available targets)
	@awk '/^[a-zA-Z\-\_0-9]+:/ {                                   \
	  nb = sub( /^## /, "", helpMsg );                             \
	  if(nb == 0) {                                                \
		helpMsg = $$0;                                             \
		nb = sub( /^[^:]*:.* ## /, "", helpMsg );                  \
	  }                                                            \
	  if (nb)                                                      \
		printf "\033[1;31m%-" width "s\033[0m %s\n", $$1, helpMsg; \
	}                                                              \
	{ helpMsg = $$0 }'                                             \
	width=$$(grep -o '^[a-zA-Z_0-9]\+:' $(MAKEFILE_LIST) | wc -L)  \
	$(MAKEFILE_LIST)