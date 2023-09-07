package main

import (
	"github.com/alex-sim-dh/rowserrcheck/passes/rowserr"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(rowserr.NewAnalyzer()) }
