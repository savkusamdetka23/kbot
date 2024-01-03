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

To enable gitleaks check
```
 git config hooks.gitleaks true
```

To disable gitleaks check
```
 git config hooks.gitleaks false
```


## Usage

```
Usage:
  gitleaks [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  detect      detect secrets in code
  help        Help about any command
  protect     protect secrets in code
  version     display gitleaks version

Flags:
  -b, --baseline-path string       path to baseline with issues that can be ignored
  -c, --config string              config file path
                                   order of precedence:
                                   1. --config/-c
                                   2. env var GITLEAKS_CONFIG
                                   3. (--source/-s)/.gitleaks.toml
                                   If none of the three options are used, then gitleaks will use the default config
      --exit-code int              exit code when leaks have been encountered (default 1)
  -h, --help                       help for gitleaks
  -l, --log-level string           log level (trace, debug, info, warn, error, fatal) (default "info")
      --max-target-megabytes int   files larger than this will be skipped
      --no-color                   turn off color for verbose output
      --no-banner                  suppress banner
      --redact                     redact secrets from logs and stdout
  -f, --report-format string       output format (json, csv, junit, sarif) (default "json")
  -r, --report-path string         report file
  -s, --source string              path to source (default ".")
  -v, --verbose                    show verbose output from scan

Use "gitleaks [command] --help" for more information about a command.
```


More details about gitleak:
https://github.com/gitleaks/gitleaks?tab=readme-ov-file#
