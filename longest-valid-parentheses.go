func enqueue(queue []byte, element byte) []byte {
	queue = append(queue, element)
	return queue
}

func dequeue(queue []byte) (byte, []byte, error) {
	if len(queue) > 0 {

		element := queue[0]
		if len(queue) == 1 {
			var tmp = []byte{}
			return element, tmp, nil
		}
		return element, queue[1:], nil
	} else {
		return 0, nil, errors.New("")
	}
}
func checkLeft(s string) int{
    greatest:=0
	queue := make([]byte, 0, len(s))
	streak := 0
	isStreak := false
	debt := 0
	var err error
	for _, x := range s {
		if x == '(' {
			queue = enqueue(queue, byte(x))
			if isStreak {
				debt++
			} else if !isStreak {
				debt = 0
				streak = 0
			}
		}
		if x == ')' {
			_, queue, err = dequeue(queue)
			if err == nil {
				if isStreak {
					debt--
				}
				isStreak = true
				streak++
			} else {
				isStreak = false
				streak = 0
				debt = 0
			}
		}
		if debt < 0 {
			debt = 0
		}
		if streak > greatest && debt == 0 {
			greatest = streak
		}

	}
    return greatest

}
func checkRight(s string) int{
    greatest:=0
	queue := make([]byte, 0, len(s))
	streak := 0
	isStreak := false
	debt := 0
	var err error

	for y := len(s) - 1; y >= 0; y-- {
		x := s[y]

		if x == ')' {
			queue = enqueue(queue, byte(x))
			if isStreak {
				debt++
			} else if !isStreak {
				debt = 0
				streak = 0
			}
		}
		if x == '(' {
			_, queue, err = dequeue(queue)
			fmt.Println(err)
			if err == nil {
				if isStreak {
					debt--
				}
				isStreak = true
				streak++
			} else {
				isStreak = false
				streak = 0
				debt = 0
			}
		}
		if debt < 0 {
			debt = 0
		}
		if streak > greatest && debt == 0 {
			greatest = streak
		}

	}
    return greatest
}


func longestValidParentheses(s string) int {
	l:= checkLeft(s)
	r:= checkRight(s)
    if l>r{
        return 2*l
    }else{
        return 2*r
    }
}