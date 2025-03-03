# Forum Web Application

A modern web forum built with Go that enables user communication through posts, comments, and reactions.

## Project Objectives
Registered enable user login using github or gmail

## Features

- **User Authentication**
  - Secure registration and login system
  - Email verification
  - Session management with cookies
  - Password encryption using bcrypt

- **Posts & Comments**
  - Create, read, and delete posts
  - Comment on posts
  - Associate categories with posts
  - File attachments for posts

- **Interactive Features**
  - Like/dislike posts and comments
  - Real-time updates for reactions
  - User profile management
  - Avatar/profile picture support

- **Content Organization**
  - Category-based post filtering
  - View created posts
  - View liked posts
  
- **Image Upload Constraints**
  - Maximum image size: 20 MB
  - If an image exceeds 20 MB, an error message will inform the user that the image is too large

## Technology Stack

- **Backend**
  - Go (Standard library)
  - SQLite3 for database
  - bcrypt for password encryption

- **Frontend**
  - Pure HTML/CSS/JavaScript
  - No external frameworks
  - Responsive design
  - Modern UI/UX

### Project Structure
```bash
forum-image-upload/
│
├── authentication/
├── controllers/
│   ├── categories.go
│   ├── categories_test.go
│   ├── filters.go
│   ├── image_handler.go
│   ├── post_handler.go
│   ├── post_handler_test.go
│   └── profile_handler.go
├── static/
├── templates/
├── utils/
├── Dockerfile
├── README.md
├── buildDocker.sh
├── fly.toml
├── go.mod
├── go.sum
└── main.go
```
## Prerequisites

- Go 1.19 or higher
- Docker
- SQLite3

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Bantu-art/forum.git
cd forum
```

## Prerequisites

- Docker

## Building and Running the Application

### Using the `buildDocker.sh` Script


2. **Make the script executable**:

```sh
  chmod +x buildDocker.sh
```

3. **Run the script**:

```sh
  ./buildDocker.sh
```

  This script will:
  - Build the Docker image for the application.
  - Stop and remove any existing container with the same name.
  - Run a new container with the built image.

4. **Access the application**:

  Open your web browser and navigate to `http://localhost:8000`.

### Manually Using Docker Commands

1. **Build the Docker image**:

```sh
  docker build -t forum .
```

2. **Run the Docker container**:

```sh
  docker run -d --name forum-container -p 8000:8000 forum
```


## Troubleshooting

- **Check container logs**:

  If you encounter any issues, you can check the logs of the running container:

```sh
  docker logs forum-container
```

- **Interactive mode**:

You can run the container in interactive mode to debug:

```sh
  docker run -it --name forum-container -p 8000:8000 forum /bin/sh
```

  Once inside the container, you can manually start the application and check for any errors.


## Database Schema

The application uses SQLite with the following main tables:
- Users
- Posts
- Comments
- Categories
- Reactions
- Sessions

## API Endpoints

### Authentication
- `POST /signup` - Register new user
- `POST /signin` - User login
- `POST /signout` - User logout

### Posts
- `GET /` - Get all posts
- `GET /post/{id}` - Get single post
- `POST /post` - Create new post
- `POST /post/delete` - Delete post

### Comments
- `POST /comment` - Add comment
- `POST /comment/delete` - Delete comment
- `POST /comment/edit` - Edit comment

### Reactions
- `POST /react` - Like/dislike post
- `POST /commentreact` - Like/dislike comment

### Filters
- `GET /category/{id}` - Filter posts by category
- `GET /created` - View created posts
- `GET /liked` - View liked posts

## Security Features

- Password encryption using bcrypt
- Session management with UUID
- CSRF protection
- Input validation and sanitization
- Secure cookie handling

### Running Tests
```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
