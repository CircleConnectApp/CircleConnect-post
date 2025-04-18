# CircleConnect Post Microservice

Part of the CircleConnect social networking platform. This microservice handles posts, comments, and likes functionality.

## Technology Stack

- Go (1.21+)
- Gin Web Framework
- MongoDB

## Features

- Create and view posts
- Post comments
- Like/unlike posts
- Filter posts by community or user

## Getting Started

### Prerequisites

- Go 1.21 or higher
- MongoDB

### Installation

1. Clone the repository
2. Navigate to the post-service directory
3. Copy the environment example file
   ```
   cp .env.example .env
   ```
4. Update the environment variables in `.env` file with your configuration
5. Run the service
   ```
   go run main.go
   ```

## API Endpoints

### Public Endpoints

- `GET /api/posts` - Get all posts
- `GET /api/posts/:id` - Get a specific post by ID
- `GET /api/communities/:community_id/posts` - Get all posts in a community
- `GET /api/posts/:post_id/comments` - Get all comments for a post

### Authenticated Endpoints

- `POST /api/posts` - Create a new post
- `GET /api/user/posts` - Get posts created by the authenticated user
- `POST /api/posts/:post_id/comments` - Add a comment to a post
- `POST /api/posts/:post_id/like` - Like a post
- `DELETE /api/posts/:post_id/like` - Unlike a post

## Authentication

This service uses JWT tokens for authentication. Include a valid JWT token in the Authorization header:

```
Authorization: Bearer your_jwt_token
```

## License

This project is part of CircleConnect, developed for SW Architecture and Design. 

# Post Microservice API Endpoints

## Post Endpoints

### Create Post

`POST /api/posts`

**Description**: Create a new post.

**Headers**:  
`Authorization: Bearer <token>`

**Request Body**:
```json
{
  "title": "string",
  "content": "string",
  "community_id": "number (optional)",
  "media_urls": ["string"] (optional),
  "tags": ["string"] (optional)
}
```

**Responses**:
* `201 Created`: Returns the created post
* `400 Bad Request`: If the request body is invalid
* `401 Unauthorized`: If the user is not authenticated

---

### Get All Posts

`GET /api/posts`

**Description**: Get all posts.

**Responses**:
* `200 OK`: Returns a list of posts

---

### Get Post by ID

`GET /api/posts/:id`

**Description**: Get a post by its ID.

**Responses**:
* `200 OK`: Returns the post
* `404 Not Found`: If the post is not found

---

### Get Posts by Community

`GET /api/communities/:community_id/posts`

**Description**: Get all posts in a specific community.

**Responses**:
* `200 OK`: Returns a list of posts
* `400 Bad Request`: If the community ID is invalid

---

### Get Posts by User

`GET /api/user/posts`

**Description**: Get all posts created by the authenticated user.

**Headers**:  
`Authorization: Bearer <token>`

**Responses**:
* `200 OK`: Returns a list of posts
* `401 Unauthorized`: If the user is not authenticated 