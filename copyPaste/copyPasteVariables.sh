#!/bin/bash

pathLength=0
pathLock=0
spaces=0
dash=0

typePresent=false
operPresent=false

type=""
operType=""
PropertyName=""
PropertyValue=""
PropertySection=""
PropertyOldValue=""

fileDirs=()
oldfileDirs=()
allCmds=()
types=(".txt" ".yml" ".json")
totalCmds=$#