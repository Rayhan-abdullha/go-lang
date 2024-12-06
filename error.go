// package main
// import "errors"

// import "fmt"

// func main() {
// 	res, err := divide(10, 0)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println(res)
// }
// func divide(a, b int) (int, error) {
//     if b == 0 {
//         return 0, errors.New("cannot divide by zero")
//     }
//     return a / b, nil
// }