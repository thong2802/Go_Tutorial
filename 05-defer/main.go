package main

import "fmt"

func main() {
	// Trì hoãn (defer) là một khái niệm khá mới trong điều khiển luồng.
	//Nó cho phép một câu lệnh được gọi ra nhưng không thực thi ngay mà hoãn lại đến khi các lệnh xung quanh trả về kết quả.
	defer fmt.Println("World")
	fmt.Println("Hello")

	// Các lệnh được gọi qua từ khóa defer sẽ được đưa vào một stack,
	// tức là hoạt động theo cơ chế vào sau ra trước (last-in-first-out).
	// Lệnh nào defer sau sẽ được thực thi trước, giống như xếp 1 chồng đĩa
	// thì chiếc đĩa sau cùng (ở trên cùng) sẽ được lấy ra trước.
}
