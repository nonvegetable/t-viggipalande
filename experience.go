package main

func experienceSection() portfolioSection {
	return portfolioSection{
		key:         "experience",
		title:       "Experience",
		summary:     "Internships and production work that sharpened my backend and frontend instincts.",
		description: "A mix of migration work, API validation, ERP integration, frontend flows, and backend logic.",
		items: []portfolioItem{
			{title: "Technical Consultant Intern • Axe Finance", meta: "Jan 2026 - Present", description: "Worked on enterprise migration flows, API validation, ERP integration, and documentation. Kept legacy systems moving without drama."},
			{title: "Software Developer Intern • Hindalco Industries", meta: "May 2025 - Jul 2025", description: "Built Oracle PL/SQL procedures and supported blockchain integration. Tightened backend logic and data handling."},
			{title: "Frontend Developer Intern • KJSCE", meta: "Jul 2024 - Dec 2024", description: "Built the Minors-Honors and Open-Elective portals. Improved navigation, reduced friction, and shipped cleaner flows."},
			{title: "Digital Marketing Intern • KJSCE", meta: "Jun 2023 - Jul 2024", description: "Created 120+ posts across social platforms. Learned consistency, iteration, and how small improvements add up."},
		},
	}
}
