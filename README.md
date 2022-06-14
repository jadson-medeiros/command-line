# Project developed with the course: [Write Professional Command-line Programs in Go](https://www.educative.io/courses/prof-command-line-programs-go)

# multi-git
A little Go command-line program that manages a list of git repos and runs git commands on all repos.

This program supports the article: [Let's Go: Command-line Programs with Golang](https://code.tutsplus.com/tutorials/lets-go-command-line-programs-with-golang--cms-26341).

# Command-line Arguments
It accepts two command-line arguments:

* --command : the git command (wrap in double quotes for multi-arguments 
commands)
* --ignore-errors: keeps going through the list of repos even the git command
 failed on some of them

# Environment variables
The list of repos is controlled by two the environment variables:

* MG_ROOT : the path to a root directory that contains all the repos
* MG_REPOS : the names of all managed repos under MG_ROOT