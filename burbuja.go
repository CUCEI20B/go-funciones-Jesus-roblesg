package main

func Burbuja(s []int64)  {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				temp := s[i]
				s[i] = s[j]
				s[j] = temp
			}
		}
	}

}

func main()  {
}