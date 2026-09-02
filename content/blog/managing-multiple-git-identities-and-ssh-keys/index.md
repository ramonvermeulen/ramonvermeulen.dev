+++
title = "Managing Multiple Git Identities and SSH Keys for Engineers Who Often Switch Hats"
date = 2026-09-02T00:18:00+02:00
description = "Stop accidentally using the wrong Git email or SSH key. With a few smart config tweaks, your project directory automatically picks the right identity, so you never have to think about it again."
image = "/img/blog/multiple-hats.jpg"
draft = false
toc_inline = true
tags = ["git", "ssh", "github", "productivity"]
+++

If you're an engineer, especially in consultancy, managing multiple Git and SSH
identities is probably a regular part of your workflow. If you've ever pushed a
commit with your personal email to a client's repository, raised a pull request
with your work email to an open‑source project, or been locked out of GitHub
because you're using the wrong SSH key, this post is for you.

Think about the different contexts you operate in: personal projects, open‑source
contributions, client work, work for your direct employer. Each one isn't just a
different codebase; it's a completely separate identity consisting of:

- Git name and email address
- Account on a service like GitHub or GitLab
- SSH key for authentication
- GPG or SSH key to sign your commits

Manually switching between these is possible, but it's also the kind of thing
you'll inevitably forget. The result? You push code with the wrong author, or
you waste time debugging SSH rejection messages. This post shows you how to make
your computer automatically choose the right identity based on where your project
lives, so you can stop thinking about it and focus on code.
{{< toc >}}

## The solution

We'll configure two separate layers:

1. **Git** – uses the correct `user.name`, `user.email`, and GPG/SSH signing key.
2. **SSH** – uses the correct private key when you push to remotes like GitHub.

The trick is to make both tools **directory‑aware**. When you navigate into a
project directory, your system detects which identity you're using and applies
the right settings. No manual switching, no custom shell aliases, no mental
overhead.

### Simple example

Here's the directory structure we'll use in this guide. Feel free to adapt it to
your own workflow:

```sh {lineNos=false copy=false}
~/projects/
├── personal/           # Personal projects, open‑source contributions
└── work/
    ├── employer/       # Your day job / consultancy firm
    ├── client-a/       # Client A
    └── client-b/       # Client B
```

When you `cd` into `~/projects/personal/`, your system uses your personal Git
identity and SSH key. Move to `~/projects/work/client-a/`, and it automatically
switches to Client A's credentials. The location of the project provides all the
context.

## Making Git directory-aware

Git supports conditional includes through the [`includeIf`](https://git-scm.com/docs/git-config#_includes)
directive. The condition we care about is [`gitdir:...`](https://git-scm.com/docs/git-config#Documentation/git-config.txt-gitdir).
The pattern after `gitdir:` is treated as a [glob](<https://en.wikipedia.org/wiki/Glob_(programming)>).
If the location of the `.git` directory matches the pattern, the include condition
is met.

And here's the key detail: **_"if the pattern ends with `/`, Git treats it as matching
that directory and everything inside it"_**.

Point it at, for example, `~/projects/work/client-a/`, and it'll apply to any
repository nested under that path.

Let's see how to set this up.

### Step 1: Create your profile files

First, create a small configuration file for each identity.

Your personal profile:

```ini {filename="~/.gitconfig_personal" lineNos=false}
[user]
 name = "Your Name"
 email = your.personal@email.com
 signingkey = /Users/yourname/.ssh/id_ed25519-personal-signing.pub
```

Your work profile:

```ini {filename="~/.gitconfig_employer" lineNos=false}
[user]
 name = "Your Name"
 email = your.name@company.com
 signingkey = /Users/yourname/.ssh/id_ed25519-employer-signing.pub
```

Client A's profile:

```ini {filename="~/.gitconfig_client_a" lineNos=false}
[user]
 name = "Your Name"
 email = your.name@client-a.com
 signingkey = /Users/yourname/.ssh/id_ed25519-client-a-signing.pub
```

Client B's profile:

```ini {filename="~/.gitconfig_client_b" lineNos=false}
[user]
 name = "Your Name"
 email = your.name@client-b.com
 signingkey = /Users/yourname/.ssh/id_ed25519-client-b-signing.pub
```

And so on for each client or organization you work with.

> **Note:** You can use either a GPG key or an SSH key for signing. The example
> above uses SSH keys (which is simpler to set up if you already have SSH keys).
> If you prefer GPG, replace the `signingkey` path with your GPG key ID, like
> `signingkey = ABC123DEF456`.

For more information about commit signing, see [Git Commit Signing](https://git-scm.com/book/en/v2/Git-Tools-Signing-Your-Work).
Also, GitHub has an in-depth guide on how to sign commits with either GPG or
SSH keys: [Commit signature verification](https://docs.github.com/en/authentication/managing-commit-signature-verification/about-commit-signature-verification).

### Step 2: Tell Git when to use each profile

Now edit your global `~/.gitconfig` file and add an `includeIf` section for each
relevant directory:

```ini {filename="~/.gitconfig" lineNos=false}
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

- The `[include]` at the top loads your personal profile by default for every
  repository.
- The `[includeIf]` sections that follow override those settings when you're
  inside a specific directory.
- Git processes configuration files **in order**. If the same setting appears
  multiple times, **the last value Git sees wins**.

### Step 3: (Optional) No default fallback

If you prefer a stricter approach where Git doesn't fall back to your personal
identity, you can use `includeIf` for everything and skip the default `[include]`:

```ini {filename="~/.gitconfig" lineNos=false}
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

With [`useConfigOnly = true`](https://git-scm.com/docs/git-config#Documentation/git-config.txt-useruseConfigOnly),
Git will refuse to create a commit if your repository isn't in one of the
configured directories. This is great for avoiding accidental commits with the
wrong identity, but it does mean you have to be explicit about where you store
your projects.

## Making SSH directory-aware

Git now knows which identity to use. But SSH may still use the wrong key when you
push to GitHub or another SSH remote.

**The distinction**: Git controls what's written in your commits (locally). SSH
controls which account you authenticate as when you push. Changing `user.email`
doesn't change your SSH key.

This is true even if you use SSH keys for signing. Signing happens locally to
verify authorship; authentication happens remotely when you push. They're separate
concerns.

> **Note:** This guide assumes you use SSH keys to authenticate with your remote,
> see [Connecting to GitHub with SSH](https://docs.github.com/en/authentication/connecting-to-github-with-ssh).

The examples below use GitHub because it is a common case, but the same pattern
works for any SSH-based Git remote. Replace `github.com` with the host you use,
such as your company Git server or a self-hosted Git service.

### The usual approach: SSH host aliases

A common solution is to create SSH aliases for each account:

```text {filename="~/.ssh/config" lineNos=false}
Host github-client-a
  HostName github.com
  User git
  IdentityFile ~/.ssh/id_ed25519-client-a
```

This works, but your repository URLs become ugly: `git@github-client-a:client-a/repo.git`.

I wanted to keep using the normal remote host, `git@github.com`, and let the
directory determine the SSH key.

### The better approach: Match exec

OpenSSH's [`Match`](https://man.openbsd.org/ssh_config#Match) directive supports
an exec condition that executes a command under your shell. **If the command returns
a zero exit status, the configuration block applies**. We'll use it to test the
current working directory via `pwd | grep -q`; the command succeeds only when the
current path matches the specified directory.

```text {filename="~/.ssh/config" lineNos=false}
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

> **Important:** Unlike Git, SSH's Match exec doesn't reliably expand `~` or `$HOME`.
> Always use the full absolute path for SSH conditions. Replace `/Users/yourname/`
> with your actual home directory path. You can find it with: `echo $HOME`.

#### Why `IdentitiesOnly yes`?

Without this setting, SSH may try all the keys loaded in your SSH agent before
finding the right one. Some servers limit authentication attempts, and offering
too many keys can cause the connection to fail. [`IdentitiesOnly yes`](https://man.openbsd.org/ssh_config#IdentitiesOnly)
tells SSH to only use the key you explicitly specified.

#### Why order matters

SSH uses the **first matching setting**. Git uses the **last matching setting**.
This means your specific directory rules must come before the fallback, otherwise
SSH will always use the fallback key.

### Testing your SSH configuration

To check which key SSH would use for the example GitHub host:

```bash {lineNos=false}
ssh -G github.com | grep identityfile
```

To test authentication with GitHub:

```bash {lineNos=false}
ssh -T git@github.com
```

## Testing everything together

Inside a repository, run these checks:

```bash {lineNos=false}
git config user.name
git config user.email
ssh -G github.com | grep identityfile
```

If everything works, your directory is now controlling both:

- The identity written into your commits.
- The SSH key used to push to your git remote.

## One limitation to keep in mind

SSH's `Match exec` checks the current working directory of the process that starts
SSH. Some tools, like certain IDEs, scripts, or background processes, may start
Git from a different directory. If a tool doesn't respect the repository directory,
`Match exec` won't match. If you run into issues with a specific tool, SSH host
aliases are the safer fallback option.

## Security note

The command inside `Match exec` runs whenever SSH evaluates your config. Keep it
simple, fast, and local. Don't use it to check anything from untrusted sources,
this is a convenience, not a security boundary.

Always protect your private keys on the local system:

```bash {lineNos=false}
chmod 700 ~/.ssh
chmod 600 ~/.ssh/config
chmod 600 ~/.ssh/*id_ed25519* # or whatever your private key filenames are
```

For an additional layer of defense, consider storing your keys in a password
manager's SSH agent, like [**1Password**](https://www.1password.dev/ssh/agent)
or [**Bitwarden**](https://bitwarden.com/help/ssh-agent/), instead of static files.
The agent decrypts your key directly into memory without writing it to disk,
and the `Match exec` logic works unchanged. Just swap `IdentityFile`
for `IdentityAgent`. It's a convenient way to centralize and sync keys, but with
a strong passphrase and full-disk encryption, the standard file-based approach is
already fairly secure.

## Closing thoughts

With these two configurations in place, your computer now knows exactly who you
are based on where you're working. The directory you're in tells both Git and SSH
which identity to use, so you no longer have to remember or manually switch
anything.

No more debugging rejected SSH keys. No more rewriting commit history to fix the
wrong author. Just you, your code, and the right credentials, automatically.
