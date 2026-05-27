package main

func contactSection() portfolioSection {
	return portfolioSection{
		key:         "contact",
		title:       "Contact",
		summary:     "Where to reach me when a project needs a builder.",
		description: "A short contact card so the next conversation is one click away.",
		items: []portfolioItem{
			{title: "Email", meta: "Direct", description: "pvighnesh81203@gmail.com", linkText: "Send email", linkURL: "mailto:pvighnesh81203@gmail.com"},
			{title: "GitHub", meta: "Code and experiments", description: "github.com/nonvegetable", linkText: "Open GitHub", linkURL: "https://github.com/nonvegetable"},
			{title: "LinkedIn", meta: "Professional profile", description: "linkedin.com/in/vighnesh-palande", linkText: "Open LinkedIn", linkURL: "https://www.linkedin.com/in/vighnesh-palande/"},
			{title: "Resume", meta: "Full history", description: "A clean summary of internships, projects, and achievements.", linkText: "Open resume", linkURL: "https://drive.google.com/file/d/1Tzh-pc6sZeDywWG6QVOvirKNhLphOgwV/view?usp=sharing"},
		},
	}
}
