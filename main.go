// func Milord(art string) bool {
//   stack := []rune{}

//   for _, char := range art {
//     if char == '('  char == '{'  char == '[' {
//       stack = append(stack, char)
//     } else if char == ')' {
//       if len(stack) == 0  stack[len(stack)-1] != '(' {
//         return false
//       }
//       stack = stack[:len(stack)-1]
//     } else if char == '}' {
//       if len(stack) == 0  stack[len(stack)-1] != '{' {
//         return false
//       }
//       stack = stack[:len(stack)-1]
//     } else if char == ']' {
//       if len(stack) == 0 || stack[len(stack)-1] != '[' {
//         return false
//       }
//       stack = stack[:len(stack)-1]
//     }
//   }

//   return len(stack) == 0
// }

// func main() {
//   art := "((()))"
//   n := Milord(art)
//   fmt.Println(n)
// }


