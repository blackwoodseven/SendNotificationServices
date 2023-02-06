#!/bin/bash
# Checks if an image with a tag already exists.
# If it exits and we are not on the main branch, we can override the tag.
# if we are on main branch, we return an error.
# Example:
#    ./check-ecr-image.sh foo/bar mytag

set -u

if [[ $# -lt 2 ]]; then
    echo "Usage: $( basename ${0} ) <repository-name> <image-tag>"
    exit 1
fi

IMAGE_META="$( aws ecr describe-images --repository-name=${1} --image-ids=imageTag=${2} 2> /dev/null )"

if [[ $? == 0 ]]; then
    IMAGE_TAGS="$( echo ${IMAGE_META} | jq '.imageDetails[0].imageTags[0]' -r )"
    echo "${1}:${2} found."
    # if branch is main, we can NOT override the tag
    if [ ${2} == "main" ]; then
        echo "${2} is main branch. Tag can not be overridden."
        exit 1
    else
        exit 0
    fi
else
    echo "${1}:${2} not found."
    exit 0
fi
