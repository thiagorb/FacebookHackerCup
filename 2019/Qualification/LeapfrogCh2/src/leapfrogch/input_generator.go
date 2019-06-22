package leapfrogch

func inputGenerator() chan string {
	var _generate func() chan string
	_generate = func() chan string {
		ch := make(chan string)
		go func() {
			ch <- ""

			g := _generate()

			for true {
				c := <-g
				ch <- c + "."
				ch <- c + "B"
			}
		}()

		return ch
	}

	return _generate()
}
