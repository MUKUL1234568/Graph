package main
import "fmt"
// Calculating the GCD of the two jug capacities
func GCD(a int, b int) int {
    if b == 0 {
        return a
    }
    return GCD(b, a%b)
}

// Pour from larger jug to smaller jug
func pour1(l_jug int, s_jug int, d int) int {
    from := l_jug
    to := 0
    step := 1
    fmt.Printf("Step: %v\n Larger: %v\t Smaller Jug: %v\n",step,from,to)
        for from != d {
        // Find the maximum amount that can be poured
        temp := min(from, s_jug-to)
        to += temp
        from -= temp
        step++
        fmt.Printf("Step: %v\n Larger: %v\t Smaller Jug: %v\n",step,from,to)
          if from == d {
            break
        }
// If the larger jug becomes empty, fill it again
        if from == 0 {
            from = l_jug
            step++
            fmt.Printf("Step: %v\n Larger: %v\t Smaller Jug: %v\n",step,from,to)
        }
 // If the smaller jug is full, empty it
        if to == s_jug {
            to = 0
            step++
            fmt.Printf("Step: %v\n Larger: %v\t Smaller Jug: %v\n",step,from,to)
         }
    }
    return step  }

// Pour from smaller jug to larger jug
func pour2(l_jug int, s_jug int, d int) int {
    from := s_jug
    to := 0
    step := 1
    fmt.Printf("Step: %v\n Larger: %v\t Smaller Jug: %v\n",step,to,from)
    for to != d {
        // Find the maximum amount that can be poured
        temp := min(from, l_jug-to)
        to += temp
        from -= temp
        step++
        fmt.Printf("Step: %v\n Larger: %v\t Smaller Jug: %v\n",step,to,from)
        if to == d {
            break
        }

        // If the smaller jug becomes empty, fill it again
        if from == 0 {
            from = s_jug
            step++
            fmt.Printf("Step: %v\n Larger: %v\t Smaller Jug: %v\n",step,to,from)
        }

        // If the larger jug is full, empty it
        if to == l_jug {
            to = 0
            step++
            fmt.Printf("Step: %v\n Larger: %v\t Smaller Jug: %v\n",step,to,from)
        }
    }
    return step
}

// Function to calculate minimum steps
func minSteps(l_jug, s_jug, d int) int {
    
    if d > l_jug {
        return -1
    }


    if d % GCD(l_jug, s_jug) != 0 {
        return -1
    }

     fmt.Printf("step by Solution 1(From large to small)\n\n")
     var s1=pour1(l_jug, s_jug, d)
     fmt.Printf("step by Solution 2(From small to large)\n\n")
     var s2=pour2(l_jug, s_jug, d)
    return min(s1, s2)
    // return pour2(l_jug,s_jug,d)
}

 
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    var l_jug, s_jug, d int
    fmt.Println("Enter values for larger jug, smaller jug, and desired amount (l_jug, s_jug, d):")
    fmt.Scan(&l_jug, &s_jug, &d)
    
    result := minSteps(l_jug, s_jug, d)
    if result == -1 {
        fmt.Println("It's not possible to measure the desired amount of water")
    } else {
        fmt.Printf("Minimum number of steps required: %d\n", result)
    }
}
