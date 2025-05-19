type PaymentEvent struct {
    PaymentID   string     `json:"payment_id"`
    UserID      int        `json:"user_id"`
    Amount      float64    `json:"amount"`
    Status      string     `json:"status"`
    Timestamp   string     `json:"timestamp"`
    MethodType  string     `json:"method_type"`
}
