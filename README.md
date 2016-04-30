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
