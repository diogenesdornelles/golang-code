package signin


type User struct {
    Email      string `validate:"required,email"`
    Password   string `validate:"gte=3"`
}