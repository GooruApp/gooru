#!/usr/bin/env bash

set -Ee
export APP_NAME="gooru"
export COMPOSE_PROJECT_NAME=${APP_NAME}

is_windows=false
if [ `uname` != "Darwin" ] && [ `uname` != "Linux" ]; then
    is_windows=true
fi

if [ is_windows ]; then
    export VITE_USE_POLLING="true"
    export AIR_START_COMMAND="air -c .air.win.toml -- start"
fi

case $1 in
    start )
        if [ "$2" = "cloud" ]; then
            echo "Starting cloud ${APP_NAME} stack..."
            export DATABASE_BACKEND="postgres"
            docker compose up -d --remove-orphans client server postgres
        else
            echo "Starting local ${APP_NAME} stack..."
            docker compose up -d --remove-orphans client server
        fi
        ./run.sh attach
        ;;

    stop )
        echo "Stopping ${APP_NAME} stack..."
        docker compose down --remove-orphans
        docker volume prune -f
        ;;

    restage )
        echo "Restaging ${APP_NAME} stack..."
        ./run.sh stop && ./run.sh start $2
        ;;

    restart )
        if [ ! -z "$2" ]; then
            docker restart $2
        else
            docker compose restart
        fi
        ;;

    attach )
        echo "Attaching to ${APP_NAME} stack..."
        docker compose logs -f --tail 1000
        ;;

    desktop )
        go install github.com/wailsapp/wails/v2/cmd/wails@latest
        
        if [ "$2" = "dev" ]; then
            ./run.sh stop || true

            printf "\nLaunching Wails app in dev mode..."
            cd app
            if [ is_windows ]; then
                sed -i "2s@.*@  \"reloaddirs\": \"$(find ../server -type d -print | sort | paste -sd "," -)\",@" wails.json
            else
                sed -i '' "2s@.*@  \"reloaddirs\": \"$(find ../server -type d -print | sort | paste -sd "," -)\",@" wails.json
            fi
            wails dev || echo "Unable to launch Wails - is GOPATH set correctly?"
        elif [ "$2" = "build" ]; then
            echo "Building Wails app..."
            cd app
            wails build || echo "Unable to launch Wails - is GOPATH set correctly?"
        else
            echo "Invalid option provided."
            exit 1
        fi
        ;;

    help )
        echo -e "\nUsage:"
        echo -e "\t start [cloud] \t\t Creates, runs, and attaches to the dev stack. Starts cloud version if specified."
        echo -e "\t stop \t\t\t Stops and removes existing containers for dev stack."
        echo -e "\t restage [cloud] \t Runs the stop command followed by the start command. Fully tears down containers."
        echo -e "\t restart [<container>] \t Restarts the given container, or all containers if not specified."
        echo -e "\t attach  \t\t Attaches to the running stack."
        echo -e "\t desktop [dev|build] \t Runs a development Wails app in 'dev' mode, or builds a Wails executable in 'build' mode."
        ;;
    
    * )
        echo "Invalid option provided."
        ./run.sh help
        exit 1
        ;;
esac

