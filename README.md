# Triangle Classifier

`triangleClassifier` reads the arguments passed from command line making sure the correct amount of arguments
are passed as well as if they can parsed to the expected type.
If they are Ok then a Triangle struct is created passing into it the values from the arguments.
After it, the values are converted to big.Float to try to make easier to handle corner cases of small
or large numbers. 
Only one function from the classifier package is exported as I wanted to only use one single
function to check the type and make all validations internally..

# Use

Simply use the provided `Makefile` to run the different commands:

- `make` will run the directives `clean` and `build` where then the binary `triangleClassifier` can be found in the
root folder.

- `make test` will run the different tests available.

Once we have our binary `triangleClassifier` we can for example simply run `./triangleClassifier 21 21 21` from the root folder.

- `make install` will install the package in `$GOPATH/bin` and then we can simply run `triangleClassifier`.