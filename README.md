<h3 align=center>
  Runefetch
</h3>
<p align=center>
  A command-line tool for displaying Old-School Runescape Hiscores written in Golang.
</p>

Fetch your Old-School Runescape hiscores and display them with Runefetch! Inspired by tools like [fastfetch](https://github.com/fastfetch-cli/fastfetch) and [neofetch](https://github.com/dylanaraps/neofetch), Runefetch is designed to flex your level 99 skills and thousands of boss completions. Mainly written in Golang for performance with an easy to customize json config file.

> [!NOTE]
> Runefetch has only been tested on different flavors of Arch Linux. However, it should be compatible with most Linux distributions following the XDG Base Directory Specification. Runefetch has not been tested on MacOS or Windows.

### Installation and Usage
1. Clone this github repository `git clone https://github.com/emersonds/runefetch`.
2. Copy one of the config files located in `runefetch/presets` to your config directory, `~/.config/runefetch`.
3. Edit the "name" and "mode" fields to match your Old-School Runescape player name and game mode.
4. Change directories to `runefetch/src` and run `./main` to run it in your terminal.

### Customization
Runefetch supports all activities and skills in Old-School Runescape. Simply add or remove skills, bosses, and minigames from the `modules` section in your `config.json`. Modules are case-insenstive.

### To-Do
- [ ] Set up commands with [Cobra](https://github.com/spf13/cobra).
- [ ] Cache HTTP Request for performance.
- [ ] Add colors and other customization options.
- [ ] Create more preset configs.
- [ ] Test functionality on MacOS and Windows.
