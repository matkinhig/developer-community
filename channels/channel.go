package channels

func OK(ch <-chan bool) bool {
	select {
	case ok := <-ch:
		if ok {
			return true
		}
		return false
	}
}
