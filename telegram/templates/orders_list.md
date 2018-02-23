Active orders in this chat:
{{range .orderList}}
  {{.Name}}
  Order ID: {{.Id}}
  CreatedBy: {{.UserId}}
{{end}}
