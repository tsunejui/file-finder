## File-Finder
File Finder is a CLI tool for finding the path of a file in a complex directory tree

## CLI Overview

```
This tool provides an easy and extensible way to finding the file.

Usage:
  finder [command]

Available Commands:
  help        Help about any command
  search      Search the path of the specified file

Flags:
  -h, --help   help for finder

Use "finder [command] --help" for more information about a command.
```

## Table of Contents

- [About](#about)
  - [BFS, Breadth-first Search](#bfs)
  - [DFS, Depth-first Search](#dfs)
- Usage
- Copyright notice

### About

This tool implement the BFS algorithm and the DFS algorithm in Golang, they are popular search algorithms that can be used for both tree and graph data structures.

<a id="bfs"></a>

### BFS, Breadth-first Search

Breadth-first search (BFS) is an algorithm for searching a tree data structure for a node that satisfies a given property. It starts at the tree root and explores all nodes at the present depth prior to moving on to the nodes at the next depth level. Extra memory, usually a queue, is needed to keep track of the child nodes that were encountered but not yet explored.

reference: https://en.wikipedia.org/wiki/Breadth-first_search

<a id="dfs"></a>

### DFS, Depth-first Search

Depth-first search (DFS) is an algorithm for traversing or searching tree or graph data structures. The algorithm starts at the root node (selecting some arbitrary node as the root node in the case of a graph) and explores as far as possible along each branch before backtracking.

reference: https://en.wikipedia.org/wiki/Depth-first_search
