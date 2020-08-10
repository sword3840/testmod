package testmod
import "fmt"
func Hi(name string) string {
  return fmt.Sprintf("Hi, %s", name)
}

func HiMother(name string) string {
  return fmt.Sprintf("Hi mother, %s", name)
}
