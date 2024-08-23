package main

import "bookutil/author"

func main() {
	// Create an Author instance.
	authorInstance := author.NewAuthor("Jane Doe", "jane@example.com")

	// Write and review a Chapter
	chapterTitle := "Introduction to Go modules"
	chapterContent := "Go modules provide a structured way to manage dependencies and improve code maintainability."

	authorInstance.WriteChapter(chapterTitle, chapterContent)
	authorInstance.ReviewChapter(chapterTitle, "This chapter looks great, but let's add some more examples.")
	authorInstance.FinalizeChapter(chapterTitle)
}
