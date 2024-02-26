## Online Judge Engine Server

---
**Command Line Interface tool follow the Linux POSIX Standard**

---

## Basic setup

1. Install golang dependencies
    
    ```shell
   go mod download
   ```

2. `Ent` generate
   
    ```shell
  go generate ./ent
   ```

3. Compile Application
   
    ```shell
   go build -o app
   ```

## Command for run application

```shell
./app run
```
**Available Flags**

- `-p`: Specify application port

## Command for docker config

### View docker-compose project service list

```shell
./app docker list
```

### Build project docker image

```shell
./app docker build
```

### Docker compose up

```shell
./app docker up
```

**Available Flags**

- `-s`: Select specific service to run

### Docker compose down

```shell
./app docker down
```

- `-s`" Select specific service to stop
