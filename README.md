## Endpoints Specification

### Authentication Controller (`auth_controller.go`)

- `POST /api/users/register` - Register a new user.
- `POST /api/users/login` - User login (generating JWT).
- `POST /api/users/reset-password` - Request a password reset.
- `POST /api/users/verify-email` - Verify user's email.

### User Controller (`user_controller.go`)

- `GET /api/users/{username}` - Get user profile.
- `PUT /api/users/{username}` - Update user data (e.g., username, email, bio, profile picture).
- `DELETE /api/users/{username}` - Delete the user's account.

### Opinion Controller (`opinion_controller.go`)

- `POST /api/opinions` - Create a new opinion (title and body).
- `GET /api/opinions/{opinionID}` - Get a specific opinion.
- `GET /api/opinions` - Get a list of all opinions.
- `PUT /api/opinions/{opinionID}` - Update an opinion.
- `DELETE /api/opinions/{opinionID}` - Delete an opinion.
- `GET /api/users/{username}/opinions` - Get opinions for a specific user.
- `GET /api/opinions/{opinionID}/likes` - Get a list of users who liked a specific opinion.

### Follow Controller (`follow_controller.go`)

- `POST /api/follow/{username}` - Follow another user.
- `DELETE /api/follow/{username}` - Unfollow another user.

### Like Controller (`like_controller.go`)

- `POST /api/likes/{opinionID}` - Like an opinion.
- `DELETE /api/likes/{opinionID}` - Unlike an opinion.
