# **Go API**

<p align="center">
    <img src="https://i.morioh.com/201003/aa184196.webp" alt="GO" />
</p>

---

## How To Use ⬇️

- First you need to install the required dependencies (You will find it in main.go)
- Create .env file and add this line to connect to your mongodb atlas
- Change the collection name to the collection you set it, in the configs/setup.go at GetCollection function.

```bash
MONGOURI = "Your Mongo database connect URI"
```

Finally use the command below to launch the server in the terminal and you ready to go

```bash
go run main.go
```

---

## Restful routes:

![41241](https://user-images.githubusercontent.com/96744413/171936617-9cf51561-3614-4e8a-992e-0789da00d416.png)

## User scheme:

![25335](https://user-images.githubusercontent.com/96744413/171936646-6d79c0f6-b108-43a9-ad3a-7d5f6d453862.png)

---
