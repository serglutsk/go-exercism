package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card {
        case "ace":
        return 11
        case "two":
        return 2
        case "three":
        return 3
        case "four":
        return 4
        case "five":
        return 5
        case "six":
        return 6
        case "seven":
        return 7
        case "eight":
        return 8
        case "nine":
        return 9
        case "ten", "jack", "queen", "king":
        return 10
        default:
        return 0
    }
	panic("Please implement the ParseCard function")
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	pScore := ParseCard(card1) + ParseCard(card2)
	dScore := ParseCard(dealerCard)

	switch {
	
	case card1 == "ace" && card2 == "ace":
		return "P"

	
	case pScore == 21:
		if dScore < 10 { 
			return "W"
		}
		return "S" 

	
	case pScore >= 17 && pScore <= 20:
		return "S"

	
	case pScore >= 12 && pScore <= 16:
		if dScore >= 7 { 
			return "H"
		}
		return "S" 

	
	default:
		return "H"
	}
}
