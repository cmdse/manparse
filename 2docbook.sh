#!/bin/bash

# This script generate one docbook xml file for each program name given as operand, or for every program in
# /usr/share/man/man1/ folder when given the --all option
# -o or --cwd option must be given to define an output directory

# Synopsis
#
# docbook2man {-o <dir> | --cwd}  PROGRAM...
# docbook2man {-o <dir> | --cwd} {-a | --all}

# Dependencies
#
# gunzip, doclifter

manvol=""
mandir=""
pkgmgr=""
outputdir=""
genall=false
programs=()

setManVol() {
  manvol="$1"
  mandir="/usr/share/man/man$1/"
}

initPkgMgr() {
  if [ ! -z "$(which yum)" ]; then
  	pkgmgr="yum install"
  fi
  if [ ! -z "$(which apt-get)" ]; then
  	pkgmgr="apt-get install"
  fi
}

checkpkg() {
	if [[ $(which "$1") == "" ]]; then
		echo -n "Package '$1' not found!"
    if [[ ! $pkgmgr == "" ]]; then
      echo "Attempt installation? (y/n)"
      read -r -n1 answer
      echo
      case $answer in
        y) $pkgmgr "$1"
        ;;
        n) echo -n "Proceed anyway? (y/n) "
        read -r -n1 answer2
        echo
        if [[ "$answer2" == "n" ]] ; then exit
        fi
        ;;
      esac
    fi
	fi
}

echoerr() {
  # shellcheck disable=SC1117
  echo -e "\033[0;31m$1\033[0m"  1>&2
}

storeArguments() {
while [[ $# -gt 0 ]]
do
key="$1"

case $key in
    -o|--output-dir)
    outputdir="$2"
    shift # past argument
    shift # past value
    ;;
    --cwd)
    outputdir=$(pwd)
    shift # past argument
    ;;
    -a|--all)
    genall=true
    shift # past argument
    ;;
    *)    # unknown option
    programs+=("$1") # save it in an array for later
    shift # past argument
    ;;
esac
done
}

checkOutput() {
  if [[ $outputdir = "" ]] ; then
    echoerr "missing output directory ; set it with --output-dir or --cwd options"
    exit 1
  fi
  if [[ ! -d $outputdir ]] ; then
    echoerr "$outputdir is not a directory"
    exit 1
  fi
}

checkPrograms() {
  if [[ $genall = false ]] ; then
    for prog in "${programs[@]}" ; do
      manpath="$mandir/$prog.$manvol.gz"
      if [[ ! -e $manpath ]] ; then
        echoerr "Couldn't find a manpage for program '$prog' at '$manpath'"
        exit 1
      fi
    done
  fi
}

checkArguments() {
  checkOutput
  checkPrograms
}

# $1 : program name
generateDocBook() {
  progname=$(basename "$1")
  progname=${progname%.*}
  gunzip -c "$1" | doclifter -v -e US-ASCII > "$outputdir/$progname.xml"
}

genFiles() {
  if [[ $genall = false ]] ; then
    for prog in "${programs[@]}" ; do
      manpath="$mandir/$prog.$manvol.gz"
      generateDocBook "$manpath"
    done
  else
    for file in $mandir/* ; do
      generateDocBook "$file"
    done
  fi
}

setManVol 1
initPkgMgr
checkpkg doclifter
checkpkg gunzip
storeArguments "$@"
checkArguments
genFiles
