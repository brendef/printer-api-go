
# Brendan de Faria's Tall Order Test 

A REST API built in GoLang for TallOrder. 

**Domain:** http://brendans-tallorder-test.com:8080

**PLEASE NOTE**: This app does not run over https just yet. So please allow your browser to visit the website. It may give you a warning that your connection is not seccure.



## Tech Stack

**Server:** GoLang

**Client:** Next.js 13, TailwindCSS

**Cloud:** AWS RDS, AWS EC2

**IDE/Editor:** VIM (Lunar Vim - NVIM), VS Code
## End Points

#### (GET) /Printers
Returns all printers as an array of JSON.

#### (GET) /printer/**xxx.xxx.x.xx**
Returns all the printer information related to a specific ip address.

#### (POST) /add_printer
 Add a new printer to the database. 

#### (GET) /update_name/**xxx.xxx.x.xx**/?name=**xxxxxx**
Change the name of the printer specified in the path.

#### (GET) /change_status/**xxx.xxx.x.xx**
Change the status of the printer specified in the path.

#### (DELETE) /remove_printer/**xxx.xxx.x.xx**
Remove the printer specified by ip in the path.
## Lessons Learned

1. I learned more about GoLang. I appreciatet the language so much more.

2. Learned how to host GoLang GIN sevrer



## Run Locally

Clone the project

```bash
  https://github.com/brendef/tallorder-printer-api.git
```

Go to the project directory

```bash
  cd tallorder-printers-api
```

Install dependencies

```bash
  go get ./
```

Start the server

```bash
    go run .
```


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`AWS_DATABASE_NAME`

`AWS_DATABASE_USER`

`AWS_DATABASE_PASSWORD`

`AWS_DATABASE_URL`

`AWS_DATABASE_PORT`
## Related

The frontend to view that interacts with this API can be found at:
[Printers Next JS App](https://github.com/brendef/tallorder-printers-crud)


## FAQ

#### Why does this not work over SSL/HTTPS

I just need to install an SSL Certificate and configure nginx or apache then it should be good to go.

#### Why GoLang

It is an amazing language that is super easy to build and run as a server on a Linux machine.

#### Why AWS EC2

My goal was to get something up and running ASAP. EC2s are super simple make a perfect dev environmnet. I do belive a serverless product such as *Amplify* for the front end and *lambda* would have been better. I will migrate this later one to use these technologies for sure!
## Feedback

I really appreciate feedback. Please reach out to me at brendan.defaria@gmail.com


## License

[MIT](https://choosealicense.com/licenses/mit/)


## Authors

- [@brendef](https://www.github.com/brendef)


## 🚀 About Me
I am a developer who is looking to one day create complex scripts and machine learning algorithms.


## 🛠 Skills
Javascript, Python, MySql, NoSql, PHP, Docker, AWS

