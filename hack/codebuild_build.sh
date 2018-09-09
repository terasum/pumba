#!/bin/sh

CODEBUILD_GIT_BRANCH="$(git symbolic-ref HEAD --short 2>/dev/null)"
if [ "$CODEBUILD_GIT_BRANCH" = "" ]; then
  CODEBUILD_GIT_BRANCH="$(git branch -a --contains HEAD | sed -n 2p | awk '{ printf $1 }')";
  CODEBUILD_GIT_BRANCH=${CODEBUILD_GIT_BRANCH#remotes/origin/};
fi

readonly account=$(aws sts get-caller-identity --query 'Account' --output text)
readonly repo=${account}.dkr.ecr.${AWS_REGION}.amazonaws.com/${REPO_NAME}
readonly branch=${CODEBUILD_GIT_BRANCH}
readonly commit=${CODEBUILD_SOURCE_VERSION}
readonly version=$(cat VERSION)
readonly build_id=${CODEBUILD_BUILD_ID}
readonly build_url=https://$AWS_REGION.console.aws.amazon.com/codebuild/home?region=$AWS_REGION#/builds/${build_id}/view/new


echo "=> current branch:commit = ${branch}:${commit}"

# Attempt to pull existing builder image
echo "=> try getting previous builder image: ${repo}:builder-${branch}" 
if docker pull ${repo}:builder-${branch}; then
    # Update builder image
    echo "=> updating builder image: ${repo}:builder-${branch}" 
    docker build -t ${repo}:builder-${branch} --target builder --cache-from ${repo}:builder-${branch} -f docker/Dockerfile .
else
    # Create new builder image
    echo "=> creating new builder image: ${repo}:builder-${branch}" 
    docker build -t ${repo}:builder-${branch} --target builder -f docker/Dockerfile .
fi

# Attempt to pull latest branch target image
echo "=> try getting previous final image: ${repo}:${branch}" 
docker pull ${repo}:${branch} || true

# Build and push target image
echo "=> try getting previous final image: ${repo}:${branch}" 
docker build -t ${repo}:${branch} --cache-from ${repo}:builder-${branch} --cache-from ${repo}:${branch} \
  --build-arg GH_SHA=${commit} \
  --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
  --build-arg CODECOV_TOKEN=$CODECOV_TOKEN \
  --build-arg VCS_COMMIT_ID=${commit} \
  --build-arg VCS_BRANCH_NAME=${branch} \
  --build-arg CI_BUILD_ID=${build_id} \
  --build-arg CI_BUILD_URL=${build_url} \
  -f docker/Dockerfile .

echo "=> push final image: ${repo}:${branch}" 
docker push ${repo}:${branch}

# for master push versioned image too
if [ "${branch}" == "master" ]; then
    echo "=> push final versioned image: ${repo}:${version}" 
    docker push ${repo}:${version}
fi

# Push builder image
echo "=> push builder image for builder cache: ${repo}:builder-${branch}" 
docker push ${repo}:builder-${branch}