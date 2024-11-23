# 🎵 Song Management API

A REST API for managing songs, lyrics, and music-related content. This service provides comprehensive endpoints for creating, retrieving, updating, and deleting song information.

## ✨ Features

- Full CRUD operations for songs
- Pagination support for large collections
- Flexible filtering options (by group, song name, release date, etc.)
- Verse-level access to song lyrics
- Comprehensive error handling
- Input validation

## 🚀 API Endpoints

### Songs Management

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/song` | Create a new song |
| GET | `/songs` | Get paginated list of songs with filters |
| GET | `/song-verse` | Get song with verse-level pagination |
| PATCH | `/song/{id}` | Update existing song |
| DELETE | `/song/{id}` | Delete a song |

### Filtering Options

The `/songs` endpoint supports multiple filtering parameters:

- `page`: Page number for pagination
- `limit`: Number of items per page
- `group`: Filter by artist/group name
- `song`: Filter by song title
- `releaseDate`: Filter by release date
- `text`: Search within song lyrics
- `link`: Filter by associated links

## 📝 Request Examples

### Creating a New Song

```json
POST /song
{
    "group": "Muse",
    "song": "Supermassive Black Hole"
}
```

### Updating a Song

```json
PATCH /song/{id}
{
    "group": "Eminem",
    "song": "Song Title",
    "text": "Lyrics content...",
    "releaseDate": "DD.MM.YYYY",
    "link": "http://example.com"
}
```

## 📊 Response Examples

### Successful Song Creation

```json
{
    "group": "Muse",
    "song": "Supermassive Black Hole",
    "text": "Ooh baby, don't you know I suffer?...",
    "releaseDate": "16.07.2006",
    "link": "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
}
```

### Paginated Songs Response

```json
{
    "page": 1,
    "limit": 10,
    "total": 31,
    "total_pages": 4,
    "songs": [
        {
            "group": "Muse",
            "song": "Supermassive Black Hole",
            "text": "...",
            "releaseDate": "16.07.2006",
            "link": "..."
        }
        // ... more songs
    ]
}
```

## ⚠️ Error Handling

The API implements comprehensive error handling with appropriate HTTP status codes:

- `400`: Bad Request - Invalid input data
- `404`: Not Found - Requested resource doesn't exist
- `422`: Unprocessable Entity - Validation failed
- `500`: Internal Server Error

Each error response includes a consistent structure:

```json
{
    "code": 404,
    "message": "Not Found - Song not found"
}
```


## 🛠️ Technical Details

- API Version: Swagger 2.0
- Content Type: application/json
- Response Format: JSON
- Fiber v.2



---

Made with ❤️ by Zelmoron
