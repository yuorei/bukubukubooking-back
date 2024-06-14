table "AUTHORS" {
  schema = schema.bukubukubooking
  column "id" {
    null = false
    type = varchar(255)
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  primary_key {
    columns = [column.id]
  }
  index "name" {
    unique  = true
    columns = [column.name]
  }
}
table "BOOKS" {
  schema = schema.bukubukubooking
  column "id" {
    null = false
    type = varchar(255)
  }
  column "title" {
    null = false
    type = varchar(255)
  }
  column "publisher" {
    null = true
    type = varchar(255)
  }
  column "owner" {
    null = false
    type = varchar(255)
  }
  column "status" {
    null = false
    type = varchar(255)
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  column "deleted_at" {
    null = true
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "BOOKS_ibfk_1" {
    columns     = [column.owner]
    ref_columns = [table.USERS.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "owner" {
    columns = [column.owner]
  }
}
table "BOOKS_AUTHORS" {
  schema = schema.bukubukubooking
  column "id" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "book_id" {
    null = true
    type = varchar(255)
  }
  column "author_id" {
    null = true
    type = varchar(255)
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "BOOKS_AUTHORS_ibfk_1" {
    columns     = [column.book_id]
    ref_columns = [table.BOOKS.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "BOOKS_AUTHORS_ibfk_2" {
    columns     = [column.author_id]
    ref_columns = [table.AUTHORS.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "author_id" {
    columns = [column.author_id]
  }
  index "book_id" {
    columns = [column.book_id]
  }
}
table "BOOKS_TAGS" {
  schema = schema.bukubukubooking
  column "id" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "book_id" {
    null = true
    type = varchar(255)
  }
  column "tag_id" {
    null = true
    type = varchar(255)
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "BOOKS_TAGS_ibfk_1" {
    columns     = [column.book_id]
    ref_columns = [table.BOOKS.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "BOOKS_TAGS_ibfk_2" {
    columns     = [column.tag_id]
    ref_columns = [table.TAGS.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "book_id" {
    columns = [column.book_id]
  }
  index "tag_id" {
    columns = [column.tag_id]
  }
}
table "BOOK_LIKES" {
  schema = schema.bukubukubooking
  column "id" {
    null = false
    type = varchar(255)
  }
  column "book_id" {
    null = true
    type = varchar(255)
  }
  column "user_id" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "BOOK_LIKES_ibfk_1" {
    columns     = [column.book_id]
    ref_columns = [table.BOOKS.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "BOOK_LIKES_ibfk_2" {
    columns     = [column.user_id]
    ref_columns = [table.USERS.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "book_id" {
    columns = [column.book_id]
  }
  index "user_id" {
    columns = [column.user_id]
  }
}
table "MESSAGES" {
  schema = schema.bukubukubooking
  column "id" {
    null = false
    type = varchar(255)
  }
  column "request_id" {
    null = true
    type = varchar(255)
  }
  column "sender_id" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "MESSAGES_ibfk_1" {
    columns     = [column.request_id]
    ref_columns = [table.REQUESTS.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "MESSAGES_ibfk_2" {
    columns     = [column.sender_id]
    ref_columns = [table.USERS.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "request_id" {
    columns = [column.request_id]
  }
  index "sender_id" {
    columns = [column.sender_id]
  }
}
table "REQUESTS" {
  schema = schema.bukubukubooking
  column "id" {
    null = false
    type = varchar(255)
  }
  column "book_id" {
    null = false
    type = varchar(255)
  }
  column "sender_user_id" {
    null = false
    type = varchar(255)
  }
  column "receiver_user_id" {
    null = false
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  column "finished_at" {
    null = true
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "REQUESTS_ibfk_1" {
    columns     = [column.book_id]
    ref_columns = [table.BOOKS.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "REQUESTS_ibfk_2" {
    columns     = [column.sender_user_id]
    ref_columns = [table.USERS.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "REQUESTS_ibfk_3" {
    columns     = [column.receiver_user_id]
    ref_columns = [table.USERS.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "book_id" {
    columns = [column.book_id]
  }
  index "receiver_user_id" {
    columns = [column.receiver_user_id]
  }
  index "sender_user_id" {
    columns = [column.sender_user_id]
  }
}
table "TAGS" {
  schema = schema.bukubukubooking
  column "id" {
    null = false
    type = varchar(255)
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  primary_key {
    columns = [column.id]
  }
  index "name" {
    unique  = true
    columns = [column.name]
  }
}
table "USERS" {
  schema = schema.bukubukubooking
  column "id" {
    null = false
    type = varchar(255)
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  column "email" {
    null = true
    type = varchar(255)
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  column "deleted_at" {
    null = true
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  index "email" {
    unique  = true
    columns = [column.email]
  }
}
schema "bukubukubooking" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
