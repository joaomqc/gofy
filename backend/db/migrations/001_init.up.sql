CREATE TABLE artist (
    id INTEGER PRIMARY KEY,
    spotify_id TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    monitor BOOLEAN NOT NULL,
    date_added DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE album (
    id INTEGER PRIMARY KEY,
    spotify_id TEXT NOT NULL UNIQUE,
    artist_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    album_type TEXT NOT NULL,
    monitor BOOLEAN NOT NULL,
    release_date DATETIME NOT NULL,
    date_added DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (artist_id) REFERENCES artist(artist_id)
);

CREATE TABLE track (
    id INTEGER PRIMARY KEY,
    spotify_id TEXT NOT NULL UNIQUE,
    album_id INTEGER,
    playlist_id INTEGER,
    title TEXT NOT NULL,
    monitor BOOLEAN NOT NULL,
    date_added DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (album_id) REFERENCES album(album_id)
    FOREIGN KEY (playlist_id) REFERENCES playlist(playlist_id)
);

CREATE TABLE playlist (
    id INTEGER PRIMARY KEY,
    spotify_id TEXT NOT NULL UNIQUE,
    title TEXT NOT NULL,
    monitor BOOLEAN NOT NULL,
    date_added DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);