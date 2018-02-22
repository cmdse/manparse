# A list of spotted edge-cases that require special handling

When such a case occurs, it is considered ambiguous and `manparse` should always prompt the user for a confirmation.

### A. In an OPTIONS section, an option synopsis has more than two coma separated values

Context is a sub-command and a reference to an other sub-command appears in description.

> `manparse` should be able to fetch the corresponding options in the sibling sub-command model.

#### Example (`go-get`)
**this**
```
-a, -n, -v, -x, -p
      The -a, -n, -v, -x, and -p flags have the same meaning as in 'go build'
      and 'go install'.  See go-build(1).
```

**should be interpreted as**
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

### B. In an OPTIONS section, an option synopsis has two option descriptions with only one having a value assignment and the other being a short POSIX switch

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