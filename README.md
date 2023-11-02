# Distributed Systems Programming Laboratory Works

### University: _Technical University of Moldova_
### Faculty: _Computers, Informatics and Microelectronics_
### Department: _Software Engineering and Automatics_
### Author: _Mocanu Liviu_

----

## Abstract
&ensp;&ensp;&ensp; This repository contains the laboratory work tasks on the PAD subject at TUM.

## Running the Project
1. Start Your Kubernetes Cluster
```bash
minikube start
```
2. Deploy the Services

Dockerfiles are already pushed to https://hub.docker.com/repositories/emporion97
```bash
# Gateway Service
kubectl apply -f gateway/redis-deployment.yaml
kubectl apply -f gateway/gateway-deployment.yaml

# Playback Service
kubectl apply -f playback_service/p-postgres-deployment.yaml
kubectl apply -f playback_service/playback-service-deployment.yaml
kubectl apply -f playback_service/playback-service-svc.yaml

# Track Service
kubectl apply -f track_service/t-postgres-deployment.yaml
kubectl apply -f track_service/track-service-deployment.yaml
kubectl apply -f track_service/track-service-svc.yaml

# User Service
kubectl apply -f user_service/mysql-deployment.yaml
kubectl apply -f user_service/user-service-deployment.yaml
kubectl apply -f user_service/user-service-svc.yaml

# Load Balancing
kubectl apply -f gateway/gateway-destinationrule.yaml
kubectl apply -f user_service/user-service-destinationrule.yaml
kubectl apply -f track_service/track-service-destinationrule.yaml
kubectl apply -f playback_service/playback-service-destinationrule.yaml
```
3. Monitor the Services

Wait a bit until services are running.
```bash
kubectl get pods
kubectl get services

kubectl get destinationrule
```

4. Get gateway URL

Copy the URL from:
```bash
minikube service gateway-service
```
in order to know on what address to send requests.

5. Clean Up
```bash
# Gateway Service
kubectl delete -f gateway/gateway-deployment.yaml
kubectl delete -f gateway/redis-deployment.yaml

# Playback Service
kubectl delete -f playback_service/playback-service-svc.yaml
kubectl delete -f playback_service/playback-service-deployment.yaml
kubectl delete -f playback_service/p-postgres-deployment.yaml

# Track Service
kubectl delete -f track_service/track-service-svc.yaml
kubectl delete -f track_service/track-service-deployment.yaml
kubectl delete -f track_service/t-postgres-deployment.yaml

# User Service
kubectl delete -f user_service/user-service-svc.yaml
kubectl delete -f user_service/user-service-deployment.yaml
kubectl delete -f user_service/mysql-deployment.yaml

# Load Balancing
kubectl delete -f gateway/gateway-destinationrule.yaml
kubectl delete -f user_service/user-service-destinationrule.yaml
kubectl delete -f track_service/track-service-destinationrule.yaml
kubectl delete -f playback_service/playback-service-destinationrule.yaml
```

## Lab 2 System Architecture
![Diagram2](https://github.com/liviumocanup/pad_web_proxy/blob/lab2/diagrams/lab2_diagram.png)

## Endpoints mini-documentation

The endpoints will be described in the order they are expected to be called.

User Service first as we need a Bearer Token for the endpoints in other services.
Track Service is expected second as we need tracks to efficiently use Playback Service.

With that being said, all the endpoints will work as expected, no matter the order.

### User Service Endpoints:

#### 1. FindAll
- **Endpoint**: `GET /user/find_all`
- **Description**: Retrieves a list of all users.
- **Parameters**: None
- **Body**: None

#### 2. FindByUsername
- **Endpoint**: `GET /user/?username={username}`
- **Description**: Fetches a user by username.
- **Parameters**: `username`
- **Body**: None

#### 3. FindByID
- **Endpoint**: `GET /user/{id}`
- **Description**: Obtains user information by ID.
- **Parameters**: `id` (user ID)
- **Body**: None

#### 4. Register
- **Endpoint**: `POST /user/register`
- **Description**: Registers a new user.
- **Parameters**: None
- **Body**: JSON with `username` and `password`

#### 5. Login
- **Endpoint**: `POST /user/login`
- **Description**: Authenticates a user.
- **Parameters**: None
- **Body**: JSON with `username` and `password`

#### 6. Validate
- **Endpoint**: `POST /user/validate`
- **Description**: Validates a user's session token. (Exposed for commodity)
- **Parameters**: None
- **Body**: JSON with `token`

#### 7. DeleteByID
- **Endpoint**: `DELETE /user/{id}`
- **Description**: Deletes a user by ID.
- **Parameters**: `id` (user ID)
- **Body**: None
- **Authorization**: Bearer token required



### Track Service Endpoints:
#### 1. FindAll
- **Endpoint**: `GET /track/find_all`
- **Description**: Retrieves all tracks.
- **Parameters**: None
- **Body**: None

#### 2. GetInfoByID
- **Endpoint**: `GET /track/{id}`
- **Description**: Fetches track info by ID.
- **Parameters**: `id` (track ID)
- **Body**: None

#### 3. Upload Track
- **Endpoint**: `POST /track/upload`
- **Description**: Uploads a track, file in S3 and info in DB.
- **Parameters**: None
- **Body**: `formdata` with mp3 file and JSON track info
- **Authorization**: Bearer token required

#### 4. Edit Track Info
- **Endpoint**: `PUT /track/edit`
- **Description**: Edits a track's metadata.
- **Parameters**: None
- **Body**: JSON with `trackId` and `metadata`
- **Authorization**: Bearer token required

#### 5. DeleteByID
- **Endpoint**: `DELETE /track/{id}`
- **Description**: Deletes a track by ID.
- **Parameters**: `id` (track ID)
- **Body**: None
- **Authorization**: Bearer token required



### Playback Service Endpoints:

#### 1. Create Playlist
- **Endpoint**: `POST /playback/create`
- **Description**: Creates a new playlist.
- **Parameters**: None
- **Body**: JSON with `name` of the playlist
- **Authorization**: Bearer token required

#### 2. Remove Playlist
- **Endpoint**: `DELETE /playback/remove`
- **Description**: Removes a playlist.
- **Parameters**: None
- **Body**: JSON with `playlistId`
- **Authorization**: Bearer token required

#### 3. Add Tracks
- **Endpoint**: `POST /playback/add_tracks`
- **Description**: Adds tracks to a playlist.
- **Parameters**: None
- **Body**: JSON with `playlistId` and array of `trackIds`
- **Authorization**: Bearer token required

#### 4. Remove Tracks
- **Endpoint**: `DELETE /playback/remove_tracks`
- **Description**: Removes tracks from a playlist.
- **Parameters**: None
- **Body**: JSON with `playlistId` and array of `trackIds`
- **Authorization**: Bearer token required

#### 5. GetPlaylistByID
- **Endpoint**: `GET /playback/{id}`
- **Description**: Retrieves a playlist by ID.
- **Parameters**: `id` (playlist ID)
- **Body**: None
- **Authorization**: Bearer token required

#### 6. Play Playlist
- **Endpoint**: `GET /playback/play/{id}`
- **Description**: Plays a playlist by ID. Creates a downloads folder in gateway and stores the mp3 files for the playlist, categorized by user-named folders and playlists.
- **Parameters**: `id` (playlist ID)
- **Body**: None
- **Authorization**: Bearer token required


## Lab 1 Checkpoint 1:

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

### Conceptually Explained Endpoints (Lab 1 and Deprecated for Lab 2)
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