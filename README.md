
# Derpy Writer 2
Welcome to Derpy Writer 2. DW2 is a smallish program that takes text and makes more text out of it.

## Running Derpy Writer
Derpy Writer 2 currently is intended to be used as a command line utility (may or may not get a GUI at some point). The ideal way to run DW2 is via your favorite terminal.

### Flags
DW2 uses flags to specify how you want your text.

* len - Specify the length of the generated text. Actual function differes based on output logic but this is commonly the number of tokens generated.
* n - Depth used. Default 1. Higher makes for more accurate text but slower text generation.
* ol - Output logic used. Defaults to 2.
	0. All tokens from any n are equally weighted. Text most odd. 
	1. Double frequency of tokens which occur in n>1 but also in n==1. Text mostly realish but still odd.
	2. Remove all tokens which occur in n>1 but not in n==1. Text most realish.

## Compiling
Derpy Writer 2 is written in Go; as such, you will need a Go compiler. At this time, there are no external dependencies.

A simple `go build` should spit out an executable.

## License
Derpy Writer 2 is licensed under an MIT license.

## The Future of DW2
The following are a few things I want to get done with DW2. These are likely done between class (or in class), so there is not really a planned date to finish these.

* Accept more input and output formats. Currently it accepts only plaintext.
* Optional grammar feature to provide more accurate sentences.
* Premade dictionaries for faster fun.

