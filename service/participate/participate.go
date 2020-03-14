package participate

import (
	"ecobook/repository/participate"
)

func Participate(eventID int, userID int) (int, error) {
	return participate.Participate(eventID, userID)
}
