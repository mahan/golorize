/*
  Go adaptation of Colorize lib for Node.JS: https://github.com/mattpat/colorize
*/

package golorize

import (
  "regexp"
  "strings"
)

var ansicodes map[string]string = map[string]string{
  "reset": "\033[0m",
  "bold": "\033[1m",
  "italic": "\033[3m",
  "underline": "\033[4m",
  "blink": "\033[5m",
  "black": "\033[30m",
  "red": "\033[31m",
  "green": "\033[32m",
  "yellow": "\033[33m",
  "blue": "\033[34m",
  "magenta": "\033[35m",
  "cyan": "\033[36m",
  "white": "\033[37m",
}

var tag *regexp.Regexp
var Enabled bool = true

func Ansify(str string) string {
  cstack := make([]string, 3)

  for tag.MatchString(str) {
    matches := tag.FindStringSubmatch(str)
    orig := matches[0]

    if Enabled {
      if orig == "]" {
        cstack = cstack[:len(cstack)-2]
      } else {
        name := matches[1]
        if code, exists := ansicodes[name]; exists {
          cstack = append(cstack, code)
        }

      }
      //str = str.replace(orig, this.ansicodes.reset + cstack.join(''));
      str = strings.Replace(str, orig, ansicodes["reset"] + strings.Join(cstack, ""), -1)
    } else {
      str = strings.Replace(str, orig, "", -1)
    }
  } //for
  return str
}

func init() {
  tag = regexp.MustCompile(`#([a-z]+)\[|\]`)
}