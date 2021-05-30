package calculate

func calculate(s string) int {
	// 遍历s, 遇到数字追加到当前tempNum
	// sign 保存上一次的出现的符号，用于与当前数字进行组合，默认'+'
	// stack 用来保存所有数字与操作符的组合 +4 -3 *2 +1
	// 遇到 + - 将tempNum入栈，
	// 遇到 * / 将栈顶取出, 与当前数结合，再入栈
	// 将所有栈结果累加

	sign := '+'
	tempNum := 0
	stack := make([]int, 0)

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			tempNum = tempNum*10 + int(c-'0')
		} else {
			// 遇上操作符，需要将数字入栈，同时更新操作符缓存
			switch sign {
			case '+':
				stack = append(stack, tempNum)
			case '-':
				stack = append(stack, -tempNum)
			case '*':
				//top := stack[len(stack)-1]
				//stack[len(stack)-1] = top*tempNum
				stack[len(stack)-1] *= tempNum
			case '/':
				stack[len(stack)-1] /= tempNum
			}

			// 更新sign tempNum
			sign = rune(c)
			tempNum = 0
		}
	}
	res := 0
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}

func calculateV1(s string) int {

}

func calc(s string, i *int) int {
	// 计算不包含括号的表达式

	res := 0
	sign := '+'
	tempNum := 0

	for *i < len(s) {
		c := s[*i]
		if c >= '0' && c <= 9 {
			tempNum = tempNum*10 + int(c-'0')
		}

		if c == '+' || c == '-' || *i == len(s)-1 {
			switch sign {
			case '+':
				res += tempNum
			case '-':
				res -= tempNum
			}
			sign = rune(c)
			tempNum = 0
		}

		*i++
		if c == '(' {
			tempNum = calc(s, i)
		}
		if c == ')' {
			break
		}
	}
	return res
}
