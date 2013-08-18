package main

import "github.com/mahan/golorize"

func main() {
  println(golorize.Ansify("I #green[heard you liked #blink[tagging], so I added #yellow[tags to #bold[tags]]] so you can #magenta[tag while you #blue[tag]]"))
  golorize.Enabled = false
  println(golorize.Ansify("I #green[heard you liked #blink[tagging], so I added #yellow[tags to #bold[tags]]] so you can #magenta[tag while you #blue[tag]]"))
  golorize.Enabled = true
  println(golorize.Ansify("I #green[heard you liked #blink[tagging], so I added #yellow[tags to #bold[tags]]] so you can #magenta[tag while you #blue[tag]]]]]]]]]]"))
}