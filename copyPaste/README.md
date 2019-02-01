# CopyPaste.sh

Devops' work can be bored in some cases, agents or servers always have complicated environmental settings that have to congifured properly. Massive copy-paste of environmental variables to a number of files are required sometimes when changes, updatings or migrations are on the way. 

I wrote this script in [bash] to automate massive copy-paste in seconds, we are able to do 
* add new property
* modify property
* delete property
* add under other property

## How to use

>` sh copyPaste.sh [PATH] [OPERATIONS] `: Run copyPaste.sh, then enter path of files, and what kinds of operation you want to perform

## Parameters

>`[PATH]`: Path supports "--all", which represents all folders or all files. It will fail the program we have both directories and files

>`[OPERATIONS]`: Operations can be "--addnew", "--modify", "--delete", "--addunder", "--spaces", "--dashes"

## Examples

>` sh copyPaste.sh a/bc D --all --thisFile.yml --spaces --addnew USERNAME 123`: The script will find all directories underneath a/bc/D, then add property "USERNAME" which has value "123" to file "thisFile.yml", this "USERNAME" will have a space before it. For example you can add to a/bc/D/first/thisFile.yml, a/bc/D/second/thisFile.yml, a/bc/D/third/thisFile.yml, a/bc/D/fourth/thisFile.yml .....

>` sh copyPaste.sh a/bc/D/first --thisFile.yml --delete USERNAME 123`: This script will find file a/bc/D/first/thisFile.yml, and delete property "USERNAME" which has value "123"

>` sh copyPaste.sh a/bc D --all --thisFile.yml --spaces --modify PASSWORD 111 999`: This script will find file a/bc/D/first/thisFile.yml, and modify property "PASSWORD", change its value from "111" to "999"