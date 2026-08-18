+++
title = "Managing Multiple Git Identities and SSH Keys for Engineers Who Often Switch Hats"
date = 2026-05-28T00:00:00+02:00
description = "Use directory-aware Git and SSH configuration to automatically select the right identity and SSH key for personal projects, work, clients, and open source."
image = "/img/blog/multiple-hats.jpg"
draft = false
toc_inline = true
tags = ["git", "ssh", "github", "productivity"]
+++
If you're a software engineer, especially if you do any consulting or freelance work, you probably wear a few different 
"hats". One day you're working on a personal project, the next you're contributing to a client's codebase, and then 
you're pushing a fix to an open-source library.

Each of these "hats" often comes with its own:

- Git name and email address.
- Account on a service like GitHub or GitLab.
- SSH key to authenticate with remote servers.
- GPG or SSH key to cryptographically sign your commits.

Switching between these manually is doable, but it's also something you'll inevitably forget. The result? You push code 
with your personal email to a client's repository, or you spend ten minutes scratching your head over why GitHub is 
rejecting your SSH key. This post shows you how to automate this process so you never have to worry about it again.

{{< toc >}}

## The solution

We're going to solve this by making your computer smart about which "hat" you're wearing, based on where your project 
is stored. The goal is a setup where moving into a project directory is all it takes:

```text {filename="Example project directory structure"}
~/projects/
├── personal/           # Personal projects, open-source contributions
└── work/
    ├── employer/       # Your day job / consultancy firm
    ├── client-a/       # Client A
    └── client-b/       # Client B
```

When you navigate to `~/projects/personal/`, your system automatically uses your personal Git name, email, SSH key, 
and commit-signing key. When you move to `~/projects/work/client-a/`, it switches to Client A's credentials instead.

No manual switching. No custom shell aliases to remember. No "which account am I currently using?" moments.
You just move into a project directory and start working.

To make that happen, we need to configure two separate things:

- Git must use the correct author identity.
- SSH must use the correct private key.

## Making Git directory-aware

Git supports conditional configuration through [`includeIf`](https://git-scm.com/docs/git-config#_includes). This lets 
you load different settings based on where your repository lives.

### Step 1: Create your profile files

First, create a small configuration file for each identity.

Your personal profile:

```toml {filename="~/.gitconfig_personal"}
[user]
  name = Your Name
  email = your.personal@email.com
  signingkey = /Users/yourname/.ssh/id_ed25519-personal-signing.pub
```

Your work profile:

```toml {filename="~/.gitconfig_employer"}
[user]
  name = Your Name
  email = your.name@company.com
  signingkey = /Users/yourname/.ssh/id_ed25519-employer-signing.pub
```

Client A's profile:
```toml {filename="~/.gitconfig_client_a"}
[user]
  name = Your Name
  email = your.name@client-a.com
  signingkey = /Users/yourname/.ssh/id_ed25519-client-a-signing.pub
```

Client B's profile:
```toml {filename="~/.gitconfig_client_b"}
[user]
  name = Your Name
  email = your.name@client-b.com
  signingkey = /Users/yourname/.ssh/id_ed25519-client-b-signing.pub
```


And so on for each client or organization you work with.

> **Note:** You can use either a GPG key or an SSH key for signing. The example above uses SSH keys (which is simpler 
> to set up if you already have SSH keys). If you prefer GPG, replace the `signingkey` path with your GPG key ID, like 
> `signingkey = ABC123DEF456`.

For more information about commit signing, see [Git Commit Signing](https://git-scm.com/book/en/v2/Git-Tools-Signing-Your-Work).
Also, GitHub has an in-depth guide on how to sign commits with either GPG or SSH keys: [Commit signature verification](https://docs.github.com/en/authentication/managing-commit-signature-verification/about-commit-signature-verification).

### Step 2: Tell Git when to use each profile

Now edit your global `~/.gitconfig` file and add an `includeIf` section for each relevant directory:

```toml {filename="~/.gitconfig"}
# Load your personal profile as the default fallback
[include]
path = ~/.gitconfig_personal

# Override with specific profiles based on directory
[includeIf "gitdir:~/projects/work/employer/"]
path = ~/.gitconfig_employer

[includeIf "gitdir:~/projects/work/client-a/"]
path = ~/.gitconfig_client_a

[includeIf "gitdir:~/projects/work/client-b/"]
path = ~/.gitconfig_client_b

# General settings that apply everywhere
[core]
autocrlf = input
excludesFile = ~/.gitignore # global gitignore file

[gpg]
format = ssh

[commit]
gpgsign = true

[tag]
gpgSign = true

[pull]
rebase = true
```

Here's how this works:

- The `[include]` at the top loads your personal profile by default for every repository.
- The `[includeIf]` sections that follow override those settings when you're inside a specific directory.
- Git processes configuration files in order. If the same setting appears multiple times, the last value Git sees wins.

For example, if you're in `~/projects/work/client-a/`, Git loads:

1. Your personal profile (`.gitconfig_personal`)
2. Then Client A's profile (`.gitconfig_client-a`)

Since Client A's profile is loaded after your personal one, its email and signingkey settings take precedence. This 
gives you a clean fallback—personal is the default, and client settings override it when needed.

### Step 3: (Optional) No default fallback

If you prefer a stricter approach where Git doesn't fall back to your personal identity, you can use `includeIf` for 
everything and skip the default `[include]`:

```toml {filename="~/.gitconfig"}
[user]
  useConfigOnly = true

[includeIf "gitdir:~/projects/personal/"]
  path = ~/.gitconfig_personal

[includeIf "gitdir:~/projects/work/employer/"]
  path = ~/.gitconfig_employer

[includeIf "gitdir:~/projects/work/client-a/"]
  path = ~/.gitconfig_client_a

[includeIf "gitdir:~/projects/work/client-b/"]
  path = ~/.gitconfig_client_b
```

With `useConfigOnly = true`, Git will refuse to create a commit if your repository isn't in one of the configured 
directories. This is great for avoiding accidental commits with the wrong identity, but it does mean you have to be 
explicit about where you store your projects.

## Making SSH directory-aware

Git now knows which identity to use. But SSH may still use the wrong key when you push to GitHub.

**The distinction:** Git controls what's written in your commits (locally). SSH controls which account you authenticate 
as when you push. Changing `user.email` doesn't change your SSH key. 

> **Note:** This tutorial assumes you already use SSH Keys to authenticate with GitHub. If you don't, see [Connecting to GitHub with SSH](https://docs.github.com/en/authentication/connecting-to-github-with-ssh)
> or keep on using HTTPS.

### The usual approach: SSH host aliases

A common solution is to create SSH aliases for each account:

```sshconfig {filename="~/.ssh/config"}
Host github-client-a
  HostName github.com
  User git
  IdentityFile ~/.ssh/id_ed25519-client-a
```

This works, but your repository URLs become ugly: `git@github-client-a:client-a/repo.git`. 

I wanted to keep using `git@github.com` and let the directory determine the SSH key. Here's how.

### The better approach: Match exec
OpenSSH has a Match feature that can execute a command and apply settings only when the command succeeds. We will use it
to check the current working directory.

```text {filename="~/.ssh/config"}
# Personal
Match host github.com exec "pwd | grep -q '^/Users/yourname/projects/personal/'"
  IdentityFile ~/.ssh/id_ed25519-personal
  IdentitiesOnly yes

# Employer
Match host github.com exec "pwd | grep -q '^/Users/yourname/projects/work/employer/'"
  IdentityFile ~/.ssh/id_ed25519-employer
  IdentitiesOnly yes

# Client A
Match host github.com exec "pwd | grep -q '^/Users/yourname/projects/work/client-a/'"
  IdentityFile ~/.ssh/id_ed25519-client-a
  IdentitiesOnly yes

# Client B
Match host github.com exec "pwd | grep -q '^/Users/yourname/projects/work/client-b/'"
  IdentityFile ~/.ssh/id_ed25519-client-b
  IdentitiesOnly yes

# Fallback (must come last)
Host github.com
  HostName github.com
  User git
  IdentityFile ~/.ssh/id_ed25519
  IdentitiesOnly yes
```

> **Important:** Unlike Git, SSH's Match exec doesn't reliably expand `~` or `$HOME`. Always use the full absolute path 
> for SSH conditions. Replace /Users/yourname/ with your actual home directory path. You can find it with: `echo $HOME`.

#### Why `IdentitiesOnly yes`?
Without this setting, SSH may try all the keys loaded in your SSH agent before finding the right one. Some servers limit 
authentication attempts, and offering too many keys can cause the connection to fail. `IdentitiesOnly yes` tells SSH to 
only use the key you explicitly specified.

#### Why order matters
SSH uses the **first matching setting**. Git uses the **last matching setting**. This means your specific directory 
rules must come before the fallback, otherwise SSH will always use the fallback key.

### Testing your SSH configuration
To check which key SSH would use:
```bash
ssh -G github.com | grep identityfile
```

To test authentication with GitHub:
```bash
ssh -T git@github.com
```

## Testing everything together
Inside a repository, run these checks:

```bash
git config user.name
git config user.email
ssh -G github.com | grep identityfile
```

If everything works, your directory is now controlling both:

- The identity written into your commits.
- The SSH key used to push to GitHub.

## One limitation to keep in mind
Git's `includeIf` checks the repository location. This is reliable.
SSH's `Match exec` checks the current working directory of the process that starts SSH. Some tools—like certain IDEs, 
scripts, or background processes, may start Git from a different directory. If a tool doesn't respect the repository 
directory, `Match exec` won't match. If you run into issues with a specific tool, SSH host aliases are the safer fallback 
option.

## Security note
The command inside `Match exec` runs whenever SSH evaluates your config. Keep it simple, fast, and local. Don't use it 
to check anything from untrusted sources, this is a convenience, not a security boundary.

Always protect your private keys:

```bash
chmod 700 ~/.ssh
chmod 600 ~/.ssh/config
chmod 600 ~/.ssh/*id_ed25519* # or whatever your private key filenames are
```

## Closing thoughts
Whether you're consulting, balancing personal and company projects, or contributing to open source, you'll probably 
switch hats from time to time. With directory-aware Git and SSH configuration, the location of a project provides the 
context. You enter the repository, and the appropriate identity follows automatically. 

No manual switching. No custom shell aliases. No "which account am I currently using?" moments.

Just you, your code, and the right hat for the job.