# Hex Plugin - Validate

Plugin to validate input from the user.

```
{
  "rule": "example input validation for arguments",
  "match": "build*",
  "help": "build <branch> <version>"
  "actions": [
    {
      "type": "hex-validate",
      "command": "${hex.input.1}",
      "config": {
        "type": "string",
        "match_re": "/.+/",
        "match_list": "develop,master,release",
        "success": "Looks good!",
        "failure": "Validation failed :("
      }
    }
  ]
}
```
