# Test AWS Beanstalk with Golang

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://img.shields.io/github/go-mod/go-version/otaviobaldan/test-beanstalk)
[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)

Just a simple test to create AWS Beanstalk with Golang

### How to use
- Configure [AWS Credentials](https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/setup-credentials.html)
- Install [EB Cli](https://docs.aws.amazon.com/es_en/elasticbeanstalk/latest/dg/eb-cli3-install.html)
- Create a [RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CreateDBInstance.html) - remember to leave it with public access enabled
  * Connect to instance
  * Create a database
  * Execute `users.sql` file
- Execute `eb create` to create the infrastructure - this step will fail because you don't have app yet
- Configure the environment variables
    * `DB_USERNAME` - put your RDS username
    * `DB_PASSWORD` - the password that you setup to your user
    * `DB_HOST` - the URL to access your RDS Database
    * `DB_SCHEMA` - the name of your database
- Go to your database and add the security groups that AWS Beanstalk created
- On your terminal, execute `sh ./deploy.sh` - it will deploy a version on Beanstalk