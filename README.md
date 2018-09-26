<h1 align="center">Paradigm Gw2 </h1>
<h4 align="center">Version 0.1 </h4>

[![GoDoc](https://godoc.org/github.com/paradigm-network/paradigm-gw2?status.svg)](https://godoc.org/github.com/paradigm-network/paradigm-gw2)
[![Go Report Card](https://goreportcard.com/badge/github.com/paradigm-network/paradigm-gw2)](https://goreportcard.com/report/github.com/paradigm-network/paradigm-gw2)
[![Travis](https://travis-ci.org/paradigm/paradigm.svg?branch=master)](https://travis-ci.org/paradigm/paradigm)
[![Discord](https://img.shields.io/discord/102860784329052160.svg)](https://discord.gg/kU3ewZQ)

Welcome to Paradigm Network !

The code is currently alpha quality, but is in the process of rapid development. The master code may be unstable; stable versions can be downloaded in the release page.


## Paradigm-Gw2 

Paradigm employs DAG-based structure to estabilish the FAAS architecture. In combination with the FAAS structure, there are two similarities between FAAS and DAG. 

>(1) **Scalability**: DAG is of good scalability and unlimited node growth. In FAAS, one function does only one thing, and can expands the volume independently without worrying about affecting other functions. And it can expands faster because of its smaller granularity. 

>(2) **Event driven**: In FAAS, functions does not publish any services, and does not consume any resources when no request is proposed.  Resources will be consumed to respond only when there is a request, and release immediately after the service. Due to this, FAAS naturally applies to any event-driven business scenario, such as  advertisement bidding, identity verification, timed tasks, and the emerging IoT applications. Messages are also transmitted by events in DAG network, which can be used as the event propagation channel of FAAS.


Function computing is naturally an event driven service. The core components of the function computing is the Function Execution Engine (FEE). FEE can be event-driven, that is, when an event occurs, the event triggers the execution of the function. Here, function computing supports API gateway as an event source. When the request setup function calculates the API for the back-end service, the API gateway triggers the corresponding function, which returns the execution result to the API gateway. It can be shown as follow. 


![images](https://github.com/shiningmyheart/paradigm-gw2/blob/master/images/pic-gw.png)


API gateway connects to the Function Execution Engine, you can safely open your functions in the form of API, and realize the functions of authentication, flow control, data conversion and other issues.


When an API gateway calls a function computing service, it converts the relevant data of the API into a Map format and passes it to the function computing service. After the function calculates the service, it returns the related data, such as statusCode, headers, body, in the form of Output Format. The API gateway then maps the content returned by the function calculation to the statusCode, header, body, etc. and returns it to the client.


## Parameter Format

The parameters are transmitted in the following fomat from the API gateway to the FEE. When the FEE is used as the back-end service of the API gateway, the API gateway transmits the request parameters to the input of the event computed by the FFE through a fixed Map structure. FEE obtains the required parameters through the following structure, and then processes.

	{
   	 "path":"api request path",
   	 "httpMethod":"request method name",
   	 "headers":{all headers,including system headers},
   	 "queryParameters":{query parameters},
   	 "pathParameters":{path parameters},
   	 "body":"string of request payload",
 	   "isBase64Encoded":"true|false, indicate if the body is Base64-encode"
	}



After the calculation of the FEE, it returns the parameter in the following JSON format. FEE returns the output to API gateway for furthering execution.

	{
    	 "isBase64Encoded":true|false,
	 "statusCode":httpStatusCode,
         "headers":{response headers},
         "body":"..."
	}
	
## Get Paradigm-gw2
### Get from source code

Clone the Paradigm-VM repository into the appropriate $GOPATH/src/github.com/paradigm-network directory.

```
$ git clone --recursive https://github.com/paradigm-network/paradigm-gw2.git
```
or
```
$ go get github.com/paradigm-network/paradigm-gw2
```
Fetch the dependent third party packages with glide.

```
$ cd $GOPATH/src/github.com/paradigm-network/paradigm-gw2

$ make install
```

Build the source code with make.

```
$ make all
```

After building the source code sucessfully, you should see an executable program:

- `paradigm`: the command line program for Paradigm FEE control



## Build development environment
The requirements to build Paradigm-gw2 are:

- Golang version 1.9 or later
- Glide (a third party package management tool)
- Properly configured Go language environment
- Golang supported operating system


## Contributions

Please open a pull request with a signed commit. We appreciate your help! You can also send your code as emails to the developer mailing list. You're welcome to join the Paradigm mailing list or developer forum.

Please provide detailed submission information when you want to contribute code for this project. The format is as follows:

Header line: explain the commit in one line (use the imperative).

Body of commit message is a few lines of text, explaining things  in more detail, possibly giving some background about the issue  being fixed, etc.

The body of the commit message can be several paragraphs. Please do proper word-wrap and keep columns shorter than 74 characters or so. That way "git log" will show things  nicely even when it is indented.

Make sure you explain your solution and why you are doing what you are  doing, as opposed to describing what you are doing. Reviewers and your  future self can read the patch, but might not understand why a  particular solution was implemented.

Reported-by: whoever-reported-it &
Signed-off-by: Your Name [youremail@yourhost.com](mailto:youremail@yourhost.com)



## Open source community
### Site

- <https://paradigm.fund/>

### Developer Discord Group

- <https://discord.gg/kU3ewZQ/>

## License

The Paradigm-vm library is licensed under the GNU Lesser General Public License v3.0, read the LICENSE file in the root directory of the project for details.