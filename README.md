# Event-Management-API

## URL Deploy
- **Public:** event-management-api-production-c173.up.railway.app

## Daftar Endpoint

### 1. Authentication

- **POST** `/user`
  - Deskripsi: Register account user dengan email dan password.
  - Request Body:
    ```json
    {
      "email": "yourname@mail.com",
      "pass_word": "yourpassword"
    }
    ```

- **POST** `/login`
  - Deskripsi: Login menggunakan email dan password yang telah dibuat.
  - Token : memberikan token JWT untuk mengakses endpoint berikutnya
  - Request Body:
    ```json
    {
      "email": "yourname@mail.com",
      "pass_word": "yourpassword",
    }
    ```

### 2. Event

- **POST** `/event`
  - Deskripsi: menambahkan data event
  - Request Body:
    ```json
    {
        "title": "title event",
        "description_event": "description event",
        "location_event": "Location event",
        "date_event": "2025-03-12T15:15:00-07:00"
    }
    ```

- **GET** `/event`
  - Deskripsi: Menampilkan semua event yang telah dibuat.
 
- **PUT** `/event/:id`
  - Deskripsi: Update event berdasarkan id event yang ingin di edit
  - Request Body:
    ```json
    {
        "title": "title event",
        "description_event": "description event",
        "location_event": "Location event",
        "date_event": "2025-03-12T15:15:00-07:00"
    }
    
- **DELETE** `/event/:id`
  - Deskripsi: Delete event berdasarkan id event yang ingin dihapus


  
### 3. Ticket

- **POST** `/ticket`
  - Deskripsi: menambahkan data tiket untuk menghadiri event, status tiket berupa Regular ataupun VIP
  - Request Body:
    ```json
    {
        "event_id" : 1, 
        "price" : 25000,
        "status" : "Regular", 
        "quantity" : 2
    }
    ```

- **GET** `/ticket`
  - Deskripsi: Menampilkan semua data tiket termasuk informasi event yang ada pada ticket
 
- **PUT** `/ticket/:id`
  - Deskripsi: Update tiket berdasarkan id tiket yang ingin di edit
  - Request Body:
    ```json
    {
    {
        "event_id" : 1, 
        "price" : 900000,
        "status" : "VIP", 
        "quantity" : 5
    }
    }
    
- **DELETE** `/ticket/:id`
  - Deskripsi: Delete tiket berdasarkan id tiket yang ingin dihapus

### 4. Person/Participant

- **POST** `/person`
  - Deskripsi: menambahkan data partisipan yang ingin mengikuti event, status participant berupa {registered, pending, canceled}.
  Setiap partisipan dapat memiliki banyak tiket
  - Request Body:
    ```json
    {
        "full_name" : "nama participant",
        "address_person" : "alamat participant",
        "status" : "registered",
        "ticket_ids": [1, 2]
    }
    ```

- **GET** `/person`
  - Deskripsi: Menampilkan semua data partisipan termasuk data tiket dan event yang terdaftar

- **PUT** `/person/:id`
  - Deskripsi: Update partisipan berdasarkan id partisipan yang ingin di edit
  - Request Body:
    ```json
    {
    {
        "full_name" : "nama participant",
        "address_person" : "alamat participant",
        "status" : "pending",
        "ticket_ids": [1, 2]
    }
    }
    
- **DELETE** `/person/:id`
  - Deskripsi: Delete partisipan berdasarkan id partisipan yang ingin dihapus
