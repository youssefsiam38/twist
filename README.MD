# Twist (Simple UI testing tool)

<img align="right" width="209px" src="./logo/twistLogo.png">

Twist is a simple UI testing CLI tool written in Go (Golang). Its main purpose is to help developers write simple and easy read & write YAML-based instructions and assertions to drive the browser in no-time. (I built this tool to test my personal projects. Use it in your own cautious)

This tool is not like selenium or its alternatives, it's less advanced (has only two assertions!). It is built over [chromedp](https://github.com/chromedp/chromedp) which built over DevTools Protocol (can not test firefox).

## Contents

- [Twist](#twist-simple-ui-testing-tool)
  - [Contents](#contents)
  - [Installation](#installation)
  - [Quick start](#quick-start)
  - [The Structure](#the-structure)
  - [Config file options](#config-file-options)
    - [execute](#execute)
    - [timeout](#timeout)
    - [output](#output)
    - [order](#order)
  - [Story file options](#story-file-options)
    - [start](#start)
    - [timeout](#timeout)
    - [headless](#headless)
    - [instructions](#instructions)
  - [instructions](#instructions)
  - [assertions](#assertions)

## Installation

To install Twist package, you need to install Go and set your Go workspace first. (the compiled binaries will be added to this repo later)

1. The first need [Go](https://golang.org/) installed (**version 1.13+ is required**), then you can use the below Go command to install Twist.

```sh
$ git clone https://github.com/youssefsiam38/twist.git
```

2. make it available everywhere in your machine:

```sh
$ cd twist
$ go install
```

## Quick start

```sh
# this will make a twist folder for you in the current directory with config file and stories folder with the initial values
$ twist init
```
```sh
$ twist
```

## The Structure

The tool read from two types of yml files and writes only to the `twist/report`. The `twist/config.yml` file which you can provide configurations about the entire flow of the tests (e.g timeout of the flow and the order of the stories in which twist should follow), and the `twist/stories/{your story name}.story.yml` include configurations, instructions and assertions for each user story in your application

## Config file options

### execute

```yaml
execute: in order # in parallel support will be added soon
```

### timeout

```yaml
timeout: 30m # the timeout of the context of all stories
```

### output

```yaml
output: stdout # file reporting will be added soon
```

### order

```yaml
# The order in which the tests will be executed, the name you put here is the first level extention of the file in the stories folder. e.g blog here will execute {blog}.story.yml
order: 
- blog 
- documentation
```

## Story file options

### start

```yaml
start: https://youssefsiam.me # the url that the story will begin from
```

### timeout

```yaml
timeout: 2m # the timeout of the this specific story
```

### headless

```yaml
headless: false # choose the mode of the user story test (default: false)
```

### instructions

```yaml
# The list of instructions and assertions in order
instructions:
- waitFor: 
    selector: "#about"
- click: "#about"
- asserPathIs: /about
```
## instructions

### **click**

```yaml
instructions:
- click: 
    selector: "#about"
```
or, since *selector* is the only paramater
```yaml
instructions:
- click: "#about"
```

### **doubleClick**

```yaml
instructions:
- doubleClick: 
	selector: "#home"
```
or
```yaml
instructions:
- doubleClick: "#home"
```

### **rightClick**

```yaml
instructions:
- rightClick: 
	selector: "#home"
```
or
```yaml
instructions:
- rightClick: "#home"
```

### **navigate**

```yaml
instructions:
- navigate: "https://youssefsiam.me"
- assertPathIs: "https://youssefsiam.me/"
```

### **refresh**

```yaml
instructions:
- do: "refresh"
```

### **saveScreenshotTo**

```yaml
instructions:
# give it the name of the file that you like to save to in the twist/report folder
- saveScreenshotTo: about 
```

### **submit**

```yaml
instructions:
- submit: 
	selector: "#home"
```
or
```yaml
instructions:
- submit: "#home"
```

### **type**

```yaml
instructions:
- type:
    selector: "input[type=\"email\"]"
    text: example@example.com
- submit: "input[type=\"email\"]"
```

### **waitFor**

```yaml
instructions:
- waitFor: 
    selector: "#about"
```
or, since *selector* is the only paramater
```yaml
instructions:
- waitFor: "#about"
```

### **waitUntilMissing**

```yaml
instructions:
- waitUntilMissing: 
    selector: "#about"
```
or, since *selector* is the only paramater
```yaml
instructions:
- waitUntilMissing: "#about"
```

## assertions

### **assertPathIs**

```yaml
instructions:
- assertPathis: 
	name: "Blog home page assertion"
    expect: "https://youssefsiam.me"
```
or
```yaml
instructions:
- assertPathis: "https://youssefsiam.me"
```
this way the assertion name will be assertPathIs

### **assertText**

```yaml
instructions:
- assertText: 
	name: "The text is `about` assertion"
    expect: "about"
	selector: "button#about"
```