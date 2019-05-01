package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/awalterschulze/gographviz"
)

type node map[string]string

func (n node) addNode(g *gographviz.Graph, parent string) error {
	for k, v := range n {
		if err := g.AddNode(parent, v, map[string]string{"label": "\"" + strings.Replace(k, "@", "\\n", -1) + "\""}); err != nil {
			return err
		}
	}
	return nil
}

type edge struct {
	left  string
	right string
}

func (e edge) String() string {
	return fmt.Sprintf("%s->%s", e.left, e.right)
}

var (
	mapNode = node{}
	edges   = []edge{}
)

func split(text string, count int) int {
	ss := strings.Split(text, " ")
	for _, s := range ss {
		if _, ok := mapNode[s]; !ok {
			count++
			mapNode[s] = fmt.Sprintf("N%d", count)
		}
	}
	e := edge{left: mapNode[ss[0]], right: mapNode[ss[1]]}
	edges = append(edges, e)
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		count = split(scanner.Text(), count)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	parent := "gomod"
	graph := gographviz.NewGraph()
	if err := graph.SetName(parent); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := graph.SetDir(true); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := mapNode.addNode(graph, parent); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, e := range edges {
		if err := graph.AddEdge(e.left, e.right, true, nil); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}
	fmt.Println(graph)
}

/* Copyright 2019 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
