package utils

import "github.com/df-mc/dragonfly/server/player"

func RemoveIndex(s []*player.Player, index int) []*player.Player {
	return append(s[:index], s[index+1:]...)
}
