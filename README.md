# Git Credential Switcher

This script allows you to switch between different GitLab and GitHub accounts by updating the Git credentials dynamically. It reads credentials from an environment file and sets them using the `git credential approve` command.

## Setup Instructions

### Step 1: Rename the Example Environment File

1. Rename the `.env.example` file to `.env`.
2. Add your GitLab and GitHub credentials to the `.env` file in the following format:

```plaintext
# GitLab credentials
gitlab_1="username1:password1"
gitlab_2="username2:password2"

# GitHub credentials
github_1="username1:password1"
github_2="username2:password2"
```

### Step 2: Create the Environment File Directory

1. Go to your home directory.
2. Create a folder named `gitswitch`.
3. Move the [`.env`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FC%3A%2FRepo%2Fgitscript%2F.env%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%5D "c:\\Repo\gitscript.env") file into the `gitswitch` folder.

#### Commands for Linux/Mac:

```sh
mkdir -p ~/gitswitch
mv .env ~/gitswitch/
```

#### Commands for Windows:

```sh
mkdir %USERPROFILE%\gitswitch
move .env %USERPROFILE%\gitswitch\
```

### Step 3: Build the Script

1. Create a folder named `gitswitch` wherever you want to store the executable. For example, `C:\Program Files\gitswitch` on Windows.
2. Build the script to the desired path.

#### Commands for Linux/Mac:

```sh
mkdir -p /usr/local/bin/gitswitch
go build -o /usr/local/bin/gitswitch/gitswitch main.go
```

#### Commands for Windows:

```sh
mkdir "C:\Program Files\gitswitch"
go build -o "C:\Program Files\gitswitch\gitswitch.exe" main.go
```

### Step 4: Add the Executable to the System PATH

#### Linux/Mac:

1. Add the following line to your `~/.bashrc`, `~/.zshrc`, or `~/.profile` file:
   ```sh
   export PATH=$PATH:/usr/local/bin/gitswitch
   ```
2. Reload the shell configuration:
   ```sh
   source ~/.bashrc  # or source ~/.zshrc or source ~/.profile
   ```

#### Windows:

1. Open the Start Search, type in "env", and select "Edit the system environment variables".
2. In the System Properties window, click on the "Environment Variables" button.
3. In the Environment Variables window, select the `Path` variable in the "System variables" section and click "Edit".
4. Click "New" and add the path to the `gitswitch` executable, e.g., `C:\Program Files\gitswitch`.
5. Click "OK" to close all windows.

### Step 5: Test the Script

You can now run the script from any directory to switch Git credentials.

#### Example Commands:

```sh
gitswitch gitlab 1
gitswitch github 2
```

This will switch the Git credentials for GitLab or GitHub based on the provided arguments.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
