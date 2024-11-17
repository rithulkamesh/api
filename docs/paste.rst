Pastebin Service
==============

A secure pastebin service for storing and retrieving text snippets.

API Endpoints
-----------

Create Paste
~~~~~~~~~~~

Create a new paste with your content.

**Endpoint:** ``POST /api/paste/``

**Headers:**

- ``Authorization``: Your API token
- ``Content-Type``: ``application/json``

**Request Body:**

The content is expected to be base64 encoded.

.. code-block:: json

    {
        "content": "Your paste content here"
    }


**Response:**

.. code-block:: json

    {
        "id": "paste_id",
        "url": "http://localhost:8080/api/paste/paste_id"
    }

Retrieve Paste
~~~~~~~~~~~~

Retrieve an existing paste by its ID.

**Endpoint:** ``GET /api/paste/:id``

**Parameters:**

- ``:id``: The unique identifier of the paste

**Response:**

.. code-block:: json

    {
        "content": "Your paste content",
        "created_at": "2024-11-17T12:00:00Z"
    }

Example Usage
-----------

Creating a paste:

.. code-block:: bash

    curl -X POST \
      -H "Authorization: your-api-token" \
      -H "Content-Type: application/json" \
      -d "{\"content\": \"$(echo "Hello, World!" | base64)\"}" \
      http://localhost:8080/api/paste/

Retrieving a paste:

.. code-block:: bash

    curl http://localhost:8080/api/paste/paste_id

Error Handling
-------------

The API returns standard HTTP status codes:

- ``200``: Success
- ``401``: Unauthorized (invalid or missing API token)
- ``404``: Paste not found
- ``500``: Server error
