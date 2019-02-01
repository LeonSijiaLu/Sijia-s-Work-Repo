#!/bin/bash

checkType(){
    for i in $2;
    do
        if [[ $type == $i ]]; then
            return 0
        fi
    done
    return 1
}
# Check if we can recognize this type

checkOperType(){
    if [[ $1 == '--delete' || $1 == '-d' || $1 == '--addnew' ]]; then
        if [[ $# != 3 ]]; then
            echo [ERROR]: Please check your parameters, it should have 3 parameters
            exit 1
        fi
        PropertyName=$2
        PropertyValue=$3
    elif [[ $1 == '--modify' || $1 == '-m' || $1 == '--addunder' ]]; then
        if [[ $# != 4 ]]; then
            echo [ERROR]: Please check your parameters, it should have 4 parameters
            exit 1
        fi
        if [[ $1 == '--modify' ]]; then
            PropertyName=$2
            PropertyOldValue=$3
            PropertyValue=$4
        elif [[ $1 == '--addunder' ]]; then
            PropertyName=$2
            PropertyValue=$3
            PropertySection=$4
        fi
    fi
}

checkPath(){
    while (($#));
    do
        if [[ ! `find $1` ]]; then
            echo [ERROR]: Please enter a correct path
            exit 1
        fi
        shift
    done
}

checkFiles(){
    while (($#));
    do
        if [[ ! -f $1 ]]; then
            echo Please enter files only
            echo $1 is not a file
            exit 1
        fi
        shift
    done
}

findallFoldersOrFiles(){
    oldfileDirs=("${fileDirs[@]}") # These are dirs from last round
    for j in ${fileDirs[@]};
    do
        fileDirs+=(`find $j -mindepth 1 -maxdepth 1 -type $1`)
    done
    fileDirs=($(printf "%s\n" "${fileDirs[@]}" "${oldfileDirs[@]}" | sort | uniq -u))  # Remove A values from B
    unset oldfileDirs
}

addProperty(){
    # sh chenvironment.sh wfm-environment\ wfm atl prod sched manifest.yml --addnew CF_STARTUP_TIMEOUT 4000 (add new property called CF_STARTUP_TIMEOUT which has value 4000)
    echo adding new property $2 has value $3 into file $1 ......
    if [[ $spaces == 1 && $dash == 0 ]]; then
        printf "\n\t $2: $3" >> $1
    elif [[ $spaces == 0 && $dash == 0 ]]; then
        printf "\n $2: $3" >> $1
    elif [[ $spaces == 1 && $dash == 1 ]]; then
        printf "\n\t - $2: $3" >> $1
    elif [[ $spaces == 0 && $dash == 1 ]]; then
        printf "\n - $2: $3" >> $1
    fi
    echo Done !!
    echo
    echo
}

deleteProperty(){
    # sh chenvironment.sh wfm-environment\ wfm atl prod sched manifest.yml --delete CF_STARTUP_TIMEOUT 5500
    echo deleting target $2 has value $3 from file $1 ......
    commentOurVariable="#$2"
    sed -ri "s/^(\s*)($2\s*:\s*$3\s*$)/\1$commentOurVariable: $3/" $1
    echo Done !!
    echo
    echo
}

modifyProperty(){
    # sh chenvironment.sh wfm-environment\ wfm atl prod sched manifest.yml --modify CF_STARTUP_TIMEOUT 4000 5500 (change from 4000 to 5500)
    echo modifying target $2, change its value from $3 to $4 inside file $1 ......
    sed -ri "s/^(\s*)($2\s*:\s*$3\s*$)/\1$2: $4/" $1
    echo Done !!
    echo
    echo
}

addUndersection(){
    echo adding new property $2 with value $3 under section $4 ......
    if [[ $spaces == 1 && $dash == 0 ]]; then
        sed -ir "/$4:/a\  $2: $3" $1
    elif [[ $spaces == 0 && $dash == 0 ]]; then
        sed -ir "/$4:/a  $2: $3" $1
    elif [[ $spaces == 1 && $dash == 1 ]]; then
        sed -ir "/$4:/a\  - $2: $3" $1
    elif [[ $spaces == 0 && $dash == 1 ]]; then
        sed -ir "/$4:/a- $2: $3" $1
    fi
    echo Done !!
    echo
    echo
}

manipulateFiles(){
    echo -e "\033[32mFile Name\033[m" is "\033[31m$1\033[m"
    if [[ $operType == '--delete' ]]; then
        echo -e "\033[32mDeleting property $PropertyName ......\033[m"
        echo -e "\033[32mIt has value $PropertyValue\033[m"
        echo "Preparing deleting property $PropertyName with value $PropertyValue from file $1 ......"
        deleteProperty $1 $PropertyName $PropertyValue

        elif [[ $operType == '--modify' ]]; then
            echo "Preparing modifying property $PropertyName from value $PropertyOldValue to value $PropertyValue into file $1 ......"
            modifyProperty $1 $PropertyName $PropertyOldValue $PropertyValue

        elif [[ $operType == '--addnew' ]]; then
            echo -e "\033[32mAdding property $PropertyName ......\033[m"
            echo -e "\033[32mIt has value $PropertyValue\033[m"
            echo "Preparing adding property $PropertyName with value $PropertyValue into file $1 ......"
            addProperty $1 $PropertyName $PropertyValue

        elif [[ $operType == '--addunder' ]]; then
            echo "Preparing adding property $PropertyName with value $PropertyValue under property $PropertySection into file $1 ......"
            addUndersection $1 $PropertyName $PropertyValue $PropertySection   
    fi
}