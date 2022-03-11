package main

import (
  "testing"
)

func TestCreateSlugFromTitle(t *testing.T) {
  results := map[string]string {
    "My good days with life": "my-good-days-with-life",
    "[4] Making a compiler": "4-making-a-compiler",
    "$life is shit$": "life-is-shit",
    "XoXo baba is dead": "xoxo-baba-is-dead",
  }

  for input, expectedOutput := range results {
    output := CreateSlugFromTitle(input)
    if output == expectedOutput {
      continue
    }
    t.Errorf("For input='%s', expected output='%s', but recieved='%s'", input, expectedOutput, output)
  }
}
