package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var length int                // dùng để lưu độ dài của chuỗi con
	var max int                   // dùng để lưu độ dài chuỗi dài nhất
	var start int                 // dùng để lưu vị trí bắt đầu của chuỗi con
	var end int                   // dùng để lưu vị trí kết thúc của chuỗi con
	m := make(map[byte]int)       // khởi tạo map , map này dùng để lưu vị trí xuất hiện của các ký tự trong chuỗi
	for i := 0; i < len(s); i++ { // duyệt qua chuỗi
		if _, ok := m[s[i]]; ok { // nếu ký tự đã xuất hiện trong chuỗi
			if m[s[i]] >= start { // nếu vị trí xuất hiện của ký tự đó lớn hơn vị trí bắt đầu của chuỗi con
				start = m[s[i]] + 1 // thì cập nhật lại vị trí bắt đầu của chuỗi con
			}
		}
		m[s[i]] = i              // cập nhật lại vị trí xuất hiện của ký tự
		end = i                  // cập nhật lại vị trí kết thúc của chuỗi con
		length = end - start + 1 // tính độ dài của chuỗi con +1 là vì vị trí bắt đầu và kết thúc là vị trí của ký tự
		if length > max {        // nếu độ dài của chuỗi con lớn hơn độ dài của chuỗi dài nhất
			max = length // thì cập nhật lại độ dài của chuỗi dài nhất
		}
	}
	return max
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))

}
