package discord

type Webhook struct {
	Username  string  `json:"username,omitempty"`
	AvatarURL string  `json:"avatar_url,omitempty"`
	Content   string  `json:"content,omitempty"`
	Embeds    []Embed `json:"embeds,omitempty"`
	TTS       bool    `json:"tts,omitempty"`
}

type Embed struct {
	Author      Author  `json:"author,omitempty"`
	Title       string  `json:"title,omitempty"`
	URL         string  `json:"url,omitempty"`
	Description string  `json:"description,omitempty"`
	Color       int64   `json:"color,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
	Thumbnail   Image   `json:"thumbnail,omitempty"`
	Image       Image   `json:"image,omitempty"`
	Footer      Footer  `json:"footer,omitempty"`
}

type Author struct {
	Name    string `json:"name,omitempty"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type Field struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline"`
}

type Footer struct {
	Text    string `json:"text,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type Image struct {
	URL string `json:"url,omitempty"`
}

type webhookBuilder struct {
	webhook Webhook
}

func (wb *webhookBuilder) Username(username string) *webhookBuilder {
	wb.webhook.Username = username
	return wb
}

func (wb *webhookBuilder) AvatarURL(avatarUrl string) *webhookBuilder {
	wb.webhook.AvatarURL = avatarUrl
	return wb
}

func (wb *webhookBuilder) Content(content string) *webhookBuilder {
	wb.webhook.Content = content
	return wb
}

func (wb *webhookBuilder) Embed(embed Embed) *webhookBuilder {
	wb.webhook.Embeds = append(wb.webhook.Embeds, embed)
	return wb
}

func (wb *webhookBuilder) Webhook() Webhook {
	return wb.webhook
}

func WebhookBuilder() *webhookBuilder {
	return &webhookBuilder{webhook: Webhook{}}
}

type embedBuilder struct {
	embed Embed
}

func (eb *embedBuilder) Author(author Author) *embedBuilder {
	eb.embed.Author = author
	return eb
}

func (eb *embedBuilder) Title(title string) *embedBuilder {
	eb.embed.Title = title
	return eb
}

func (eb *embedBuilder) URL(url string) *embedBuilder {
	eb.embed.URL = url
	return eb
}

func (eb *embedBuilder) Description(description string) *embedBuilder {
	eb.embed.Description = description
	return eb
}

func (eb *embedBuilder) Color(color int64) *embedBuilder {
	eb.embed.Color = color
	return eb
}

func (eb *embedBuilder) Field(field Field) *embedBuilder {
	eb.embed.Fields = append(eb.embed.Fields, field)
	return eb
}

func (eb *embedBuilder) Thumbnail(thumbnail Image) *embedBuilder {
	eb.embed.Thumbnail = thumbnail
	return eb
}

func (eb *embedBuilder) Image(image Image) *embedBuilder {
	eb.embed.Image = image
	return eb
}

func (eb *embedBuilder) Footer(content []string) *embedBuilder {
	if len(content) == 1 {
		eb.embed.Footer = Footer{Text: content[0]}
	} else if len(content) > 1 {
		eb.embed.Footer = Footer{Text: content[0], IconURL: content[1]}
	}
	return eb
}

func (eb *embedBuilder) Embed() Embed {
	return eb.embed
}

func EmbedBuilder() *embedBuilder {
	return &embedBuilder{embed: Embed{}}
}
