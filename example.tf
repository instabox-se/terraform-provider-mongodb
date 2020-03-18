provider "mongodb" {
  // Please note that you need a user with userAdmin role or similar to
  // be able to use this provider. It is not recommended to manage this user
  // through Terraform :)
  connection_string = "mongodb://localhost:27017"
}

resource "mongodb_user" "user1" {
  allow_password_update = false
  username = "user1"
  db       = "admin"
  name     = "User1 Userson"
  password = "s3cr37" // This is stored as SHA-512 hash in state

  role {
    role = mongodb_role.role1.name
    db   = "admin"
  }
}

resource "mongodb_role" "role1" {
  name = "role1"
  db   = "admin"

  role {
    role = "read" // This is a built-in role in MongoDB
    db   = "production"
  }
  role {
    role = "readWrite" // This is a built-in role in MongoDB
    db   = "development"
  }

  privilege {
    cluster = true
    actions = ["listDatabases"]
  }

  privilege {
    db = "admin"
    collection = "*"
    actions = ["find"]
  }
}
