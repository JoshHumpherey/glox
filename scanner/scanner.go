package scanner

type Scanner struct {
	Source string
}

type Token struct {
}

func (s *Scanner) ScanTokens() []Token {
	return []Token{}
}
