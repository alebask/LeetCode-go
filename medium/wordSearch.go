package medium

import (
	"sync"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
func wordSearch(board [][]byte, word string) bool {

	m := len(board)
	n := len(board[0])
	wl := len(word)

	if wl > m*n {
		return false
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var findWord func(i, j, pos int) bool
	findWord = func(i, j, pos int) bool {
		if pos == wl {
			return true
		}

		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || word[pos] != board[i][j] {
			return false
		}

		visited[i][j] = true

		if findWord(i+1, j, pos+1) || findWord(i-1, j, pos+1) ||
			findWord(i, j+1, pos+1) || findWord(i, j-1, pos+1) {
			return true
		}

		visited[i][j] = false
		return false
	}

	for i, row := range board {
		for j, val := range row {
			if val == word[0] && findWord(i, j, 0) {
				return true
			}
		}
	}

	return false
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func wordSearchAsync(board [][]byte, word string) bool {

	m := len(board)
	n := len(board[0])
	wl := len(word)

	if wl > m*n {
		return false
	}

	var findWord func(i, j, pos int, visited [][]bool) bool
	findWord = func(i, j, pos int, visited [][]bool) bool {
		if pos == wl {
			return true
		}

		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || word[pos] != board[i][j] {
			return false
		}

		visited[i][j] = true

		if findWord(i+1, j, pos+1, visited) || findWord(i-1, j, pos+1, visited) ||
			findWord(i, j+1, pos+1, visited) || findWord(i, j-1, pos+1, visited) {
			return true
		}

		visited[i][j] = false
		return false
	}

	res := false
	var mu sync.Mutex
	var wg sync.WaitGroup
	done := make(chan struct{})

	for i, row := range board {
		for j, val := range row {
			if val == word[0] {
				wg.Add(1)
				go func(i, j int) {
					select {
					case <-done:
						return
					default:
						visited := make([][]bool, m)
						for i := range visited {
							visited[i] = make([]bool, n)
						}
						if findWord(i, j, 0, visited) {
							mu.Lock()
							res = true
							mu.Unlock()
							close(done)
						}
					}
					wg.Done()
				}(i, j)
			}
		}
	}

	wg.Wait()

	return res
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func wordSearchChannels(board [][]byte, word string) bool {

	m := len(board)
	n := len(board[0])
	wl := len(word)

	if wl > m*n {
		return false
	}

	var findWord func(i, j, pos int, visited [][]bool) bool
	findWord = func(i, j, pos int, visited [][]bool) bool {
		if pos == wl {
			return true
		}

		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || word[pos] != board[i][j] {
			return false
		}

		visited[i][j] = true

		if findWord(i+1, j, pos+1, visited) || findWord(i-1, j, pos+1, visited) ||
			findWord(i, j+1, pos+1, visited) || findWord(i, j-1, pos+1, visited) {
			return true
		}

		visited[i][j] = false
		return false
	}

	scanBoard := func(board [][]byte, word string, quit <-chan struct{}) <-chan [2]int {
		wordStream := make(chan [2]int)
		go func() {
			for i, row := range board {
				for j, val := range row {
					if val == word[0] {
						select {
						case wordStream <- [2]int{i, j}:
						case <-quit:
							close(wordStream)
							return
						}
					}
				}
			}
			close(wordStream)
		}()
		return wordStream
	}

	findWords := func(wordStream <-chan [2]int, quit <-chan struct{}) <-chan bool {
		foundStream := make(chan bool)
		var wg sync.WaitGroup
		for pos := range wordStream {
			wg.Add(1)
			go func(pos [2]int) {
				visited := make([][]bool, m)
				for i := range visited {
					visited[i] = make([]bool, n)
				}
				if findWord(pos[0], pos[1], 0, visited) {
					select {
					case foundStream <- true:
					case <-quit:
						wg.Done()
						close(foundStream)
						return
					}
				}
				wg.Done()
			}(pos)
		}
		go func() {
			wg.Wait()
			close(foundStream)
		}()

		return foundStream
	}

	quitStream := make(chan struct{})
	defer close(quitStream)

	wordStream := scanBoard(board, word, quitStream)
	foundStream := findWords(wordStream, quitStream)

	return <-foundStream
}
