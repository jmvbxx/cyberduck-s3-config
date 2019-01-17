# Cyberduck S3 Configuration Tool

## Purpose

To facilitate the production, and management, of the required Cyberduck S3
configuration file when `assume role` is required for Amazon AWS
[Security Token Service](https://docs.aws.amazon.com/STS/latest/APIReference/Welcome.html).

## Usage (output from --help)
```bash
Available Commands:
  append      A brief description of your command
  delete      To delete an existing config file
  generate    To generate
  help        Help about any command

Flags:
  -a, --account string   The account number you wish to access
  -e, --email string     Your email
  -h, --help             help for cyberduck-s3-config
  -m, --mfa string       The account number you'd like to MFA from
  -n, --name string      The name of the profile you'd like to create
  -r, --role string      Your role
  ```

