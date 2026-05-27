package main

func hackathonSection() portfolioSection {
	return portfolioSection{
		key:         "hackathons",
		title:       "Hackathons",
		summary:     "Fast builds, short deadlines, and a lot of iteration.",
		description: "Evidence that I can ship under pressure and still keep the details intact.",
		items: []portfolioItem{
			{title: "Codenovate 2024 • Hyderabad", meta: "Best in Category", description: "24-hour AI healthcare build."},
			{title: "Navonmesh 2.0 • Raipur", meta: "Top 30 out of 500+", description: "Cleaner healthcare platform rework."},
			{title: "Parul-Hackverse • Vadodara", meta: "Finalist", description: "IoT water management system with sensor integration."},
		},
	}
}

func otherSection() portfolioSection {
	return portfolioSection{
		key:         "other",
		title:       "Other",
		summary:     "A few human details and off-hours projects.",
		description: "Interests and side quests that keep the work grounded.",
		items: []portfolioItem{
			{title: "Open to Collaboration", description: "If the problem is interesting, I am usually in."},
			{title: "MFA League Winner • 2019", description: "Football taught resilience, teamwork, and how to keep moving after a bad first half."},
			{title: "Self-Hosted Systems", description: "Set up a personal NAS with a Raspberry Pi and a 1TB drive."},
			{title: "Guitar Enthusiast", description: "Guitar is where consistency gets tested without a compiler yelling back."},
			{title: "Geography Buff", description: "Can point out every country in Europe. Curiosity has a few practical uses."},
			{title: "Fast Execution", description: "Built multiple full-stack apps in 24 to 48 hour hackathons."},
		},
	}
}
