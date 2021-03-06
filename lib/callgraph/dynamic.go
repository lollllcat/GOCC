//Copyright (c) 2020 Uber Technologies, Inc.
//
//Licensed under the Uber Non-Commercial License (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at the root directory of this project.
//
//See the License for the specific language governing permissions and
//limitations under the License.
package callgraph

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/google/pprof/profile"
)

func BuildDynamicCG(profileFile string, buildCallTree bool) *Graph {
	log.Println("Building dynamic callgraph...")
	log.Println("Build calltree: ", buildCallTree)
	log.Println("Profile file: ", profileFile)

	log.Println("Reading profile data file...")
	data, err := ioutil.ReadFile(profileFile)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("          done")
	log.Println("Byte array len = ", len(data))

	log.Println("Parsing profile data...")
	prof, err := profile.ParseData(data)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("          done")
	log.Println("Building callgraph from profile data...")
	graph := New(prof, &Options{
		CallTree:    buildCallTree,
		SampleValue: func(v []int64) int64 { return v[1] },
	})
	log.Println("          done")
	log.Println("  # of nodes = ", len(graph.Nodes))
	return graph
}

func IsBuiltin(funcName string) bool {
	if _, ok := builtins[funcName]; ok {
		return true
	}
	return strings.HasPrefix(funcName, "syscall.") ||
		strings.HasPrefix(funcName, "sync.") ||
		strings.HasPrefix(funcName, "time.") ||
		strings.HasPrefix(funcName, "math.") ||
		strings.HasPrefix(funcName, "os.")
}

func IsAutoGenerated(ni NodeInfo) bool {
	return ni.File == "<autogenerated>"
}

func IsAsm(funcName string) bool {
	_, ok := asmFunctions[funcName]
	return ok
}

func PrintDynamicGraphStats(g *Graph) {
	fmt.Println("  # of nodes = ", len(g.Nodes))
	outEdges := 0
	for _, n := range g.Nodes {
		outEdges += len(n.Out)
	}
	fmt.Println("  # of edges = ", outEdges)
}
