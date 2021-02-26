package feb

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, len(pushed))
	for i := range stack {
		stack[i] = -1
	}

	stackTop := -1
	pushInd := 0
	popInd := 0

	for i := 0; i < len(pushed)+len(popped); i++ {
		if stackTop == -1 || stack[stackTop] != popped[popInd] {
			if pushInd >= len(pushed) {
				return false
			}

			stackTop++
			stack[stackTop] = pushed[pushInd]
			pushInd++
		} else {
			if popInd >= len(popped) {
				return false
			}

			stackTop--
			popInd++
		}

	}

	return true
}
