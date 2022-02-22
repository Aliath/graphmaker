# graphmaker

Tiny project to build graph layout by edge weights.


```golang
edges := []graphmaker.Edge{
    {source: "A", target: "B", weight: 5},
    {source: "A", target: "C", weight: 5 * math.Sqrt(2)},
    {source: "A", target: "D", weight: 5},

    {source: "B", target: "C", weight: 5},
    {source: "B", target: "D", weight: 5 * math.Sqrt(2)},

    {source: "C", target: "D", weight: 5},
}

nodeIdentifiers := []string{"A", "B", "C", "D"}

result, _ := BuildGraph(edges, nodeIdentifiers)

fmt.Println(result) // [{0 0 A} {5 0 B} {5.000000000000001 -5 C} {7.105427357601002e-16 -5.000000000000001 D}]
```