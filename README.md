![Build Status](https://github.com/tsaisean/iHR/actions/workflows/ci.yml/badge.svg) 
![Coverage](https://img.shields.io/badge/dynamic/json?url=https://tsaisean.github.io/coverage.json&label=Coverage&query=$.coverage&color=brightgreen)

## iHR
iHR is a feature-rich and user-friendly Human Resource (HR) management system designed to streamline HR processes. This open-source project empowers organizations to manage employee data, track attendance, and much more efficiently. Built with Go and powered by the Gin web framework, iHR delivers performance and scalability for modern HR management needs.

## Getting started
### Prerequisites
1. iHR requires Go version 1.22 or above.
2. Generate the secret using the command:
```bash
make gen_secret
```
3. Fill in the secret and other required fields in the **config.toml** file.

### Running the services
This project is built with Docker, allowing you to easily set up the environment with a single command.
``` bash
make docker_run_local
```

### Dependent container services
### Local
| Service      | Version |
| ------------ |---------|
| Redis        | 7.4.2   |
| MySQL        | 8.4     |       

You can also update the config file if you are hosting your own Redis and MySQL services on the cloud platforms like AWS or GCP.


## License
iHR is licensed under the GNU General Public License v3.0. You may use, modify, and distribute this software in accordance with the terms of the license.

See the [LICENSE](./LICENSE) file for more details.
