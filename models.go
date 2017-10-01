package foldit

import "github.com/jinzhu/gorm"

type Puzzles struct {
	gorm.Model
	PuzzleID   int
	SolutionID int
}

type Scores struct {
	gorm.Model
	SolutionID int
	Score      float32
	HistoryID  int
}

type Histories struct {
	gorm.Model
	HistoryID  int
	HistoryIX  int
	SolutionID int
}

type Actions struct {
	gorm.Model
	ActionID    int
	ActionType  string
	ActionCount int
}

type Participation struct {
	gorm.Model
	UserID    int
	PuzzleID  int
	ActionID  int
	HistoryID int
}
