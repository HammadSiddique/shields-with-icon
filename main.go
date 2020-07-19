package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	c "github.com/progfay/shields-with-icon/color"
	i "github.com/progfay/shields-with-icon/icon"
)

var (
	white = c.Color{
		Red:   1,
		Green: 1,
		Blue:  1,
		Code:  "FFFFFF",
	}
	black = c.Color{
		Red:   34.0 / 255.0,
		Green: 34.0 / 255.0,
		Blue:  34.0 / 255.0,
		Code:  "222222",
	}
)

func FormatShield(icon i.Icon) (string, error) {
	color, err := c.NewColor(icon.Hex)
	if err != nil {
		return "", err
	}

	var foreground, background c.Color

	if c.GetContrastRatio(white, *color) >= 2.5 {
		foreground = white
		background = *color
	} else {
		foreground = *color
		background = black
	}

	return fmt.Sprintf("![%v](https://img.shields.io/static/v1?style=for-the-badge&message=%s&color=%v&logo=%s&logoColor=%s&label=)",
		icon.Title,
		url.QueryEscape(icon.Title),
		url.QueryEscape(background.Code),
		url.QueryEscape(icon.Title),
		url.QueryEscape(foreground.Code),
	), nil
}

func main() {
	icons, err := i.GetIcons()
	if err != nil {
		log.Panicln(err)
	}

	for _, icon := range icons {
		shield, err := FormatShield(icon)
		if err != nil {
			log.Panicln(err)
		}
		fmt.Fprintln(os.Stdout, shield)
		fmt.Fprintf(os.Stderr, "## %s\n```markdown\n%s\n```\n", shield, shield)
	}
}