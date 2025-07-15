# Mangaz scraper in Go (grj1234 fork)

## Overview
This project is an open-source tool developed in Golang for extracting mangaz data.  
This repository is a fork of [johnbalvin/mangaz](https://github.com/johnbalvin/mangaz).

## Modified portions
* Save only one book specified as argument, instead of trying to scrape all the books
* Single-thread main routine instead of multi-thread one
	* Note that image processing routine is still multi-thread
* Save image files as PNG instead of JPG (to avoid degradation of image quality)
* The log message is modified

## Install
You need to have Golang install:

[golang](https://go.dev/dl)

```bash
 git clone https://github.com/grj1234/mangaz-single.git
```

## Run

```bash
 go run . <manga id>
```

## Output
by default images will be saved into result_imgs folder and metadata on result_meta folder  
You can change the folder and other values in the file start.go
