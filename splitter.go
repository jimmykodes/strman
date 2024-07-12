package strman

import "strings"

type splitState int

const (
	stateUnknown splitState = iota
	stateLower
	stateUpper
	stateNum
	stateSymbol
	stateOther
)

type splitter struct {
	src      string
	startPos int
	readPos  int
	hasLower bool
	done     bool
}

func (s *splitter) readState() splitState {
	c := s.src[s.readPos]
	switch {
	case c == 0:
		return stateUnknown
	case isNum(c):
		return stateNum
	case isLower(c):
		s.hasLower = true
		return stateLower
	case isUpper(c):
		return stateUpper
	case isSymbol(c):
		return stateSymbol
	default:
		return stateOther
	}
}

func (s *splitter) next() (string, bool) {
	if s.done {
		return "", true
	}
	currentState := s.readState()
	for currentState == stateOther {
		// eat any characters we don't care about
		s.readPos++
		currentState = s.readState()
	}
	// record start of current "word"
	s.startPos = s.readPos
	hasUpper := currentState == stateUpper
	for {
		s.readPos++
		if s.readPos >= len(s.src) {
			s.done = true
			break
		}
		newState := s.readState()
		hasUpper = hasUpper || newState == stateUpper
		if newState != currentState {
			if newState == stateLower && currentState == stateUpper {
				// transitioning from a capital to a lowercase, generally the same word
				// ie: NewWord - N -> e is upper to lower
				// need to set state to lower so we catch the w in New
				currentState = stateLower
			} else {
				break
			}
		} else if newState == stateUpper && s.hasLower {
			// consider subsequent uppercase letters as separate tokens if the word contains lowercase letters
			// thisIsATest -> this is a test
			// instead of
			// thisIsATest -> this is atest
			break
		}
	}
	if !hasUpper {
		// if the current word has no uppercase, skip the `ToLower` call, since that
		// inherently allocates a new string
		return s.src[s.startPos:s.readPos], false
	}
	return strings.ToLower(s.src[s.startPos:s.readPos]), false
}
