#!/usr/bin/env bash
set -ex
docker run --rm -v "${PWD}":/src openapitools/openapi-generator-cli generate -i /src/openapi.yaml -o /src -g go --git-repo-id=tenor-go --git-user-id=warrengray

# remove unnecessary files
rm -rf api
rm -rf gen
