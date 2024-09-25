#!/bin/bash

# Check if an argument is provided
if [ -z "$1" ]; then
  echo "Please provide a folder name."
  exit 1
fi

# Extract folder name from input, remove the first 2 characters (03)
FOLDER_NAME="$1"
MODULE_NAME="${FOLDER_NAME:2}"

# Create the directory with the provided name
mkdir -p "$FOLDER_NAME"

# Navigate into the folder
cd "$FOLDER_NAME" || exit

# Create a main.go file
touch main.go

# Initialize a Go module with the module name
go mod init "$MODULE_NAME"

echo ""
echo "Folder '$FOLDER_NAME' created."
