# Hospital-Management

## Overview
Hospital Management focuses on key aspects such as patient management, appointment scheduling, billing, and staff management.
The systeam supports the following functionalities:
-	Patient management: Register patients, update patient records, manage medical history, retrieve patient details
-	Appointment scheduling: allow patients to book, reschedule, and cancel appointments with doctors, update appointment status and manage doctor availability.
-	Billing: handle patient billing.
-	Staff management: Manage doctor information, scheduling, and assigning tasks.

     Note: “Tasks” are internal work for the doctors, tasks are not appointment.
     For simplicity, all doctors are assume to work from 7 – 11 am and 13 – 17 pm. Appointments last 15 minutes each and be scheduled at even 15-minute intervals within an hour: 0–15 minutes, 15–30 minutes, 30–45 minutes, or 45–60 minutes.
---

## User stories
Patient:
-	As a patient, I need to register an account.
-	As a patient, I need to change my account’s information.
-	As a patient, I need to see my medical bill history.
-	As a patient, I need to book an appointment with a doctor.
-	As a patient, I need to reschedule my appointments.
-	As a patient, I need to cancel my appointments.
-	As a patient, I need to view doctor’s availability.
-	As a patient, I need to view my appointments.

Doctors:
-	As a doctor, I need to update my account’s information.
-	As a doctor, I need to view my appointments.
-	As a doctor, I need to cancel an appointment with my patient.
-	As a doctor, I need to generate a treatment plan invoice.
     Cashing officer:
-	As a billing officer, I need to change the status of a bill from unpaid to paid.
     
Manager:
-	As a manager, I need to register a new account for a doctor.
-	As a manager, I need to update the doctor account’s information.
-	As a manager, I need to delete an account of a doctor.
-	As a manager, I need to assign a task to a doctor.
-	As a manager, I need to view the staff’s workload.
-	As a manager, I need to cancel a task of my doctor.
---

## APIs
User APIs:
-	API to register new account (patients can register for themselves, only managers can register for staff)
-	API to view account information (patients can view themselves, staff can view themselves and their patients, manager can view themselves and their doctors)
-	API to update account information (every users can update themselves).
-	API to delete account. (only managers can delete staff’s account)
     Patient management:
-	API to retrieve patient medical history.
     Appointment scheduling:
-	API to create an appointments with doctors. (for patients)
-	API to update appointment time. (only patients can change the time, if the doctor is available)
-	API to update appointment status. (Scheduled, Completed, canceled)
-	API to delete an appointment. (both patients and doctors can delete)
-	API to retrieve appointments of a patient.
-	API to retrieve appointments of a doctor.
-	API to check doctor’s availability.
     Billing:
-	API to create a patient treatment plan invoice. (only doctors can create)
-	API to update patient bill status. (only cashing officer can change)
     Staff management:
-	API to assign a task to a doctor. (only manager can assign)
-	API to view the tasks assigned to a doctor and their appointments. (doctors can view themselves, managers can view their doctors’ tasks)
-	API to delete a task of a doctor (only managers can delete)
---

## Technical choices

- Go v1.24+
- Gin – HTTP framework
- GORM – ORM for PostgreSQL
- Docker + Render

---

## How to run

### Option 1: Docker Compose

1. Clone the repo:
```bash
git clone https://github.com/tienkhoa03/Hospital-Management.git
```
2. Run Docker Compose

```bash
docker compose up --build
```
This command:
- Builds the Go API server
- Runs the service on localhost:8080
- Automatically initialize the database with default schema/config

### Option 2: Manual

1. Clone the repo:
```bash
git clone https://github.com/tienkhoa03/Hospital-Management.git
```
2. Set up dependencies
```bash
cd be
go mod tidy
```
3. Run the APIs:
```bash
cd be/cmd/server
go run main.go
```

- API runs on: http://localhost:8080
- You must manually ensure the database is running and accessible (configured via .env)

### Option 3: Dockerfile (Single-container build)
1. Build the Docker image:
```bash
docker build -t hospital-api .
```
2. Run the container:
```bash
docker run -p 8080:8080 --env-file .env hospital-api
```
- API runs on: http://localhost:8080
- You must manually ensure the database is running and accessible (configured via .env)

---

## API Endpoints


---

## Project Structure

```
be
├── .github/workflows/   # CI/CD workflows (e.g., GitHub Actions)
├── .vscode/             # VSCode editor settings and workspace configuration
├── api/                 # API layer: routing, handlers, middleware
│   ├── handler/         # HTTP request handlers
│   ├── middleware/      # HTTP middlewares (e.g., logging, auth)
│   └── router/          # API route definitions
├── cmd/server/          # Main application entry point (e.g., main.go)
├── config/              # Application configurations and environment loading
├── constant/            # Constant values, enums, status codes
├── internal/            # Internal application logic (domain-driven design)
│   ├── domain/
│   │   ├── dto/         # Data Transfer Objects (used between layers)
│   │   └── entity/      # Core business entities / database models
│   ├── repository/      # Repository interfaces and their implementations
│   └── service/         # Business logic and use cases
├── pkg/                 # Reusable helper packages (e.g., JWT, hashing, utils)
├── .env.template        # Template for environment variables
├── .gitignore           # Git ignore rules
├── docker-compose.yml   # Docker multi-service configuration
├── Dockerfile           # Docker build instructions for the app
├── go.mod               # Go module file (dependencies and module name)
└── go.sum               # Go checksum file (dependency verification)

```

---

## Environment Variables

| Key          | Description        |
| ------------ | ------------------ |
| PORT         | Port server        |
| DB\_USER     | Database username  |
| DB\_PASSWORD | Database password  |
| DB\_NAME     | Database name      |
| DB\_HOST     | Database host      |


Create `.env` file base on `.env.template`.

---

## Testing
- Unit tests cover all core services and edge cases
- Run tests with:
```bash
cd be
go test ./...
```

---

## Deployment

Deploy using **Railway**:

1. Connect your GitHub repository.
2. Set up environment variables according to `.env.template`.
3. Railway will automatically build and deploy the container.

---

## Production

The API has been deployed and is publicly available at:



---
## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

---
