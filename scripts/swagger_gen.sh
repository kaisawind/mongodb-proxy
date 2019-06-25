#!/bin/bash
set -e -o pipefail

# Generates server & client code based on swagger spec. Uses existing model's structures.
# the spec references the model.json which is generated from code using generate-models.sh
# For server, spec must be flattened first, otherwise the embedded spec still contains
# references to external file and does not validate properly when services are started.

: ${WORKDIR:="./"}
: ${QUIET:="-q"}
: ${SWAGGER:="swagger.yaml"}
: ${MODELS_SWAGGER:="models.json"}
: ${MODELS_PACKAGE:="github.com/kaisawind/mongodb-proxy/api/v1"}
: ${APP:="api"}
: ${TARGET:=""}
: ${PACKAGE:="v1"}

#export GOROOT="/usr/lib/go-1.10"

MODEL_COMMAND="pushd ${WORKDIR} && \
                swagger generate model ${QUIET} -t ./${PACKAGE} -f ./swagger/${MODELS_SWAGGER} --model-package=v1 && \
                popd"

SPEC_COMMAND="pushd ${WORKDIR} && \
                CGO_ENABLED=0 swagger generate spec ${QUIET} -o ./swagger/${MODELS_SWAGGER} -b ./${PACKAGE} -m && \
                popd"

SERVER_COMMAND="pushd ${WORKDIR} && \
                swagger generate server ${QUIET} -A ${APP} -P ${PACKAGE}.Principal -t ./${TARGET} -f ./swagger/${SWAGGER} --existing-models=${MODELS_PACKAGE} --model-package=v1 && \
                popd"

CLIENT_COMMAND="pushd ${WORKDIR} && \
                swagger generate client ${QUIET} -A ${APP} -P ${PACKAGE}.Principal -t ./${TARGET} -f ./swagger/${SWAGGER} --existing-models=${MODELS_PACKAGE} --model-package=v1 && \
                popd"

CLEAN_COMMAND="pushd ${WORKDIR} && \
                rm -rf ./client && \
                rm -rf ./cmd-server && \
                rm -rf ./restapi/operations && \
                rm -rf ./restapi/doc.go && \
                rm -rf ./restapi/embedded_spec.go && \
                rm -rf ./restapi/server.go && \
                rm -rf ./v1 && \
                rm -rf ./swagger/models.json && \
               popd"

bash -c "${CLEAN_COMMAND}"
#bash -c "${MODEL_COMMAND}"
bash -c "${SERVER_COMMAND}"
bash -c "${CLIENT_COMMAND}"
#bash -c "${SPEC_COMMAND}"
