# A list of spotted guesses that require special handling

When such a case occurs, it is considered ambiguous and `manparse` should always prompt the user for a confirmation.

## Options synopsis parsing

Option synopsis are found in the OPTIONS section of a manpage.

### REMOTE_OPT_SYNOPSIS

**An option synopsis has more than two coma separated values and a reference to a sub-command appears in the description.**

It is likely that the option descriptions should be replaced after extraction of the sub-command PIM.

> `manparse` should be able to fetch the corresponding options in the sibling sub-command model.

#### Example (`go-get`)
**this**
```
-a, -n, -v, -x, -p
      The -a, -n, -v, -x, and -p flags have the same meaning as in 'go build'
      and 'go install'.  See go-build(1).
```

**should be interpreted as** (from go-build(1))
``` xml
-a
      force rebuilding of packages that are already up-to-date.
-n
      print the commands but do not run them.
-v
      print the names of packages as they are compiled.
-x
      print the commands.
-p n
      the number of builds that can be run in parallel. The default is the number of CPUs available.
```

### SUGGESTED_POSIX_IMPLICIT_ASSIGNMENT

**An option synopsis has two option descriptions with the latest having a value assignment and the foremost being a short POSIX switch**

It's likely that the intended POSIX switch is an assignment with the second option description assignment value.

> `manparse` should consider the first option description as a POSIX short assignment with assignment value name taken from the second option description.

#### Example (`mv`)

**this**
```
-t, --target-directory=DIRECTORY
      move all SOURCE arguments into DIRECTORY
```

**should be interpreted as**
```
-t DIRECTORY, --target-directory=DIRECTORY
      move all SOURCE arguments into DIRECTORY
```

### OPTIONAL_EXPLICIT_ASSIGNMENT

**An option synopsis holds an optional explicit assignment**

> `manparse` should split the option synopsis in two, with one flag and one option assignment.
> The description fields should be reviewed by user.


#### Example (`git`)

**this**
```
--exec-path[=<path>]
      Path to wherever your core Git programs are installed...
      If no path is given, git will print the current setting and then exit.
```

**should be interpreted as** (with user feedback for description parts)
```
--exec-path
      Path to wherever your core Git programs are installed...
--exec-path=<path>
      Print the current exec-path setting and exit.
```

### OPTIONAL_IMPLICIT_ASSIGNMENT

**An option synopsis holds an optional implicit assignment**

> `manparse` should split the option synopsis in two, with one flag and one option assignment.
> The description fields should be reviewed by user.


#### Example (invented)

**this**
```
--context [CTX]
      provide CTX, default to ...
```

**should be interpreted as**
```
--context
      provide default context ...
--context CTX
      provide context ...
```
