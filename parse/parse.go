package parse

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/emicklei/dot"
	"github.com/spiegel-im-spiegel/ggm/config"
	"github.com/spiegel-im-spiegel/ggm/errs"
)

type nodes struct {
	count   int
	mapNode map[string]string
	conf    *config.Config
}

func newNodes(conf *config.Config) *nodes {
	return &nodes{count: 0, mapNode: map[string]string{}, conf: conf}
}

func (ns *nodes) getNodeFrom(graph *dot.Graph, label string) dot.Node {
	if ns == nil {
		return dot.Node{}
	}
	if n, ok := ns.mapNode[label]; !ok {
		ns.count++
		n = fmt.Sprintf("N%d", ns.count)
		ns.mapNode[label] = n
		return ns.addNodeAttr(graph.Node(n)).Attr("label", strings.ReplaceAll(label, "@", "\n"))
	} else {
		return graph.Node(n)
	}
}

func (ns *nodes) addNodeAttr(n dot.Node) dot.Node {
	if ns == nil || ns.conf == nil || ns.conf.Node == nil {
		return n
	}
	for k, v := range ns.conf.Node {
		n.Attr(k, v)
	}
	return n
}

func (ns *nodes) addEdgeAttr(n dot.Edge) dot.Edge {
	if ns == nil || ns.conf == nil || ns.conf.Edge == nil {
		return n
	}
	for k, v := range ns.conf.Edge {
		n.Attr(k, v)
	}
	return n
}

//Cxt is context class for parsing
type Cxt struct {
	graph    *dot.Graph
	nodeList *nodes
}

//New returns new Cxt instance
func New(r io.Reader) (*Cxt, error) {
	var conf *config.Config
	var err error
	if r != nil {
		conf, err = config.Decode(r)
		if err != nil {
			return nil, errs.Wrap(err, "error in parse.New() function")
		}
	}
	c := &Cxt{graph: dot.NewGraph(dot.Directed).ID("G"), nodeList: newNodes(conf)}
	return c, nil
}

//Do returns result parsing data
func (c *Cxt) Do(r io.Reader) error {
	if c == nil || c.graph == nil {
		return errs.Wrap(errs.ErrNullPointer, "error in parse.Cxt.Do() function")
	}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), " ")
		if len(strs) == 2 {
			c.nodeList.addEdgeAttr(c.nodeList.getNodeFrom(c.graph, strs[0]).Edge(c.nodeList.getNodeFrom(c.graph, strs[1])))
		}
	}
	if err := scanner.Err(); err != nil {
		return errs.Wrap(err, "error in parse.Cxt.Do() function")
	}
	return nil
}

func (c *Cxt) Write(w io.Writer) error {
	if c == nil || c.graph == nil {
		return errs.Wrap(errs.ErrNullPointer, "error in parse.Cxt.Do() function")
	}
	c.graph.Write(w)
	return nil
}

func (c *Cxt) String() string {
	if c == nil || c.graph == nil {
		return ""
	}
	return c.graph.String()
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
