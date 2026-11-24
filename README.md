# Floyd–Warshall

> Striver SDE Sheet 167 · Day 24 Graph II — DP over allowed intermediate vertices: three loops, k outermost, and the matrix converges to all-pairs shortest paths. Go, stdlib only, tested.

## Problem

Shortest distances between **every** pair of vertices at once, negative edges allowed. Running Dijkstra from each vertex breaks on negatives and costs more code; this is one triple loop.

## Approach

The DP state is sneaky: after iteration k, `d[i][j]` is the shortest path using only intermediates from {0..k}. Each round asks one question — does routing through k help? — via `d[i][k] + d[k][j] < d[i][j]`. The loop order is the whole algorithm: **k must be outermost**; swap it inward and the invariant collapses. Two robustness details worth stealing: INF is set to *half* of MaxInt so `INF + weight` can't overflow into a fake short path (with an explicit skip when `d[i][k]` is INF), and after convergence a negative diagonal entry is the negative-cycle detector — vertex i reaching itself at negative cost.

**Complexity:** O(V³) time, O(V²) space — fine to ~500 vertices, which is its niche.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
