package domain

type Notification struct {
    Message string `json:"message"`
    Data    interface{} `json:"data"`
}

func NewNotification(message string, data interface{}) *Notification {
    return &Notification{
        Message: message,
        Data:    data,
    }
}
