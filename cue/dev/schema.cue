package dev

// Define the configuration for different tools and environments
config: {
    database: [string]: #Database,
    artifactory: [string]: #Artifactory,
    ci_tools: [string]: #CITools,
}

// Database Schema
#Database: {
    kind:      "database",
    type:      "mongodb" | "postgres" | "oracle",  // Limit to these database types
    host:      string & =~"^.*\\.example\\.com$",  // Host should match example.com domain
    port:      int & >=0 & <=65535,               // Port should be valid range (0-65535)
    credentials: {
        type:  "file" | "env",                   // Limit to these credential types
        name:  string,
        env:   "prod" | "staging" | "dev",       // Restrict valid environments
    },
}

// Artifactory Schema
#Artifactory: {
    kind: "artifactory",
    url:  string & =~"^https://.*$",              // URL should be HTTPS
    repo: string,
    auth?: {
        token?: string,
        user?: string,
        pass?: string,
    },
}

// CI/CD Tools Schema
#CITools: {
    kind:      "ci-tools",
    type:      "jenkins" | "gh-actions",  // Limit to these CI tools
    endpoint:  string & =~"^https://.*$",              // Endpoint should be HTTPS
    auth?: {
        method: "token" | "basic",                     // Limit to "token" or "basic" auth
        token?: string,
        user?: string,
        pass?: string,
    },
}