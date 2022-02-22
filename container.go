package main

type Container interface {
	Valid(string) bool
}

type container struct {
	brackets [][2]string // len is set as 2 as one pair of brackets can only have 2 symbols
}

func NewContainer(brackets [][2]string) Container {
	return container{brackets: brackets}
}

func (c container) Valid(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	for i := 0; i < len(s)-1; i++ {
		for _, bracket := range c.brackets {
			if string(s[i]) == bracket[0] && string(s[i+1]) == bracket[1] {
				return c.Valid(s[:i] + s[i+2:]) // remove matched bracket and recursively check valid
			}
		}
	}

	return false
}
