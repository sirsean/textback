
```
docker run --name redis -d redis
```

```
docker build -t textback .
```

```
docker run -p 2345:80 --link redis:redis -e TWILIO_ACCOUNT_SID="xxx" -e TWILIO_AUTH_TOKEN="xxx" -e TWILIO_FROM="xxx" --name textback textback
```
