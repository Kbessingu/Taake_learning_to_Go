# Phase 2 Go CLI Tools & Utilities Roadmap

## CLI Fundamentals

1. **Basic CLI Applications**
   - [ ] Command-line arguments
   - [ ] Flag parsing (flag package)
   - [ ] Reading user input (bufio)
   - [ ] Output formatting (fmt, tabwriter)<br>
         ***Project:***
     - [ ] Simple calculator or unit converter CLI

2. **File Operations & Processing**
   - [ ] Reading and writing files
   - [ ] Directory traversal
   - [ ] File searching and filtering
   - [ ] Batch file operations<br>
         ***Project:***
     - [ ] File renamer or organizer tool

## Data Processing Tools

3. **Working with Structured Data**
   - [ ] CSV parsing and generation
   - [ ] JSON processing
   - [ ] Text parsing and regex
   - [ ] Data transformation and filtering<br>
         ***Project:***
     - [ ] Log analyzer or data converter

4. **Advanced CLI with Cobra**
   - [ ] Cobra framework basics
   - [ ] Subcommands and nested commands
   - [ ] Persistent and local flags
   - [ ] Help generation and documentation<br>
         ***Project:***
     - [ ] Multi-command CLI tool (git-style)

## Interactive & Automation Tools

5. **Interactive CLI Menus**
   - [ ] Prompt-based interactions
   - [ ] Menu systems and navigation
   - [ ] Input validation
   - [ ] Persistent data storage<br>
         ***Project:***
     - [ ] Game catalog manager or collection tracker

6. **Automation Utilities**
   - [ ] Running external commands (os/exec)
   - [ ] Task automation and scripting
   - [ ] File watching and monitoring
   - [ ] Scheduling and background tasks<br>
         ***Project:***
     - [ ] Backup tool or file sync utility
---

**Key Libraries:**
- flag (standard library)
- Cobra (spf13/cobra)
- Viper (spf13/viper) - configuration
- survey (AlecAivazis/survey) - interactive prompts
- color (fatih/color) - colored output

**CLI Tool Ideas:**
- File organizer by type/date
- Duplicate file finder
- Text search and replace across files
- Game/movie/book collection manager
- Todo list with priorities
- Note-taking system
- Backup and sync tools

---

**Note:** When it comes to projects, build something that's fun and you can use!
