# Scratch Data Demo Project

Hi Jay, 

Here is my submission for the Scratch Data take-home
assignment. Please read the notes below about my approach.

## Project Notes

### Backend
* The backend server is written entirely in `Go`.
* `SQLite` is used as requested.
* The `/query` endpoint behaves as requested.
* The `/data` endpoint can take one or more rows of data at once.
  * It can accept a single standalone object `{col1: val, col2: val...}`
    in a request body to ingest a single row.
  * Additionally, multiple objects can be sent as an array `[{row1...}, {row2...}, ...]` to ingest many rows in a single request.
* As of now, testing is __NOT__ at an acceptable state. 
  The project contains simple tests for the core backend functionalities.
  Given more time, my first priority would be writing more comprehensive backend
  tests.
* You mentioned that my server should be able to handle ~5000 requests per
  second. I believe my server should be able to handle such a load. However, 
  I have not load tested the server, so I cannot promise this as of now.
* The server does have a simple shutdown procedure which closes the database
  before exiting.

#### Build and Run

```bash
# Compile and build the server
cd backend/cmd
go build -o sd

# Run server (listens on port 3000)
./sd        
```

### Frontend

* The frontend is built with `svelte` and does not need a backend server.
* I planned to use `TailwindCSS`, but had faced difficulties integrating it
  with `svelte`. Most of the approaches I saw used `svelteKit`, which I believe
  required a backend server. Ultimately, I decided to use vanilla `CSS`. Given more
  time, I would like to solve this problem.
* The UI has two components: one for running queries, and one for uploading data.

#### Build and Run
```bash
cd frontend
npm install
npm run build

# After these commands static HTML can be found at...
# frontend/public/index.html
```

### Deployment
  
Both the frontend and backend deploy to DigitalOcean servers using
`Ansible`.

```
# Deploy frontend.
ansible-playbook -i deploy/hosts/inventory.ini deploy/frontend.yaml

# Deploy backend.
ansible-playbook -i deploy/hosts/inventory.ini deploy/frontend.yaml
```

__NOTE__: Both playbooks require an sshkey file which is __NOT__ in this repo.
