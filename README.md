<img src="icon.png" width="89" height="130" align="right" alt="reloadly-golang-icon"/>

# Reloadly SDK for Golang

[![CircleCI][circle-ci-badge]][circle-ci-url] [![MIT][mit-badge]][mit-url] [![codecov][codecov-badge]][codecov-url] [![Go Reference][golang-reference-badge]][golang-reference-url]
The **Reloadly SDK for Golang** enables Go developers to easily work with [Reloadly Services][reloadly-main-site] and build scalable solutions. You can get started in minutes if you have Go 1.15+ installed on your machine.


* [SDK Homepage][sdk-website] (Coming soon)
* [Sample Code][sample-code]
* [API Docs][docs-api]
* [Issues][sdk-issues]
* [Giving Feedback](#giving-feedback)
* [Getting Help](#getting-help)

## Getting Started

#### Sign up for Reloadly ####   

Before you begin, you need a Reloadly account. Please see the [Sign Up for Reloadly][reloadly-signup-help] section of
the knowledge-base for information about how to create a Reloadly account and retrieve
your [Reloadly APIs credentials][api-credentials-help].

#### Minimum requirements ####   

**Go 1.15+** installed

#### Installation

`$ go get -u github.com/reloadly/reloadly-sdk-golang`

## Usage

#### Authentication

The Authentication module is implemented based on
the [Authentication API Docs](https://developers.reloadly.com/#authentication-api). This module has a `GetAccessToken`
function which is used to derive an accessToken. The access Token which should be used as a bearer token when making
authenticated requests to the Reloadly API. **This module should be used only when you just need the access token, but
do not intend to use this library for other interactions with Reloadly API**.

```Go  
authClient,err := authentication.NewAuthClient(clientID, clientSecret, false)  
  
if err != nil {  
   fmt.Println(err)  
   return  
}  
  
accessResponse, err := authClient.GetAccessToken()  
  
if err != nil {  
   fmt.Println(err)  
   return  
}  
  
accessToken := accessResponse.AccessToken  
expiresIn := accessResponse.ExpiresIn  
fmt.Println(accessToken,expiresIn)  
```  

## Getting Help

GitHub [issues][sdk-issues] is the preferred channel to interact with our team. Also check these community resources
for getting help:

* Checkout & search our [knowledge-base][reloadly-knowledge-base]
* Talk to us live on our chat tool on our [website][reloadly-main-site] (bottom right)
* Ask a question on [StackOverflow][stack-overflow] and tag it with `reloadly-golang-sdk`
* Articulate your feature request or upvote existing ones on our [Issues][features] page
* Take a look at our [youtube series][youtube-series] for plenty of helpful walkthroughs and tips
* Open a case via with the [Reloadly Support Center][support-center]
* If it turns out that you may have found a bug, please open an [issue][sdk-issues]

## Documentation

Please see the [Godoc](https://pkg.go.dev/github.com/Ghvstcode/reloadly) for the most up-to-date documentation.

## Giving Feedback

We need your help in making this SDK great. Please participate in the community and contribute to this effort by    
submitting issues, participating in discussion forums and submitting pull requests through the following channels:

* Submit [issues][sdk-issues] - this is the preferred channel to interact with our team
* Come join the Reloadly Golang community chat on [Gitter][gitter]
* Articulate your feature request or upvote existing ones on our [Issues][features] page
* Send feedback directly to the team at oss@reloadly.com

## License

This project is licensed under the MIT license. See the [LICENSE](LICENSE) file for more info.

[reloadly-main-site]: https://www.reloadly.com/

[reloadly-signup-help]: https://faq.reloadly.com/en/articles/2307724-how-do-i-register-for-my-free-account

[api-credentials-help]: https://faq.reloadly.com/en/articles/3519543-locating-your-api-credentials

[sdk-issues]: https://github.com/ghvstcode/reloadly/issues

[sdk-license]: http://www.reloadly.com/software/apache2.0/

[gitter]: https://gitter.im/reloadly/reloadly-sdk-golang

[sample-code]: https://github.com/reloadly/reloadly-sdk-golang/blob/main/SAMPLE_CODE.MD

[docs-api]: https://developers.reloadly.com

[features]: https://github.com/reloadly/reloadly-sdk-golang/issues?q=is%3Aopen+is%3Aissue+label%3A%22feature-request%22

[api-docs]: https://developers.reloadly.com

[godoc]: https://pkg.go.dev/github.com/reloadly/reloadly-sdk-golang

[lombok]: https://projectlombok.org

[lombok-plugins]: https://projectlombok.org/setup/overview

[mit-badge]: http://img.shields.io/:license-mit-blue.svg?style=flat

[mit-url]: https://github.com/reloadly/reloadly-sdk-golang/raw/main/LICENSE

[circle-ci-badge]: https://circleci.com/gh/Reloadly/reloadly-sdk-golang.svg?style=svg&circle-token=8f018250b6732bd0be3b183cb09c94942f800b0a

[circle-ci-url]: https://app.circleci.com/pipelines/github/Reloadly/reloadly-sdk-golang

[codecov-badge]: https://codecov.io/gh/Reloadly/reloadly-sdk-golang/branch/main/graph/badge.svg?token=SUV66Q3J2Y

[codecov-url]: https://codecov.io/gh/Reloadly/reloadly-sdk-golang

[youtube-series]: https://www.youtube.com/watch?v=TbXC4Ic8x30&t=141s&ab_channel=Reloadly

[reloadly-knowledge-base]: https://faq.reloadly.com

[stack-overflow]: http://stackoverflow.com/questions/tagged/reloadly-reloadly-sdk

[support-center]: https://faq.reloadly.com/en/articles/3423196-contacting-support

[sdk-website]: https://www.reloadly.com

[golang-reference-badge]: https://pkg.go.dev/badge/github.com/reloadly/reloadly-sdk-golang.svg

[golang-reference-url]: https://pkg.go.dev/github.com/reloadly/reloadly-sdk-golang