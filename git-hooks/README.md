# Using gitleaks pre-commit hook
To install pre-commit git-hook, you can just execute next commands
```
 cp  git-hooks/pre-commit .git/hooks/ 
 git config core.hooksPath .git/hooks
 #Linux
 #chmod +x .git/hooks/pre-commit
 #Windows
 #icacls .git\hooks\pre-commit /grant Everyone:RX
```

More details about gitleak:
https://github.com/gitleaks/gitleaks?tab=readme-ov-file#
