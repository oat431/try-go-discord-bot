package bot

import "github.com/bwmarrin/discordgo"

var session *discordgo.Session

func setSession(s *discordgo.Session) {
	session = s
}

func GetSession() *discordgo.Session {
	return session
}
