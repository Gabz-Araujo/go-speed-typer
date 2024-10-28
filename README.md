# Go Speed Typer

Go Speed Typer is a command-line interface (CLI) typing game built with Go. It's designed to help users improve their typing speed and accuracy in both Portuguese and English.

## Features

- Command-line interface for a distraction-free typing experience
- Support for both Portuguese and English languages
- Real-time WPM (Words Per Minute) and accuracy calculation
- Colorful interface highlighting correct and incorrect characters
- Multiple game states: Start screen, Typing session, and Result screen
- Option to replay the same text or start a new session
- Efficient input handling using a custom ring buffer implementation

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go 1.16 or higher installed on your system
- A terminal that supports UTF-8 encoding

## Installation

To install Go Speed Typer, follow these steps:

1. Clone the repository:

   ```
   git clone https://github.com/GabzAraujo/go-speed-typer.git
   ```

2. Navigate to the project directory:

   ```
   cd go-speed-typer
   ```

3. Build the project:

   ```
   go build -o speed-typer cmd/go-speed-typer/main.go
   ```

## Usage

To start the typing game, run the following command in your terminal:

```
./speed-typer
```

Once the game starts, you'll see the start screen. Press Enter to begin a typing session or Q to quit.

During a typing session:

- Type the text displayed on the screen
- Your current WPM and accuracy are shown in real-time
- Correctly typed characters are highlighted in green
- Incorrectly typed characters are highlighted in red

After completing a typing session, you'll see your final WPM and accuracy on the result screen. From here, you can:

- Press Shift+Tab to start a new session with a different text
- Press Shift+R to replay the same text
- Press Shift+Q to quit the game

## Contributing

Contributions to Go Speed Typer are welcome! Here are a few ways you can help:

1. Report bugs
2. Suggest new features
3. Add support for more languages
4. Improve the text generation algorithm
5. Enhance the UI/UX

To contribute:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/AmazingFeature`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
5. Push to the branch (`git push origin feature/AmazingFeature`)
6. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) for the TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) for terminal styling

## Contact

If you have any questions or feedback, please open an issue on the GitHub repository.

Happy typing!
