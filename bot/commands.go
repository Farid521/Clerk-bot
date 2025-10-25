package bot

import (
	"github.com/bwmarrin/discordgo"
)

var	Commands = map[string]discordgo.ApplicationCommand {
	"lorem-ipsum" : {
		Name: "lorem-ipsum",
		Description: "command khusus untuk testing server",
		Type: discordgo.ChatApplicationCommand,
	},

	"jadwal" : {
		Name: "jadwal",
		Description: "Menampilkan jadwal kelas",
		Type: discordgo.ChatApplicationCommand,
	},

	"materi" : {
		Name: "materi",
		Description: "menampilkan materi apa saja yang dipelajari berupa link video dan catatan (jika tersedia)",
		Type: discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name: "fisika 1 ",
				Type: discordgo.ApplicationCommandOptionString,
				Description: "Menampilkan materi fisika 1",
				Required: true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name: "[20/10/2025]",
					},
				},
			},
		},
	},

	"absensi" : {
		Name: "absensi",
		Description: "(coming soon) mengecek apakah terdapat absensi yang aktif",
	},

	"sys-status" : {
		Name: "sys-status",
		Description: "check the system status and health",
	},
}
