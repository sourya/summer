# Summer

HTTP <-> File System Adapter

Interact with Linux file system via REST API calls

## Usage

#### `GET /path/to/file`

Serve a file to the client as a stream.

#### `GET /directory/path/with/slash/`

Serve a directory listing as a JSON document.

#### `PUT /path/to/file`

Recieve a file from the client and save it to the vfs.  The file body is streamed.

#### `PUT /directory/path/with/slash/`

Create a directory

#### `DELETE /path/to/file`

Delete a file.

#### `DELETE /directory/path/with/slash/`

Delete a directory (Recursive)


#### `POST /path/to/target`

Rename (or) copy a file or folder (Recursive)

The client sends a JSON body containing the request information.

 - {"renameFrom": from} - rename a file from `from` to `target`.
 - {"copyFrom": from} - copy a file from `from` to `target`.

## Error handling

| Error Code    | Description                        |
| ------------- |:----------------------------------:|
| 0             | Operation Successful               |
| 1024          | Error reading folder               |
| 1025          | Error reading file                 |
| 1026          | Error creating folder              |
| 1027          | Error creating/writing file        |
| 1028          | Error reading source file          |
| 1029          | Error creating destination file    |
| 1030          | Error copying file                 |
| 1031          | Error reading folder stats         |
| 1032          | Is a file                          |
| 1033          | Error reading destination folder   |
| 1034          | Error creating destination folder  |
| 1035          | Request body not found             |
| 1036          | Error renaming file/folder         |
| 1037          | Unsupported keys in request body   |
