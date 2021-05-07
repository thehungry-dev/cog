package main

import "github.com/thehungry-dev/log"

func main() {
	log.
		Tags("foo", "bar").
		Data(
			log.String("name", "aName"),
			log.String("lastName", "aLastName"),
		).
		Info("This is the body of a message")

	log.Error("Some error")

	log.
		Data(
			log.Int("aNumber", 123),
		).
		Fatald()

	log.
		TagsWriter.
		Tags("foo", "bar").
		Data(
			log.String("name", "aName"),
			log.String("lastName", "aLastName"),
		).
		Info("This is the body of a message with tags")

	log.
		JSONWriter.
		Tags("foo", "bar").
		Data(
			log.String("name", "aName"),
			log.String("lastName", "aLastName"),
			log.Int("age", 123),
		).
		Info("This is the body of a message with tags")
}
