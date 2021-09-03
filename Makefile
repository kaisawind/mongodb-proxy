IMPORT_PATH := github.com/kaisawind/mongodb-proxy

V := 1 # When V is set, print commands and build progress.
M := 1 # When M is set, build with -mod vendor.

# Space separated patterns of packages to skip in list, test, format.
IGNORED_PACKAGES := vendor

# Docker informations
DOCKER_REPO := 
DOCKER_NAMESPACE := kaisawind

.PHONY: all
all: build

.PHONY: build
build: apiserver
	@echo go install $(FLAGS)
	@echo "building completed!" 

.PHONY: apiserver
apiserver:
	CGO_ENABLED=0 go build $(FLAGS)cmd-server ./cmd/api-server
	@echo "building apiserver completed!" 

.PHONY: docker
docker: apiserver
	docker build -t $(DOCKER_REPO)$(DOCKER_NAMESPACE)/mongodbproxy .

.PHONY: docker-push
docker-push:
	docker push $(DOCKER_REPO)$(DOCKER_NAMESPACE)/mongodbproxy

##### ^^^^^^ EDIT ABOVE ^^^^^^ #####

##### =====> Internals <===== #####

# 版本号 v1.0.3-6-g0c2b1cf-dev
# 1、6:表示自打tag v1.0.3以来有6次提交（commit）
# 2、g0c2b1cf：g 为git的缩写，在多种管理工具并存的环境中很有用处
# 3、0c2b1cf：7位字符表示为最新提交的commit id 前7位
# 4、如果本地仓库有修改，则认为是dirty的，则追加-dev，表示是开发版：v1.0.3-6-g0c2b1cf-dev
VERSION          := $(shell git describe --tags --always --dirty="-dev")

# 时间
DATE             := $(shell date -u '+%Y-%m-%d-%H%M UTC')

# 版本标志
VERSION_FLAGS    := -ldflags='-X "main.Version=$(VERSION)" -X "main.BuildTime=$(DATE)"'

# 输出文件夹
OUTPUT_DIR       := -o ./bin/

# 标志
FLAGS            := $(if $V,-v) $(if $M,-mod vendor) $(VERSION_FLAGS) $(OUTPUT_DIR)

