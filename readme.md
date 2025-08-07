# Let's Go follow along (Snippetbox)

Welcome to the practice repository for the book **[Let's Go](https://lets-go.alexedwards.net)** by **Alex Edwards**. This book is a practical introduction to building modern, secure web applications with Go (Golang).

**Let's Go** is aimed at developers who are familiar with the basics of Go and want to learn how to build robust web applications. It covers:

- HTTP servers
- Routing
- Middleware
- Working with HTML templates
- Sessions and cookies
- User authentication
- Using MySQL with Go
- Secure password handling
- Testing

> The book emphasizes writing clean, idiomatic, and secure Go code, and uses a hands-on, project-based approach to teach you web development.

# 📚 Snippetbox – Route Table

This table lists all HTTP routes used in the "Let's Go" web application.

| Method | Path               | Description                        | Access        | Working |
| ------ | ------------------ | ---------------------------------- | ------------- | ------- |
| GET    | /                  | Home page (latest snippets)        | Public        | ✅      |
| GET    | /snippet/view/:id  | View a single snippet by ID        | Public        | ✅      |
| GET    | /snippet/create    | Show snippet creation form         | Authenticated | ✅      |
| POST   | /snippet/create    | Submit new snippet                 | Authenticated | ✅      |
| GET    | /user/signup       | Show user signup form              | Public        | ✅      |
| POST   | /user/signup       | Process user signup                | Public        | ✅      |
| GET    | /user/login        | Show login form                    | Public        | ✅      |
| POST   | /user/login        | Process login                      | Public        | ✅      |
| POST   | /user/logout       | Log out the user                   | Authenticated | ✅      |
| GET    | /ping              | Healthcheck endpoint               | Internal/Test | ❌      |
| GET    | /static/\*filepath | Serve static files (CSS, JS, etc.) | Public        | ✅      |

> Note:
>
> - `:id` is a placeholder for the snippet ID (e.g., `/snippet/view/1`)
> - Authenticated routes require a user to be logged in
> - Static files are typically served under `/static/`
