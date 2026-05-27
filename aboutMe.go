package main

func aboutSection() portfolioSection {
	return portfolioSection{
		key:         "about",
		title:       "About Me",
		summary:     "How I build, think, and keep improving.",
		description: "Engineering background, internships, and shipping real work instead of pretending specs are enough.",
		items: []portfolioItem{
			{title: "Engineering Background", description: "Built through coursework, internships, and shipped projects with real users in mind."},
			{title: "Always Learning", description: "Learns by building. Try, break, fix, improve, repeat."},
			{title: "Football Mindset", description: "Football taught discipline, pressure handling, and staying calm when the plan changes mid-match."},
			{title: "Curious by Default", description: "Interested in geography, systems, and the patterns that show up before they become obvious."},
			{title: "Pattern Thinking", description: "Likes finding structure in messy problems. Usually before coffee. Not always by much."},
			{title: "Built to Improve", description: "Not trying to be loud. Just better than the last version."},
		},
	}
}
