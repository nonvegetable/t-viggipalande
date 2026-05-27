package main

func skillsSection() portfolioSection {
	return portfolioSection{
		key:         "skills",
		title:       "Skills",
		summary:     "Languages, frameworks, data stores, and tools I use to ship.",
		description: "A practical stack with enough range to move from interface work to backend logic and integrations.",
		items: []portfolioItem{
			{title: "Programming Languages", description: "Java, Python, JavaScript, TypeScript, C, HTML, CSS, PL/SQL. The regular set."},
			{title: "Frontend Frameworks", description: "React.js for clean interfaces and predictable component structure."},
			{title: "Backend Technologies", description: "Node.js, Express.js, Flask for APIs, services, and backend logic that needs to hold up."},
			{title: "Database Systems", description: "MySQL, MongoDB, PostgreSQL, Oracle. Comfortable with SQL and data flow."},
			{title: "Development Tools", description: "Git, GitHub, VS Code, PyCharm, Docker."},
			{title: "Integration & APIs", description: "REST APIs, Razorpay, Google APIs, Supabase, and third-party integrations."},
		},
	}
}
