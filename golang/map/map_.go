package main

import "fmt"

// map遍历
func MapRange() {
    blogArticleViews := map[string]int{
        "unix":         0,
        "python":       1,
        "go":           2,
        "javascript":   3,
        "testing":      4,
        "philosophy":   5,
        "startups":     6,
        "productivity": 7,
        "hn":           8,
        "reddit":       9,
        "C++":          10,
    }
    for key, views := range blogArticleViews {
        fmt.Println("There are", views, "views for", key)
    }
    // 结果随机
    // There are 9 views for reddit
    // There are 1 views for python
    // There are 4 views for testing
    // There are 5 views for philosophy
    // There are 7 views for productivity
    // There are 8 views for hn
    // There are 0 views for unix
    // There are 2 views for go
    // There are 3 views for javascript
    // There are 6 views for startups
    // There are 10 views for C++
}

func main() {
    MapRange()
}
