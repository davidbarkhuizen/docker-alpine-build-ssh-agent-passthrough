 #!/usr/bin/env bash

dockerfile_path=Dockerfile
image_name=alpine-build-ssh-agent-passthrough
image_tag=latest
container_name=agent-passthrough

function configure_buildkit {

    export DOCKER_BUILDKIT=1;
    export BUILDKIT_PROGRESS=plain;
}

function build {

    configure_buildkit
    docker build -f $dockerfile_path --tag $image_name:$image_tag --ssh default .
}

function bash {
    docker exec -it $container_name /bin/sh;
}

function run {
    docker run --name=$container_name -d $image_name:$image_tag tail -f /dev/null 
}

function kill {
    docker stop $container_name;
    docker rm $container_name;
}

function usage {
    echo "go build|run|bash|kill"
}

case $1 in 

    "build")
        build
    ;;

    "run")
        run
    ;;

    "kill")
        kill
    ;;

    "bash")
        bash
    ;;

    *)
        usage
    ;;
esac