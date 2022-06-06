# **Go API**

<p align="center">
    <img src="https://i.morioh.com/201003/aa184196.webp" alt="GO" />
</p>

---

## How To Use ⬇️

- First you need to install the required dependencies (You will find it in main.go)
- Create .env file and add this line to connect to your mongodb atlas

```bash
MONGOURI = "Your Mongo database connect URI"
```

- Change the collection name to the collection name you set it on atlas, in the `configs/setup.go` at GetCollection function.
- Server is running on port 6000, you can change the port at `main.go`.

Finally use the command below to launch the server in the terminal and you ready to go

```bash
go run main.go
```

---

## Restful routes:

```go
func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)       // Create a new user
	app.Get("/user/:id", controllers.GetUser)       // Get a user by id
	app.Get("/user", controllers.GetAllUsers)       // Get all users
	app.Put("/user/:id", controllers.EditUser)      // Edit a user
	app.Delete("/user/:id", controllers.DeleteUser) // Delete a user
	app.Delete("/users", controllers.PurgeUsers)    // Delete all users
}
```
## User scheme:

```go
type User struct {
	Id        primitive.ObjectID `json:"id omitempty"`
	Name      string             `json:"name,omitempty,validate:"required""`
	Age       int                `json:"age",omitempty,validate:"required"`
	Email     string             `json:"email", omitempty, validate:"required"`
	Gender    string             `json:"gender",omitempty, default:"x"`
	Location  Location           `json:"location", omitempty, default:"Empty"`
	Title     string             `json:"title", omitempty, validate:"required"`
	CreatedAt time.Time          `json:"created_at", omitempty`
	UpdatedAt time.Time          `json:"updated_at", omitempty`
}

type Location struct {
	Street    string    `json:"street,omitempty"`
	City      string    `json:"city,omitempty"`
	State     string    `json:"state,omitempty"`
	VisitedAt time.Time `json:"visitedAt,omitempty"`
}
```

---
![Go_8001611039611515](https://user-images.githubusercontent.com/96744413/172076517-71e0a3c4-7f41-496b-b42e-6ac6ccac45ad.gif)



