package services

import "fmt"

func main() {

	userRouter := UserServiceRouter()
	productRouter := ProductServiceRouter()
	go func() {
		if err := userRouter.Run(":8080"); err != nil {
			fmt.Printf("user failed to start", err)
		}
	}()
	go func() {
		if err := productRouter.Run(":8081"); err != nil {
			fmt.Printf("product failed to start", err)
		}
	}()
	select {}
}
