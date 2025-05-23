package dev

config: config & {
    database: {
        mongo01: {
            kind: "database",
            type: "mongodb",
            host: "mongo01.example.com",
            port: 27017,
            credentials: {
                type: "file",
                name: ".env",
                env: "prod",
            },
        },
        mongo02: {
            kind: "database",
            type: "mongodb",
            host: "mongo02.example.com",
            port: 27017,
            credentials: {
                type: "file",
                name: ".env",
                env: "prod",
            },
        },
        postgres01: {
            kind: "database",
            type: "postgres",
            host: "postgres01.example.com",
            port: 5432,
            credentials: {
                type: "env",
                name: "DB_SECRET",
                env: "staging",
            },
        },
    },

    artifactory: {
        artifactory01: {
            kind: "artifactory",
            type: "jfrog",
            url: "https://artifactory.example.com",
            repo: "repo01",
            auth: {
                user: "admin",
                pass: "password123",
            },
        },
    },

    ci_tools: {
        jenkins01: {
            kind: "ci-tools",
            type: "jenkins",
            endpoint: "https://jenkins.example.com",
            auth: {
                method: "token",
                token: "jenkins-secret-token",
            },
        },
        github_actions: {
            kind: "ci-tools",
            type: "gh-actions",
            endpoint: "https://github.com/actions",
            auth: {
                method: "basic",
                user: "github-user",
                pass: "secure-password",
            },
        },
    },
}