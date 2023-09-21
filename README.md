# Distributed Systems Programming Laboratory Works

### University: _Technical University of Moldova_
### Faculty: _Computers, Informatics and Microelectronics_
### Department: _Software Engineering and Automatics_
### Author: _Mocanu Liviu_

----

## Abstract
&ensp;&ensp;&ensp; This repository contains the laboratory work tasks on the PAD subject at TUM.

## Checkpoint 1:

### Application Suitability:
The music streaming service consists implementation through distributed systems is suitable since :
* It consists of multiple components that have a distinct functionality and can be developed, deployed, scaled, and maintained independently.
* Different components of the application may experience different load levels (the playback service might face heavy loads during peak times or when a new song goes viral).
* As each service is independent, it can be deployed without affecting others, which promotes continuous integration and deployment practices.

Real-life examples:
- ***Spotify*** - The microservices architecture enables Spotify to handle a large traffic volume and multiple teams simultaneously working on different services. Microservices have been instrumental in Spotify's success as a music streaming service. The company has over 1,600 production services running on Kubernetes with around 100 microservices being added per month, all being monitored with Backstage.
- ***SoundCloud*** - Originally built as a monolithic application, SoundCloud moved to a microservices architecture to cope with their growing user base and to streamline their development process. This allowed them to deploy features more rapidly and scale individual components as needed.

### Service Boundaries:
- **User Service** - Responsible for user authentication, as well as getting or deleting user data.
- **Track Service** - Responsible for the management and storage of track metadata and the actual audio files.
- **Playback Service** - Responsible for handling the playback functionalities, including playlists.
- **API Gateway** -  Serves as the single entry point for all client requests. It routes the request to the appropriate service and aggregates the responses if necessary. 
- **Service Discorvery** - Keeps a list of all service instances, their locations, and their health status. It helps the API Gateway route client requests to the appropriate service instance.
- **Cache** - Used to store frequently accessed data, reducing the load on the databases and services and increasing response times. For example, frequently accessed tracks or user profiles can be cached.
- **Load Balancer** - Distributes incoming client requests across multiple instances of the services to ensure even load distribution and high availability. While the API Gateway can handle basic load balancing, a dedicated load balancer can offer more advanced strategies and handle higher loads.

![Diagram](https://github.com/liviumocanup/pad_web_proxy/blob/main/diagrams/no_lb_diagram.png)

### Technology Stack 
- **User Service**:
    * Language: Go
    * Framework: Gin, gorm
    * Database: MySQL
- **Track Service**: 
    * Language: Go
    * Framework: Gin, gorm
    * Database: PostgreSQL,S3
- **Playback Service**:
    * Language: Go
    * Framework: Gin, gorm
    * Database: PostgreSQL
- **API Gateway**: 
    * Language: Python
    * Framework: FastAPI
- **Service Discorvery**:
    * Language: Python
    * Framework: FastAPI
- **Cache**:
    * Technology: Redis
- **Load Balancer**:
    * Technology: Nginx

### Communication Patterns
- **Client to Gateway:** RESTful APIs. Given that clients might be diverse (web apps, mobile apps, etc.), REST provides a simple and universally supported interface.
- **Inter-Service Communication:** gRPC. By using gRPC, services can invoke methods in other services as if they were local procedures. This can simplify communication patterns and ensure type safety.

### Data Management
1. **Database per Service**: Each microservice will manage its own database. This ensures loose coupling, as each service has full control over its data model and is not dependent on other services.
2. **API for Data Access**: Services will not access each other's databases directly. Instead, they will use the defined APIs to request any required data from another service. This maintains encapsulation and ensures changes to one service's data model don't impact other services.

### Endpoints
* **User Service**:
    1. User Registration:
        - Endpoint: `/users/register`
        - Method: `POST`
        - Request: 
            ```bash
            {
                "username": "liviu",
                "password": "password123",
                "email": "liviu@email.com"
            }
            ```
        - Response:
            ```bash
            {
                "user_id": "12345"
            }
            ```
    2. User Login:
        - Endpoint: `/users/login`
        - Method: `POST`
        - Request: 
            ```bash
            {
                "username": "liviu",
                "password": "password123"
            }
            ```
        - Response:
            ```bash
            {
                "token": "jwt_token_here"
            }
            ```

    3. Get User Info:
        - Endpoint: `/users/{userID}`
        - Method: `GET`
        - Response:
            ```bash
            {
                "user_id": "12345",
                "username": "liviu",
                "email": "liviu@email.com",
                "created_at": "2023-09-19T10:00:00Z"
            }
            ```

    4. Delete User:
        - Endpoint: `/users/{userID}`
        - Method: `DELETE`
        - Response:
            ```bash
            {
                "message": "User deleted successfully"
            }
            ```

* **Track Service**:
    1. Upload Track:
        - Endpoint: `/tracks/upload`
        - Method: `POST`
        - Request: 
            * Metadata part:
                ```bash
                {
                    "title": "Song Name",
                    "artist": "Artist Name",
                    "album": "Album Name",
                    "genre": "Pop"
                }
                ```
            * File part: The actual music file (MP3).
        - Response:
            ```bash
            {
                "track_id": "6789"
            }
            ```
    2. Get Track Info:
        - Endpoint: `/tracks/{trackID}`
        - Method: `GET`
        - Response:
            ```bash
            {
                "track_id": "6789",
                "title": "Song Name",
                "artist": "Artist Name",
                "album": "Album Name",
                "genre": "Pop",
                "upload_date": "2023-09-19T11:00:00Z"
            }
            ```

    3. Get the Song File:
        - Endpoint: `/tracks/{trackID}/file`
        - Method: `GET`
        - Response: Stream of the music file. This will serve the raw file, allowing other services to retrieve it.

    4. Edit Track Info:
        - Endpoint: `/tracks/upload`
        - Method: `PATCH`
        - Request: 
            * Metadata part:
                ```bash
                {
                    "title": "Song Name",
                    "artist": "Artist Name",
                    "album": "Album Name",
                    "genre": "Pop"
                }
                ```
            * File part: The actual music file (MP3).
        - Response:
            ```bash
            {
                "message": "Track info updated successfully"
            }
            ```

    5. Delete Track:
        - Endpoint: `/tracks/{trackID}`
        - Method: `DELETE`
        - Response:
            ```bash
            {
                "message": "Track deleted successfully"
            }
            ```

* **Playback Service**:
    1. Play Track:
        - Endpoint: `/play/{trackID}`
        - Method: `GET`
        - Response: The service will first fetch the song file from the `Track Service` and then stream it to the user for playback. The response header would contain the MIME type of the music file (like `audio/mpeg` for MP3s) to facilitate the streaming.

    2. Add to Playlist:
        - Endpoint: `/playlist/add`
        - Method: `POST`
        - Request: 
            ```bash
            {
                "user_id": "12345",
                "track_id": "6789"
            }
            ```
        - Response:
            ```bash
            {
                "message": "Track added to playlist",
                "track_info": {
                    "track_id": "6789",
                    "title": "Song Name",
                    "artist": "Artist Name",
                    "album": "Album Name",
                    "duration": "3:45",
                    "url": "streaming_url"
                }
            }
            ```

    3. Remove from the Playlist:
        - Endpoint: `/playlist/remove`
        - Method: `DELETE`
        - Request: 
            ```bash
            {
                "user_id": "12345",
                "track_id": "6789"
            }
            ```
        - Response:
            ```bash
            {
                "message": "Track removed from the playlist"
            }
            ```

### Deployment and Scaling

1. *Containerization with **Docker***: Build Docker images for each service and push them to a container registry (Docker Hub). Go applications can be packaged as lightweight, statically compiled binaries, and Python services can be containerized using a suitable base image.
2. *Orchestration with **Kubernetes***: Given the need for scaling (especially for a music streaming service that might face varying loads) and easy management, deploying with Kubernetes would be best. The `Track Service` might face higher loads during track upload sprees or when multiple services fetch song data. The `Playback Service`, on the other hand, would likely experience spikes during peak listening hours or during specific promotional events. Kubernetes' auto-scaling can adjust the number of pods for each service based on the incoming traffic and load.

This setup will ensure scalability, reliability, and ease of management.