#!/bin/sh

readonly repo=${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${REPO_NAME}
readonly branch=$(git rev-parse --abbrev-ref HEAD 2>/dev/null)
readonly commit=$(git rev-parse --short HEAD 2>/dev/null)
readonly version=$(cat VERSION)
readonly build-id=${CODEBUILD_BUILD_ID}
readonly build-url=https://$AWS_REGION.console.aws.amazon.com/codebuild/home?region=$AWS_REGION#/builds/${build-id}/view/new

# Attempt to pull existing builder image
if docker pull ${repo}:builder-${branch}; then
    # Update builder image
    docker build -t ${repo}:builder-${branch} --target builder --cache-from ${repo}:builder-${branch} -f docker/Dockerfile .
else
    # Create new builder image
    docker build -t ${repo}:builder-${branch} --target builder -f docker/Dockerfile .
fi

# Attempt to pull latest branch target image
docker pull ${repo}:${branch} || true

# Build and push target image
docker build -t ${repo}:branch --cache-from ${repo}:builder-${branch} --cache-from ${repo}:${branch} \
  --build-arg GH_SHA=${commit} \
  --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
  --build-arg CODECOV_TOKEN=$CODECOV_TOKEN \
  --build-arg VCS_COMMIT_ID=${commit} \
  --build-arg VCS_BRANCH_NAME=${branch} \
  --build-arg CI_BUILD_ID=${build-id} \
  --build-arg CI_BUILD_URL=${build-url} \
  -f docker/Dockerfile .

docker push ${repo}:${branch}

# for master push versioned image too
if [ "${branch}" == "master" ]; then
    docker push ${repo}:${version}
fi

# Push builder image
docker push ${repo}:builder-${branch}