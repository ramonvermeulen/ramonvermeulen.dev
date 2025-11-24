package data

import (
	"sort"
	"time"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
)

// techRegistry central repository of technologies and their URLs.
var techRegistry = map[string]string{
	"Go":                  "https://go.dev/",
	"Python":              "https://python.org/",
	"GCP":                 "https://cloud.google.com/",
	"AWS":                 "https://aws.amazon.com/",
	"Terraform":           "https://www.terraform.io/",
	"Kubernetes":          "https://kubernetes.io/",
	"Docker":              "https://www.docker.com/",
	"CI/CD":               "https://en.wikipedia.org/wiki/CI/CD",
	"React":               "https://react.dev/",
	"BigQuery":            "https://cloud.google.com/bigquery",
	"Apache Airflow":      "https://airflow.apache.org/",
	"dbt":                 "https://www.getdbt.com/",
	"PostgreSQL":          "https://www.postgresql.org/",
	"REST APIs":           "https://en.wikipedia.org/wiki/Representational_state_transfer",
	"React Native":        "https://reactnative.dev/",
	"TypeScript":          "https://www.typescriptlang.org/",
	"Redux":               "https://redux.js.org/",
	"iOS":                 "https://developer.apple.com/ios/",
	"Android":             "https://developer.android.com/",
	"Web Development":     "https://en.wikipedia.org/wiki/Webdevelopment",
	"Bash/Shell":          "https://en.wikipedia.org/wiki/Shell_script",
	"Jenkins":             "https://www.jenkins.io/",
	"Pub/Sub":             "https://cloud.google.com/pubsub",
	"MongoDB":             "https://www.mongodb.com/",
	"Linux":               "https://www.linux.org/",
	"Airbyte":             "https://airbyte.com/",
	"Apache Spark":        "https://spark.apache.org/",
	"SQL":                 "https://en.wikipedia.org/wiki/SQL",
	"Looker":              "https://looker.com/",
	"IaC":                 "https://en.wikipedia.org/wiki/Infrastructure_as_code",
	"Network Engineering": "https://en.wikipedia.org/wiki/Computer_networking",
	"Github Actions":      "https://github.com/features/actions",
	"gRPC":                "https://grpc.io/",
	"Grafana":             "https://grafana.com/",
}

// Tech returns a Technology struct for a registered name.
func Tech(name string) models.Technology {
	return models.Technology{Name: name, URL: techRegistry[name]}
}

// sortTechs alphabetically by Name (in-place).
func sortTechs(techs []models.Technology) []models.Technology {
	sort.Slice(techs, func(i, j int) bool { return techs[i].Name < techs[j].Name })
	return techs
}

func ptr(t time.Time) *time.Time { return &t }

// Positions is an immutable list of career positions used for the experience page.
var Positions = []*models.Position{
	{
		CompanyName:    "Xebia",
		CompanyWebsite: "https://xebia.com/",
		StartDate:      time.Date(2025, time.November, 1, 0, 0, 0, 0, time.UTC),
		EndDate:        nil,
		Title:          "Software & Cloud Engineer",
		Location:       "Hilversum, the Netherlands",
		Description:    "I work at the intersection of Software and Cloud Engineering, helping clients solve complex challenges using a variety of technologies. Whether it's building custom tooling, designing cloud-native platforms, or optimizing existing systems, I focus on creating solutions that are both technically sound and practical to operate. My background in data engineering gives me a unique perspective on building reliable, data-intensive systems. I often enjoy projects that involve Software Engineering, Cloud Infrastructure, and DevOps practices to deliver high-quality solutions.",
		Technologies: sortTechs([]models.Technology{
			Tech("Go"), Tech("Python"), Tech("GCP"), Tech("Terraform"), Tech("Kubernetes"), Tech("Docker"), Tech("CI/CD"), Tech("REST APIs"), Tech("BigQuery"), Tech("Bash/Shell"), Tech("Pub/Sub"), Tech("Linux"), Tech("SQL"), Tech("IaC"), Tech("Network Engineering"), Tech("Github Actions"), Tech("gRPC"),
		}),
	},
	{
		CompanyName:    "Xebia",
		CompanyWebsite: "https://xebia.com/",
		StartDate:      time.Date(2021, time.December, 1, 0, 0, 0, 0, time.UTC),
		EndDate:        ptr(time.Date(2025, time.November, 1, 0, 0, 0, 0, time.UTC)),
		Title:          "Data Engineer",
		Location:       "Amsterdam, the Netherlands",
		Description:    "As a Data Engineer I worked on full-time project basis for various clients (Enza Zaden, ING, ASML, Rituals). My work varied from building reliable data pipelines to designing and implementing whole analytical data platforms on GCP. I collaborated closely with cross-functional teams to ensure data quality, scalability, and performance, enabling clients to make data-driven decisions. My role often involved bringing Software Engineering best practices to Data Engineering, including Infrastructure as Code (IaC), CI/CD, and automated testing.",
		Technologies: sortTechs([]models.Technology{
			Tech("Python"), Tech("GCP"), Tech("Terraform"), Tech("Kubernetes"), Tech("Docker"), Tech("CI/CD"), Tech("BigQuery"), Tech("Airflow"), Tech("dbt"), Tech("PostgreSQL"), Tech("Bash/Shell"), Tech("Linux"), Tech("Airbyte"), Tech("Apache Spark"), Tech("SQL"), Tech("Looker"), Tech("IaC"), Tech("Github Actions"),
		}),
	},
	{
		CompanyName:    "ShoppingMinds",
		CompanyWebsite: "https://shoppingminds.com/",
		StartDate:      time.Date(2020, time.November, 1, 0, 0, 0, 0, time.UTC),
		EndDate:        ptr(time.Date(2021, time.November, 1, 0, 0, 0, 0, time.UTC)),
		Title:          "Software Engineer",
		Location:       "Utrecht, the Netherlands",
		Description:    "Returning after a successful internship, I worked on the data platform that powered personalized recommendations for e-commerce and travel clients. I developed and maintained core features for the customer-facing data management platform, with a focus on Python back-end services. This included designing and building robust connectors to external APIs to automate marketing workflows and enable real-time audience data synchronization.",
		Technologies: sortTechs([]models.Technology{
			Tech("Python"), Tech("GCP"), Tech("Kubernetes"), Tech("Docker"), Tech("REST APIs"), Tech("BigQuery"), Tech("PostgreSQL"), Tech("Web Development"), Tech("Bash/Shell"), Tech("Jenkins"), Tech("Pub/Sub"), Tech("MongoDB"), Tech("Linux"), Tech("Grafana"),
		}),
	},
	{
		CompanyName:    "Coffee IT",
		CompanyWebsite: "https://coffeeit.nl/",
		StartDate:      time.Date(2019, time.February, 1, 0, 0, 0, 0, time.UTC),
		EndDate:        ptr(time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC)),
		Title:          "Software Engineer (Mobile Apps)",
		Location:       "Utrecht, the Netherlands",
		Description:    "I started my professional career in mobile development, shipping React Native apps to both the iOS and Android App Store. Learned the importance of user experience, performance optimization, standardized API interfaces, and the complete software delivery lifecycle. Worked on apps for clients such as: Knaek, WieBetaaltWat, ROC Midden Nederland, and ExamenOverzicht.",
		Technologies: sortTechs([]models.Technology{
			Tech("Docker"), Tech("CI/CD"), Tech("React"), Tech("TypeScript"), Tech("Redux"), Tech("REST APIs"), Tech("React Native"), Tech("iOS"), Tech("Android"), Tech("Web Development"), Tech("Bash/Shell"), Tech("Linux"),
		}),
	},
}
