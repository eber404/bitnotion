# Bitcoin Price to Notion

This Go application fetches the current Bitcoin price from a cryptocurrency API and automatically updates a specified Notion page with this data every 15 minutes.

## Features

*   **Automated Price Fetching:** Retrieves Bitcoin prices for various currencies.
*   **Notion Integration:** Updates a dedicated Notion page with the latest price information.
*   **Scheduled Updates:** Runs every 15 minutes to keep the Notion page current.
*   **Containerized:** Easy deployment and management using Docker.

## Setup

### Prerequisites

*   Go (version 1.21 or higher)
*   Docker and Docker Compose
*   A Notion integration token and the ID of the Notion page you wish to update.

### Configuration

1.  Create a `.env` file in the project root with the following variables:

    ```
    NOTION_API_KEY="your_notion_integration_token"
    # The Notion Page ID to update. You can find this in the page URL.
    # Example: https://www.notion.so/yourworkspace/Page-Title-2a35403ae33981cfa7f4eb25bd713204
    # The ID is "2a35403ae33981cfa7f4eb25bd713204"
    NOTION_PAGE_ID="your_notion_page_id"
    ```
    **Note:** The `NOTION_PAGE_ID` is currently hardcoded in `internal/services/notion_service.go` as `pageId = "2a35403ae33981cfa7f4eb25bd713204"`. You will need to update this constant in the code or modify the `notion_service.go` to read it from the environment variable.

## Usage

To build and run the application using Docker Compose:

```bash
docker-compose up --build -d
```

This command will:
1.  Build the Docker image for the application.
2.  Start the application service in detached mode (runs in the background).
3.  The application will fetch and update the Notion page once on startup, and then every 15 minutes thereafter.

### Viewing Logs

To view the application logs:

```bash
docker-compose logs -f
```

### Stopping the Service

To stop and remove the running container:

```bash
docker-compose down
```
