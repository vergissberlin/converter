---
title: Gonverter
description: A simple web service to convert between different file formats
tags:
  + json
  + golang
  + protobuf
---

# Gonverter

This simple mirco service is used to convert data formats.

## ✨ Features

* json to protobuf

## 💁‍♀️ How to use

* Connect to your Railway project `railway link`
* Start the development server `railway run go run app/api/*.go`

## 📝 Notes

Send a request to the service with the following format:

```bash
curl --location 'http://YOURDOMAIN.railway.com/json2protobuf' \ # replace with your railway url
--header 'Content-Type: application/json' \
--data './tests/fixtures/person.json'
```

## 🛠️ Development

### Run the service

```bash
docker build -t gonverter .
docker run -d -p 8080:8080 gonverter
```

### Send a request

```bash
curl --location 'http://0.0.0.0:8080/json2protobuf' \
--header 'Content-Type: application/json' \
--data './tests/fixtures/person.json'
```

#### Response

 `Content-Type: application/octet-stream`

 `Body: test test@example.com`

## Know how

* Read <https://seb-nyberg.medium.com/customizing-protobuf-json-serialization-in-golang-6c58b5890356>
