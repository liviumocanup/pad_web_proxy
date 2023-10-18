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
```
3. Monitor the Services

Wait a bit until services are running.
```bash
kubectl get pods
kubectl get services
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
```

## Endpoints (postman.json)
```json
{
	"info": {
		"_postman_id": "28151a59-633d-477d-a359-75fd06cf361a",
		"name": "PAD",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "18071230"
	},
	"item": [
		{
			"name": "user_service",
			"item": [
				{
					"name": "FindAll",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://{{gateway}}/user/find_all",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"user",
								"find_all"
							]
						}
					},
					"response": []
				},
				{
					"name": "FindByUsername",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://{{gateway}}/user/?username=test2",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"user",
								""
							],
							"query": [
								{
									"key": "username",
									"value": "test2"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "FindByID",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://{{gateway}}/user/2",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"user",
								"2"
							]
						}
					},
					"response": []
				},
				{
					"name": "Register",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"username\": \"test\",\n    \"password\": \"1234\"\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://{{gateway}}/user/register",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"user",
								"register"
							]
						}
					},
					"response": []
				},
				{
					"name": "Login",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"username\": \"test\",\n    \"password\": \"1234\"\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://0.0.0.0:8000/user/login/",
							"protocol": "http",
							"host": [
								"0",
								"0",
								"0",
								"0"
							],
							"port": "8000",
							"path": [
								"user",
								"login",
								""
							]
						}
					},
					"response": []
				},
				{
					"name": "Validate",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc2Nzg3MjQsInVzZXJJZCI6NH0.xKkfjFflvTlg9LZ5d1vayRzNAAg-yDy406yQfNkmrVU\"\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://{{gateway}}/user/validate",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"user",
								"validate"
							]
						}
					},
					"response": []
				},
				{
					"name": "DeleteByID",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc2Nzg3MjQsInVzZXJJZCI6NH0.xKkfjFflvTlg9LZ5d1vayRzNAAg-yDy406yQfNkmrVU",
									"type": "string"
								}
							]
						},
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "http://{{gateway}}/user/4",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"user",
								"4"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "track_service",
			"item": [
				{
					"name": "FindAll",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://0.0.0.0:8000/track/find_all",
							"protocol": "http",
							"host": [
								"0",
								"0",
								"0",
								"0"
							],
							"port": "8000",
							"path": [
								"track",
								"find_all"
							]
						}
					},
					"response": []
				},
				{
					"name": "GetInfoByID",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://0.0.0.0:8000/track/6",
							"protocol": "http",
							"host": [
								"0",
								"0",
								"0",
								"0"
							],
							"port": "8000",
							"path": [
								"track",
								"6"
							]
						}
					},
					"response": []
				},
				{
					"name": "Upload Track",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTczNzc2MjQsInVzZXJJZCI6MX0.RZgn5rC7-kr0nIVlSy4upJ4-6NW14oFBA2HeTO157-s",
									"type": "string"
								}
							]
						},
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"title\": \"Pre-Test Song 1\",\n    \"artist\": \"Pre-Test Dude\",\n    \"album\": \"Pre-test Album\",\n    \"genre\": \"tests\"\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://0.0.0.0:8000/track/upload/",
							"protocol": "http",
							"host": [
								"0",
								"0",
								"0",
								"0"
							],
							"port": "8000",
							"path": [
								"track",
								"upload",
								""
							]
						}
					},
					"response": []
				},
				{
					"name": "Edit Track Info",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTczNzc2MjQsInVzZXJJZCI6MX0.RZgn5rC7-kr0nIVlSy4upJ4-6NW14oFBA2HeTO157-s",
									"type": "string"
								}
							]
						},
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"trackId\": \"5\",\n    \"metadata\": {\n        \"title\": \"Pre-Test Edit Song 1\",\n        \"artist\": \"Pre-Test Dude\",\n        \"album\": \"Pre-test Edit Album\",\n        \"genre\": \"tests\"\n    }\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://0.0.0.0:8000/track/edit_info/",
							"protocol": "http",
							"host": [
								"0",
								"0",
								"0",
								"0"
							],
							"port": "8000",
							"path": [
								"track",
								"edit_info",
								""
							]
						}
					},
					"response": []
				},
				{
					"name": "DeleteByID",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTczNzc2MjQsInVzZXJJZCI6MX0.RZgn5rC7-kr0nIVlSy4upJ4-6NW14oFBA2HeTO157-s",
									"type": "string"
								}
							]
						},
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "http://0.0.0.0:8000/track/delete_by_id/?track_id=5",
							"protocol": "http",
							"host": [
								"0",
								"0",
								"0",
								"0"
							],
							"port": "8000",
							"path": [
								"track",
								"delete_by_id",
								""
							],
							"query": [
								{
									"key": "track_id",
									"value": "5"
								}
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "playback_service",
			"item": [
				{
					"name": "Create Playlist",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc3MjYyODcsInVzZXJJZCI6MX0.u-ehaXPT__HQVJqTHGrBAiFnQ82HQs37W-NZ7A5PljA",
									"type": "string"
								}
							]
						},
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"name\": \"Working Playlist 1\"\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://{{gateway}}/playback/create",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"playback",
								"create"
							]
						}
					},
					"response": []
				},
				{
					"name": "Remove Playlist",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc2NjM5MjUsInVzZXJJZCI6MX0.QzD7S87JUeEu7GmK32U95F3E4NiUikza4Ku-oQthsYM",
									"type": "string"
								}
							]
						},
						"method": "DELETE",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"playlistId\": \"8\"\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://{{gateway}}/playback/remove",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"playback",
								"remove"
							]
						}
					},
					"response": []
				},
				{
					"name": "Add Tracks",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc2NjM5MjUsInVzZXJJZCI6MX0.QzD7S87JUeEu7GmK32U95F3E4NiUikza4Ku-oQthsYM",
									"type": "string"
								}
							]
						},
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"playlistId\": \"1\",\n    \"trackIds\": [\n        \"3\"\n    ]\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://{{gateway}}/playback/add_tracks",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"playback",
								"add_tracks"
							]
						}
					},
					"response": []
				},
				{
					"name": "Remove Tracks",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc2NjM5MjUsInVzZXJJZCI6MX0.QzD7S87JUeEu7GmK32U95F3E4NiUikza4Ku-oQthsYM",
									"type": "string"
								}
							]
						},
						"method": "DELETE",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"playlistId\": \"8\",\n    \"trackIds\": [\n        \"20\"\n    ]\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://{{gateway}}/playback/remove_tracks",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"playback",
								"remove_tracks"
							]
						}
					},
					"response": []
				},
				{
					"name": "GetPlaylistByID",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc2NTUwMDEsInVzZXJJZCI6Mn0.TsQGQXKoQhl4fVAydfBMc-rod6P88tZdTuMklD3gGRc",
									"type": "string"
								}
							]
						},
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://0.0.0.0:8000/playback/1",
							"protocol": "http",
							"host": [
								"0",
								"0",
								"0",
								"0"
							],
							"port": "8000",
							"path": [
								"playback",
								"1"
							]
						}
					},
					"response": []
				},
				{
					"name": "Play Playlist",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc2NjM5MjUsInVzZXJJZCI6MX0.QzD7S87JUeEu7GmK32U95F3E4NiUikza4Ku-oQthsYM",
									"type": "string"
								}
							]
						},
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://{{gateway}}/playback/play/8",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"playback",
								"play",
								"8"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "status",
			"item": [
				{
					"name": "service discovery",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://{{gateway}}/status/discovery",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"status",
								"discovery"
							]
						}
					},
					"response": []
				},
				{
					"name": "gateway",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://{{gateway}}/status",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"status"
							]
						}
					},
					"response": []
				},
				{
					"name": "user",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://{{gateway}}/user/status",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"user",
								"status"
							]
						}
					},
					"response": []
				},
				{
					"name": "track",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://{{gateway}}/track/status",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"track",
								"status"
							]
						}
					},
					"response": []
				},
				{
					"name": "playback",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://{{gateway}}/playback/status",
							"protocol": "http",
							"host": [
								"{{gateway}}"
							],
							"path": [
								"playback",
								"status"
							]
						}
					},
					"response": []
				}
			]
		}
	],
	"variable": [
		{
			"key": "gateway",
			"value": "192.168.58.2:30168"
		}
	]
}
```
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