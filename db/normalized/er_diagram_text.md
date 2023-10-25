erDiagram
    WORKSPACE { 
        int id PK
        string name
        string description
        string thumbnail_url
        timestamp date_created
    }
    BOARD {
        int id PK
        int id_workspace FK
        string name
        string description
        string thumbnail_url
        timestamp date_created
    }
    COLUMN {
        int id PK
        int id_board FK
        string name
        string description
        string thumbnail_url
        timestamp date_created
    }
    TASK {
        int id PK
        int id_column FK
        string name
        string description
        string thumbnail_url
        timestamp date_created
        timestamp start
        timestamp end
        int list_position
    }
    ROLE {
        int id PK
        string name
        string description
    }
    TAG {
        int id PK
        string name
        string color
    }
    USER {
        int id PK
        string email
        string password
        string name
        string surname
        string avatar_url
        string description
    }
    TASK_TEMPLATE {
        int id
        json data
    }
    BOARD_TEMPLATE {
        int id
        json data
    }
    CHECKLIST {
        int id PK
        int id_task FK
        string name
        int list_position
    }
    CHECKLIST_ITEM {
        int id PK
        int id_checklist FK
        string name
        bool done
        int list_position
    }
    USER_WORKSPACE {
        int id_user PK, FK
        int id_workspace PK, FK
        int id_role FK
    }
    BOARD_USER {
        int id_user PK, FK
        int id_board PK, FK
    }
    TASK_USER {
        int id_user PK, FK
        int id_task PK, FK
    }
    TASK_EMBEDDING {
        int id PK
        int id_user FK
        int id_task FK
        string url
    }
    SESSION {
        string token PK
        int id_user FK
        timestamp expiration_date
    }
    TAG_TASK {
        int id_tag PK, FK
        int id_task PK, FK
    }
    COMMENT {
        int id
        int id_task
        int id_user
        content content
        timestamp date_created
    }
    COMMENT_REPLY {
        int id_reply PK, FK
        int id_comment FK
    }
    FAVOURITE_BOARDS {
        int id_board PK, FK
        int id_user PK, FK
    }
    USER_TASK_TEMPLATE {
        int id_user PK, FK
        int id_template PK, FK
    }   
    USER_BOARD_TEMPLATE {
        int id_user PK, FK
        int id_template PK, FK
    }   
    REACTION {
        int id PK
        int id_comment FK
        int id_user FK
        string content
    }
    COMMENT_EMBEDDING {
        int id PK
        int id_comment FK
        int id_user FK
        content url
    }

    WORKSPACE ||--o{ BOARD : contains
    WORKSPACE ||--o{ USER_WORKSPACE : associated_with

    BOARD ||--o{ COLUMN : contains
    BOARD ||--o{ BOARD_USER : associated_with
    BOARD ||--o{ FAVOURITE_BOARDS : included_in

    COLUMN ||--o{ TASK : contains

    TASK ||--o{ CHECKLIST : contains
    TASK ||--o{ TASK_EMBEDDING : has
    TASK ||--o{ TAG_TASK : marked_by
    TASK ||--o{ TASK_USER : assigned_to
    TASK ||--o{ COMMENT : has

    CHECKLIST ||--o{ CHECKLIST_ITEM : contains

    ROLE ||--o{ USER_WORKSPACE : used_in

    TAG ||--o{ TAG_TASK : used_by

    USER ||--o{ USER_WORKSPACE : has_access_to
    USER ||--o{ BOARD_USER : has_access_to
    USER ||--o{ TASK_USER : assigned_to
    USER ||--o{ FAVOURITE_BOARDS : has
    USER ||--o{ USER_TASK_TEMPLATE : uses
    USER ||--o{ USER_BOARD_TEMPLATE : uses
    USER ||--o{ TASK_EMBEDDING : uploaded
    USER ||--o{ COMMENT_EMBEDDING : uploaded
    USER ||--o{ COMMENT : made
    USER ||--o{ REACTION : made
    USER ||--|| SESSION : has

    COMMENT ||--o{ COMMENT_EMBEDDING : has
    COMMENT ||--o{ REACTION : has
    COMMENT ||--o{ COMMENT_REPLY : replied_to
    COMMENT ||--o{ COMMENT_REPLY : replies_to

    TASK_TEMPLATE ||--o{ USER_TASK_TEMPLATE : used_by

    BOARD_TEMPLATE ||--o{ USER_BOARD_TEMPLATE : used_by