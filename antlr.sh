#!/bin/bash

export CLASSPATH=".:${HOME}/Documents/lib/antlr-4.7.1-complete.jar:$CLASSPATH"

antlr4() {
	java org.antlr.v4.Tool $@
}
grun() {
	java org.antlr.v4.gui.TestRig $@
}

sudo rm -fr parser/ && antlr4 -Dlanguage=Go -o parser Gocc.g4