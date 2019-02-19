# EverybodyVotes-GO
A custom server for the Everybody Votes Channel on the PSP.

## API Endpoints
- **GET** `/polls` - Responds with the latest poll. TODO: Add fetching via the ID.
- **POST** `/register` - Registers a console.
- **POST** `/unregister` - Deletes the account of a console. __Cannot be undone.__
- **POST** `/vote` - Let's a console vote on a poll.

#### Note: __All__ endpoints require a Token, except `/register`.