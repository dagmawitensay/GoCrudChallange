package Domain


type Person struct {
    Id      string
    Name    string `json:"name" validate:"required,min=5,max=10"`
    Age     int    `json:"age" validate:"required,min=1"`
    Hobbies []string `json:"hobbies"`
}