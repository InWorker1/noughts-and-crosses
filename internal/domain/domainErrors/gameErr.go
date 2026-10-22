package domainErrors

import "errors"

var ErrGameNotFound = errors.New("game not found")
var ErrInvalidBoard = errors.New("invalid board")
var ErrNotYourGame = errors.New("not your game")
var ErrGameWaiting = errors.New("game is waiting for second player")
var ErrGameFinished = errors.New("game is finished")
var ErrNotYourTurn = errors.New("not your turn")
