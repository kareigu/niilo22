module main

go 1.16

require (
	commands v0.0.0 //indirect
	github.com/bwmarrin/discordgo v0.23.3-0.20210529215543-f5bb723db8d9
	github.com/joho/godotenv v1.3.0
)

replace commands v0.0.0 => ./src/commands
