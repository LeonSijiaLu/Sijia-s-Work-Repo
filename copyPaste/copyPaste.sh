#!/bin/bash

source ./copyPasteVariables.sh
source ./copyPasteFunctions.sh

for i in $@;
do
    allCmds+=($i)
done

for ((index=0; index <= ${#allCmds[@]}; index++)); 
do
    if [[ ${allCmds[index]} == "--"* && ${allCmds[index]} != "--all" ]]; then
        pathLock=1 # Which means path area has ended
        if [[ ${allCmds[index]} == "--type" ]]; then
            typePresent=true # Indicate if we have type in the command
            type=${allCmds[index+1]} # What type we entered
            checkType $type ${types[@]} # Check if this type fits what we have
            if [[ ! $? ]]; then
                echo [ERROR]: "Please enter correct type"
                exit 1
            else
                echo -e "\033[32mSearching $type files ......\033[m"
            fi
        elif [[ ${allCmds[index]} == '--delete' || ${allCmds[index]} == '-d' ]]; then
            operPresent=true
            operType=${allCmds[index]}
            checkOperType $operType ${allCmds[index+1]} ${allCmds[index+2]}
        elif [[ ${allCmds[index]} == '--addnew' ]]; then
            operPresent=true
            operType=${allCmds[index]}
            checkOperType $operType ${allCmds[index+1]} ${allCmds[index+2]}
        elif [[ ${allCmds[index]} == '--addunder' ]]; then
            operPresent=true
            operType=${allCmds[index]}
            checkOperType $operType ${allCmds[index+1]} ${allCmds[index+2]} ${allCmds[index+3]}
        elif [[ ${allCmds[index]} == '--modify' || ${allCmds[index]} == '-m' ]]; then
            operPresent=true
            operType=${allCmds[index]}
            checkOperType $operType ${allCmds[index+1]} ${allCmds[index+2]} ${allCmds[index+3]}
        elif [[ ${allCmds[index]} == '--spaces' ]]; then
            spaces=1
        elif [[ ${allCmds[index]} == '--dash' ]]; then
            dash=1
        fi
    fi
    if [[ $pathLock == 0 ]]; then
        pathLength=$((pathLength+1))
    fi
done

if [[ $operPresent == false ]]; then
    echo [ERROR]: "Please enter an operation type, it can be --modify, --addnew, --delete or --addunder"
    exit 1
fi

for ((index=0; index < $pathLength; index++));
do
    if [[ ${allCmds[index]} == *"/" ]]; then
        allCmds[index]=${allCmds[index]::-1}
    fi
    if [[ ${allCmds[index]} == '--all' ]]; then
        if [[ $index == $((pathLength-1)) ]]; then
            findallFoldersOrFiles "f"
        else
            findallFoldersOrFiles "d"
        fi
    else
        #fileDirs+=(`find ${allCmds[index]} -mindepth 1 -maxdepth 1 -type d`)
        if [[ $index == 0 ]]; then
            fileDirs+=${allCmds[index]}
        else
            fileDirs=( "${fileDirs[@]/%//${allCmds[index]}}" )
            checkPath ${fileDirs[@]}
        fi
    fi
done

checkFiles ${fileDirs[@]}

for i in ${fileDirs[@]};
do
    if [[ $typePresent == true ]]; then
        if [[ $i == *$type ]]; then
            manipulateFiles $i
        fi
    else
        manipulateFiles $i
    fi
done
