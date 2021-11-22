package message

func BuildTextMessagePayload(text string) map[string]interface{} {
	return map[string]interface{}{
		"msg_type": "text",
		"content": map[string]interface{}{
			"text": text,
		},
	}
}
