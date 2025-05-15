# GFC Discord Bot

A Discord bot designed to manage AWS Lightsail instances for the Greater Faith Church website. This bot allows authorized Discord users to control server operations and monitor website status directly from Discord.

## Features

- **Server Management**:

  - Start the AWS Lightsail instance
  - Stop the AWS Lightsail instance
  - Reboot the AWS Lightsail instance
  - Check server status
  - Verify website availability (HTTP status check)

- **Access Control**:

  - Role-based permission system
  - Only users with the configured admin role can use the bot commands

- **Bot Administration**:
  - Check bot status

## Prerequisites

- Go 1.16 or higher
- Discord Bot Token
- AWS credentials with Lightsail permissions
- Discord Server (Guild) with administrative access

## Installation

### Manual Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/glucktek/gfc-d-bot.git
   cd gfc-d-bot
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Build the application:
   ```bash
   go build -o gfc-d-bot
   ```

## Configuration

The bot requires the following environment variables:

| Variable                | Description                                         |
| ----------------------- | --------------------------------------------------- |
| `DISCORD_BOT_TOKEN`     | Your Discord bot token                              |
| `DISCORD_ADMIN_ROLE`    | Role ID that is allowed to use the bot commands     |
| `DISCORD_GUILD_ID`      | ID of your Discord server                           |
| `AWS_ACCESS_KEY_ID`     | AWS access key with Lightsail permissions           |
| `AWS_SECRET_ACCESS_KEY` | AWS secret key                                      |
| `AWS_REGION`            | AWS region where your Lightsail instance is located |

You can set these in a `.env` file or export them directly in your environment.

Example `.env` file:

```
DISCORD_BOT_TOKEN=your_bot_token_here
DISCORD_ADMIN_ROLE=your_admin_role_id
DISCORD_GUILD_ID=your_guild_id
AWS_ACCESS_KEY_ID=your_aws_access_key
AWS_SECRET_ACCESS_KEY=your_aws_secret_key
AWS_REGION=us-east-1
```

## Usage

### Running the Bot

After configuring the environment variables, run the bot:

```bash
./gfc-d-bot
```

### Bot Commands

The bot uses Discord's slash commands system. Available commands:

- `/gfcbot server start` - Start the Lightsail instance
- `/gfcbot server stop` - Stop the Lightsail instance
- `/gfcbot server reboot` - Reboot the Lightsail instance
- `/gfcbot server status` - Check the current state of the Lightsail instance
- `/gfcbot server check-website` - Verify if the website is returning a 200 status code
- `/gfcbot bot status` - Check if the bot is running correctly

Only users with the configured admin role can execute these commands.

## Docker Deployment

The project includes Docker support for easy deployment.

### Building the Docker Image

```bash
docker build -t gfc-d-bot .
```

### Running with Docker Compose

1. Make sure your `.env` file is configured with all required variables
2. Run the bot using Docker Compose:

```bash
docker-compose up -d
```

This will start the bot as a detached service. You can check logs with:

```bash
docker-compose logs -f
```

To stop the bot:

```bash
docker-compose down
```

Created by [Glucktek](https://github.com/glucktek)
