# Git Cheat Sheet

## 1. My normal evening

```bash
git status                       # what changed?
git switch -c learn/integers     # make a new branch (once per chapter)

# repeat many times
git add .                        # put changes in the box
git commit -m "test(hello): add failing test for French"   # close the box

# end of the evening
git push -u origin learn/integers   # first push of a new branch
git push                            # next pushes
```

When the chapter is done:
1. Open a Pull Request on GitHub (`learn/...` into `main`).
2. Click "Create a merge commit" (not Squash).
3. Then:

```bash
git switch main
git pull
git switch -c learn/next-chapter
```

## 2. How to write a commit message

```
type(area): what this commit does
```

| Type | Use it when |
|---|---|
| `feat` | new behavior |
| `fix` | bug fix |
| `refactor` | better code, same behavior |
| `test` | only tests |
| `docs` | notes, README |
| `chore` | small chores (.gitignore, go.mod) |

Rules:
- English only.
- Start with a verb: `add`, `fix`, `use` (not `added`).
- No dot at the end.
- Short: about 50 characters.
- One idea per commit.

## 3. Look around

```bash
git status                       # what is going on?
git diff                         # changes not added yet
git diff --staged                # changes that will go into the commit
git log --oneline                # short history
git log --oneline --graph --all  # history with branches
git branch                       # list my branches (* = current)
git show <hash>                  # what was in one commit
```

## 4. Problems and fixes

### Wrong message in my last commit (not pushed yet)
```bash
git commit --amend -m "new message"
```

### I forgot a file in my last commit (not pushed yet)
```bash
git add forgotten.go
git commit --amend --no-edit
```

### Throw away my changes in a file
```bash
git restore hello/hello.go
```
Careful: the changes are gone forever.

### I did `git add` by mistake
```bash
git restore --staged hello/hello.go
```
The file is not in the box anymore. My changes stay.

### Undo my last commit, keep my code (not pushed yet)
```bash
git reset --soft HEAD~1
```

### Undo a commit that is already pushed
```bash
git revert <hash>
git push
```
This makes a new commit that cancels the old one. History is safe.

### I need to switch branch, but my work is not ready
```bash
git stash          # hide my changes
git switch main
# do what I need...
git switch learn/hello
git stash pop      # bring my changes back
```

### I committed to `main` by mistake (not pushed yet)
```bash
git switch -c learn/my-work     # new branch keeps my commits
git switch main
git reset --hard origin/main    # main is clean again
```

### Get new changes from GitHub
```bash
git fetch     # download only
git pull      # download and merge
```

### Merge a branch into main (without a Pull Request)
```bash
git switch main
git pull
git merge learn/hello
git push
```

### Merge conflict
1. Run `git status`. It shows the files with a conflict.
2. Open the file. Find `<<<<<<<`, `=======`, `>>>>>>>`.
3. Keep the code I want. Delete these marker lines.
4. Run `go test`.
5. Finish:
```bash
git add <file>
git commit
```

I changed my mind? Run `git merge --abort`.

### Delete a branch I created by mistake
First, leave the branch (you cannot delete the branch you are on):
```bash
git switch main
```
Delete it on my computer:
```bash
git branch -d learn/wrong-name    # safe: refuses if not merged
git branch -D learn/wrong-name    # force: deletes anyway
```
If I already pushed it, delete it on GitHub too:
```bash
git push origin --delete learn/wrong-name
```

### Who changed this line, and when?
```bash
git blame hello/hello.go
```

### I committed a file I do not want (like .DS_Store)
```bash
echo ".DS_Store" >> .gitignore
git rm --cached .DS_Store
git commit -m "chore: ignore .DS_Store"
```

### I committed a password or a token
Treat it as stolen. Change it right now. Deleting the commit is not enough if I already pushed.

## 5. Golden rules

1. Commit often. Push at least once per evening.
2. Before commit: `go test` is green, and `git status` looks normal.
3. Do not work directly in `main`. Use a branch and a Pull Request.
4. Do not rewrite history that is already pushed.
5. Not sure? Run `git status` and `git log --oneline` first.
6. Never commit secrets or keys. Put `.env` in `.gitignore`.