# go-org

[![Go Report Card](https://goreportcard.com/badge/github.com/manfromth3m0oN/go-org)](https://goreportcard.com/report/github.com/manfromth3m0oN/go-org)

A text editor agnostic org style organisational tool 

---

## The backstory

I've always been a vim user but I really liked the look of org mode for emacs.
So, much to that end, I decided to give emacs a try. Long story short: it isn't for me.

Well, what now?

That's where the idea came up to set up a distributed note taking system that is agnostic of your frontend.
In theory, you could write a mobile app that works with this system. But 'system' is a very vague term, 
what exactly what are we making?

## How to work with this tool

Principally, one starts with a TODO.txt and branches out from there.
You wouldn't write the body of your todo items in the todo file, those are their own
text file which the daemon will create as the todo is added. 

Once signed in, integration with google calendar will add the relevant events to your todo for the current day.
Both yesterday and tomorrows sections will be populated also (maybe idk).

### Sections

Sections are denoted with fields wrapped in ==. Known sections are:
* `==TODAY==` - Todos to be done today, becomes yesterday is then archived
* `==PROJECTS==` - Where todos for one's projects live, they can be given dates with the `[[]]`

### Referencing

You can reference a todo item with the syntax `%<project name>.<heading level>.<subheading level>`

## The plan

So the plan is to make a note taking, outlining and writing system that combines the best features from
services like google docs and keep with the power and versatility of org-mode

### Local files

Locally each user will have a directory in which they write and take notes, the exact directory structure
is still up for debate, and those files will contain some markdown like syntax based around the org-mode
asterisk (*) notation. E.g:

```
* Heading 1
	* Subheading 1.1
	Some information on the subheading
		* Subsubheading 1.1.1
		The main body of the subsubheading
		* Another subsubheading 1.1.2
	* Subheading 1.2
	Another subheadings content
* Heading 2
A second heading
```

Parsing will accomplish:
1. Creating & verifying the file structure
2. creating desired markup
3. include code snippets either locally or from github

### Sharing and serving

Users can run servers to serve thier notes on the web or to access and edit
on other machines or even to work as a team. This may implement git so will
probably a late feature.

## TODO

1. Write the parser for the document format
2. Make the local daemon. 
