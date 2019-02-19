# EverybodyVotes-GO
A custom server for the Everybody Votes Channel on the PSP.

## API Endpoints
| Type     | Endpoint      | Requires Token | Description                                                        |
|----------|---------------|----------------|--------------------------------------------------------------------|
| **GET**  | `/polls`      | Yes            | Responds with the latest poll. **TODO: Add fetching via the ID.**  |
| **POST** | `/register`   | Uses MAC       | Registers a console.                                               |
| **POST** | `/unregister` | Yes            | Deletes the account of a console. __Cannot be undone.__            |
| **POST** | `/vote`       | Yes            | Let's a console vote on a poll.                                    |
