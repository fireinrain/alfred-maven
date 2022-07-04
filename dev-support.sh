#!/bin/zsh
# move target compile bin to alfred workflow dir

# get current_dir
cd "$(dirname "$0")"

current_dir="$(pwd)"
#echo "$current_file" &&

# exec_file bin name
exec_file="$current_dir"/alfred-maven

# alfred workflow dir you create
# shellcheck disable=SC2089
alfred_workflow_dir=/Users/sunrise/Library/Application\ Support/Alfred/Alfred.alfredpreferences/workflows/user.workflow.1DA5192E-9A44-4EEA-8F12-1A412AC6B505
# if not exists alfred-maven bin, run build.sh to create bin file
if [ -f "$exec_file" ]; then
  ./build.sh
  cp "$exec_file" "$alfred_workflow_dir"
  echo "Done! move alfred-maven bin to $alfred_workflow_dir"

else
  echo "alfred-maven bin not exists, run build.sh to create bin file"
  ./build.sh
  cp "$exec_file" "$alfred_workflow_dir"
  echo "Done! move alfred-maven bin to $alfred_workflow_dir"
fi
# warning! this script is just for me to copy project compile bin to workflow dir
