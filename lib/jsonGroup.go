package lib

type name struct {
    Name string `json:"name"`
}

type output struct {
    Code string `json:"code"`
    Prod []name  `json:"name"`
    // Orders              int `json:"id"`
    // AmountDiscounted    string `json:"name"`
    // Prod1               string `json:"image"`
    // Prod2               string `json:"image"`
    // Prod3               string `json:"image"`
    // Total               string `json:"image"`
}
