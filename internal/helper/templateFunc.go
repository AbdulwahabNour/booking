package helper


func seq(start, end int) []int {
    if start > end {
        return nil
    }
    s := make([]int, end-start+1)
    for i := range s {
        s[i] = start + i
    }
    return s
}


func sub(a, b int) int {
    return a - b
}