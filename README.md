# OpenTaaS

OpenTaaS (Tool-as-a-Service) is a platform that provides DevOps tools on demand (lego style), relieving you from manually configuring and tinkering with the tools and their config files. It should also deploy, monitor, and secure them for you on the platform of choice (e.g., hosted ~ [k8s, compose, swarm, etc] or cloud-based ~ [AWS, OpenStack, Azure, etc]). The goal of OpenTaaS is to provide a market place of tools and frameworks with built-in monitoring and security based on common best practices on-demand.

## Overview

OpenTaaS provides three important workflows:

- [x] Create configuration files (including Dockerfiles) for supported tools.
- [x] Compose multiple tools together as a service (composing a compose file).
- [ ] Automate the deployment of the compose services locally or to the cloud.
- [ ] Monitor deployed tools.
- [ ] Provide Built-in security for deployed tools.


**Note:** This project is still under developement and not meant for production in its current state.


## Getting Started

To get started with OpenTaaS, you can download the corresponding binary for your OS (Darwin, Linux, Windows). Or you can clone this repository and build the project locally.

### Clone and build

You need to install `make` and *optionally* `Go` on your system before proceeding.

```bash
git clone https://github.com/zanetworker/opentaas.git
cd opentaas

# build OpenTaaS binary if you have go installed
make OS=<darwin|linux|windows> install

# execute taas for command overview
taas

# build OpenTaaS binary if you don't have go installed
make OS=<darwin|linux|windows> dry

# execute taas command for overview
./taas
```

### Usage

Let's go through the workflows mentioned above to understand how to use OpenTaaS: 

- [x] Create configuration files (including Dockerfiles) for supported tools (e.g, Jenkins, Nginx, Goss)

```bash 
taas create jenkins -u user -p pass
taas create nginx --frontend "jenkins:8081" --backend "jenkins:8080"
taas create goss --conn "tcp:jenkins:8080" --conn "tcp:nginx:8081"
```

> This will create all the configuration files needed for all the tools under configs/tools_name/out (e.g., taas create goss) creates config files and Dockerfile for the goss tool.

```
├── Dockerfile
├── out
│   └── gossconfig.yml
```

> Now that we have all the configuration file we need, we can compose our tools into a serivce. 

- [x] Compose multiple tools together as a service (composing a compose file).

```bash
taas compose --jenkins --nginx --goss
# or
taas compose -j -n -g
```

> This will create (compose) a compose file with all the services specified on the command line

- [ ] Automate the deployment of the compose services locally or to the cloud. 
- [ ] Monitor deployed tools.
- [ ] Provide Built-in security for deployed tools.

## Running the tests

To run the tests you can simply use `make test`. 

## Contributing

<!-- [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) -->
Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

Coming soon!

<!-- We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).  -->

## Authors

See also the list of [contributors](https://github.com/zanetworker/opentaas/graphs/contributors) who participated in this project.

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

The idea for this project was inspired to me by reviewing this repo https://github.com/michaellihs/jenkins-swarm. 