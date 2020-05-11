package saying

import "fmt"

// Greet func return a phrase saying hello!
func Greet(s string) string {
	return fmt.Sprintf("Hello my dear %v", s)
}
